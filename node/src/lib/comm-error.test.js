const uid = require('uuid');
const errors = require('./comm-error');

function mockServiceCode(retryNumIn) {
  const retryNum = retryNumIn + 1;
  try {
    const uuid = uid.v7();
    console.log('retryNum', retryNum);
    if (retryNum >= 3) {
      console.log('执行成功');
      return { code: 0, msg: 'ok' };
    }
    throw errors.DBError(uuid, '数据库插入异常');
    // throw errors.ThreePartError(uuid, "三方接口异常");
  } catch (e) {
    e.printStackError();
    if (e.getCode() === errors.T_ERROR.CODE_THREE_PART) {
      return { code: errors.T_ERROR.CODE_THREE_PART };
    } if (e.getCode() === errors.T_ERROR.CODE_DB_ERROR) {
      return mockServiceCode(retryNum);
    }
  }
  return { code: 0, msg: 'ok' };
}

const servicResp = mockServiceCode(0);
console.log('resp', servicResp);
