<template>
  <div class="space-y-6">
    <!-- API Key Management Section -->
    <div class="bg-gradient-to-br from-blue-50 to-indigo-50 dark:from-blue-900/10 dark:to-indigo-900/10 rounded-xl border border-blue-200 dark:border-blue-800 p-6">
      <div class="flex items-center gap-3 mb-6">
        <div class="w-12 h-12 bg-gradient-to-br from-blue-500 to-indigo-600 rounded-lg flex items-center justify-center">
          <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 7a2 2 0 012 2m4 0a6 6 0 01-7.743 5.743L11 17H9v2H7v2H4a1 1 0 01-1-1v-2.586a1 1 0 01.293-.707l5.964-5.964A6 6 0 1721 9z"></path>
          </svg>
        </div>
        <div>
          <h3 class="text-xl font-bold text-gray-900 dark:text-white">API Key Management</h3>
          <p class="text-sm text-gray-600 dark:text-gray-400">
            Secure keys for integrating ErrorZen with your application
          </p>
        </div>
      </div>

      <!-- Current API Key -->
      <div class="bg-white dark:bg-gray-800 rounded-lg border border-gray-200 dark:border-gray-700 p-4 mb-6">
        <div class="flex items-center justify-between mb-3">
          <h4 class="text-sm font-semibold text-gray-700 dark:text-gray-300">Current API Key</h4>
          <div class="flex items-center gap-2 text-xs">
            <div class="w-2 h-2 bg-green-500 rounded-full animate-pulse"></div>
            <span class="text-green-600 dark:text-green-400 font-medium">Active</span>
          </div>
        </div>
        
        <div class="relative group">
          <input
            type="password"
            :value="maskedSecret"
            readonly
            class="w-full bg-gray-50 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg px-4 py-3 pr-32 text-sm font-mono text-gray-900 dark:text-white focus:outline-none focus:ring-2 focus:ring-blue-500 cursor-pointer"
            @click="toggleSecretVisibility"
            @focus="handleFocus"
          />
          
          <div class="absolute right-2 top-1/2 transform -translate-y-1/2 flex gap-1">
            <button
              class="px-3 py-1.5 bg-gray-600 hover:bg-gray-700 text-white text-xs font-medium rounded-md transition-colors flex items-center gap-1"
              @click="toggleSecretVisibility"
              type="button"
            >
              <svg v-if="!isSecretVisible" class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
              </svg>
              <svg v-else class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L3 3m6.878 6.878L21 21"></path>
              </svg>
              {{ isSecretVisible ? 'Hide' : 'Show' }}
            </button>
            
            <button
              class="px-3 py-1.5 bg-blue-600 hover:bg-blue-700 text-white text-xs font-medium rounded-md transition-colors flex items-center gap-1"
              @click="copySecret"
              type="button"
            >
              <svg v-if="!copied" class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
              </svg>
              <svg v-else class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
              </svg>
              {{ copied ? 'Copied!' : 'Copy' }}
            </button>
          </div>
        </div>
        
        <p class="mt-2 text-xs text-gray-500 dark:text-gray-400">
          Keep this key secure. It provides full access to your project's error tracking data.
        </p>
      </div>

      <!-- Key Information -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
        <div class="bg-white dark:bg-gray-800 rounded-lg border border-gray-200 dark:border-gray-700 p-4">
          <div class="flex items-center gap-2 mb-2">
            <svg class="w-4 h-4 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
            <span class="text-sm font-medium text-gray-700 dark:text-gray-300">Created</span>
          </div>
          <p class="text-sm text-gray-600 dark:text-gray-400">{{ formatDate(keyCreatedDate) }}</p>
        </div>
        
        <div class="bg-white dark:bg-gray-800 rounded-lg border border-gray-200 dark:border-gray-700 p-4">
          <div class="flex items-center gap-2 mb-2">
            <svg class="w-4 h-4 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
            </svg>
            <span class="text-sm font-medium text-gray-700 dark:text-gray-300">Last Used</span>
          </div>
          <p class="text-sm text-gray-600 dark:text-gray-400">{{ formatDate(lastUsedDate) }}</p>
        </div>
        
        <div class="bg-white dark:bg-gray-800 rounded-lg border border-gray-200 dark:border-gray-700 p-4">
          <div class="flex items-center gap-2 mb-2">
            <svg class="w-4 h-4 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
            <span class="text-sm font-medium text-gray-700 dark:text-gray-300">API Calls</span>
          </div>
          <p class="text-sm text-gray-600 dark:text-gray-400">{{ apiCallsCount.toLocaleString() }} this month</p>
        </div>
      </div>

      <!-- Actions -->
      <div class="flex flex-wrap gap-3">
        <button 
          class="flex items-center gap-2 px-4 py-2.5 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-lg transition-colors shadow-sm hover:shadow-md"
          @click="generateNewKey"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
          </svg>
          Generate New Key
        </button>
        
        <button 
          class="flex items-center gap-2 px-4 py-2.5 bg-gray-100 hover:bg-gray-200 dark:bg-gray-700 dark:hover:bg-gray-600 text-gray-700 dark:text-gray-300 font-medium rounded-lg transition-colors"
          @click="downloadKeyBackup"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          Download Backup
        </button>
        
        <button 
          class="flex items-center gap-2 px-4 py-2.5 bg-red-50 hover:bg-red-100 dark:bg-red-900/20 dark:hover:bg-red-900/30 text-red-600 dark:text-red-400 font-medium rounded-lg transition-colors border border-red-200 dark:border-red-800"
          @click="showDeleteConfirmation = true"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
          </svg>
          Revoke Key
        </button>
      </div>
    </div>

    <!-- Security Notice -->
    <div class="bg-amber-50 dark:bg-amber-900/10 border border-amber-200 dark:border-amber-800 rounded-lg p-4">
      <div class="flex items-start gap-3">
        <svg class="w-5 h-5 text-amber-600 dark:text-amber-400 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.934-.833-2.464 0L4.35 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
        </svg>
        <div>
          <h4 class="font-semibold text-amber-800 dark:text-amber-300 mb-1">Security Best Practices</h4>
          <ul class="text-sm text-amber-700 dark:text-amber-400 space-y-1">
            <li>• Store your API key securely and never expose it in client-side code</li>
            <li>• Rotate your API key regularly for enhanced security</li>
            <li>• Monitor API usage and revoke keys immediately if compromised</li>
            <li>• Use environment variables to store keys in your deployment</li>
          </ul>
        </div>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <div v-if="showDeleteConfirmation" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white dark:bg-gray-800 rounded-xl shadow-xl max-w-md w-full mx-4 p-6">
        <div class="flex items-center gap-3 mb-4">
          <div class="w-10 h-10 bg-red-100 dark:bg-red-900/20 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-red-600 dark:text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.934-.833-2.464 0L4.35 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-bold text-gray-900 dark:text-white">Revoke API Key</h3>
        </div>
        
        <p class="text-gray-600 dark:text-gray-400 mb-6">
          Are you sure you want to revoke this API key? This action cannot be undone and will immediately stop all API access using this key.
        </p>
        
        <div class="flex gap-3">
          <button 
            class="flex-1 px-4 py-2.5 bg-red-600 hover:bg-red-700 text-white font-medium rounded-lg transition-colors"
            @click="deleteKey"
          >
            Yes, Revoke Key
          </button>
          <button 
            class="flex-1 px-4 py-2.5 bg-gray-200 hover:bg-gray-300 dark:bg-gray-700 dark:hover:bg-gray-600 text-gray-700 dark:text-gray-300 font-medium rounded-lg transition-colors"
            @click="showDeleteConfirmation = false"
          >
            Cancel
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue'
import { API_URLS } from '@/controller/services/service.ts'
import { ElMessage } from 'element-plus'

