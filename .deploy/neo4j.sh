docker run \
    --detach \
    --publish=7473:7473 --publish=7474:7474 --publish=7687:7687 \
    --volume=$HOME/neo4j/data:/data --volume=$HOME/neo4j/logs:/logs \
    --env=NEO4J_dbms_memory_pagecache_size=4G \
    --name=neo4j \
    neo4j