#! /bin/sh

sleep 1
COMMIT_MSG="$1"
if [ -z "$COMMIT_MSG" ]; then
  echo "Commit message not provided! Aborting..."
  exit 1
fi

echo "git add -A"
git add -A

sleep 1
echo "git commit -m" \""$COMMIT_MSG"\"
git commit -m "$COMMIT_MSG"
