<script setup lang="ts">
import { onUpdated, ref } from 'vue'

const props = defineProps<{
  modalDisplayed: boolean
  modalDisplaySetter: (status: boolean) => any
  confirmCallback: Function
}>()

function handleConfim(): void {
  props.confirmCallback()
  props.modalDisplaySetter(false)
}

const refClear = ref<HTMLButtonElement>()

onUpdated((): void => {
  refClear.value?.focus()
})

function tabConfirm(): void {
  refClear.value?.focus()
}
</script>

<template>
  <div id="confirm-modal" v-show="modalDisplayed">
    <span>确认操作吗？</span>

    <div class="button-group">
      <button ref="refClear" @click="() => modalDisplaySetter(false)">
        取消
      </button>

      <button @click="handleConfim" @blur="tabConfirm">确认</button>
    </div>
  </div>
</template>

<style scoped>
#confirm-modal {
  align-items: center;
  background: #eee;
  border-radius: 5px;
  display: flex;
  flex-direction: column;
  height: 10rem;
  justify-content: center;
  position: fixed;
  width: 20rem;
}

.button-group {
  display: flex;
  justify-content: space-between;
  margin: 1rem 0 0;
  width: 10rem;
}
</style>
