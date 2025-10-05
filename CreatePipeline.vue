<script lang="ts">
import { defineComponent, ref } from 'vue'
import AdminLayout from '@/components/layout/AdminLayout.vue'

export default defineComponent({
  name: 'PipelineSteps',
  components: { AdminLayout },
  setup() {
    const buildTool = ref('')
    const testCommands = ref('')
    const deploymentTarget = ref('')
    const deploymentRegion = ref('')
    const ciCdTool = ref('')
    const envVars = ref([
      { name: 'API_KEY', value: 'your_api_key_here' },
      { name: 'DATABASE_URL', value: 'your_database_url_here' },
      { name: 'ENVIRONMENT', value: 'production' }
    ])
    const addEnvVar = () => {
      envVars.value.push({ name: '', value: '' })
    }
    const removeEnvVar = (index: number) => {
      envVars.value.splice(index, 1)
    }

    // Jobs block
    const jobs = ref([
      { name: 'Install Dependencies', stage: 'Build', status: 'Pending' },
      { name: 'Run Tests', stage: 'Test', status: 'Pending' },
      { name: 'Deploy', stage: 'Deploy', status: 'Pending' }
    ])
    const addJob = () => {
      jobs.value.push({ name: '', stage: '', status: 'Pending' })
    }
    const removeJob = (index: number) => {
      jobs.value.splice(index, 1)
    }


    return { buildTool, ciCdTool,testCommands, envVars, deploymentTarget, deploymentRegion , addEnvVar, removeEnvVar , jobs, addJob, removeJob}
  }
})
</script>

