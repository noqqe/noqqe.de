---
title: MongoDB Chunks
date: 2014-04-01T15:49:11.000000
tags: 
- Databases
- MongoDB
---


#### Get Collection stats

~~~ { .json }
mongos> db.testcol.stats()
{
        "sharded" : true,
        "ns" : "devopstest.testcol",
        "count" : 21912105,
        "numExtents" : 33,
        "size" : 876484264,
        "storageSize" : 1487011840,
        "totalIndexSize" : 2015465760,
        "indexSizes" : {
                "_id_" : 991340000,
                "_id_hashed" : 1024125760
        },
        "avgObjSize" : 40.00000292076001,
        "nindexes" : 2,
        "nchunks" : 138,
        "shards" : {
                "rs0" : {
                        "ns" : "devopstest.testcol",
                        "count" : 9719051,
                        "size" : 388762072,
                        "avgObjSize" : 40.00000329250253,
                        "storageSize" : 629612544,
                        "numExtents" : 16,
                        "nindexes" : 2,
                        "lastExtentSize" : 168730624,
                        "paddingFactor" : 1,
                        "systemFlags" : 1,
                        "userFlags" : 0,
                        "totalIndexSize" : 905254896,
                        "indexSizes" : {
                                "_id_" : 450007040,
                                "_id_hashed" : 455247856
                        },
                        "ok" : 1
                },
                "rs1" : {
                        "ns" : "devopstest.testcol",
                        "count" : 12193054,
                        "size" : 487722192,
                        "avgObjSize" : 40.00000262444503,
                        "storageSize" : 857399296,
                        "numExtents" : 17,
                        "nindexes" : 2,
                        "lastExtentSize" : 227786752,
                        "paddingFactor" : 1,
                        "systemFlags" : 1,
                        "userFlags" : 0,
                        "totalIndexSize" : 1110210864,
                        "indexSizes" : {
                                "_id_" : 541332960,
                                "_id_hashed" : 568877904
                        },
                        "ok" : 1
                }
        },
        "ok" : 1
}
~~~


## Get Chunksize

~~~
mongos> use config
mongos> db.settings.findOne({_id:"chunksize"})
 { "_id" : "chunksize", "value" : 64 }
~~~


## Get Sharding Status

