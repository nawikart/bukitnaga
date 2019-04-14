<template>
  <v-flex xs12 grey lighten-3 py-5 px-1>
    <v-flex px-2 xs12 text-xs-center font-weight-thin display-2 mt-2 mb-2>OUR PRESS</v-flex>
    <v-flex px-3 text-xs-center font-weight-thin subheading>
      For media enquiries, please contact: eMail: info@villabukitnaga.com, Tel - Bali // +62 811 388 11 22
    </v-flex>
    <v-layout id="press" mt-4 row wrap font-weight-light subheading>
      <v-flex item xs12 sm6 md4 py-4>
        <img src="/press/press_baccarat_2009.jpg" width="100" style="float:left;margin-right:20px;margin-top:9px;" />
        <h3 style="margin:10px 0;">BACCARAT - Jun 2009</h3>
        <p style="font-size:13px;">
          villa bukit naga. serene estate of six spacious suites based around a spiral staircased main house. volcanic views, literally.
          <br>bukitnaga.com sleeps 12 / 10 mins from ubud.
        </p>
        <v-btn small outline color="info" href="/press/baccarat_2009.pdf" style="margin-top:15px;" >download pdf</v-btn>
      </v-flex>
      <v-flex item xs12 sm6 md4 py-4>
        <img src="/press/timeout.jpg" width="100" style="float:left;margin-right:20px;margin-top:9px;" />
        <h3 style="margin:10px 0;">TimeOut - Nov 2009</h3>
        <p style="font-size:13px;">

          Villa Bukit Naga's decor is naturecentric with earthy colours and natural balinese materials. The lush green woodlands serve
        </p>
        <v-btn small outline color="info" href="/press/TimeOut_2009.pdf" style="margin-top:15px;">download pdf</v-btn>
      </v-flex>
      <v-flex item xs12 sm6 md4 py-4>
        <img src="/press/elle_2010.jpg" width="100" style="float:left;margin-right:20px;margin-top:9px;" />
        <h3 style="margin:10px 0;">ELLE - April 2010</h3>
        <p style="font-size:13px;">


          Villa Bukit Naga selected as Bali's most beautiful estate.
        </p>
        <v-btn small outline color="info" href="/press/baccarat_2009.pdf" style="margin-top:15px;">download pdf</v-btn>
      </v-flex>

      <v-flex item xs12 sm6 md4 py-4>
        <img src="/press/elle.jpg" width="100" style="float:left;margin-right:20px;margin-top:9px;" />
        <h3 style="margin:10px 0;">ELLE - April 2010</h3>
        <p style="font-size:13px;">


          Elle China selected Villa Bukit Naga as the most beautiful villa in Bali, and shot 20 pages of fashion and beauty at our estate.

          A behind-the-scenes video will also be shown on ellechina.com.
        </p>
        <v-btn small outline color="info" href="/press/ELLE_2010.pdf" style="margin-top:15px;">download pdf</v-btn>
      </v-flex>


      <v-flex item xs12 sm6 md4 py-4>
        <img src="/press/TRAVELLEISUREJuly2010.jpg" width="100" style="float:left;margin-right:20px;margin-top:9px;" />
        <h3 style="margin:10px 0;">TRAVEL + LEISURE, July 2010</h3>
        <p style="font-size:13px;">
          Villa Bukit Naga leads this feature on best villas in Asia for luxury holidays with groups of friends or families.
        </p>
        <v-btn small outline color="info" href="/press/TRAVELLEISURE July 2010.pdf.pdf" style="margin-top:15px;">download pdf</v-btn>
      </v-flex>
      <v-flex item xs12 sm6 md4 py-4>
        <img src="/press/NewYorkTimes.jpg" width="100" style="float:left;margin-right:20px;margin-top:9px;" />
        <h3 style="margin:10px 0;">New York Times online, January 2011</h3>
        <p style="font-size:13px;">


          International Herald Tribune print edition, January 2011.
        </p>
        <v-btn small outline color="info" href="/press/New York Times online January 2011.pdf" style="margin-top:15px;">download pdf</v-btn>
      </v-flex>
    </v-layout>
  </v-flex>
</template>

<style>
  #press {
    max-width: 1200px;
    margin: 0 auto;
  }

  @media only screen and (min-width: 1025px) {
    #press .item:nth-child(1) {
      border-bottom: 1px solid #ffffff;
    }

    #press .item:nth-child(2) {
      border-left: 1px solid #ffffff;
      border-bottom: 1px solid #ffffff;
      border-right: 1px solid #ffffff;
    }

    #press .item:nth-child(3) {
      border-bottom: 1px solid #ffffff;
    }

    #press .item:nth-child(5) {
      border-left: 1px solid #ffffff;
      border-right: 1px solid #ffffff;
    }
  }

  @media only screen and (min-width: 601px) {
    #press .item {
      padding: 30px;
    }
  }

  @media only screen and (min-width: 601px) and (max-width: 1024px) {
    #press .item:nth-child(1) {
      border-bottom: 1px solid #ffffff;
    }

    #press .item:nth-child(2) {
      border-left: 1px solid #ffffff;
      border-bottom: 1px solid #ffffff;
    }

    #press .item:nth-child(3) {
      border-right: 1px solid #ffffff;
      border-bottom: 1px solid #ffffff;
    }

    #press .item:nth-child(4) {
      border-bottom: 1px solid #ffffff;
    }

    #press .item:nth-child(5) {
      border-right: 1px solid #ffffff;
    }
  }

  @media only screen and (max-width: 600px) {
    #press .item {
      border-bottom: 1px solid #ffffff;
      padding: 15px;
    }
  }
</style>
