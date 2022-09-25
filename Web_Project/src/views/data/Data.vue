<template>
    <div>
        <div class="background">
            <img :src="imgSrcBK" height="100%" width="100%"/>
        </div>

        <div class="dataView">
        
            <div class="content">
            
                <div class="leftView">
                    <e-charts :option="opApi4"/>
                </div>
                <div class="rightView">
                    <div id="scroll_in2" class="htp-list_scroll_in" style="height: 100%;overflow-y:hidden;font-size:20px;">

                        <table>
                            <tr>
                                <td style="height: 20%;"> 品牌 </td>
                                <td> 时间 </td>
                                <td> 轮次 </td>
                                <td> 投资商 </td>
                            </tr>
                            <tr v-for="data in rightValue">
                                <td style="height: 10%;">
                                    {{data.Brand}}
                                </td>
                                <td>
                                    {{data.Time}}
                                </td>
                                <td>
                                    {{data.FundingRound}}
                                </td>
                                <td>
                                    {{data.Investor}}
                                </td>
                            </tr>
                        </table>
                    </div>
                </div>
                    
            </div>


            <div class="indexes">
                <e-charts :option="opApi4"/>
                
            </div>
        
        </div>
    </div>
    

</template>

<script>
    
    import axios from 'axios'
    
    export default{
        name: "login",

        data(){
            return{
                // 背景
                imgSrcBK:require('../../assets/login/login/background.jpeg'),
                Username: "",
                Password: "",
                RealName:"",
                Identity:"",
                BeginYear:2018,
                EndYear:2021,
                rightValue:{},
                showLeftValue: false,
                opApi4 : {},
            }
        },
        methods:{
            getIdentity(){
                if ( this.Identity == "0" ) {
                    return "管理员"
                }
                else if ( this.Identity == "1" ) {
                    return "投资方"
                }
                else if ( this.Identity == "2" ) {
                    return "从业者"
                }
                else {
                    return "游客"
                }
            },
            Api4(){
                let that = this
                axios.post(
                    "http://127.0.0.1:1432/api/v1/api4"
                ).then( function(res) {
                    if (res.data.Code == 0) {
                        that.opApi4 = {
                            title: {
                                text: '中国火锅品牌梯队分布'
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
                                    position: 'inside'
                                },
                                labelLine: {
                                    length: 10,
                                    lineStyle: {
                                    width: 1,
                                    type: 'solid'
                                    }
                                },
                                itemStyle: {
                                    borderColor: '#fff',
                                    borderWidth: 1
                                },
                                emphasis: {
                                    label: {
                                    fontSize: 20
                                    }
                                },
                                data: [
                                    { value: 20, name: res.data.Data[0] },
                                    { value: 40, name: res.data.Data[1] },
                                    { value: 60, name: res.data.Data[2] },
                                    { value: 80, name: res.data.Data[3] }
                                ]
                                }
                            ]
                            };
                    }
                }).catch(function(err){
                    console.log(err)
                })
            },
            Api5(){
                let that = this
                axios.post(
                    "http://127.0.0.1:1432/api/v1/api5"
                ).then(function(res){
                    if ( res.data.Code == 0 ) {
                        that.rightValue = res.data.Data
                    }
                }).catch(function(err){
                    console.log(err)
                })
            },    // 添加自动滚动
            /* 
            Id   需要滚动的区域 id名称
            */
            autoSroll(Id) {
            // flag 为true时停止滚动
                var flag = false;
                // 定时器
                var timer;
                function roll() {
                    var h = -1;
                    timer = setInterval(function() {
                    flag = true;
                    //获取当前滚动条高度
                    var current = document.getElementById(Id).scrollTop;
                    if (current == h) {
                        //滚动到底端,返回顶端
                        h = 0;
                        document.getElementById(Id).scrollTop = h + 1;
                    } else {
                        //以25ms/3.5px的速度滚动
                        h = current;
                        document.getElementById(Id).scrollTop = h + 1;
                    }
                    }, 50);
                }
                // setTimeout(function() {
                    //滚动区域内单击鼠标 滚动暂停 再次点击鼠标 继续滚动
                    document.getElementById(Id).onclick = () => {
                    console.log("点击了", timer, flag);
                    if (flag) {
                        flag = false;
                        clearInterval(timer);
                    } else {
                        roll();
                    }
                    };
                    roll();
                // }, 2000);
                },    
        },
        mounted(){
            this.Username = this.$route.query.Username,
            this.RealName = this.$route.query.RealName,
            this.Identity = this.$route.query.Identity,
            this.Api4(),
            this.Api5(),
            this.autoSroll("scroll_in2");
        },
    }

