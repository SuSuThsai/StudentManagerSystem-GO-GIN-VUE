<template>
  <div id="container"></div>
</template>
<script>

export default {
  data() {
    return {
    }
  },
  methods: {// 实例化点标记
    MapInit() {
      let a=this.axios
      var map = new AMap.Map('container', {
        resizeEnable: true
      });
      AMap.plugin('AMap.Geolocation', function () {
        var geolocation = new AMap.Geolocation({
          enableHighAccuracy: true,//是否使用高精度定位，默认:true
          timeout: 10000,          //超过10秒后停止定位，默认：5s
          buttonPosition: 'RB',    //定位按钮的停靠位置
          buttonOffset: new AMap.Pixel(10, 20),//定位按钮与设置的停靠位置的偏移量，默认：Pixel(10, 20)
          zoomToAccuracy: true,   //定位成功后是否自动调整地图视野到定位点
        });

        map.addControl(geolocation);
        geolocation.getCurrentPosition(function (status, result) {
          if (status == 'complete') {


            fa(result.position.lng,result.position.lat)


            onComplete(result)
          } else {
            onError(result)
          }
        });

      });

      function fa(lng,lat){
        let e = new URLSearchParams()
        e.append("lng", lng)
        e.append("lat", lat)
        console.log("发fff：",lng+"_"+lat)
        const api = 'http://localhost:3000/api/v1/student/rtslupdate'
        a.post(api, e).then(res => {
          console.log("更新成功！！！！")
          this.$message({
            message: "更新定位成功:" + res.data.message,
            type: 'success'
          })
        }).catch(err => {
          this.$message({
            message: "更新定位失败:" + err.response.data.message,
            type: 'error'
          })
          console.log('err:', err.response.data)
        })
      }

      //解析定位结果
      function onComplete(data) {
        document.getElementById('status').innerHTML = '定位成功'
        var str = [];
        str.push('定位结果：' + data.position);
        str.push('定位类别：' + data.location_type);
        if (data.accuracy) {
          str.push('精度：' + data.accuracy + ' 米');
        }//如为IP精确定位结果则没有精度信息
        str.push('是否经过偏移：' + (data.isConverted ? '是' : '否'));
        document.getElementById('result').innerHTML = str.join('<br>');
      }

      //解析定位错误信息
      function onError(data) {
        document.getElementById('status').innerHTML = '定位失败'
        document.getElementById('result').innerHTML = '失败原因排查信息:' + data.message;
      }

    },
  },
  mounted() {
    this.MapInit()
    window.setInterval(this.MapInit,20*1000)
  }

}
</script>

<style>
html,
body,
#container {
  width: 100%;
  height: 100%;
}
</style>
