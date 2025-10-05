import { defineStore } from 'pinia'
const nextCall: string = ''

export const useAccessStore = defineStore('access', {
  state: () => ({
    nextCall: nextCall,
  }),
  actions: {
    setNextCall(call: string) {
      this.nextCall = call
    },
    clearNextCall() {
      this.nextCall = ''
    },
  },
  getters: {
    getNextCall: (state) => state.nextCall,
  }
})
