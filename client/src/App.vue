<script setup lang="ts">
import ConfirmModal from './features/confirm-modal/components/ConfirmModal.vue'

import { ref, watch } from 'vue'
import http from './lib/axios'

const refInput = ref<HTMLInputElement>()

const balance = ref<number>(0)
const setBalance = (money: number) => (balance.value = money)

watch<number>(balance, () => {
  http({ method: 'GET' })
  console.log('post')
})

function changeBalance(): void {
  const money = Number(refInput.value?.value ?? 0)
  setBalance(balance.value ? balance.value - money : money)
  refInput.value && (refInput.value.value = '')
}

function clearBalance(): void {
  setBalance(0)
}

const confirmModalDisplayed = ref<boolean>(false)

const setConfirmModalDisplayed = (status: boolean) =>
  (confirmModalDisplayed.value = status)

const confirmCallback = ref<Function>(() => {})

const setConfirmCallback = (callback: Function) =>
  (confirmCallback.value = callback)

function confirmClear(): void {
  setConfirmCallback(clearBalance)
  setConfirmModalDisplayed(true)
}

function confirmChange(): void {
  setConfirmCallback(changeBalance)
  setConfirmModalDisplayed(true)
}
</script>

<template>
  <div class="app">
    <ConfirmModal
      :modal-displayed="confirmModalDisplayed"
      :modal-display-setter="setConfirmModalDisplayed"
      :confirm-callback="confirmCallback"
    />

    <div class="main-wrapper">
      <span class="balance">{{ balance }} JPY</span>

      <input
        type="number"
        ref="refInput"
        class="inputter"
        @keypress.enter="confirmChange"
      />

      <div class="button-group">
        <button @click="confirmClear">清空</button>
        <button @click="confirmChange">确认</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.app {
  align-items: center;
  background: #888;
  display: flex;
  flex-direction: column;
  height: 100vh;
  justify-content: center;
}

.main-wrapper {
  display: flex;
  flex-direction: column;
  width: 10rem;
}

.balance {
  align-self: flex-end;
}

.inputter {
  box-sizing: border-box;
  margin: 1rem 0 0;
}

.button-group {
  display: flex;
  justify-content: space-between;
  margin: 1rem 0 0;
}
</style>
