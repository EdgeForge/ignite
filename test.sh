export PATH=$PWD/bin:$PATH

VMNAME=${VMNAME:-my-vm-v3}

ignite run weaveworks/ignite-ubuntu \
    --cpus 2 \
    --memory 1GB \
    --ssh \
    --name ${VMNAME}

