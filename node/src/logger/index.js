// "winston": "^3.17.0",
// "winston-daily-rotate-file": "^5.0.0"
/**
 * @file logger.js
 * @description 日志模块
 */
let DailyRotateFile = require("winston-daily-rotate-file");
let { createLogger, format } = require("winston");
let { cwd } = require("process");
let { resolve } = require("path");

function Logger(logDir) {
  this.logDir = logDir;
}

Logger.prototype.create = function (logName) {
  let logDir = this.logDir;
  let filePath = resolve(logDir, `${logName}.log`);
  return createLogger({
    level: "info", // 设置日志级别
    format: format.combine(
      format.timestamp({
        format: "YYYY-MM-DD HH:mm:ss",
      }),
      format.printf(
        (info) => `${info.timestamp} ${info.level}: ${info.message}`
      )
    ),
    transports: [
      new DailyRotateFile({
        filename: filePath.replace(/\.log$/, "-%DATE%.log"),
        datePattern: "YYYY-MM-DD",
        zippedArchive: true,
        maxSize: "20m",
        maxFiles: "14d",
      }),
    ],
  });
};

let logger = new Logger(resolve(cwd(), "log"));
Logger.app = logger.create("app");

module.exports = logger;
