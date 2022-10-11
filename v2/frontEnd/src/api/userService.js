// /$$
//  $ @Author: WuShaobei
//  $ @Date: 2022-10-09 18:32:17
//  $ @LastEditTime: 2022-10-09 18:32:23
//  $ @FilePath: /Screen-Display/v2/frontEnd/src/api/user.js
//  $ @Description: 用户相关 axios
//  $/

import axios from 'axios'

function postApiWhoAmI(id, callback){
    axios.post(
        "http://127.0.0.1:1432/api/user/whoAmI?",{                        
            id : id
        }
    ).then(function(res){
        if (res.data.Code == 0 ){
            let data = {}
            data.id  = id,
            data.username = res.data.Data.username
            data.realName = res.data.Data.realName
            data.identity = res.data.Data.identity
            callback({data:data})
        } else {
            callback({data:{}})
        }
    }).catch(function(err){
        console.log(err)
    })
}

function postApiLoginByPassword(Username, Password, callback){
    axios.post(
        "http://127.0.0.1:1432/api/user/loginByPassword?",{                        
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
        "http://127.0.0.1:1432/api/user/loginBySession?",{                        
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