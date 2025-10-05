<template>
  <section class="space-y-3 relative">
    <h3 class="text-lg font-semibold">Pipeline Association</h3>
    <div class="flex space-x-2 items-center">
      <input
        v-model="newPipeline"
        type="text"
        placeholder="Type pipeline name"
        class="flex-1 border rounded-md px-3 py-2 text-sm"
        @input="onInput"
        @focus="onInput"
        @blur="hideSuggestions"
        autocomplete="off"
      />
      <button
        class="text-white px-4 py-2 rounded-md"
        style="background-color: #548FC7;"
        @click.prevent="associatePipeline"
      >
        Associate Pipeline
      </button>
    </div>
    <div
      v-if="showSuggestions && filteredOrgPipelines.length"
      class="bg-white border rounded-md mt-2 p-2 absolute z-10 w-1/2 shadow-lg"
      style="max-height: 200px; overflow-y: auto; top: 100%;"
    >
      <ul>
        <li
          v-for="pipeline in filteredOrgPipelines"
          :key="pipeline.id"
          class="py-1 px-2 hover:bg-gray-100 cursor-pointer"
          @mousedown="selectSuggestion(pipeline)"
        >
          {{ pipeline.name }}
        </li>
      </ul>
    </div>
    <div v-if="projectPipelines.length" class="mt-4">
      <table class="min-w-full border rounded-md">
        <thead>
          <tr>
            <th class="px-4 py-2 border">Pipeline Name</th>
            <th class="px-4 py-2 border">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="pipeline in projectPipelines" :key="pipeline.id">
            <td class="px-4 py-2 border">{{ pipeline.name }}</td>
            <td class="px-4 py-2 border">
              <button
                class="bg-red-500 hover:bg-red-600 text-white px-2 py-1 rounded"
                @click="removePipeline(pipeline.id)"
              >
                Remove
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </section>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'
import { API_URLS } from '@/controller/services/service.ts'

export default defineComponent({
  name: 'ProjectSPipline',
  props: {
    projectId: { type: String, required: true },
    organizationId: { type: String, required: true }
  },
  setup(props) {
    const newPipeline = ref('')
    const orgPipelines = ref<any[]>([])
    const projectPipelines = ref<any[]>([])
    const filteredOrgPipelines = ref<any[]>([])
    const showSuggestions = ref(false)
    const selectedPipeline = ref<any | null>(null)

    const fetchOrganizationPipelines = async () => {
      // ...same as before, but set orgPipelines.value
      const token = localStorage.getItem('token')
      if (!token) return
      const url = API_URLS.BASE_URL + API_URLS.GET_PIPELINES_BY_ORG + encodeURIComponent(props.organizationId)
      const response = await fetch(url, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
      })
      if (response.ok) {
        const data = await response.json()
        console.log("Organization pipelines data:", data)
        orgPipelines.value = Array.isArray(data) ? data : []
      }
    }

    const fetchProjectPipelines = async () => {
      const token = localStorage.getItem('token')
      if (!token) return
      const url = API_URLS.BASE_URL + API_URLS.GET_PROJECT_PIPELINES + encodeURIComponent(props.projectId)
      const response = await fetch(url, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
      })
      if (response.ok) {
        const data = await response.json()
        console.log("Project pipelines data:", data)
        projectPipelines.value = Array.isArray(data) ? data : []
      } else {
        projectPipelines.value = []
      }
    }

    const setPipelineToProject=async (pipelineId: string) => {
      const token = localStorage.getItem('token')
      if (!token) return
      const url = API_URLS.BASE_URL + API_URLS.GET_PROJECT_PIPELINES + encodeURIComponent(props.projectId)
      const response = await fetch(url, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify({ pipeline_config_id: pipelineId }),
      })
      if (response.ok) {
        await fetchProjectPipelines()
      }
    }

    const onInput = () => {
      if (newPipeline.value.trim()) {
        filteredOrgPipelines.value = orgPipelines.value
          .filter(p =>
            p.name.toLowerCase().includes(newPipeline.value.toLowerCase()) &&
            !projectPipelines.value.some(pp => pp.id === p.id)
          )
        showSuggestions.value = filteredOrgPipelines.value.length > 0
      } else {
        showSuggestions.value = false
        filteredOrgPipelines.value = []
      }
      selectedPipeline.value = null
    }

    const hideSuggestions = () => {
      setTimeout(() => {
        showSuggestions.value = false
      }, 200)
    }

    const selectSuggestion = (pipeline: any) => {
      newPipeline.value = pipeline.name
      selectedPipeline.value = pipeline
      showSuggestions.value = false
    }

    const associatePipeline = async () => {
      if (selectedPipeline.value) {
        await setPipelineToProject(selectedPipeline.value.id)
        newPipeline.value = ''
        selectedPipeline.value = null
        // Optionally, call fetchProjectPipelines() again
      } else if (newPipeline.value.trim()) {
        const found = orgPipelines.value.find(p =>
          p.name.toLowerCase() === newPipeline.value.toLowerCase()
        )
        if (found && !projectPipelines.value.some(p => p.id === found.id)) {
          await setPipelineToProject(found.id)
          newPipeline.value = ''
        } else if (!found) {
          alert('Pipeline not found in organization')
        }
      }
    }

    const removePipeline = (pipelineId: string) => {
      projectPipelines.value = projectPipelines.value.filter(p => p.id !== pipelineId)
      // Call your API to remove pipeline from project here
    }

    onMounted(async () => {
      await fetchOrganizationPipelines()
      await fetchProjectPipelines()
    })

    return {
      newPipeline,
      orgPipelines,
      projectPipelines,
      filteredOrgPipelines,
      showSuggestions,
      onInput,
      hideSuggestions,
      selectSuggestion,
      associatePipeline,
      removePipeline
    }
  }
})
</script>
