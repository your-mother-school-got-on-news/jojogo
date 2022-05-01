import { login, register} from '@/api/user'
import { defineStore } from 'pinia'
import { ref} from 'vue'
import router from '@/router/index'
// import { useRouterStore } from './router'

export const useUserStore = defineStore('user', () => {
    // const userInfo = ref({
    //   uuid: '',
    //   nickName: '',
    //   headerImg: '',
    //   authority: {},
    //   sideMode: 'dark',
    //   activeColor: '#4D70FF',
    //   baseColor: '#fff'
    // })
    // const setUserInfo = (val) => {
    //     userInfo.value = val
    // }
      const setToken = (val) => {
        token.value = val
    }
    const token = ref(window.localStorage.getItem('token') || '')
    // login
    const LoginIn = async(loginInfo) => {
      const res = await login(loginInfo)
      // console.log(res.status)
      if (res.status === 200) {
          // console.log(res.status, userInfo, res.data.JSON)
          const resJson = JSON.parse(JSON.stringify(res.data))
          // console.log(resJson)
        // setUserInfo(res.data.ㄎ)
        // console.log(resJson["data"]["String"])
        setToken(resJson["data"]["String"])
        // const routerStore = useRouterStore()
        // await routerStore.SetAsyncRouter()
        // const asyncRouters = routerStore.asyncRouters
        // asyncRouters.forEach(asyncRouter => {
        //   router.addRoute(asyncRouter)
        // })
        router.push({ name: 'home' })
        return true
      }
    }

    const Register = async(registerInfo) => {
      const res = await register(registerInfo)
      console.log(res.status)
    }
  
    return {
      token,
      LoginIn,
      Register
    }
})