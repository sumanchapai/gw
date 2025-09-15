#! /bin/sh

echo "Reading Commit Msg..."
sleep 1

COMMIT_MSG="$1"
if [ -z "$COMMIT_MSG" ]; then
  echo "Commit message not provided! Aborting..."
  exit 1
fi

echo "Staging..."
sleep 1

git add -A

echo "Committing..."
sleep 1
git commit -m "$COMMIT_MSG"
