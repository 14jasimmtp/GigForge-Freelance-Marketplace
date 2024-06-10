function getQueryParam(param) {
    const urlParams = new URLSearchParams(window.location.search);
    return urlParams.get(param);
}

// Render PayPal buttons
function renderPayPalButtons() {
  paypal.Buttons({
      createOrder: function(data, actions) {
        const ID = getQueryParam('orderID')
          return fetch(`https://gigforge.jasim.online/client/payment/project/execute?orderID=${ID}`, {
              method: 'POST',
              headers: {
                  'Content-Type': 'application/json'
              }
          }).then(function(response) {
              return response.json();
          }).then(function(orderData) {
              if (orderData.orderID) {
                  return orderData.orderID;
              } else {
                  throw new Error('Order ID not found');
              }
          }).catch(function(err) {
              console.error('Create Order Error:', err);
              alert('Failed to create order');
          });
      },
      onApprove: function(data, actions) {
        const orderID = getQueryParam('orderID')
        console.log(orderID,data.orderID)
          return fetch(`https://gigforge.jasim.online/client/payment/project/capture?paymentID=${data.orderID}&orderID=${orderID}`, {
              method: 'POST',
              headers: {
                  'Content-Type': 'application/json'
              },
          }).then(function(response) {
              return response.json();
          }).then(function(orderData) {
              alert('Transaction completed by ' + orderData);
              showPaymentSuccessModal();
          }).catch(function(err) {
              console.log(orderID,data.orderID)
              console.error('Capture Order Error:', err);
              alert('Failed to capture order');
          });
      },
      onError: function(err) {
          console.error('An error occurred during the transaction:', err);
          alert('An error occurred during the transaction: ' + err.message);
      }
  }).render('#paypal-button-container');
}

// Show payment success modal
function showPaymentSuccessModal() {
  var modal = document.getElementById("paymentSuccessModal");
  var span = document.getElementsByClassName("close")[0];
  modal.style.display = "block";
  span.onclick = function() {
      modal.style.display = "none";
  }
  window.onclick = function(event) {
      if (event.target == modal) {
          modal.style.display = "none";
      }
  }
}

// Load PayPal SDK and render buttons
function loadPayPalSDK(clientId, merchantIDs) {
  const script = document.createElement('script');
  script.src = `https://www.paypal.com/sdk/js?client-id=${clientId}&merchant-id=${merchantIDs.join(',')}`;
  script.onload = renderPayPalButtons;
  document.head.appendChild(script);
}

// Fetch order details and merchant IDs, then load PayPal SDK
const ID = getQueryParam('orderID')
fetch(`https://gigforge.jasim.online/client/payment/project/execute?orderID=${ID}`, {
  method: 'POST',
  headers: {
      'Content-Type': 'application/json'
  }
}).then(function(response) {
  return response.json();
}).then(function(orderData) {
  if (orderData.orderID && orderData.merchantIDs) {
      loadPayPalSDK('ARh_tJkDPzL6OkIUdjKEMyxg8t_ZKiZ_sYm7Sapv4x9NTsPxjQqKAGyEUcpsyT_7_MdZeYUTM40o7oLl', orderData.merchantIDs);
  } else {
      throw new Error('Order ID or merchant IDs not found');
  }
}).catch(function(err) {
  console.error('Error fetching order details:', err);
  alert('Failed to initialize PayPal payment');
});

// Add event listener to PayPal button for payment success modal
document.getElementById("paypal-button").addEventListener("click", function() {
  showPaymentSuccessModal();
});
