import { createStore } from 'vuex'
import createPersistedState from "vuex-persistedstate";
export default createStore({
  plugins: [createPersistedState({
    storage: sessionStorage,
  })],
  strict: true, // process.env.NODE_ENV !== 'production',
  state: {
    isSidebarMinimized: false,
    userName: 'Vasili S'
  },
  mutations: {
    updateSidebarCollapsedState(state, isSidebarMinimized) {
      state.isSidebarMinimized = isSidebarMinimized
    },
    changeUserName(state, newUserName) {
      state.userName = newUserName
    }
  },
})
