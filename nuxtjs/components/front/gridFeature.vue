
<template>
    <v-flex grey lighten-3  >      
        <v-layout id="grid6" row wrap py-5 font-weight-thin subheading >
          <v-flex xs12 text-xs-center font-weight-medium display-1 mb-5 style="color:#444;">Our Recommended Events</v-flex>
          <v-flex item xs12 sm6 md4 :key="1" px-4 py-4>
            <v-icon medium color="teal lighten-1" style="float:left;margin-right:10px;">home</v-icon>
            <h3>WEDDINGS</h3>
            <p style="font-size:14px;margin-top:10px;">Villa Bukit Naga is an ideal setting for Weddings & Engagements. Set against a breathtaking view of Gianyar forest and Mount Batukau Where vows can exchange in the enchanting Secret Garden “A Perfect Match” For an intimate wedding ceremony. The enchanting Secret Garden can accommodate 40 guests comfortably</p>
            <p><v-btn flat color="grey" href="weddings">learn more</v-btn></p>
          </v-flex>
          <v-flex item xs12 sm6 md4 :key="2" px-4 py-4>
            <v-icon medium color="teal lighten-1" style="float:left;margin-right:10px;">event</v-icon>
            <h3>BIRTHDAYS</h3>
            <p style="font-size:14px;margin-top:10px;">Planning to celebrate your birthday in Bali? Forget noisy restaurants and have an exclusive birthday party at Villa Bukit Naga. The villa is perfectly suited for birthdays, whether you decide to be seated in air-conditioned comfort or in the outdoors at the poolside terrace</p>
            <p><v-btn flat color="grey" href="birthday">learn more</v-btn></p>
          </v-flex>
          <v-flex item xs12 sm6 md4 :key="3" px-4 py-4>
            <v-icon medium color="teal lighten-1" style="float:left;margin-right:10px;">group</v-icon>
            <h3>MEETINGS</h3>
            <p style="font-size:14px;margin-top:10px;">Looking for business meeting venue? Villa Bukit Naga is exclusive, spacious and totally private. Please contact our Reservation Manager Erma for details Email: reservation@villabukitnaga.com Mobile/WhatsApp : +62 811 388 1122 </p>
            <p><v-btn flat color="grey" href="meetings">learn more</v-btn></p>
          </v-flex>
          <v-flex item xs12 sm6 md4 :key="4" px-4 py-4>
            <v-icon medium color="teal lighten-1" style="float:left;margin-right:10px;">business</v-icon>
            <h3>CORPORATE FUNCTIONS</h3>
            <p style="font-size:14px;margin-top:10px;">Perfectly suited for corporate functions. With a capacity of up to 100 persons for sit-down party, the design of the estate has unobtrusive view of the horizon and the grandeur of the tropical valley </p>
            <p><v-btn flat color="grey" href="corporate">learn more</v-btn></p>
          </v-flex>
          <v-flex item xs12 sm6 md4 :key="5" px-4 py-4>
            <v-icon medium color="teal lighten-1" style="float:left;margin-right:10px;">sentiment_very_satisfied</v-icon>
            <h3>REUNIONS</h3>
            <p style="font-size:14px;margin-top:10px;">Villa Bukit Naga makes a perfect venue for any gathering. With a capacity of up to 100 persons for sit-down party the estate has been designed for intimate social gatherings set against a lush forest backdrop </p>
            <p><v-btn flat color="grey" href="reunions">learn more</v-btn></p>
          </v-flex>
          <v-flex item xs12 sm6 md4 :key="6" px-4 py-4>
            <v-icon medium color="teal lighten-1" style="float:left;margin-right:10px;">cloud_circle</v-icon>

            <h3>SPECIAL EVENTS</h3>
            <p style="font-size:14px;margin-top:10px;">The estate is splendid venue for any function. With a capacity of up to 100 persons for sit-down party, the estate has been designed for intimate social gatherings and functions set against a magnificent view of Gianyar forest and Mount Batukau </p>
            <p><v-btn flat color="grey" href="special">learn more</v-btn></p>
          </v-flex>
        </v-layout>      
    </v-flex>
 
</template>

<style>
  @media only screen and (min-width: 1025px) {
    #grid6 {
      max-width: 1280px;
      margin: auto;
    }

      #grid6 .item {
        padding: 38px !important;
      }

        #grid6 .item:nth-child(2) {
          border-bottom: 1px solid #ffffff;
        }

        #grid6 .item:nth-child(3) {
          border-left: 1px solid #ffffff;
          border-bottom: 1px solid #ffffff;
          border-right: 1px solid #ffffff;
        }

        #grid6 .item:nth-child(4) {
          border-bottom: 1px solid #ffffff;
        }

        #grid6 .item:nth-child(6) {
          border-left: 1px solid #ffffff;
          border-right: 1px solid #ffffff;
        }
  }

  @media only screen and (min-width: 601px) and (max-width: 1024px) {
    #grid6 .item:nth-child(2) {
      border-bottom: 1px solid #ffffff;
    }

    #grid6 .item:nth-child(3) {
      border-left: 1px solid #ffffff;
      border-bottom: 1px solid #ffffff;
    }

    #grid6 .item:nth-child(4) {
      border-right: 1px solid #ffffff;
      border-bottom: 1px solid #ffffff;
    }

    #grid6 .item:nth-child(5) {
      border-bottom: 1px solid #ffffff;
    }

    #grid6 .item:nth-child(6) {
      border-right: 1px solid #ffffff;
    }
  }

  @media only screen and (max-width: 600px) {
    #grid6 .item {
      border-bottom: 1px solid #ffffff;
    }
  }
</style>

