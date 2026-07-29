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
# "$@" passes through whatever args/CMD the container was started with,
# so this doesn't change how you invoke the binary otherwise.
exec ./server "$@"