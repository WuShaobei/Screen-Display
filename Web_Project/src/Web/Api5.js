import axios from 'axios'

function getApi5Data(callback){
    axios.post(
        "http://127.0.0.1:1432/api/v1/api5"
    ).then(function(res){
        if ( res.data.Code == 0 ) {
            callback({Data : res.data.Data})
        }
    }).catch(function(err){
        console.log(err)
    })
}

export{
    getApi5Data
}