<template>
  <router-view/>
</template>

<script lang="ts">
import Vue from 'vue'
import { Component, Prop } from 'vue-property-decorator'
import { Location } from 'vue-router'

interface DrawerLink {
  icon: string,
  title: string,
  to: Location,
}

@Component
export default class App extends Vue {
  @Prop({
    type: String
  })
  source!: string
  drawer: boolean = false
  navigation: DrawerLink[] = [
    {
      icon: 'move_to_inbox',
      title: 'Home',
      to: {
        name: 'home'
      }
    },
    {
      icon: 'monetization_on',
      title: 'Calculator',
      to: {
        name: 'calculator'
      }
    },
    {
      icon: 'account_box',
      title: 'Accounts',
      to: { name: 'accounts' }
    },
    {
      icon: 'group',
      title: 'Groups',
      to: { name: 'groups' }
    },
    {
      icon: 'account_balance_wallet',
      title: 'Transactions',
      to: { name: 'transactions' }
    }]
}

</script>

<style lang="scss">
  .md-app {
    min-height: 100vh;
    border: 1px solid rgba(#000, .12);
  }

  .md-app-toolbar {
    height: 196px;
  }

  .md-app-container {
    overflow: hidden;
  }
</style>
