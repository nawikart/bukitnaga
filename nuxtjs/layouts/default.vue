<template>
  <v-app>
    <v-toolbar flat color="grey lighten-5" style="max-width:1136px;margin:0 auto;">
      <v-toolbar-side-icon></v-toolbar-side-icon>
      <v-toolbar-title>Bukit Naga</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items class="hidden-sm-and-down">
        <v-btn flat href="/"><nuxt-link to="/">Home</nuxt-link></v-btn>
        <v-btn flat href="/villa"><nuxt-link to="/villa">Villa</nuxt-link></v-btn>
        <v-btn flat href="/about-bali"><nuxt-link to="/about-bali">About Bali</nuxt-link></v-btn>
        <v-btn flat href="/faqs"><nuxt-link to="/faqs">FAQs</nuxt-link></v-btn>
        <v-btn flat href="/contact"><nuxt-link to="/contact">Contact</nuxt-link></v-btn>
        <v-btn flat href="/press"><nuxt-link to="/press">Press</nuxt-link></v-btn>
        <v-btn flat href="/reservation"><nuxt-link to="/reservation">Reservation</nuxt-link></v-btn>
      </v-toolbar-items>
    </v-toolbar>    
    <nuxt />
    <v-flex xs12 white darken-2>
      <v-container>
        <v-layout row wrap black--text font-weight-light subheading py-3>
          <v-flex xs12 sm6 md4>Copyright © 2019 Villa Bukit Naga Bali. <br />All rights reserved.</v-flex>
          <v-flex xs12 sm6 md4></v-flex>
          <v-flex xs12 sm6 md4>
            <v-img src="/link.png" width="250" height="50" />
          </v-flex>
        </v-layout>
      </v-container>
    </v-flex>
  </v-app>
</template>

<script>
  export default {
    data() {
      return {
        clipped: false,
        drawer: true,
        fixed: false,
        items: [
          { icon: 'apps', title: 'Welcome', to: '/' },
          { icon: 'bubble_chart', title: 'Inspire', to: '/inspire' }
        ],
        miniVariant: false,
        right: true,
        rightDrawer: false,
        title: 'Vuetify.js'
      }
    }
  }
</script>
