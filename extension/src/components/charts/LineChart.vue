<template>
  <div class="h-[140px] bg-gray-100 rounded-md">
    <VisXYContainer
      :data="data"
      :padding="{ bottom: 0, top: 10, left: 10, right: 10 }"
      width="99%"
      class="h-[130px]"
      :style="{
        opacity: 0.9
      }"
    >
      <VisStackedBar
        :x="(d: Data, i: number) => i"
        :y="(d: Data) => d.value"
        color="#000000"
        :bar-padding="0.1"
        :rounded-corners="10"
      />
      <VisAxis type="x" :numTicks="7" :tickFormat="(idx) => data[idx].day" />
    </VisXYContainer>
  </div>
</template>
<script setup lang="ts">
import { VisStackedBar, VisXYContainer, VisAxis } from '@unovis/vue'

type Data = (typeof data)[number]
const data = [
  { value: 400, day: 'S' },
  { value: 300, day: 'M' },
  { value: 200, day: 'T' },
  { value: 300, day: 'W' },
  { value: 200, day: 'T' },
  { value: 278, day: 'F' },
  { value: 189, day: 'S' }
]
</script>
