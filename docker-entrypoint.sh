#!/bin/sh
set -e

# Applies any pending "up" migrations (from Go migration files compiled into
# the binary, and/or from pb_migrations/*.js on disk) before the server
# starts serving traffic. This assumes your main.go follows the standard
# PocketBase pattern:
#
#   app := pocketbase.New()
#   migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{ ... })
#   ...
#   app.Start()
#
# which is what registers the `migrate` subcommand. If your app doesn't
# register migratecmd, this step will fail — see DEPLOY.md.
echo "Applying PocketBase migrations..."
./server migrate up

echo "Starting server..."
# Binds on 0.0.0.0 (required so Coolify's proxy/other containers can reach
# it — PocketBase's default 127.0.0.1 bind is not reachable from outside
# the container) on the port from POCKETBASE_PORT (falls back to 8096).
# "$@" still passes through any extra args/flags you add via CMD.
exec ./server serve --http="0.0.0.0:${POCKETBASE_PORT:-8096}" "$@"