const fetchProjectSecret = async (projectId: string) => {
  try {
    const token = localStorage.getItem('token')
    if (!token) {
      throw new Error('No authentication token found')
    }
    const url = API_URLS.BASE_URL + API_URLS.GET_PROJECT_KEY + encodeURIComponent(projectId)
    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
    })
    if (!response.ok) {
      throw new Error('Failed to fetch project secret')
    }
    const data = await response.json()
    return data.project_secret
  } catch (error) {
    console.error('Error fetching project secret:', error)
    return null
  }
}

export default defineComponent({
  name: 'ProjectsSecrets',
  props: {
    projectId: {
      type: String,
      required: false,
      default: 'demo-project-id'
    },
  },
  setup(props) {
    const secret = ref<string>('ez_sk_test_1234567890abcdefghijklmnopqrstuvwxyz1234567890')
    const isSecretVisible = ref<boolean>(false)
    const copied = ref<boolean>(false)
    const showDeleteConfirmation = ref<boolean>(false)
    const keyCreatedDate = ref<Date>(new Date('2024-01-15'))
    const lastUsedDate = ref<Date>(new Date('2024-09-30'))
    const apiCallsCount = ref<number>(15248)

    const maskedSecret = ref<string>('')

    const updateMaskedSecret = () => {
      if (isSecretVisible.value) {
        maskedSecret.value = secret.value
      } else {
        const visiblePart = secret.value.slice(0, 12)
        const hiddenPart = '•'.repeat(Math.max(0, secret.value.length - 12))
        maskedSecret.value = visiblePart + hiddenPart
      }
    }

    const toggleSecretVisibility = () => {
      isSecretVisible.value = !isSecretVisible.value
      updateMaskedSecret()
    }

    const handleFocus = (e: Event) => {
      const target = e.target as HTMLInputElement | null
      target?.select()
    }

    const copySecret = async () => {
      if (secret.value) {
        try {
          await navigator.clipboard.writeText(secret.value)
          copied.value = true
          ElMessage.success('API key copied to clipboard')
          setTimeout(() => {
            copied.value = false
          }, 2000)
        } catch {
          ElMessage.error('Failed to copy API key')
        }
      }
    }

    const generateNewKey = () => {
      ElMessage.warning('This feature will generate a new API key. Please implement the API call.')
    }

    const downloadKeyBackup = () => {
      const dataStr = JSON.stringify({
        projectId: props.projectId,
        apiKey: secret.value,
        createdDate: keyCreatedDate.value,
        downloadedAt: new Date().toISOString()
      }, null, 2)
      
      const dataBlob = new Blob([dataStr], { type: 'application/json' })
      const url = URL.createObjectURL(dataBlob)
      const link = document.createElement('a')
      link.href = url
      link.download = `errorzen-api-key-backup-${props.projectId}.json`
      document.body.appendChild(link)
      link.click()
      document.body.removeChild(link)
      URL.revokeObjectURL(url)
      
      ElMessage.success('API key backup downloaded')
    }

    const deleteKey = () => {
      ElMessage.warning('This feature will revoke the API key. Please implement the API call.')
      showDeleteConfirmation.value = false
    }

    const formatDate = (date: Date) => {
      return date.toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'short',
        day: 'numeric'
      })
    }

    onMounted(async () => {
      if (props.projectId && props.projectId !== 'demo-project-id') {
        const fetchedSecret = await fetchProjectSecret(props.projectId)
        if (fetchedSecret) {
          secret.value = fetchedSecret
        } else {
          console.warn('No secret returned from API')
        }
      }
      updateMaskedSecret()
    })

    return {
      secret,
      isSecretVisible,
      copied,
      showDeleteConfirmation,
      keyCreatedDate,
      lastUsedDate,
      apiCallsCount,
      maskedSecret,
      toggleSecretVisibility,
      handleFocus,
      copySecret,
      generateNewKey,
      downloadKeyBackup,
      deleteKey,
      formatDate,
    }
  },
})
</script>

<style scoped>
</style>
