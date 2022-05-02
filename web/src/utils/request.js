import axios from 'axios' // 引入axios
import { ElMessage, ElMessageBox } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import { emitter } from '@/utils/bus.js'
import router from '@/router/index'

const service = axios.create({
  baseURL: "http://localhost:8080/",
  timeout: 99999
})
let acitveAxios = 0
let timer
// const showLoading = () => {
//   acitveAxios++
//   if (timer) {
//     clearTimeout(timer)
//   }
//   timer = setTimeout(() => {
//     if (acitveAxios > 0) {
//       emitter.emit('showLoading')
//     }
//   }, 400)
// }

const closeLoading = () => {
  acitveAxios--
  if (acitveAxios <= 0) {
    clearTimeout(timer)
    emitter.emit('closeLoading')
  }
}
// http request 拦截器
service.interceptors.request.use(
  config => {
    // if (!config.donNotShowLoading) {
    //   showLoading()
    // }
    // // const userStore = useUserStore()
    // config.headers = {
    //   'Content-Type': 'application/json',
    //   // 'x-token': userStore.token,
    //   // 'x-user-id': userStore.userInfo.ID,
    //   ...config.headers
    // }
    return config
  },
  error => {
    closeLoading()
    ElMessage({
      showClose: true,
      message: error,
      type: 'error'
    })
    return error
  }
)

// http response 拦截器
service.interceptors.response.use(
  response => {
    // console.log(JSON.stringify(response.data))
    // console.log(response.status)
    return response;
    // console.log(response)
    // const userStore = useUserStore()
    // closeLoading()
    // if (response.headers['new-token']) {
    //   userStore.setToken(response.headers['new-token'])
    // }
    // if (response.data.code === 0 || response.headers.success === 'true') {
    //   if (response.headers.msg) {
    //     response.data.msg = decodeURI(response.headers.msg)
    //   }
    //   return response.data
    // } else {
    //   ElMessage({
    //     showClose: true,
    //     message: response.data.msg || decodeURI(response.headers.msg),
    //     type: 'error'
    //   })
    //   if (response.data.data && response.data.data.reload) {
    //     userStore.token = ''
    //     localStorage.clear()
    //     router.push({ name: 'Login', replace: true })
    //   }
    //   console.log(response)
    //   return response.data.msg ? response.data : response
    // }
  },
  error => {
    closeLoading()
    switch (error.response.status) {
      case 500:
        ElMessageBox.confirm(`
        <p>${error}</p>
        <p>錯誤碼<span style="color:red"> 500 </span>：伺服器忙碌中</p>
        `, '接口報錯', {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: '清理暫存',
          cancelButtonText: '取消'
        })
          .then(() => {
            const userStore = useUserStore()
            userStore.token = ''
            localStorage.clear()
            router.push({ name: 'Login', replace: true })
          })
        break
      case 403:
          ElMessageBox.confirm(`
            <p>${error}</p>
            <p>錯誤碼<span style="color:red"> 403 </span>：帳號已存在，請再試一次。</p>
            `, '接口報錯', {
            dangerouslyUseHTMLString: true,
            distinguishCancelAndClose: true,
            confirmButtonText: '我知道了',
            cancelButtonText: '取消'
          })
          break
      case 404:
        ElMessageBox.confirm(`
          <p>${error}</p>
          <p>錯誤碼<span style="color:red"> 404 </span>：找不到帳號，您可能未註冊。</p>
          `, '接口報錯', {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: '我知道了',
          cancelButtonText: '取消'
        })
        break
    }

    return error
  }
)
export default service