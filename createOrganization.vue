// in this vue and I want  to create an organization with field name then select plan to pay for
<template>
  <FullScreenLayout>
    <div class="relative p-6 bg-white z-1 dark:bg-gray-900 sm:p-0">
      <NoSignUpHeaderMenu />
      <div
        class="flex flex-col items-start justify-center min-h-screen dark:bg-gray-900 p-2 sm:p-0"
      >
        <!-- Main Content Container -->
        <div class="w-full max-w-3xl mx-auto">
          <!-- Centered Header -->
          <div class="mb-8">
            <h1 class="text-2xl font-semibold text-gray-800 dark:text-white/90">
              Create Your Organization
            </h1>
          </div>

          <!-- Form Container -->
          <form @submit.prevent="handleSubmit" class="space-y-6">
            <div class="space-y-4 flex flex-col items-start">
              <!-- Organization Name -->
              <div class="w-[40%]">
                <label
                  for="organizationName"
                  class="mb-1 block text-sm font-medium text-gray-700 dark:text-gray-400"
                >
                  Organization Name<span class="text-error-500">*</span>
                </label>
                <input
                  v-model="organizationName"
                  type="text"
                  id="organizationName"
                  name="organizationName"
                  placeholder="Enter your organization name"
                  class="dark:bg-dark-900 h-11 w-full rounded-lg border border-gray-300 bg-transparent px-4 py-2.5 text-sm text-gray-800 shadow-theme-xs placeholder:text-gray-400 focus:border-brand-300 focus:outline-hidden focus:ring-3 focus:ring-brand-500/10 dark:border-gray-700 dark:bg-gray-900 dark:text-white/90 dark:placeholder:text-white/30 dark:focus:border-brand-800"
                />
              </div>

              <!-- Plan Selection -->
              <div class="w-[85%]">
                <label class="mb-1 block text-sm font-medium text-gray-700 dark:text-gray-400">
                  Select Plan<span class="text-error-500">*</span>
                </label>
                <div class="space-y-2 py-2">
                  <div
                    v-for="plan in plans"
                    :key="plan.id"
                    class="flex items-center p-4 my-3 rounded-lg border border-gray-200 dark:border-gray-700"
                  >
                    <input
                      type="radio"
                      :id="`plan-${plan.id}`"
                      :value="plan.id"
                      v-model="selectedPlan"
                      name="plan"
                      class="h-4 w-4 text-brand-600 focus:ring-brand-500 border-gray-300 dark:border-gray-700 dark:bg-gray-900"
                    />
                    <label
                      :for="`plan-${plan.id}`"
                      class="ml-3 block text-sm text-gray-800 dark:text-white/90"
                    >
                      <div>
                        <span class="font-medium">{{ plan.name }}</span> - ${{ plan.price }}/month
                      </div>
                      <div class="text-xs text-gray-500 dark:text-gray-400">
                        {{ plan.description }}
                      </div>
                    </label>
                  </div>
                </div>
              </div>

              <!-- Terms and Conditions -->
              <div class="flex items-center w-full">
                <input
                  v-model="termsAccepted"
                  type="checkbox"
                  id="terms"
                  name="terms"
                  class="h-4 w-4 rounded border-gray-300 text-brand-600 focus:ring-brand-500 dark:border-gray-700 dark:bg-gray-900 dark:focus:ring-brand-800"
                />
                <label for="terms" class="ml-2 block text-sm text-gray-700 dark:text-gray-400">
                  I accept the
                  <router-link
                    to="/terms"
                    class="text-brand-500 hover:text-brand-600 dark:text-brand-400"
                    >Terms and Conditions</router-link
                  >
                </label>
              </div>

              <!-- Submit Button -->
              <div class="w-full flex justify-center">
                <BaseButton
                  type="submit"
                  style="
                    background-color: #2563eb;
                    color: #ffffff;
                    border-radius: 0.375rem;
                    padding: 0.5rem 1rem;
                    font-size: 1rem;
                    font-weight: 600;
                    transition: background-color 0.2s ease;
                    width: 40%;
                  "
                >
                  Create Organization
                </BaseButton>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  </FullScreenLayout>
</template>
<script setup lang="ts">
import FullScreenLayout from '@/components/layout/FullScreenLayout.vue'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import NoSignUpHeaderMenu from '@/components/layout/header/noSignUpHeaderMenu.vue'
import BaseButton from '@/components/common/BaseButton.vue'
import { API_URLS } from '@/controller/services/service'
import { toast } from 'vue3-toastify'
import 'vue3-toastify/dist/index.css'
import { useOrganizationStore } from '@/stores/organization_store.ts'

