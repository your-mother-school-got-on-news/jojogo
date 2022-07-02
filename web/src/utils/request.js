import axios from 'axios' // 引入axios

const service = axios.create({
    baseURL: "http://localhost:8081/",
    timeout: 99999
  })

  // http request 拦截器
service.interceptors.request.use(
    config => {
      return config
    },
    error => {
      return error
    }
  )

// http response 拦截器
service.interceptors.response.use(
  response => {

    return {data: response.data, status: response.status};
  })
export default service