document.addEventListener("DOMContentLoaded", function () {
    var canvas = document.getElementById("renderCanvas");
    //var engine = new BABYLON.Engine(canvas, true);
    var engine = new BABYLON.Engine(canvas, false);

    const createScene = () => {
        const scene = new BABYLON.Scene(engine);

        // カメラとライトの設定
        const camera = new BABYLON.ArcRotateCamera("camera", -Math.PI / 2, Math.PI / 2.5, 10, new BABYLON.Vector3(0, 0, 0));
        camera.attachControl(canvas, true);
        //const light = new BABYLON.HemisphericLight("light", new BABYLON.Vector3(1, 1, 0));
        const light = new BABYLON.HemisphericLight("light", camera.position);

        BABYLON.SceneLoader.ImportMeshAsync("", ".asset/model/", "test.obj").then((result) => {
            result.meshes[0].position.y = 1;

            const material = new BABYLON.StandardMaterial("testmat", scene);
            material.diffuseColor = new BABYLON.Color3(0, 1, 0);
            result.meshes[0].diffuseColor = material;
        })

        // 地面
        const ground = BABYLON.MeshBuilder.CreateGround("ground", { width: 10, height: 10 });

        return scene;
    }

    var scene = createScene();

    engine.runRenderLoop(function () {
        scene.render();
    });

    window.addEventListener("resize", function () {
        engine.resize();
    });
});
