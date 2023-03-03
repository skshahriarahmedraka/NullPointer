
// function fetchWithAutoRetry(fetcher, maxRetryCount:number) {
//     return new Promise((resolve, reject) => {
//       let retries = 0;
//       const caller = () =>
//         fetcher().then((data) => {
//             resolve(data);
//           })
//           .catch((error) => {
//             if (retries < maxRetryCount) {
//               retries++;
//               caller();
//             } else {
//               reject(error);
//             }
//           });
//       retries = 1;
//       caller();
//     });
//   }
  
//   const fetchSouravProfile = async () => {
//     console.log("Fetching..");
//     const rawResponse = await fetch("https://api.github.com/users/sourav-singhh");
//     const jsonResponse = await rawResponse.json();
//     console.log(jsonResponse);
//     return jsonResponse;
//   };
  
//   fetchWithAutoRetry(fetchSouravProfile, 5);