import axios from 'axios'

function postApiSelectChineseCateringBrandStatistic(callback){
    axios.post(
        "http://127.0.0.1:1432/api/v1/select/ChineseCateringBrandStatistic"
    ).then( function(res) { 
        if (res.data.Code == 0) {
            let data = {
                title: {
                    text: '中国火锅品牌梯队分布',
                },
                series: [
                    {
                    type: 'funnel',
                    min: 0,
                    max: 100,
                    minSize: '0%',
                    maxSize: '100%',
                    sort: 'ascending',
                    gap: 2,
                    label: {
                        show: true,
                        position: 'inside',
                        color: '#555555',
                        fontSize:12.5,
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
                        borderWidth: 5
                    },
                    emphasis: {
                        label: {
                        fontSize: 20
                        }
                    },
                    data: [
                        { 
                            value: 20, 
                            name: res.data.Data[0],
                            itemStyle: {
                                color: "#FF9900",
                                borderColor: "none",
                            },
                        },
                        { 
                            value: 40, 
                            name: res.data.Data[1],
                            itemStyle: {
                                color: "#FF7700",
                                borderColor: "none",
                            },
                        },
                        { 
                            value: 60,
                            name: res.data.Data[2],
                            itemStyle: {
                                color: "#FF5500",
                                borderColor: "none",
                            },
                        },
                        {
                            value: 80,
                            name: res.data.Data[3],
                            itemStyle: {
                                color: "#FF3300",
                                borderColor: "none",
                            },
                         }
                    ]
                    }
                ]
            }
            callback({Data:data})
        }
    }).catch(function(err){
        console.log(err)
    })
}

export{
    postApiSelectChineseCateringBrandStatistic
}