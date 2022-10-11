<!--
 * @Author: WuShaobei
 * @Date: 2022-10-09 16:49:59
 * @LastEditTime: 2022-10-10 17:19:54
 * @FilePath: /Screen-Display/v2/frontEnd/src/views/display.vue
 * @Description: 数据显示页
-->
<template>
    <div class = "viewDisplay">
        <div class="background">
            <img :src="imgSrcBK" height="100%" width="100%"/>
        </div>

        <div class="content">
            <!-- <div class="brandStatisticView"><e-charts :option="brandStatisticOpData"/></div> 
            <div class="statisticView"><e-charts :option="statisticOpData"/></div>  -->
            <div class="fundingStatisticView"></div> 
            <div class="whoAmIView" style="display: flex;flex-direction: column;">
                <div class="tableTitle"><h1>欢迎{{whoAmITableData.identidy}}{{whoAmITableData.realName}}</h1></div>
                <div class="tableTr"><div  class="tableTdLeft">用户名</div><div  class="tableTdRight">{{whoAmITableData.username}}</div></div></div> 
            <div class="paymentView"></div> 
            <div class="onlineOrderView"></div> 
        </div>

        <div class="title" style="display: flex;flex-direction: column;">
            <div v-for="c in '中国餐饮市场数据'" style="width: 100%;">
                <h1>{{c}}</h1>
            </div>
            <div style="width:100%;">
                <button @click="logout()">注销<br></button>
            </div>
            <div style="width:100%;">
                <button @click="logout()">用户<br>管理</button>
            </div>
        </div>
    </div>
</template>

<script>
    export default{
        data () {
            return {
                // 背景
                imgSrcBK:require('../assets/background.jpeg'),
                whoAmITableData:{
                    username : "hello",
                    identidy : "管理员",
                    realName : "admin",
                }
            }
        },
        methods : {
            /**
             * @description: 切换显示的界面  
             * @param {*} todo : 将显示的界面  login / register
             * @return {*}
             */            
            changeState( todo ) {
                this.nowIs = todo
                this.username = ""
                this.password = ""
                this.realName = ""
            }
        }
    }
</script>

<style>
/* ========================================================= */
/* =======================底层界面=========================== */
/* ========================================================= */
.viewDisplay{
    height: 100%;
    width: 100%;
    top: 0%;
    left: 0%;

    /* 缩放正常显示 */
    min-width: 1000px;
    min-height: 800px;
    overflow: hidden;

    position: absolute;
}

    /* 背景设置 */
    .background{
        height: 100%;
        width: 100%;
        left: 0%;
        right: 0%;
        z-index: -1;
        position: absolute;
    }

    .content{
        opacity: 0.7;
        height: 100%;
        width: 95%;
        left: 0%;
        right: 0%;
        z-index: 1;
        position: absolute;
    }
        
    .title{
        height: 100%;
        width: 5%;
        left: 95%;
        right: 0%;
        z-index: 1;
        position: absolute;
        opacity: 0.7;
        background-color: rgb(124, 94, 16);
    }

/* ========================================================= */
/* ========================数据页============================ */
/* ========================================================= */

.brandStatisticView{
    height: 100%; top: 0%;
    width: 30%; left: 0%;
    position: absolute;
    background-color: rgb(255, 255, 255);
}
.statisticView{
    height: 50%; top: 0%;
    width: 40%; left: 30%;
    position: absolute;
    background-color: rgb(250, 250, 250);
}
.fundingStatisticView{
    height: 50%; top: 50%;
    width: 40%; left: 30%;
    position: absolute;
    background-color: rgb(230, 230, 230);
}
.whoAmIView{
    height: 20%; top: 0%;
    width: 30%; left: 70%;
    position: absolute;
    background-color: rgb(245, 245, 245);
}
.paymentView{
    height: 40%; top: 20%;
    width: 30%; left: 70%;
    position: absolute;
    background-color: rgb(240, 240, 240);
}
.onlineOrderView{
    height: 40%; top: 60%;
    width: 30%; left: 70%;
    position: absolute;
    background-color: rgb(235, 235, 235);
}
.tableTitle{
    width: 100%;
}
.tableTr{
    margin-top: 5%;
    margin-bottom: 5%;
    width: 100%;
    display: flex;
    flex-direction: row;
}
.tableTdLeft{
    left: 0;
    width: 40%;
    position: absolute;
}
.tableTdRight{
    width: 60%;
    left: 40%;
    position: absolute;
}
/* ========================================================= */
/* ==========================其他============================ */
/* ========================================================= */

#button{
    width: 40%;
    background-color: bisque;
}

</style>