</script>


<style>
    .background{
        height:100%; 
        width:100%;
        top: 0%;
        left: 0%;
        z-index: -1;
        position: absolute;
    }

    .dataView{
        height: 100%;
        width: 100%;
        top: 0%;
        left: 0%;
        z-index: 1;
        position: absolute;
        display: flex;
        
        opacity: 0.7; 
        /* background-color: antiquewhite; */
    }    
        


        .content{
            height:100%; 
            width:80%;
            display: flex;
            position: absolute;     
            /* background-color: black;    */
        }

            
            .leftView{
                height:100%; 
                width:50%;
                left: 0%;
                display: flex;    
                position: absolute;
                background-color: azure;
            }
        
            .rightView{
                height:100%; 
                width:50%;
                left: 50%;
                display: flex;    
                position: absolute;     
                background-color: green;   
            }

        .indexes{
            height:100%;
            width:20%;
            top: 0%;
            display: flex;
            position: absolute;
            left: 80%;
            background-color: rgb(124, 94, 16);
        }
.leftBox{
    top: 0%;
    height: 100%;
    width: 30%;
    position: absolute;

}
.rightBox{
    top: 0%;
    height: 100%;
    width: 70%;
    left: 30%;
    position: absolute;
}
.data0{
    top: 0%;
    height: 10%;
    width: 100%;
    position: absolute;
    display: flex;
    /* background-color: antiquewhite; */
}
.data1{
    height: 10%;
    top: 10%;
    width: 100%;
    position: absolute;
    display: flex;
    /* background-color: aqua; */
}
.data2{
    height: 10%;
    top: 20%;
    width: 100%;
    position: absolute;
    display: flex;
    /* background-color: antiquewhite; */
}
.data3{
    height: 10%;
    top: 30%;
    width: 100%;
    position: absolute;
    display: flex;
    /* background-color: aqua; */
}
.data4{
    height: 60%;
    top: 40%;
    width: 100%;
    position: absolute;
    display: flex;
    /* background-color: antiquewhite; */
}

/* text */
.text{
    height:100%;
    width: 100%;
    position: absolute;
    line-height: 100px; 
    text-align: center;  
    background-color: aquamarine;
}
.textIndex{
    height:100%;
    width: 100%;
    position: absolute;
    line-height: 100px; 
    text-align: left;  
    background-color: aquamarine;
}
.textTable{
    height:90%;
    width: 90%;
    top: 5%;
    left: 5%;
    position: absolute;
    line-height: 100px; 
    text-align: center;  
    background-color: aquamarine;
}


.table1{
    top: 0%;
    height: 20%;
    width: 100%;
    position: absolute;
    display: flex;
    background-color: antiquewhite;
}
.table2{
    top: 20%;
    height: 20%;
    width: 100%;
    position: absolute;
    display: flex;
    background-color: aqua;
}
.table3{
    top: 40%;
    height: 20%;
    width: 100%;
    position: absolute;
    display: flex;
    background-color: antiquewhite;
}
.table4{
    top: 60%;
    height: 20%;
    width: 100%;
    position: absolute;
    display: flex;
    background-color: aqua;
}
.table5{
    top: 80%;
    height: 20%;
    width: 100%;
    position: absolute;
    display: flex;
    background-color: antiquewhite;
}
.tableLeft{
    height: 100%;
    width: 30%;
    left: 0%;
    position: absolute;
    display: flex;
}
.tableCenter{
    height: 100%;
    width: 40%;
    left: 30%;
    position: absolute;
    display: flex;
}
.tableRight{
    height: 100%;
    width: 30%;
    left: 70%;
    position: absolute;
    display: flex;
}

.top{
    height: 50%;
    width: 100%;
    top:0%;
    position: absolute;
    display: flex;
    background-color: aliceblue;
}
.down{
    height: 50%;
    width: 100%;
    top:50%;
    position: absolute;
    display: flex;
    background-color: aliceblue;
}
.htp-list_scroll_in{
    height: 100%;
    width: 100%;
    top:0%;
    position: absolute;
    display: flex;
    background-color: aliceblue;
}
 #input{
    text-align: left;
    height: 80%;
    width: 80%;
    top: 10%; 
    left: 0%;
    position: absolute;
    
 }
 
</style>

