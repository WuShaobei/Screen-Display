import axios from 'axios'

function getWhoAmIData(Id, callback){
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

export{
    getWhoAmIData
}