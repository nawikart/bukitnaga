<style>
 
  .overlayed2.v-parallax {
    position: relative;
  }

    .overlayed2.v-parallax:before {
      content: '';
      height: 2000px;
      display: table;
      width: 100%;
      position: absolute;
      /* background-color: rgba(225, 225, 225, .6); */
      background-color: rgba(0, 0, 0, .7);
      z-index: 2;
    }
    
</style>
<template>
  <v-parallax class="overlayed2" src="/assets/images/slide4.jpg" height="525" >
    <v-container grid-list-xl>
      <v-layout wrap xs12>
        <v-flex display-1 xs12 text-xs-center font-weight-medium my-3>Photo Galleries</v-flex>
        <v-flex xs12
                sm12
                md6
                lg6
                d-flex>
          <v-img class="d-flex" gradient src="/assets/images/leftgallery.jpg" aspect-ratio="1.7" />
        </v-flex>



        <v-flex xs12
                sm12
                md6
                lg6
                d-flex>
          <v-layout wrap>
            <v-flex xs6>
              <v-img class="d-flex" gradient src="/assets/images/right1gallery.jpg" aspect-ratio="1.7" />
            </v-flex>
            <v-flex xs6>
              <v-img class="d-flex" gradient src="/assets/images/right2gallery.jpg" aspect-ratio="1.7" />
            </v-flex>
            <v-flex xs6>
              <v-img class="d-flex" gradient src="/assets/images/right3gallery.jpg" aspect-ratio="1.7" />
            </v-flex>
            <v-flex xs6>
              <v-img class="d-flex" gradient src="/assets/images/right4gallery.jpg" aspect-ratio="1.7" />
            </v-flex>
          </v-layout>
        </v-flex>



      </v-layout>
    </v-container>
  </v-parallax>
</template>
<script>
  export default {
    data() {
      return {
        items: [
          {
            src: '/slide1.jpg'
          },
          {
            src: '/slide2.jpg'
          },
          {
            src: '/slide3.jpg'
          },
          {
            src: '/slide4.jpg'
          }
        ]
      }
    }
  }</script>
