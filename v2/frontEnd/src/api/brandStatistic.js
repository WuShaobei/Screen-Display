///
 // @Date: 2022-10-10 17:36:47
 // @LastEditTime: 2022-10-14 17:50:01
 // @FilePath: /frontEnd/src/api/brandStatistic.js
 // @Description: 火锅品牌类接口
 ///
 
 import axios from 'axios'
 
 /******* 
  * @description:  get "http://127.0.0.1:1432/api/data/getAllDataFromBrandStatistic" 然后转化成金字塔图并返回
  * @param {*} callback
  * @return {*}
  */
 function getDataFromBrandStatisticApi(callback) {
    axios.get(
        "http://127.0.0.1:1432/api/data/getAllDataFromBrandStatistic"
    ).then( function(res) { 
        if (res.data.code == 0) {
            let data = {
                title: {
                    text: '中国火锅品牌梯队分布',
                    color: '#FFFFFF',
                },
                series: [
                    {
                    type: 'funnel',
                    min: 0,
                    max: 4,
                    minSize: '0%',
                    maxSize: '150%',
                    sort: 'ascending',
                    gap: 10,
                    label: {
                        show: true,
                        position: 'inside',
                        color: '#555555',
                        fontSize:11.5,
                    },
                    labelLine: {
                        length: 10,
                        lineStyle: {
                            width: 1,
                            type: 'solid'
                        }
                    },
                    itemStyle: {
                        borderColor: '#FF9900',
                        borderWidth: 2
                    },
                    emphasis: {
                        label: {
                            fontSize: 20
                        }
                    },
                    data: [
                        { 
                            value: 1, 
                            name: res.data.data[0],
                            itemStyle: {
                                color: "#FFDD77",
                                // borderColor: "none",
                            },
                        },
                        { 
                            value: 2, 
                            name: res.data.data[1],
                            itemStyle: {
                                color: "#FFBB77",
                                // borderColor: "none",
                            },
                        },
                        { 
                            value: 3,
                            name: res.data.data[2],
                            itemStyle: {
                                color: "#FF9977",
                                // borderColor: "none",
                            },
                        },
                        {
                            value: 4,
                            name: res.data.data[3],
                            itemStyle: {
                                color: "#FF7777",
                                // borderColor: "none",
                            },
                         }
                    ]
                    }
                ]
            }
            callback({code : res.data.code, data:data})
        } else {
            callback({code : res.data.code})
        }
    }).catch(function(err){
        console.log(err)
    })
 }

 export{
    getDataFromBrandStatisticApi
 }