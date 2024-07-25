import { reslogs } from './reslogs.js';
import { updateChart } from './timelapsechart.js';
import { getColor, getStatusText, constructStatusStream } from './utils.js';
import { normalizeData } from './dataProcessing.js';
import { create } from './domManipulation.js';
import { scrolltoright } from './scroll.js';
import { fetchUrlsConfig } from './fetchurlsconfig.js';

/**
 * 异步函数：根据 config.cfg 文件，生成所有报告
 * @param {string} urlspath - 配置文件的路径，其中包含需要生成报告的 URL 列表。
 */
export async function genAllReports() {
  const configLines = await fetchUrlsConfig();
  for (let ii = 0; ii < configLines.length; ii++) {
    const configLine = configLines[ii];
    const [key, url] = configLine.split("=");
    await genReportLog(document.getElementById("reports"), key, url);
  }
  scrolltoright();
}

/**
 * 异步生成报告日志。
 * @param {HTMLElement} container - 用于装载报告日志的容器元素。
 * @param {string} key - 报告日志的唯一标识键。
 * @param {string} url - 相关 URL，用于报告中显示。
 * @param {string} logspath - 日志文件的路径。
 */
async function genReportLog(container, key, url) {
  const response = await reslogs(key);
  let statusLines = "";
  if (response.ok) {
    statusLines = await response.text();
  }
  const normalized = normalizeData(statusLines);
  const statusStream = constructStatusStream(key, url, normalized);
  container.appendChild(statusStream);
  // 创建一个 div 来包裹 span 标签
  const divWrapper = create("div");
  divWrapper.classList.add("span-wrapper"); // 添加一个类以便在 CSS 中定位这个 div
  divWrapper.id = "status-prompt"; // 设置 div 的 ID
  // 创建并添加两个 span 标签到 divWrapper 中
  const spanLeft = create("span", "span-title");
  spanLeft.textContent = "响应时间(ms)";
  spanLeft.classList.add("align-left");
  const spanRight = create("span", "span-text");
  spanRight.classList.add("align-right");
  divWrapper.appendChild(spanLeft);
  divWrapper.appendChild(spanRight);
  // 将包含两个 span 的 div 添加到 container 中
  container.appendChild(divWrapper);
  const canvas = create("canvas", "chart");
  canvas.id = "chart_clone_" + key++;
  container.appendChild(canvas);
  updateChart(canvas, statusLines);
}

// 所有服务当天整体状态评估
export async function getLastDayStatus() {
  const configLines = await fetchUrlsConfig();
  const statusTexts = []; // 存储 statusText 的数组
  for (let ii = 0; ii < configLines.length; ii++) {
    const configLine = configLines[ii];
    const [key, url] = configLine.split("=");
    const response = await reslogs(key);
    let statusLines = "";
    if (response.ok) {
      statusLines = await response.text();
    }
    const normalized = normalizeData(statusLines);
    // 获取最后一天的状态
    const lastDayStatus = normalized[0];
    const color = getColor(lastDayStatus);
    const statusText = getStatusText(color);
    statusTexts.push(statusText); // 将 statusText 存入数组
  }
  const upCount = statusTexts.filter(text => text === 'UP').length;
  const downCount = statusTexts.filter(text => text === 'Down').length;
  const nodateCount = statusTexts.filter(text => text === 'No Data').length;

  const img = document.querySelector('#statusImg');

  const totalCount = statusTexts.length;
  const downThreshold = totalCount * 0.2;  // 有效服务 Down 20% 即整体报告为Down
  const nodateThreshold = totalCount * 0.5;  // 有效服务 No Data 50% 即整体报告为No Data

  if (upCount === totalCount) {
    img.src = './public/check/up.svg';
    img.alt = 'UP';
  } else if (nodateCount === totalCount) {
    img.src = './public/check/nodata.svg';
    img.alt = 'No Data';
  } else if (downCount >= downThreshold || nodateCount >= nodateThreshold) {
    img.src = './public/check/down.svg';
    img.alt = 'Down';
  } else {
    img.src = './public/check/degraded.svg';
    img.alt = 'Degraded';
  }
}