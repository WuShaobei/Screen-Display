import axios from 'axios'

function postApiWhoAmI(Id, callback){
    axios.post(
        "http://127.0.0.1:1432/api/v1/whoAmI?",{                        
            Id : Id
        }
    ).then(function(res){
        if (res.data.Code == 0 ){
            let data = {}
            data.Id  = Id,
            data.Username = res.data.Data.Username
            data.RealName = res.data.Data.RealName
            data.Identity = res.data.Data.Identity
            callback({Data:data})
        } else {
            callback({Data:{}})
        }
    }).catch(function(err){
        console.log(err)
    })
}

function postApiLoginByPassword(Username, Password, callback){
    axios.post(
        "http://127.0.0.1:1432/api/v1/loginByPassword?",{                        
            Username : Username, 
            Password : Password,
        }
    ).then(function(res){
        if (res.data.Code == 0 ){
            callback({
                Id : res.data.Data.Id,
                SessionKey: res.data.Data.Session
            })
        } else {
            callback()
        }
    }).catch(function(err){
        console.log(err)
        callback()
    })
}


function postApiRegister(Username, Password, RealName, Identity, callback){
    axios.post(
        "http://127.0.0.1:1432/api/v1/create",{
            Username: Username,
            Password: Password,
            Identity: Identity,
            RealName :RealName
        }
    ).then(function(res){
        if (res.data.Code == 0 ){
            callback("注册成功")
        } else {
            callback("用户已存在")
        }
    }).catch(function(err){
        console.log(err)
        callback("err")
    })
}

function postApiLoginBySession(SessionKey, callback){
    axios.post(
        "http://127.0.0.1:1432/api/v1/loginBySession?",{                        
            SessionKey: SessionKey
        }
    ).then(function(res){
        if (res.data.Code == 0 ){
            callback({Id : res.data.Data.Id})
        } else {
            callback()
        }
    }).catch(function(err){
        console.log(err)
        callback()
    })
}

export{
    postApiWhoAmI,
    postApiRegister,
    postApiLoginByPassword,
    postApiLoginBySession
}