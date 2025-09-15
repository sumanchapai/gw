// Run arbitrary git command
document
  .querySelector(".git-command-form")
  .addEventListener("submit", async (e) => {
    e.preventDefault();
    const cmd = document.getElementById("command").value.trim();
    if (!cmd) return;

    const resultBox = document.getElementById("command-result");
    resultBox.textContent = "Running...";
    try {
      const resp = await fetch("/git-command", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ args: cmd.split(" ") }),
      });
      const data = await resp.json();
      resultBox.textContent = data.output || data.error || "(no output)";
      // Refresh diff
      loadDiff();
    } catch (err) {
      resultBox.textContent = "Error: " + err.message;
      // Refresh diff
      loadDiff();
    }
  });

// Switch branch on select change
document
  .getElementById("branch-select")
  .addEventListener("change", async (e) => {
    const branch = e.target.value;
    e.target.disabled = true;
    try {
      const resp = await fetch("/git-command", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ args: ["checkout", branch] }),
      });
      const data = await resp.json();

      if (data.success) {
        alert(`Switched to branch: ${branch}`);
        loadDiff();
      } else {
        alert("Error switching branch:\n" + (data.error || data.output));
        e.target.value = "{{.CurrentBranch}}";
      }
    } catch (err) {
      alert("Request failed: " + err.message);
      e.target.value = "{{.CurrentBranch}}";
    } finally {
      e.target.disabled = false;
    }
  });

// Load git diff on page load
async function loadDiff() {
  const noDiff = "(no diff)";
  const diffBox = document.getElementById("diff-result");
  if (diffBox.textContent === noDiff) {
    diffBox.textContent = "Loading...";
  }
  try {
    const resp = await fetch("/git-command", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ args: ["diff"] }),
    });
    const data = await resp.json();
    const raw = data.output || data.error || noDiff;
    diffBox.innerHTML = highlightDiff(raw);
  } catch (err) {
    diffBox.textContent = "Error: " + err.message;
  }
}

function escapeHtml(str) {
  return str.replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;");
}

const btnCommit = document.getElementById("btn-commit");
const btnCommitPush = document.getElementById("btn-commit-push");
const btnCommitPR = document.getElementById("btn-commit-pr");
const outputBox = document.getElementById("commit-action-output");
const branchSelect = document.getElementById("branch-select");

// helper: run an action on server
async function runGitAction(payload) {
  const spinner = document.getElementById("commit-spinner");
  spinner.style.display = "block"; // show spinner
  outputBox.textContent = ""; // clear old logs

  return new Promise((resolve, reject) => {
    const ws = new WebSocket(`ws://${window.location.host}/git-action`);

    ws.onopen = () => {
      ws.send(JSON.stringify(payload));
    };

    ws.onmessage = (event) => {
      outputBox.textContent += event.data + "\n";
      outputBox.scrollTop = outputBox.scrollHeight;
    };

    ws.onerror = (err) => {
      outputBox.textContent += "\nError: " + err.message;
      spinner.style.display = "none"; // hide spinner on error
      reject(err);
    };

    ws.onclose = () => {
      spinner.style.display = "none"; // hide spinner when finished
      resolve({ success: true });
      try {
        loadDiff();
      } catch (e) {}
    };
  });
}

function setBusy(flag) {
  [btnCommit, btnCommitPush, btnCommitPR].forEach((b) => {
    if (flag) b.classList.add("busy");
    else b.classList.remove("busy");
    b.disabled = flag;
  });
}

// Button handlers:
btnCommit.addEventListener("click", async () => {
  const commitMsg = prompt("Commit message:");
  if (!commitMsg) {
    alert("Aborted: commit message required.");
    return;
  }
  const branch = branchSelect?.value || null;
  await runGitAction({
    action: "commit",
    args: [commitMsg],
  });
});

btnCommitPush.addEventListener("click", async () => {
  const commitMsg = prompt("Commit message:");
  if (!commitMsg) {
    alert("Aborted: commit message required.");
    return;
  }
  const branch = branchSelect?.value || null;
  await runGitAction({
    action: "commit-push",
    branch,
    message: commitMsg,
  });
});

btnCommitPR.addEventListener("click", async () => {
  const commitMsg = prompt("Commit message:");
  if (!commitMsg) {
    alert("Aborted: commit message required.");
    return;
  }

  // Prompt for PR title/body (default to commit message)
  const prTitle = prompt("PR title:", commitMsg) || commitMsg;
  const prBody = prompt("PR body (optional):", "") || "";

  // Optionally choose base (default 'main')
  const prBase = prompt("PR base branch (default: main):", "main") || "main";

  const branch = branchSelect?.value || null;
  await runGitAction({
    action: "commit-pr",
    branch,
    message: commitMsg,
    pr_base: prBase,
    pr_title: prTitle,
    pr_body: prBody,
  });
});

loadDiff();
function highlightDiff(diffText) {
  return diffText
    .split("\n")
    .map((line) => {
      if (line.startsWith("+") && !line.startsWith("+++")) {
        return `<span class="diff-line diff-add">${escapeHtml(line)}</span>`;
      } else if (line.startsWith("-") && !line.startsWith("---")) {
        return `<span class="diff-line diff-remove">${escapeHtml(line)}</span>`;
      } else if (line.startsWith("@@")) {
        return `<span class="diff-line diff-hunk">${escapeHtml(line)}</span>`;
      } else if (
        line.startsWith("diff ") ||
        line.startsWith("index ") ||
        line.startsWith("---") ||
        line.startsWith("+++")
      ) {
        return `<span class="diff-line diff-meta">${escapeHtml(line)}</span>`;
      } else {
        return `<span class="diff-line">${escapeHtml(line)}</span>`;
      }
    })
    .join("");
}
