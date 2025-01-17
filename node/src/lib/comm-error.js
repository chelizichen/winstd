class T_ERROR extends Error {
  static {
    T_ERROR.CODE_THREE_PART = -300;
    T_ERROR.CODE_DB_ERROR = -400;
  }

  constructor(uid, message, code) {
    super(message);
    this.code = code;
    this.uid = uid;
  }

  getCode() {
    return this.code;
  }

  printStackError() {
    console.log('T_ERROR | uid %s | code %s | stack \n%s', this.uid, this.code, this.stack);
  }
}

const ThreePartError = (uid, message) => new T_ERROR(uid, message || '三方接口异常', T_ERROR.CODE_THREE_PART);
const DBError = (uid, message) => new T_ERROR(uid, message || '数据库异常', T_ERROR.CODE_DB_ERROR);

module.exports = {
  ThreePartError,
  T_ERROR,
  DBError,
};
