const uid = require('uid')

class T_ERROR extends Error{
    static{
        T_ERROR.CODE_THREE_PART = -300
    }

    constructor(uid,message,code){
        super(message)
        this.code = code;
        this.uid = uid
    }

    getCode(){
        return this.code
    }

    printStackError(){
        console.log('T_ERROR | uid %s | code %s | stack \n%s',this.uid,this.code,this.stack);
    }
}

const ThreePartError = (uid,message) => new T_ERROR(uid, message || '三方接口异常',T_ERROR.CODE_THREE_PART)


function main(){
    try{
        const uuid = uid()
        throw ThreePartError(uuid,'三方接口异常')
    }catch(e){
        e.printStackError()
        if(e.getCode() === T_ERROR.CODE_THREE_PART){
            return { code:T_ERROR.CODE_THREE_PART }
        }
        else{
            // retry
        }
    }

}


main ()