<template>
  <AdminLayout>
    <div class="min-h-screen rounded-2xl border border-gray-200 bg-white px-5 py-7 dark:border-gray-800 dark:bg-white/[0.03] xl:px-10 xl:py-12">
      <!-- Header -->
      <h2 class="mb-2 text-2xl font-semibold text-gray-800 dark:text-white/90">
        Configure Pipeline Steps
      </h2>
      <p class="mb-8 text-sm text-gray-500 dark:text-gray-400">
        Customize the steps in your pipeline to ensure efficient error detection and resolution.
      </p>

      <!-- Build Tools -->
      <div class="mb-6">
        <label class="mb-2 block text-sm font-medium text-gray-700 dark:text-gray-300">
          Select Build Tool
        </label>
        <select v-model="buildTool" class="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm dark:border-gray-700 dark:bg-white/[0.05] dark:text-white">
          <option value="">Choose</option>
          <option value="maven">Maven</option>
          <option value="gradle">Gradle</option>
          <option value="npm">NPM</option>
          <option value="yarn">Yarn</option>
        </select>
      </div>

      <!-- Test Commands -->
      <div class="mb-6">
        <label class="mb-2 block text-sm font-medium text-gray-700 dark:text-gray-300">
          Test Commands
        </label>
        <textarea v-model="testCommands" rows="3" class="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm dark:border-gray-700 dark:bg-white/[0.05] dark:text-white"></textarea>
      </div>
      <!-- CI/CD Tools -->
      <!-- CI/CD Tool Selection -->
      <div class="mb-6">
        <label class="mb-2 block text-sm font-medium text-gray-700 dark:text-gray-300">
          Select CI/CD Tool
        </label>
        <select v-model="ciCdTool" class="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm dark:border-gray-700 dark:bg-white/[0.05] dark:text-white">
          <option value="">Choose</option>
          <option value="jenkins">Jenkins</option>
          <option value="github-actions">GitHub Actions</option>
          <option value="gitlab-ci">GitLab CI</option>
          <option value="circleci">CircleCI</option>
          <option value="travisci">Travis CI</option>
        </select>
      </div>
      <!-- Environment Variables -->
      <div class="mb-8">
        <h3 class="mb-3 text-base font-semibold text-gray-700 dark:text-gray-300">Environment Variables</h3>
        <div class="overflow-hidden rounded-xl border border-gray-200 dark:border-gray-700">
          <table class="w-full text-sm">
            <thead class="bg-gray-50 dark:bg-white/[0.05]">
            <tr>
              <th class="px-4 py-3 text-left">Variable Name</th>
              <th class="px-4 py-3 text-left">Value</th>
              <th class="px-4 py-3 text-left"></th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="(env, i) in envVars" :key="i" class="border-t border-gray-200 dark:border-gray-700">
              <td class="px-4 py-3">
                <input
                  v-model="env.name"
                  class="w-full rounded border border-gray-300 px-2 py-1 text-sm dark:border-gray-700 dark:bg-white/[0.05] dark:text-white"
                  placeholder="Variable Name"
                />
              </td>
              <td class="px-4 py-3">
                <input
                  v-model="env.value"
                  class="w-full rounded border border-gray-300 px-2 py-1 text-sm dark:border-gray-700 dark:bg-white/[0.05] dark:text-white"
                  placeholder="Value"
                />
              </td>
              <td class="px-4 py-3">
                <button
                  @click="removeEnvVar(i)"
                  class="text-red-500 hover:underline"
                  type="button"
                >Remove</button>
              </td>
            </tr>
            </tbody>
          </table>
        </div>
        <button
          @click="addEnvVar"
          class="mt-3 rounded border border-gray-300 bg-white px-3 py-1 text-sm font-medium hover:bg-gray-50 dark:border-gray-600 dark:bg-white/[0.05] dark:text-white"
          type="button"
        >
          Add Variable
        </button>
      </div>

      <!-- Deployment Targets -->
      <div class="mb-6">
        <label class="mb-2 block text-sm font-medium text-gray-700 dark:text-gray-300">
          Select Deployment Target
        </label>
        <select v-model="deploymentTarget" class="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm dark:border-gray-700 dark:bg-white/[0.05] dark:text-white">
          <option value="">Choose</option>
          <option value="aws">AWS</option>
          <option value="gcp">GCP</option>
          <option value="azure">Azure</option>
        </select>
      </div>

      <div class="mb-8">
        <label class="mb-2 block text-sm font-medium text-gray-700 dark:text-gray-300">
          Deployment Region
        </label>
        <input v-model="deploymentRegion" placeholder="Enter deployment region" class="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm dark:border-gray-700 dark:bg-white/[0.05] dark:text-white" />
      </div>
      <!-- Jobs Configuration -->
      <!-- Jobs Block -->
      <div class="mb-8">
        <h3 class="mb-3 text-base font-semibold text-gray-700 dark:text-gray-300">Pipeline Jobs</h3>
        <div class="overflow-hidden rounded-xl border border-gray-200 dark:border-gray-700">
          <table class="w-full text-sm">
            <thead class="bg-gray-50 dark:bg-white/[0.05]">
            <tr>
              <th class="px-4 py-3 text-left">Job Name</th>
              <th class="px-4 py-3 text-left">Stage</th>
              <th class="px-4 py-3 text-left">Status</th>
              <th class="px-4 py-3 text-left"></th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="(job, i) in jobs" :key="i" class="border-t border-gray-200 dark:border-gray-700">
              <td class="px-4 py-3">
                <input
                  v-model="job.name"
                  class="w-full rounded border border-gray-300 px-2 py-1 text-sm dark:border-gray-700 dark:bg-white/[0.05] dark:text-white"
                  placeholder="Job Name"
                />
              </td>
              <td class="px-4 py-3">
                <input
                  v-model="job.stage"
                  class="w-full rounded border border-gray-300 px-2 py-1 text-sm dark:border-gray-700 dark:bg-white/[0.05] dark:text-white"
                  placeholder="Stage"
                />
              </td>
              <td class="px-4 py-3">
                <select v-model="job.status" class="rounded border border-gray-300 px-2 py-1 text-sm dark:border-gray-700 dark:bg-white/[0.05] dark:text-white">
                  <option value="Pending">Pending</option>
                  <option value="Running">Running</option>
                  <option value="Successful">Successful</option>
                  <option value="Failed">Failed</option>
                </select>
              </td>
              <td class="px-4 py-3">
                <button
                  @click="removeJob(i)"
                  class="text-red-500 hover:underline"
                  type="button"
                >Remove</button>
              </td>
            </tr>
            </tbody>
          </table>
        </div>
        <button
          @click="addJob"
          class="mt-3 rounded border border-gray-300 bg-white px-3 py-1 text-sm font-medium hover:bg-gray-50 dark:border-gray-600 dark:bg-white/[0.05] dark:text-white"
          type="button"
        >
          Add Job
        </button>
      </div>
      <!-- Additional Steps -->
      <div class="mb-8 rounded-xl border border-dashed border-gray-300 p-6 text-center dark:border-gray-700">
        <h3 class="mb-2 text-base font-semibold text-gray-700 dark:text-gray-300">
          Add Additional Steps
        </h3>
        <p class="mb-4 text-sm text-gray-500 dark:text-gray-400">
          Customize your pipeline with additional steps like code analysis or security checks.
        </p>
        <button class="rounded-xl border border-gray-300 bg-white px-4 py-2 text-sm font-medium hover:bg-gray-50 dark:border-gray-600 dark:bg-white/[0.05] dark:text-white">
          Add Step
        </button>
      </div>

      <!-- Actions -->
      <div class="flex justify-end gap-3">
        <button class="rounded-xl border border-gray-300 bg-white px-4 py-2 text-sm font-medium hover:bg-gray-50 dark:border-gray-600 dark:bg-white/[0.05] dark:text-white">
          Cancel
        </button>
        <button class="rounded-xl bg-black px-4 py-2 text-sm font-medium text-white hover:bg-gray-800">
          Save Configuration
        </button>
      </div>
    </div>
  </AdminLayout>
</template>
