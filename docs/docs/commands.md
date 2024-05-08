---
sidebar_position: 2
---

# Commands

## `flow start [project] [tags]`

Starts a new flow session for the specified project.

| name | default | description                     |
| ---- | ------- | ------------------------------- |
| tags | \       | Tags to be used for the session |

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

| name              | default | description                                           |
| ----------------- | ------- | ----------------------------------------------------- |
| --format [format] | by-day  | Format of the report. Options: `by-day`, `by-project` |
| --day             | /       | Get a report for all sessions of the current day      |
| --week            | /       | Get a report for all sessions of the current week     |
| --project         | /       | Get a report for all sessions of the given project    |
| --since [date]    | /       | Get a report for all sessions since the given date    |
| --until [date]    | /       | Get a report for all sessions until the given date    |

## `flow edit [session-id (optional)]`

Open the session with given ID in the default editor. If no ID is provided, it
will open the last session

## `flow abort`

Abort the current session.
