---
sidebar_position: 2
---

# Commands

## `flow start [project] [tags]`

Starts a new flow session for the specified project.

| name | default | description                     |
| ---- | ------- | ------------------------------- |
| tags |         | Tags to be used for the session |

example:

```bash
flow start my-project +tag1 +tag2
```

## `flow stop`

Stops the current flow session.

## `flow status`

See the status of the current flow session.

## `flow report`

View a user-friendly report of sessions.

| name     | default | description                                           |
| -------- | ------- | ----------------------------------------------------- |
| --format | by-day  | Format of the report. Options: `by-day`, `by-project` |

## `flow edit [session-id]`

Open the session file in the default editor.

## `flow abort`

Abort the current session.
