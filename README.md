# go-api-mapbox

Go: API that handles requests for isochrones from Mapbox API.  The returned JSON is ready for use in LeafletJS.com

- This was written for using in a different project and the functionality is narrow in scope.
- The API returns JSON that make up the verticies of the isochrone (polygon).

__*Deployment:*__ *http://zotact1.ddns.net:8001/v1/mapbox-isochrone/{lng}/{lat}/{time}/{token}*

- __*lng*__ => longitude (decimal degrees)
- __*lat*__ => latitude (decimal degrees)
- __*time*__ => drive time polygon in minutes
- __*token*__ => Mapbox access token

Example API Call & Return Value
-  http://192.168.1.100:8004/v1/mapbox-isochrone/-95.9672927993605/36.1332946325754/1/pk.eyJ1IjoiZ2hvcm5lIiwiYSI6ImNqb2c0bDN4ODAxYzIzdmxieWRpZ2Y2N2oifQ.l3JUrPyN0wEq2Y8WBSjEwQ
-   
{"mapbox":"[[36.138500,-95.967079],[36.137348,-95.966225],[36.135742,-95.965843],[36.135441,-95.964294],[36.134048,-95.963295],[36.133293,-95.961105],[36.133175,-95.962173],[36.132786,-95.962296],[36.132484,-95.963295],[36.132988,-95.963600],[36.133018,-95.964294],[36.132103,-95.965302],[36.131721,-95.966721],[36.129093,-95.967293],[36.129822,-95.968758],[36.131367,-95.969231],[36.131561,-95.970291],[36.132866,-95.971718],[36.133293,-95.975006],[36.133945,-95.972939],[36.135586,-95.970299],[36.135563,-95.969559],[36.138725,-95.967293],[36.138500,-95.967079]]"}