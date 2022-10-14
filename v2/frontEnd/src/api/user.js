///
 // @Date: 2022-10-10 17:36:47
 // @LastEditTime: 2022-10-14 17:56:15
 // @FilePath: /frontEnd/src/api/user.js
 // @Description: 用户类 axios 请求
 ///
import axios from 'axios'

/******* 
 * @description: 获取用户信息
 * @param {*} id
 * @param {*} callback
 * @return {*}
 */
function postWhoAmIApi(id, callback) {
    axios.post(
        "http://127.0.0.1:1432/api/user/whoAmI",{
            id : id
        }
    ) .then(function(res) {
        if ( res.data.code == 0 ) {
            callback({code :  0, data : res.data.data })
        } else {
            callback({code : res.data.code})
        }
    }) .catch(function(err){
        console.log(err)
    })
}


/******* 
 * @description: 通过密码登录
 * @param {*} username
 * @param {*} password
 * @param {*} setCookie
 * @param {*} callback
 * @return {*}
 */
function postLoginByPasswordApi(username, password, setCookie, callback) {
    axios.post(
        "http://127.0.0.1:1432/api/user/loginByPassword",{
            username : username,
            password : password,
            setCookie : setCookie != ""
        }
    ) .then(function(res) {
        if ( res.data.code == 0 ) {
            callback({code : 0, data : res.data.data })
        } else {
            callback({code : res.data.code })
        }
    }) .catch(function(err){
        console.log(err)
    })
}

/******* 
 * @description: 通过 session 值登录
 * @return {*}
 */
function postLoginBySessionKeyApi(sessionKey, callback) {
    axios.post(
        "http://127.0.0.1:1432/api/user/loginBySessionKey",{
            sessionKey : sessionKey
        }
    ) .then(function(res) {
        if ( res.data.code == 0 ) {
            callback({code : 0, data : res.data.data })
        } else {
            callback({code : res.data.code })
        }
    }) .catch(function(err){
        console.log(err)
    })
}


/******* 
 * @description: 登出
 * @param {*} sessionKey
 * @param {*} callback
 * @return {*}
 */
function postLogoutApi(sessionKey, callback) {
    axios.post(
        "http://127.0.0.1:1432/api/user/logout",{
            sessionKey : sessionKey
        }
    ).then(function(res){
        callback({code : res.data.code})
    }).catch(function(err){
        console.log(err)
    })
}

/******* 
 * @description: 注册
 * @param {*} username
 * @param {*} password
 * @param {*} realName
 * @param {*} identity
 * @param {*} callback
 * @return {*}
 */
function postRegisterApi(username, password, realName, identity, callback) {
    axios.post(
        "http://127.0.0.1:1432/api/user/register",{
            username : username,
            password : password,
            realName : realName,
            identity : identity,
        }
    ).then(function(res){
        if ( res.data.code == 0 ) {            
            callback({code : 0, data : res.data.data.id})
        } else {
            callback({code : res.data.code})
        }
    }).catch(function(err){
        console.log(err)
    }) 
}


export{
    postWhoAmIApi,
    postLoginByPasswordApi,
    postLoginBySessionKeyApi,
    postRegisterApi,
    postLogoutApi
}