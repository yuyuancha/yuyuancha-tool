<template>
  <h3>Three.js 練習</h3>
  <div id="canvas" style="width: 100%; height: 500px;"></div>
</template>

<script>
import * as THREE from "three";
import {onMounted} from "vue";
import {OrbitControls} from "three/addons";

export default {
    name: "ThreeJsPractice",
    setup() {
        onMounted(() => {
            const scene = new THREE.Scene();
            const camera = new THREE.PerspectiveCamera( 75, window.innerWidth / window.innerHeight, 0.1, 1000 );

            const renderer = new THREE.WebGLRenderer();
            renderer.setSize( window.innerWidth, window.innerHeight );
            document.body.appendChild( renderer.domElement );

            const geometry = new THREE.BoxGeometry( 1, 1, 1 );
            const material = new THREE.MeshBasicMaterial( { color: 0x00ff00 } );
            const cube = new THREE.Mesh( geometry, material );
            scene.add( cube );

            camera.position.z = 5;

            function animate() {
                requestAnimationFrame( animate );

                cube.rotation.x += 0.01;
                cube.rotation.y += 0.01;

                renderer.render( scene, camera );
            }

            animate();

            renderer.shadowMap.enabled = true;
            renderer.shadowMap.type = THREE.PCFSoftShadowMap; // default THREE.PCFShadowMap

//Create a PointLight and turn on shadows for the light
            const light = new THREE.PointLight( 0xffffff, 1, 100 );
            light.position.set( 0, 10, 4 );
            light.castShadow = true; // default false
            scene.add( light );

//Set up shadow properties for the light
            light.shadow.mapSize.width = 512; // default
            light.shadow.mapSize.height = 512; // default
            light.shadow.camera.near = 0.5; // default
            light.shadow.camera.far = 500 // default

//Create a sphere that cast shadows (but does not receive them)
            const sphereGeometry = new THREE.SphereGeometry( 5, 32, 32 );
            const sphereMaterial = new THREE.MeshStandardMaterial( { color: 0xff0000 } );
            const sphere = new THREE.Mesh( sphereGeometry, sphereMaterial );
            sphere.castShadow = true; //default is false
            sphere.receiveShadow = false; //default
            scene.add( sphere );

//Create a plane that receives shadows (but does not cast them)
            const planeGeometry = new THREE.PlaneGeometry( 20, 20, 32, 32 );
            const planeMaterial = new THREE.MeshStandardMaterial( { color: 0x00ff00 } )
            const plane = new THREE.Mesh( planeGeometry, planeMaterial );
            plane.receiveShadow = true;
            scene.add( plane );

//Create a helper for the shadow camera (optional)
            const helper = new THREE.CameraHelper( light.shadow.camera );
            scene.add( helper );
        });


        return {

        }
    }
}
</script>

<style scoped>

</style>