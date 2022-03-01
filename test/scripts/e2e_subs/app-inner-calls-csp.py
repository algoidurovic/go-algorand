#!/usr/bin/env python

import more_itertools
import os
import sys
from goal import Goal

from datetime import datetime

stamp = datetime.now().strftime("%Y%m%d_%H%M%S")
print(f"{os.path.basename(sys.argv[0])} start {stamp}")

goal = Goal(sys.argv[1], autosend=True)

joe = goal.new_account()

_, err = goal.pay(goal.account, joe, amt=500_000)
assert not err, err

txinfo1, err = goal.keyreg(joe, nonpart=True)
assert not err, err

# app1 calls the clear state program of app2 (after opting into it) which issues an
# inner app call to app3. This verifies that both accessing a CSP through inner app
# calls and issuing inner app calls from a CSP is possible.
app1 = """
#pragma version 6
 txn ApplicationID
 bz end

 itxn_begin
  int appl
  itxn_field TypeEnum

  txn Applications 1
  dup
  store 0
  itxn_field ApplicationID

  int OptIn
  itxn_field OnCompletion

  txn Applications 2
  itxn_field Applications

 itxn_next
  int appl
  itxn_field TypeEnum

  load 0
  itxn_field ApplicationID

  txn Applications 2
  itxn_field Applications

  int ClearState
  itxn_field OnCompletion

 itxn_submit


end:
 int 1
"""

app2 = """
#pragma version 6
 txn ApplicationID
 bz end

 itxn_begin
  int appl
  itxn_field TypeEnum

  txn Applications 1
  itxn_field ApplicationID
 itxn_submit


 end:
  int 1
"""

app3 = """
#pragma version 6
int 1
"""

goal.autosend = True

ApiError = str
ApprovalProgram = str
ClearStateProgram = str


def create_apps(
        sender: str,
        programs: list[tuple[ApprovalProgram, ClearStateProgram]]) -> (list[dict], list[ApiError]):
    creates = [
        goal.app_create(
            sender,
            goal.assemble(ap),
            None if csp is None else goal.assemble(csp))
        for ap, csp in programs
    ]
    (ts, es) = more_itertools.partition(lambda t: t[1], creates)
    return [[t for t, _ in list(ts)], [e for _, e in list(es)]]


(tx_infos, errors) = create_apps(joe, [(app1, None), (app2, app2), (app3, None)])
assert not errors, errors

app_ids = [tx_info['application-index'] for tx_info in tx_infos]
assert len(app_ids) == 3, app_ids

(app1_id, app2_id, app3_id) = app_ids

# fund the apps
txinfo1, err = goal.pay(goal.account, goal.app_address(app1_id), amt=4_000_000)
assert not err, err

txinfo2, err = goal.pay(goal.account, goal.app_address(app2_id), amt=4_000_000)
assert not err, err

# execute c2c w/ CSP
start_app, err = goal.app_call(joe, app1_id, foreign_apps=[int(app2_id), int(app3_id)])
assert not err, err

print(f"{os.path.basename(sys.argv[0])} OK {stamp}")
