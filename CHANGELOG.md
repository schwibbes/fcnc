# Agenda for 0.1

- single-run-mode (A)
-- check all conditions
-- run all actions in order
-- assert state

- service-mode
-- run in loop and execute
-- install as service

- client-server-mode
-- fetch config
-- notify master
-- collect status
-- save in db on master

- some useful queries (B)
-- ?diskspace>(int %)
-- ?app-state(expected)
-- ?http-get-contains(string)
-- ?file-get-contains(string)
-- ?path-exists(string)
-- ?config-setting(provider, key, val)
-- ?cli-returns-contains(string)

- some useful actions (C)
-- !delete-path(string)
-- !run-cli(string)
-- !download-file(url, path)
-- !unzip(src,dst)
-- !config-setting(provider, key, val)

- notifications
-- logfile
-- mail
-- hipchat
-- irc