~~~ { .json }
mongos> use devopstest
mongos> db.printShardingStatus(true)
 shards:
        {  "_id" : "rs0",  "host" : "rs0/mongo01:27018,mongo02:27018" }
        {  "_id" : "rs1",  "host" : "rs1/mongo03:27018,mongo04:27018" }
  databases:
        {  "_id" : "devopstest",  "partitioned" : true,  "primary" : "rs1" }
                devopstest.testcol
                        shard key: { "_id" : "hashed" }
                        chunks:
                                rs1     71
                                rs0     69
                        { "_id" : { "$minKey" : 1 } } -->> { "_id" : NumberLong("-8851772785412215424") } on : rs1 { "t" : 4, "i" : 0 }
                        { "_id" : NumberLong("-8851772785412215424") } -->> { "_id" : NumberLong("-8585547353746998266") } on : rs1 { "t" : 5, "i" : 0 }
                        { "_id" : NumberLong("-8585547353746998266") } -->> { "_id" : NumberLong("-8213499752793658456") } on : rs1 { "t" : 3, "i" : 0 }
                        { "_id" : NumberLong("-8213499752793658456") } -->> { "_id" : NumberLong("-8213468441668123871") } on : rs1 { "t" : 6, "i" : 0 }
                        { "_id" : NumberLong("-8213468441668123871") } -->> { "_id" : NumberLong("-8213433528679053783") } on : rs1 { "t" : 7, "i" : 0 }
                        { "_id" : NumberLong("-8213433528679053783") } -->> { "_id" : NumberLong("-8213361472744662183") } on : rs1 { "t" : 8, "i" : 0 }
                        { "_id" : NumberLong("-8213361472744662183") } -->> { "_id" : NumberLong("-8213211554694716494") } on : rs1 { "t" : 9, "i" : 0 }
                        { "_id" : NumberLong("-8213211554694716494") } -->> { "_id" : NumberLong("-8213152340671369269") } on : rs1 { "t" : 10, "i" : 0 }
                        { "_id" : NumberLong("-8213152340671369269") } -->> { "_id" : NumberLong("-8213107409837459196") } on : rs1 { "t" : 11, "i" : 0 }
                        { "_id" : NumberLong("-8213107409837459196") } -->> { "_id" : NumberLong("-8212908001549855550") } on : rs1 { "t" : 12, "i" : 0 }
                        { "_id" : NumberLong("-8212908001549855550") } -->> { "_id" : NumberLong("-8212845814068706059") } on : rs1 { "t" : 13, "i" : 0 }
                        { "_id" : NumberLong("-8212845814068706059") } -->> { "_id" : NumberLong("-8212732277560053140") } on : rs1 { "t" : 14, "i" : 0 }
                        { "_id" : NumberLong("-8212732277560053140") } -->> { "_id" : NumberLong("-8212605945588982194") } on : rs1 { "t" : 15, "i" : 0 }
                        { "_id" : NumberLong("-8212605945588982194") } -->> { "_id" : NumberLong("-8212383859917606801") } on : rs1 { "t" : 16, "i" : 0 }
                        { "_id" : NumberLong("-8212383859917606801") } -->> { "_id" : NumberLong("-8212174651659981878") } on : rs1 { "t" : 17, "i" : 0 }
                        { "_id" : NumberLong("-8212174651659981878") } -->> { "_id" : NumberLong("-8211847520929086931") } on : rs1 { "t" : 18, "i" : 0 }
                        { "_id" : NumberLong("-8211847520929086931") } -->> { "_id" : NumberLong("-8211610536291345087") } on : rs1 { "t" : 19, "i" : 0 }
                        { "_id" : NumberLong("-8211610536291345087") } -->> { "_id" : NumberLong("-8211362944455052127") } on : rs1 { "t" : 20, "i" : 0 }
                        { "_id" : NumberLong("-8211362944455052127") } -->> { "_id" : NumberLong("-8211014871721836197") } on : rs1 { "t" : 21, "i" : 0 }
                        { "_id" : NumberLong("-8211014871721836197") } -->> { "_id" : NumberLong("-8210822576025156970") } on : rs1 { "t" : 22, "i" : 0 }
                        { "_id" : NumberLong("-8210822576025156970") } -->> { "_id" : NumberLong("-8210434016863483895") } on : rs1 { "t" : 23, "i" : 0 }
                        { "_id" : NumberLong("-8210434016863483895") } -->> { "_id" : NumberLong("-8210247967607246679") } on : rs1 { "t" : 24, "i" : 0 }
                        { "_id" : NumberLong("-8210247967607246679") } -->> { "_id" : NumberLong("-8210126852444792430") } on : rs1 { "t" : 25, "i" : 0 }
                        { "_id" : NumberLong("-8210126852444792430") } -->> { "_id" : NumberLong("-8209958346534354009") } on : rs1 { "t" : 26, "i" : 0 }
                        { "_id" : NumberLong("-8209958346534354009") } -->> { "_id" : NumberLong("-8209609943232641980") } on : rs1 { "t" : 27, "i" : 0 }
                        { "_id" : NumberLong("-8209609943232641980") } -->> { "_id" : NumberLong("-8209461759548593544") } on : rs1 { "t" : 28, "i" : 0 }
                        { "_id" : NumberLong("-8209461759548593544") } -->> { "_id" : NumberLong("-8209065996260658819") } on : rs1 { "t" : 29, "i" : 0 }
                        { "_id" : NumberLong("-8209065996260658819") } -->> { "_id" : NumberLong("-8208165217921895597") } on : rs1 { "t" : 30, "i" : 0 }
                        { "_id" : NumberLong("-8208165217921895597") } -->> { "_id" : NumberLong("-8207783574190702039") } on : rs1 { "t" : 31, "i" : 0 }
                        { "_id" : NumberLong("-8207783574190702039") } -->> { "_id" : NumberLong("-8207448066833398771") } on : rs1 { "t" : 32, "i" : 0 }
                        { "_id" : NumberLong("-8207448066833398771") } -->> { "_id" : NumberLong("-8206731309841976392") } on : rs1 { "t" : 33, "i" : 0 }
                        { "_id" : NumberLong("-8206731309841976392") } -->> { "_id" : NumberLong("-8206041491479360832") } on : rs1 { "t" : 34, "i" : 0 }
                        { "_id" : NumberLong("-8206041491479360832") } -->> { "_id" : NumberLong("-8205585376560425878") } on : rs1 { "t" : 35, "i" : 0 }
                        { "_id" : NumberLong("-8205585376560425878") } -->> { "_id" : NumberLong("-8204384437960561962") } on : rs1 { "t" : 36, "i" : 0 }
                        { "_id" : NumberLong("-8204384437960561962") } -->> { "_id" : NumberLong("-8203681637219954982") } on : rs1 { "t" : 37, "i" : 0 }
                        { "_id" : NumberLong("-8203681637219954982") } -->> { "_id" : NumberLong("-8201575869779719564") } on : rs1 { "t" : 38, "i" : 0 }
                        { "_id" : NumberLong("-8201575869779719564") } -->> { "_id" : NumberLong("-8199927048117401331") } on : rs1 { "t" : 39, "i" : 0 }
                        { "_id" : NumberLong("-8199927048117401331") } -->> { "_id" : NumberLong("-8199179367860271135") } on : rs1 { "t" : 40, "i" : 0 }
                        { "_id" : NumberLong("-8199179367860271135") } -->> { "_id" : NumberLong("-8198440375446800010") } on : rs1 { "t" : 41, "i" : 0 }
                        { "_id" : NumberLong("-8198440375446800010") } -->> { "_id" : NumberLong("-8197507699937024357") } on : rs1 { "t" : 42, "i" : 0 }
                        { "_id" : NumberLong("-8197507699937024357") } -->> { "_id" : NumberLong("-8196405706797729138") } on : rs1 { "t" : 43, "i" : 0 }
                        { "_id" : NumberLong("-8196405706797729138") } -->> { "_id" : NumberLong("-8195263906662049560") } on : rs1 { "t" : 44, "i" : 0 }
                        { "_id" : NumberLong("-8195263906662049560") } -->> { "_id" : NumberLong("-8194569637200014091") } on : rs1 { "t" : 45, "i" : 0 }
                        { "_id" : NumberLong("-8194569637200014091") } -->> { "_id" : NumberLong("-8193579967613594852") } on : rs1 { "t" : 46, "i" : 0 }
                        { "_id" : NumberLong("-8193579967613594852") } -->> { "_id" : NumberLong("-8192774609495050364") } on : rs1 { "t" : 47, "i" : 0 }
                        { "_id" : NumberLong("-8192774609495050364") } -->> { "_id" : NumberLong("-8191932736542176892") } on : rs1 { "t" : 48, "i" : 0 }
                        { "_id" : NumberLong("-8191932736542176892") } -->> { "_id" : NumberLong("-8191352439097470940") } on : rs1 { "t" : 49, "i" : 0 }
                        { "_id" : NumberLong("-8191352439097470940") } -->> { "_id" : NumberLong("-8190600814987176771") } on : rs1 { "t" : 50, "i" : 0 }
                        { "_id" : NumberLong("-8190600814987176771") } -->> { "_id" : NumberLong("-8189714575393337195") } on : rs1 { "t" : 51, "i" : 0 }
                        { "_id" : NumberLong("-8189714575393337195") } -->> { "_id" : NumberLong("-8188738026158082508") } on : rs1 { "t" : 52, "i" : 0 }
                        { "_id" : NumberLong("-8188738026158082508") } -->> { "_id" : NumberLong("-8187789726302485266") } on : rs1 { "t" : 53, "i" : 0 }
                        { "_id" : NumberLong("-8187789726302485266") } -->> { "_id" : NumberLong("-8186907339703538438") } on : rs1 { "t" : 54, "i" : 0 }
                        { "_id" : NumberLong("-8186907339703538438") } -->> { "_id" : NumberLong("-8186099173629188271") } on : rs1 { "t" : 55, "i" : 0 }
                        { "_id" : NumberLong("-8186099173629188271") } -->> { "_id" : NumberLong("-8185347640506990752") } on : rs1 { "t" : 56, "i" : 0 }
                        { "_id" : NumberLong("-8185347640506990752") } -->> { "_id" : NumberLong("-8184160180263779109") } on : rs1 { "t" : 57, "i" : 0 }
                        { "_id" : NumberLong("-8184160180263779109") } -->> { "_id" : NumberLong("-8183174617056500354") } on : rs1 { "t" : 58, "i" : 0 }
                        { "_id" : NumberLong("-8183174617056500354") } -->> { "_id" : NumberLong("-8182149563950255752") } on : rs1 { "t" : 59, "i" : 0 }
                        { "_id" : NumberLong("-8182149563950255752") } -->> { "_id" : NumberLong("-8181408316341416659") } on : rs1 { "t" : 60, "i" : 0 }
                        { "_id" : NumberLong("-8181408316341416659") } -->> { "_id" : NumberLong("-8179791006004615832") } on : rs1 { "t" : 61, "i" : 0 }
                        { "_id" : NumberLong("-8179791006004615832") } -->> { "_id" : NumberLong("-8179080962204393165") } on : rs0 { "t" : 61, "i" : 1 }
                        { "_id" : NumberLong("-8179080962204393165") } -->> { "_id" : NumberLong("-8178351216033744109") } on : rs0 { "t" : 2, "i" : 157 }
                        { "_id" : NumberLong("-8178351216033744109") } -->> { "_id" : NumberLong("-8177257178546793446") } on : rs0 { "t" : 2, "i" : 155 }
                        { "_id" : NumberLong("-8177257178546793446") } -->> { "_id" : NumberLong("-8176171925315762331") } on : rs0 { "t" : 2, "i" : 153 }
                        { "_id" : NumberLong("-8176171925315762331") } -->> { "_id" : NumberLong("-8175203055874879048") } on : rs0 { "t" : 2, "i" : 151 }
                        { "_id" : NumberLong("-8175203055874879048") } -->> { "_id" : NumberLong("-8173770360363750259") } on : rs0 { "t" : 2, "i" : 149 }
                        { "_id" : NumberLong("-8173770360363750259") } -->> { "_id" : NumberLong("-8172543986388985517") } on : rs0 { "t" : 2, "i" : 147 }
                        { "_id" : NumberLong("-8172543986388985517") } -->> { "_id" : NumberLong("-8171133082073731067") } on : rs0 { "t" : 2, "i" : 145 }
                        { "_id" : NumberLong("-8171133082073731067") } -->> { "_id" : NumberLong("-8170218965132029401") } on : rs0 { "t" : 2, "i" : 143 }
                        { "_id" : NumberLong("-8170218965132029401") } -->> { "_id" : NumberLong("-8169107854523839443") } on : rs0 { "t" : 2, "i" : 141 }
                        { "_id" : NumberLong("-8169107854523839443") } -->> { "_id" : NumberLong("-8168169358327006396") } on : rs0 { "t" : 2, "i" : 139 }
                        { "_id" : NumberLong("-8168169358327006396") } -->> { "_id" : NumberLong("-8167549901058193571") } on : rs0 { "t" : 2, "i" : 137 }
                        { "_id" : NumberLong("-8167549901058193571") } -->> { "_id" : NumberLong("-8166554509378262468") } on : rs0 { "t" : 2, "i" : 135 }
                        { "_id" : NumberLong("-8166554509378262468") } -->> { "_id" : NumberLong("-8165695875757792424") } on : rs0 { "t" : 2, "i" : 133 }
                        { "_id" : NumberLong("-8165695875757792424") } -->> { "_id" : NumberLong("-8164923866884348840") } on : rs0 { "t" : 2, "i" : 131 }
                        { "_id" : NumberLong("-8164923866884348840") } -->> { "_id" : NumberLong("-8164006361819083153") } on : rs0 { "t" : 2, "i" : 129 }
                        { "_id" : NumberLong("-8164006361819083153") } -->> { "_id" : NumberLong("-8162847861179552008") } on : rs0 { "t" : 2, "i" : 127 }
                        { "_id" : NumberLong("-8162847861179552008") } -->> { "_id" : NumberLong("-8162029678861577158") } on : rs0 { "t" : 2, "i" : 125 }
                        { "_id" : NumberLong("-8162029678861577158") } -->> { "_id" : NumberLong("-8161011359381462169") } on : rs0 { "t" : 2, "i" : 123 }
                        { "_id" : NumberLong("-8161011359381462169") } -->> { "_id" : NumberLong("-8159997280817650411") } on : rs0 { "t" : 2, "i" : 121 }
                        { "_id" : NumberLong("-8159997280817650411") } -->> { "_id" : NumberLong("-8158947738060147705") } on : rs0 { "t" : 2, "i" : 119 }
                        { "_id" : NumberLong("-8158947738060147705") } -->> { "_id" : NumberLong("-8158078439711620257") } on : rs0 { "t" : 2, "i" : 117 }
                        { "_id" : NumberLong("-8158078439711620257") } -->> { "_id" : NumberLong("-8156984808555982340") } on : rs0 { "t" : 2, "i" : 115 }
                        { "_id" : NumberLong("-8156984808555982340") } -->> { "_id" : NumberLong("-8155661451276350639") } on : rs0 { "t" : 2, "i" : 113 }
                        { "_id" : NumberLong("-8155661451276350639") } -->> { "_id" : NumberLong("-8154335012322561593") } on : rs0 { "t" : 2, "i" : 111 }
                        { "_id" : NumberLong("-8154335012322561593") } -->> { "_id" : NumberLong("-8153176132279541756") } on : rs0 { "t" : 2, "i" : 109 }
                        { "_id" : NumberLong("-8153176132279541756") } -->> { "_id" : NumberLong("-8152242903205065312") } on : rs0 { "t" : 2, "i" : 107 }
                        { "_id" : NumberLong("-8152242903205065312") } -->> { "_id" : NumberLong("-8151102798908523359") } on : rs0 { "t" : 2, "i" : 105 }
                        { "_id" : NumberLong("-8151102798908523359") } -->> { "_id" : NumberLong("-8149619516359624681") } on : rs0 { "t" : 2, "i" : 103 }
                        { "_id" : NumberLong("-8149619516359624681") } -->> { "_id" : NumberLong("-8148409976458972260") } on : rs0 { "t" : 2, "i" : 101 }
                        { "_id" : NumberLong("-8148409976458972260") } -->> { "_id" : NumberLong("-8147332794765245160") } on : rs0 { "t" : 2, "i" : 99 }
                        { "_id" : NumberLong("-8147332794765245160") } -->> { "_id" : NumberLong("-8146439479819920097") } on : rs0 { "t" : 2, "i" : 97 }
                        { "_id" : NumberLong("-8146439479819920097") } -->> { "_id" : NumberLong("-8145320450725583213") } on : rs0 { "t" : 2, "i" : 95 }
                        { "_id" : NumberLong("-8145320450725583213") } -->> { "_id" : NumberLong("-8144535012828570679") } on : rs0 { "t" : 2, "i" : 93 }
                        { "_id" : NumberLong("-8144535012828570679") } -->> { "_id" : NumberLong("-8143442520736185697") } on : rs0 { "t" : 2, "i" : 91 }
                        { "_id" : NumberLong("-8143442520736185697") } -->> { "_id" : NumberLong("-8142356725721430638") } on : rs0 { "t" : 2, "i" : 89 }
                        { "_id" : NumberLong("-8142356725721430638") } -->> { "_id" : NumberLong("-8141342217679373106") } on : rs0 { "t" : 2, "i" : 87 }
                        { "_id" : NumberLong("-8141342217679373106") } -->> { "_id" : NumberLong("-8140183809247704998") } on : rs0 { "t" : 2, "i" : 85 }
                        { "_id" : NumberLong("-8140183809247704998") } -->> { "_id" : NumberLong("-8138915441096847421") } on : rs0 { "t" : 2, "i" : 83 }
                        { "_id" : NumberLong("-8138915441096847421") } -->> { "_id" : NumberLong("-8137693969213943867") } on : rs0 { "t" : 2, "i" : 81 }
                        { "_id" : NumberLong("-8137693969213943867") } -->> { "_id" : NumberLong("-8136661189027297871") } on : rs0 { "t" : 2, "i" : 79 }
                        { "_id" : NumberLong("-8136661189027297871") } -->> { "_id" : NumberLong("-8135839293705164909") } on : rs0 { "t" : 2, "i" : 77 }
                        { "_id" : NumberLong("-8135839293705164909") } -->> { "_id" : NumberLong("-8134452081567553892") } on : rs0 { "t" : 2, "i" : 75 }
                        { "_id" : NumberLong("-8134452081567553892") } -->> { "_id" : NumberLong("-8132751197059767359") } on : rs0 { "t" : 2, "i" : 73 }
                        { "_id" : NumberLong("-8132751197059767359") } -->> { "_id" : NumberLong("-8131677445313270180") } on : rs0 { "t" : 2, "i" : 71 }
                        { "_id" : NumberLong("-8131677445313270180") } -->> { "_id" : NumberLong("-8130496714404671629") } on : rs0 { "t" : 2, "i" : 69 }
                        { "_id" : NumberLong("-8130496714404671629") } -->> { "_id" : NumberLong("-8129398416188061155") } on : rs0 { "t" : 2, "i" : 67 }
                        { "_id" : NumberLong("-8129398416188061155") } -->> { "_id" : NumberLong("-8128119220798110768") } on : rs0 { "t" : 2, "i" : 65 }
                        { "_id" : NumberLong("-8128119220798110768") } -->> { "_id" : NumberLong("-8126929551665975026") } on : rs0 { "t" : 2, "i" : 63 }
                        { "_id" : NumberLong("-8126929551665975026") } -->> { "_id" : NumberLong("-8125920269040853133") } on : rs0 { "t" : 2, "i" : 61 }
                        { "_id" : NumberLong("-8125920269040853133") } -->> { "_id" : NumberLong("-8124038628197504050") } on : rs0 { "t" : 2, "i" : 59 }
                        { "_id" : NumberLong("-8124038628197504050") } -->> { "_id" : NumberLong("-8122899172829813362") } on : rs0 { "t" : 2, "i" : 57 }
                        { "_id" : NumberLong("-8122899172829813362") } -->> { "_id" : NumberLong("-8121274148563724548") } on : rs0 { "t" : 2, "i" : 55 }
                        { "_id" : NumberLong("-8121274148563724548") } -->> { "_id" : NumberLong("-8119901074732153159") } on : rs0 { "t" : 2, "i" : 53 }
                        { "_id" : NumberLong("-8119901074732153159") } -->> { "_id" : NumberLong("-8118154521591886705") } on : rs0 { "t" : 2, "i" : 51 }
                        { "_id" : NumberLong("-8118154521591886705") } -->> { "_id" : NumberLong("-8116993711800984043") } on : rs0 { "t" : 2, "i" : 49 }
                        { "_id" : NumberLong("-8116993711800984043") } -->> { "_id" : NumberLong("-8115837489138121826") } on : rs0 { "t" : 2, "i" : 47 }
                        { "_id" : NumberLong("-8115837489138121826") } -->> { "_id" : NumberLong("-7955679762482393067") } on : rs0 { "t" : 2, "i" : 45 }
                        { "_id" : NumberLong("-7955679762482393067") } -->> { "_id" : NumberLong("-7088276312959863537") } on : rs0 { "t" : 2, "i" : 41 }
                        { "_id" : NumberLong("-7088276312959863537") } -->> { "_id" : NumberLong("-6457398946748347364") } on : rs0 { "t" : 2, "i" : 32 } jumbo
                        { "_id" : NumberLong("-6457398946748347364") } -->> { "_id" : NumberLong("-6453893125579898133") } on : rs0 { "t" : 2, "i" : 33 }
                        { "_id" : NumberLong("-6453893125579898133") } -->> { "_id" : NumberLong("-5826014850355698880") } on : rs0 { "t" : 2, "i" : 36 }
                        { "_id" : NumberLong("-5826014850355698880") } -->> { "_id" : NumberLong("-4611686018427387902") } on : rs0 { "t" : 2, "i" : 37 }
                        { "_id" : NumberLong("-4611686018427387902") } -->> { "_id" : NumberLong("-4076749157618855140") } on : rs0 { "t" : 2, "i" : 14 }
                        { "_id" : NumberLong("-4076749157618855140") } -->> { "_id" : NumberLong("-3540258727155571787") } on : rs0 { "t" : 2, "i" : 16 }
                        { "_id" : NumberLong("-3540258727155571787") } -->> { "_id" : NumberLong("-2304232382838626787") } on : rs0 { "t" : 2, "i" : 17 }
                        { "_id" : NumberLong("-2304232382838626787") } -->> { "_id" : NumberLong("-1670689140503435181") } on : rs0 { "t" : 2, "i" : 22 }
                        { "_id" : NumberLong("-1670689140503435181") } -->> { "_id" : NumberLong("-1038957101455629159") } on : rs0 { "t" : 2, "i" : 30 }
                        { "_id" : NumberLong("-1038957101455629159") } -->> { "_id" : NumberLong(0) } on : rs0 { "t" : 2, "i" : 31 }
                        { "_id" : NumberLong(0) } -->> { "_id" : NumberLong("269474596422786706") } on : rs1 { "t" : 61, "i" : 4 }
                        { "_id" : NumberLong("269474596422786706") } -->> { "_id" : NumberLong("2302928180697545647") } on : rs1 { "t" : 61, "i" : 5 }
                        { "_id" : NumberLong("2302928180697545647") } -->> { "_id" : NumberLong("2942556921558436718") } on : rs1 { "t" : 2, "i" : 18 }
                        { "_id" : NumberLong("2942556921558436718") } -->> { "_id" : NumberLong("3572026050869470840") } on : rs1 { "t" : 2, "i" : 34 }
                        { "_id" : NumberLong("3572026050869470840") } -->> { "_id" : NumberLong("4611686018427387902") } on : rs1 { "t" : 2, "i" : 35 }
                        { "_id" : NumberLong("4611686018427387902") } -->> { "_id" : NumberLong("5245700550699935301") } on : rs1 { "t" : 2, "i" : 28 }
                        { "_id" : NumberLong("5245700550699935301") } -->> { "_id" : NumberLong("5874770515804143613") } on : rs1 { "t" : 2, "i" : 42 }
                        { "_id" : NumberLong("5874770515804143613") } -->> { "_id" : NumberLong("6856054966306647663") } on : rs1 { "t" : 2, "i" : 43 }
                        { "_id" : NumberLong("6856054966306647663") } -->> { "_id" : NumberLong("7489830441767112879") } on : rs1 { "t" : 2, "i" : 26 }
                        { "_id" : NumberLong("7489830441767112879") } -->> { "_id" : NumberLong("8116920262240846262") } on : rs1 { "t" : 2, "i" : 38 }
                        { "_id" : NumberLong("8116920262240846262") } -->> { "_id" : NumberLong("8388579422741846231") } on : rs1 { "t" : 61, "i" : 2 }
                        { "_id" : NumberLong("8388579422741846231") } -->> { "_id" : { "$maxKey" : 1 } } on : rs1 { "t" : 61, "i" : 3 }
~~~

## Jumbo Chunks

Jumbo Chunks enstehen wenn zuviele Daten den gleichen Shardkey haben.
Jumbochunks können auch nicht mehr migriert werden und müssen wahlweise
gesplittet werden.
