import {defineStore} from 'pinia'


const organizationId: string = ''

const organizationName: string = ''
const organizationPlan: string = ''

export const useOrganizationStore = defineStore('organization', {
  state: () => ({
    organizationId: organizationId,
    organizationName: organizationName,
    organizationPlan: organizationPlan  
  }),
  actions: {
    setOrganizationId(id: string) {
      this.organizationId = id
    },
    setOrganizationName(name: string) {
      this.organizationName = name
    },
    setOrganizationPlan(plan: string) {
      this.organizationPlan = plan
    },
    clearOrganizationData() {
      this.organizationId = ''
      this.organizationName = ''
      this.organizationPlan = ''
    }
  },
  getters: {
    getOrganizationId: (state) => state.organizationId,
    getOrganizationName: (state) => state.organizationName,
    getOrganizationPlan: (state) => state.organizationPlan,
  }
})
