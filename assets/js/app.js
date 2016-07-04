new Vue({
  el: '#orders',
  data: {
    order: {
      type: 'order',
      attributes: {
        id: 0,
        product_title: 'test',
        product_description: 'test',
      }
    },
    orders: [],
    },
  ready: function() {
    this.orderEndpoint = this.$resource('api/v1/order/{id}');
    this.orderDeleteEndpoint = this.$resource('api/v1/order/{id}/delete');
    this.fetchOrders();
  },
  methods: {
    fetchOrders: function() {
      this.orderEndpoint.get().then(function(response) {
        this.$set('orders', response.data.data);
      }, function(response) {
        console.log(response);
      });
    },
    postOrder: function() {
        this.orderEndpoint.save({data: this.order}).then(function(response) {
            this.orders.push(response.data.data);
        }, function(response) {
            console.log(response);
        });
        this.order.type.attributes = {
            product_title: '',
            product_description: '',
        };
    },
    deleteOrder: function(index) {
        console.log("index passed in is" + index)
        console.log("id of indexed item is " + this.orders[index].attributes.id)
        this.orderDeleteEndpoint.get({id: this.orders[index].attributes.id}).then(function(response) {
            this.orders.splice(index, 1);
        }, function(response) {
            console.log(response);
        });
    }
  }
});
