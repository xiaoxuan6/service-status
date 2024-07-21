// 获取最后更新时间
async function lastUpdatedtime() {
  const configResponse = await fetch("./config.cfg");
  const configText = await configResponse.text();
  const configLines = configText.split("\n");

  let firstLine;
  for (const line of configLines) {
    // 排除前面带有 # 的注释行
    if (!line.trim().startsWith("#")) {
      firstLine = line;
      break;
    }
  }

  const key = firstLine.split("=")[0];
  const response = await fetch("logs/" + key + "_report.log");
  let statusLines = "";
  if (response.ok) {
    statusLines = await response.text();
  }
  const lines = statusLines.split("\n");

  let lastTime = ''
  if (lines.length < 3) {
    const now = new Date();

    const year = now.getFullYear();
    const month = String(now.getMonth() + 1).padStart(2, '0'); // 月份从0开始，所以需要加1，并确保两位数
    const day = String(now.getDate()).padStart(2, '0');       // 确保两位数
    const hours = String(now.getHours()).padStart(2, '0');    // 确保两位数
    const minutes = String(now.getMinutes()).padStart(2, '0'); // 确保两位数

    lastTime = `${year}-${month}-${day} ${hours}:${minutes}`;
  } else {
    lastTime = lines[lines.length - 2].split(",")[0];
  }

  // 在这里调用 updateLastUpdated 函数并传递 lastTime
  updateLastUpdated(lastTime);
}

// 将最后更新时间写到页面
function updateLastUpdated(lastUpdateTime) {
  const updateTimeElement = document.getElementById("updateTime");
  if (updateTimeElement) {
    updateTimeElement.textContent = `Last updated on : ${lastUpdateTime}`;
  }
}

// 当 DOM 加载完成后调用 lastUpdatedtime 函数
document.addEventListener("DOMContentLoaded", function () {
  lastUpdatedtime();
});