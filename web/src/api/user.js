import service from '@/utils/request'
// @Summary Login
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/login [post]
export const login = (data) => {
  return service({
    url: '/login',
    method: 'post',
    data: data
  })
}


// @Summary Register
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/register [post]
export const register = (data) => {
  return service({
    url: '/register',
    method: 'post',
    data: data
  })
}