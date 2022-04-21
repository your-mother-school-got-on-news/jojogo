import { login} from '@/api/user'
import { defineStore } from 'pinia'
import { ref} from 'vue'
// import router from '@/router/index'
// import { useRouterStore } from './router'

export const useUserStore = defineStore('user', () => {
    const userInfo = ref({
      uuid: '',
      nickName: '',
      headerImg: '',
      authority: {},
      sideMode: 'dark',
      activeColor: '#4D70FF',
      baseColor: '#fff'
    })
    // const setUserInfo = (val) => {
    //     userInfo.value = val
    // }
    //   const setToken = (val) => {
    //     token.value = val
    // }
    // const token = ref(window.localStorage.getItem('token') || '')
    /* 登录*/
    const LoginIn = async(loginInfo) => {
      const res = await login(loginInfo)
      if (res.code === 0) {
          console.log(res.code, userInfo)
        // setUserInfo(res.data.user)
        // setToken(res.data.token)
        // const routerStore = useRouterStore()
        // await routerStore.SetAsyncRouter()
        // const asyncRouters = routerStore.asyncRouters
        // asyncRouters.forEach(asyncRouter => {
        //   router.addRoute(asyncRouter)
        // })
        // router.push({ name: userInfo.value.authority.defaultRouter })
        return true
      }
    }

  
    return {
      LoginIn
    }
})