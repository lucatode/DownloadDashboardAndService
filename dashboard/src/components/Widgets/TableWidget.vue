<template>
<div >
  <h3>Top Downloads</h3>
  <vuetable ref="vuetable"
    api-url="https://gomicroservice.herokuapp.com/vueTableData"
    :fields="fields"
    :sort-order="sortOrder"
    @vuetable:pagination-data="onPaginationData"
  ></vuetable>
  <vuetable-pagination ref="pagination" @vuetable-pagination:change-page="onChangePage"></vuetable-pagination>
</div>
</template>

<script>
import accounting from 'accounting'
  import moment from 'moment'
import Vuetable from 'vuetable-2/src/components/Vuetable.vue'
import VuetablePagination from 'vuetable-2/src/components/VuetablePagination.vue'

export default {
  data(){
    return {
      fields: [
        {
          name: 'CountryDetails.name',
          title: 'Country',
          titleClass: 'text-center',
        },
        {
          name: 'CountryDetails.region',
          title: 'Region',
          
          titleClass: 'text-center',
        },
        {
          name: 'Count',
          title: 'Count',
          sortField: 'Count',
          titleClass: 'text-center',
        }
      ],
      sortOrder: [
          {
            field: 'Count',
            sortField: 'Count',
            direction: 'asc'
          }
        ]
    }
  },
  components: {
    Vuetable,
    'vuetable-pagination':VuetablePagination
  },
  methods:{
    onPaginationData (paginationData) {
      this.$refs.pagination.setPaginationData(paginationData)
    },
    onChangePage (page) {
      this.$refs.vuetable.changePage(page)
    }
  }
}
</script>

<style></style>