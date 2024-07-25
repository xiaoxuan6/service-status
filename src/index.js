import { genAllReports, getLastDayStatus } from './js/genReports.js';
import { lastupdated } from './js/lastupdated.js';
import { getclieninfo } from './js/getclieninfo.js';
import { scrollheader } from './js/scroll.js';

// 配置参数
export const maxDays = 60;
export const maxHour = 12;
export const urlspath = "./config.cfg"; // 配置文件路径,不带后/
export const logspath = "./logs";  // 日志文件路径,不带后/

// 主函数入口
async function main() {
  await lastupdated();
  await getclieninfo();
  await genAllReports();
  await getLastDayStatus();
  await scrollheader()
}

main();