const organizationStore = useOrganizationStore()
const router = useRouter()
const organizationName = ref('')
const selectedPlan = ref('')
const plans = ref([
  // Example plans, replace with your actual plans
  { id: 'free', name: 'Free Plan', price: 0, description: 'Basic features with limited usage' },
  { id: 'basic', name: 'Basic Plan', price: 10, description: 'Standard features for small teams' },
  {
    id: 'pro',
    name: 'Pro Plan',
    price: 20,
    description: 'Advanced features for growing businesses',
  },
  {
    id: 'enterprise',
    name: 'Enterprise Plan',
    price: 50,
    description: 'All features for large organizations',
  },
])
const termsAccepted = ref(false)
const createOrganization = async () => {
  if (!organizationName.value || !selectedPlan.value) {
    alert('Please fill in all fields.')
    return
  }

  fetch(API_URLS.BASE_URL + API_URLS.CREATE_ORGANIZATION, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${localStorage.getItem('token')}`,
      // Include auth token if required
    },
    body: JSON.stringify({
      name: organizationName.value,
      plan: selectedPlan.value,
    }),
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error('Network response was not ok')
      }

      toast('Wow so easy !', {
        autoClose: 1000,
      }) // ToastOptions
      // ToastOptions

      return response.json()
    })
    .then((data) => {
      console.log('Organization created successfully:', data)
      organizationStore.setOrganizationId(data.org_id)
      organizationStore.setOrganizationName(data.org_name)
      // Redirect to dashboard or another page after successful creation
      router.push('/dashboard')
    })
    .catch((error) => {
      console.error('Error creating organization:', error)
      toast('Error creating organization', { autoClose: 1000 })
    })
}

const handleSubmit = async () => {
  await createOrganization()
}
</script>
<style scoped>
.container {
  max-width: 600px;
  margin: auto;
}
input,
select {
  padding: 0.5rem;
  border: 1px solid #d1d5db; /* Tailwind gray-300 */
  border-radius: 0.375rem; /* Tailwind rounded-md */
}
button {
  transition: background-color 0.2s ease;
}
button:hover {
  background-color: #2563eb; /* Tailwind blue-600 */
}
</style>

<style scoped>
.container {
  max-width: 600px;
  margin: auto;
}
input,
select {
  padding: 0.5rem;
  border: 1px solid #d1d5db; /* Tailwind gray-300 */
  border-radius: 0.375rem; /* Tailwind rounded-md */
}
button {
  transition: background-color 0.2s ease;
}
button:hover {
  background-color: #2563eb; /* Tailwind blue-600 */
}
.dark {
  background-color: #1f2937; /* Tailwind gray-900 */
}
.dark input,
.dark select {
  background-color: #374151; /* Tailwind gray-800 */
  color: #f3f4f6; /* Tailwind gray-200 */
  border-color: #4b5563; /* Tailwind gray-600 */
}
.dark button {
  background-color: #2563eb; /* Tailwind blue-600 */
  color: #ffffff; /* Tailwind white */
}
.dark button:hover {
  background-color: #1d4ed8; /* Tailwind blue-700 */
}
.dark .text-brand-500 {
  color: #3b82f6; /* Tailwind blue-500 */
}
.dark .text-brand-600 {
  color: #2563eb; /* Tailwind blue-600 */
}
.dark .text-brand-300 {
  color: #93c5fd; /* Tailwind blue-300 */
}
.dark .text-brand-400 {
  color: #60a5fa; /* Tailwind blue-400 */
}
.dark .text-brand-800 {
  color: #1e40af; /* Tailwind blue-800 */
}
.dark .text-brand-900 {
  color: #1e3a8a; /* Tailwind blue-900 */
}
.dark .text-error-500 {
  color: #ef4444; /* Tailwind red-500 */
}
.dark .text-gray-700 {
  color: #d1d5db; /* Tailwind gray-300 */
}
.dark .text-gray-400 {
  color: #9ca3af; /* Tailwind gray-400 */
}
.dark .text-gray-800 {
  color: #f3f4f6; /* Tailwind gray-200 */
}
.dark .text-gray-900 {
  color: #e5e7eb; /* Tailwind gray-100 */
}
.dark .text-white {
  color: #ffffff; /* Tailwind white */
}
/* Base button styles */
/* Remove invalid selectors and use Tailwind utility classes directly in your template elements.
   If you need to customize dark mode for buttons, target the button or BaseButton component with a valid selector. */
</style>


function toast(arg0: string, arg1: { autoClose: number }) {
  throw new Error('Function not implemented.')
}

function toast(arg0: string, arg1: { autoClose: number }) {
  throw new Error('Function not implemented.')
}
