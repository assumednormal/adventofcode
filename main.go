package main

import (
	"fmt"

	"adventofcode2017/advent2017"
)

var (
	input01 = "5994521226795838486188872189952551475352929145357284983463678944777228139398117649129843853837124228353689551178129353548331779783742915361343229141538334688254819714813664439268791978215553677772838853328835345484711229767477729948473391228776486456686265114875686536926498634495695692252159373971631543594656954494117149294648876661157534851938933954787612146436571183144494679952452325989212481219139686138139314915852774628718443532415524776642877131763359413822986619312862889689472397776968662148753187767793762654133429349515324333877787925465541588584988827136676376128887819161672467142579261995482731878979284573246533688835226352691122169847832943513758924194232345988726741789247379184319782387757613138742817826316376233443521857881678228694863681971445442663251423184177628977899963919997529468354953548612966699526718649132789922584524556697715133163376463256225181833257692821331665532681288216949451276844419154245423434141834913951854551253339785533395949815115622811565999252555234944554473912359674379862182425695187593452363724591541992766651311175217218144998691121856882973825162368564156726989939993412963536831593196997676992942673571336164535927371229823236937293782396318237879715612956317715187757397815346635454412183198642637577528632393813964514681344162814122588795865169788121655353319233798811796765852443424783552419541481132132344487835757888468196543736833342945718867855493422435511348343711311624399744482832385998592864795271972577548584967433917322296752992127719964453376414665576196829945664941856493768794911984537445227285657716317974649417586528395488789946689914972732288276665356179889783557481819454699354317555417691494844812852232551189751386484638428296871436139489616192954267794441256929783839652519285835238736142997245189363849356454645663151314124885661919451447628964996797247781196891787171648169427894282768776275689124191811751135567692313571663637214298625367655969575699851121381872872875774999172839521617845847358966264291175387374464425566514426499166813392768677233356646752273398541814142523651415521363267414564886379863699323887278761615927993953372779567675"
	input02 = input01
	input03 = `4347	3350	196	162	233	4932	4419	3485	4509	4287	4433	4033	207	3682	2193	4223
648	94	778	957	1634	2885	1964	2929	2754	89	972	112	80	2819	543	2820
400	133	1010	918	1154	1008	126	150	1118	117	148	463	141	940	1101	89
596	527	224	382	511	565	284	121	643	139	625	335	657	134	125	152
2069	1183	233	213	2192	193	2222	2130	2073	2262	1969	2159	2149	410	181	1924
1610	128	1021	511	740	1384	459	224	183	266	152	1845	1423	230	1500	1381
5454	3936	250	5125	244	720	5059	202	4877	5186	313	6125	172	727	1982	748
3390	3440	220	228	195	4525	1759	1865	1483	5174	4897	4511	5663	4976	3850	199
130	1733	238	1123	231	1347	241	291	1389	1392	269	1687	1359	1694	1629	1750
1590	1394	101	434	1196	623	1033	78	890	1413	74	1274	1512	1043	1103	84
203	236	3001	1664	195	4616	2466	4875	986	1681	152	3788	541	4447	4063	5366
216	4134	255	235	1894	5454	1529	4735	4962	220	2011	2475	185	5060	4676	4089
224	253	19	546	1134	3666	3532	207	210	3947	2290	3573	3808	1494	4308	4372
134	130	2236	118	142	2350	3007	2495	2813	2833	2576	2704	169	2666	2267	850
401	151	309	961	124	1027	1084	389	1150	166	1057	137	932	669	590	188
784	232	363	316	336	666	711	430	192	867	628	57	222	575	622	234`
	input04 = input03
	input05 = 325489
	input06 = input05
	input07 = `bdwdjjo avricm cjbmj ran lmfsom ivsof
mxonybc fndyzzi gmdp gdfyoi inrvhr kpuueel wdpga vkq
bneh ylltsc vhryov lsd hmruxy ebnh pdln vdprrky
fumay zbccai qymavw zwoove hqpd rcxyvy
bcuo khhkkro mpt dxrebym qwum zqp lhmbma esmr qiyomu
qjs giedut mzsubkn rcbugk voxk yrlp rqxfvz kspz vxg zskp
srceh xdwao reshc shecr
dcao isz wwse bbdgn ewsw qkze pwu
lbnvl lviftmr zqiv iadanl fdhrldn dlaani lxy dhfndrl fkoukx
raovmz pdysjsw hqps orsyqw rrwnzcz vrzoam jjljt
wgt gzi icpwp qeefgbe msadakj jbbrma sbj dufuujx zex
cfzx bvyu eswr hafcfy klw bgnhynv qrf aop
rzlq atrzcpb hpl pajjw cdxep ean aptzcrb rzcbapt
xogpf ucc nsukoz umtfbw xfvth ozusnk fopxg ubp iflb
xot nqcdyu kpwix szo cyxv hpmp hwtrc zso nyuqdc aha
mkzf cat tkjprc izxdggf obspan lmlbg bsyspf twox
lfmfrd ooclx tcl clt
dxvnyd nxwojj arutn eyqocj swzao tmh juvpezm
teu eman rlmdmk xkbodv fvrcm zorgy wmwe
hmo fdayx duciqf cgt duciqf
imjnv vfmsha cyrusow xjswoq nclrmjy sjxowq ynjrcml
rwbsay alsi bmzpvw ozq aduui nihwx glwdiz ixmkgfx
vtjzc ntkh zekj qrbkjhn zekj lyfnbg
afaig jqhli oie lhwyduh kqfnraz nfrzaqk mayfg iljqh
inb zum zmu dnl zjxg vrdziq ypdnsvt
uhbzmre mpdxm alkbmsq aopjmkl mqxenry ayvkrf zxvs qkfqva
fimjr ccv cnug crdsv
bqyve lhxdj ydu qbyve vihwjr vyodhc
vmng dyttyf noagpji tdtyfy ewiest ogg
kgscfj klmsv vmksl vmlks
qlvh veo wruft wtm fbrlvjr evo wvwmny dhp bvrlrfj lvt vgzuyyw
mxuro orxmu tivu tjdq ojjvju cdd
kjexme gxbysxp yxrum hinrklv fxehytd qkt tqk umryx nim
kywnux wab tzrft dsaz jgwuw dubarmi fle wjguvr umjp uzncwj mzz
qokwh zrda xywufk tbxhhj eejqaoa hwoqk zer hwt hbjxth xyf hmh
eregs qdx tdequa agrlrg mwwpba qjie yrjvhr txujk
iyot fxwdcb zvwfv vfzwv wvkw ncwbr wdejrr ltcdza
waix eza znmonya ldfghws ialwfvc dey ubsz uhbnh svgekg nonzyam
bryz tfbqn xznfmw xiplgww wwxigpl jxzcgxl rzyb
cqvl rrcoqxs staeuqr hzzow cwv tsvol dio coc ddavii uuojy
txbn qvkkyh gbqnjtq ndpkqr srt bkpqfmm ytycev ypcv bpqmmfk
uqkjmul dour zgq ztango yrrjhrg ufxnmuw
ekxbcv vkxbec xbcevk jiq bar
wczff qdu cwffz hhk wlvyg
zjlconc osgsro dajzo hqih ehml
hnio shccluw cpu ivaby tormn vkef abv vkef ivaby
xgbdeso xiizs omqwy sbtnnt khago evviclw xyu dtvg wsyxfuc humewp
cnzu bia vdyqrf wwb qveum hmh ouupgc owli
pjpmfxa dvd lxgh ndy gwph oebfkqv vtlxdg efl ekj dyn
mvan nmdkc ucyshia mavn ecst poo
oybm pjwm bmyo wovgu xykziaq obmy eiirhqd
xkvomx yxvv oxxpth elh vxvy lhe ycn
okxglw gmaangx gnxaamg yduzrr nzwxtnd rcxcu xjjvno yat cin gaxnamg yss
oicgs rrol zvnbna rrol
abb edpnxuo peoudxn bab ceay
ncpkfz gvwunb fckpzn caafx pkcfzn tsfl
fnrt ymenkpq wodubcm niv nvi ziluu cuowbdm zocg pdakwt mlzxkex nuxqclo
uouxcgl stgua otadr ideannq wizxunv iqsdpj mxddt ldst ucxogul
rbrwyhk wqoz zqwo ikwgexl atpu iza
smo yolp pcahlu muljxkq cbkljmz zlbcmkj zvbmgz eaiv ncv zplifm yplo
ocutdhz zmnaap llgv llzpl loavju guzkfq saay rxyhng cwxzx lcv anrnzs
etyzx tcm upxrtvd imyoiu rdpj fed dmm
gonqa szteh szteh razdqh phyff upf knfqfaf knfqfaf fpsgl kakag
mcju mixskop isrwat lcr nfyi lcr aaevr nfyi pqrbk gnful
xfmr fkmnq fbnhd mxrf litniid xbae frxm zcenf
yuh lzojtj rqsh hyu
vbjgql yeshsuv lokt efqota wpwjfu ykyq rxc fxxh ycqfkk gndts vdf
wnylmr kkuruxm azr xukrkum dohkwx dmdb
bjiyrwf dvf fdv vdf gnokekr
jsaq hcww iayqtu llv gdpxdrd hwlo nosjit wpm lcab fcgwr
fxjp bys nnf xzllckh bys hvojw zcwtgwz wye ccyvjv
grafa hbb ghk wkdpsf ufa uoqmysd
yvacf kssbff iovrm mvrio cfbpb avh zzje
gqd qmsen wkvrfz vhtsa zrwfkv gul zkvwrf
hrbi svaogb aogsvb bgrk hibr jbtkr
ljl ryc mrewrge yky
fcqyymt emk qcmyytf mcfvusb luy qany cbsvumf
oknt mcozuc ccmuoz uoccmz
uziil xobutwf acnof iqgwl diso
sekq fxbtsuv ddnnqg rnemlt dngnqd hhgjfus stxvubf
lajcp qgiw fyz blrlcd pviwv buh wnkk
wolqfk nvpapfc rwcqxfz xobno yzjfz irpj wolqfk wbnwjt
vmabj noiljne hhqf holxkbk swwzx ylgj lnmxy lqodhk abjvm bmsrf
bpnp yrz pjepxxs jlmhuy vihlx zacm inuazhc xsxjepp
tryl kryh eonvaad ucevssk umkxg lqej nswopjj svkeucs bmh stosxxz
cfdwd dmfdrvm ibuhsz nwtgmb pjt dmfdrvm cqdcm fzjjz afa ibuhsz
erwp abn jwx ynmkkj rhgg abn epd atqhs rst rhgg
jtnp cegdsoy gfuvfbg gdmn ahlsc
jgrp diu jrgp onjnml nojmnl vxockc
lakqyuw khq dcpiwt ykwlqua hkq plklx ujbckec hjcvur jnp pvyf
usuvoo jkih ylafyy yhio jureyj
uazisdf cnwlfnf ewodatr woaddkd wbla qmn atdrowe
bnyepaa ntqh xppe ydtsw ppex
yewjwsp jxylmtk ijese ewry ijese kbja nfml zeuwcsh juimz
qbvmf nca zsfreo uurgaiz twe fbqmv ncwi etdcsk atowfp
jeotslx kgdpzp wxlcww pdd dcn ddp
macllv ldl kyluine lbt hbxbr wxcaspp ezwvc qxkeu
ivg gxv zsf ucr uff yrz
tdlwbny bqlrlz tbynwdl lwtbdny
tnekq pdaievs ttwpfh xfm fcaa
zqqhl zbf fbz uqrv bfz ffwavhk foccg
vcw ebqdd cwv eddbq nrmq
hpiusz sizphu xzq sgyehk wgagkv hsygek
vagkxa iou frqdnnr ipcg uxvh vvh eskf katgpiq aqktigp gzvseyi
xkwgd kzfxk pgdy fmtvq ngf rshx zti pamviob ely knz
hwo rteohu qzwoe rotuhe wzb
bsqgg tid dti gtats dit
sjtux djwxv dljwjq xwvjd xnqfvx veqdrtl uxtsj nnkjn wnhilaf unirrp
fruuqjk gtote gooklg bzwhim zfnccmm ezipnf cxwdxa wfu fdca
zcyxb byzxc cxbyz pgcqco ivlxz
wrjh zfdinsf ihw xwosiah hdg xpiabno bilyy azdeczg javuwa
rinlv dcpt qhencba mmb njxw gadc
qwcpua qzyzt cxjsgh kumh byiimas qhsgf qytzz rqqruwp ismyiba xydcxz rwkscqa
xbzefi hltca ibzxfe fkx xizbfe wvaynts
oyuce vzk ouxvj gfh efgbv ubc nyb bxnbhd mtwboe whksy ovmrt
ljrebp tacn bpjler utphw wmfw rcnha
drdnic eyodes rcnidd yseeod
umxmsf kfroz ukhjin awpnnnu ooyyohh tuv rafano jze
bakz lfzpjyg gfkqcgn kzh zwpvk gqfngck
jpaony ojpnya hmro xaaz tovary aaxz iel pbg
swvbgc bbhjp yvrcddd rhj clfu eao afrkegn qvvb yvcx nxjmdo rcvtx
conbjy jeqtri wvujt jeqtri rkhllgw tsdt zowreo qxr qbpragn kuzmplw wvujt
jrpxyp hchljy rkowqb eeaf ltllebb gtksrwx iazx vnsfmc zzrxw hlcjyh
piehb cjdzt eqn kuje rls oaewoz lrqwt lcrrq
hdjowxv uknhlv hluknv pokxg
txiqxfr fyyp pyyf xfxtrqi tvm rtvby cfx trx nwrf kqrxtat alwot
wdaadr stexpow ardawd uejqxc
wwgwjel wwgwjel mtjt wwgwjel
mczx uua lgceb dqru vkcea tcet ruz
jkt yroojr qdrtdu wze ovwz fdmqnr xxsyfd kchytwl hctlkwy gyd
eif irnrce iamhxgh bmis uxye azrwdi sznv yuowb vdlqqxu
dxdjyj hngqwzs yhwku qhsctfe rhbc rchb tqhcfse
fxyxnzs qtxevin rvtxtc iqnxtve
zgbpk mwzxx bgpkz wkpkn
rjiym iub lcyw agbtlb bzhx osv rbtf
emmyu uoflio tinih skpqaj rbor gezbhhv ine mij qlqte uuj ycns
owmwc uhxv pyho ftjh jzsg blqn bszyo bob trbycy mkru
mwgz bbqsmpp fgzs bihhg bbn pjxxexs qrqmt htsxfwo qltqp vqqaxi
lpr wcvy sxjqq ltd rftdgv pdnog ymu
qhcos shuy icdhucu lrikh rwslv yxbgibl rcomhn wakirz
civdmee owlzocl vedecim rogmjnn pix pohcmk dsjm yworm
vzdpxp lvt inufv yofqt omm qfoty qrlseqy amkt kjcvg vgkjc
huhq quhh levzsws sjuun ofgqr cjhp nfxbbft rnt wtbd tbzab
tjftkx xpfcv hvftvhw lpypbjg batrn fhwhtvv uthl arbtn brb sthv
ogr uyuxdco bpjgir edztxv sxtgu jzfmx ihnauz zwegqkr kvkw
mhxthf pervvn gshy jig ezjteq ckkcpy gww
tiljyki rpe prcojy tjkylii moxu
pjsdqc lgqydfd lohck emrtejw axwmo wuuv rfi qzyncmw gjijdfb bljfd xrs
ywjab gynzi relf kziy xmsby izyk ocwoho kqnyh bwayj
bhjlz uonz jhmzuq eiajoos zjnbj tomj bmyv hjlbz fgw jjbnz
kszz xzw xzw prtznyb
ghzk vxhwt thxwv slwpayp qxegmi dawdwo kgzh
ibpcvuf wnuwxu sbf jsj bfjynl cdp jbylnf
epaxr vfhf hvff azepadz pwf sbo pgfzya hslyo rqqj rmklw hwtta
yyolko pwbvxvg xdwl yfje hftep kzzsr kho jeyf yvslxpw kfyv
xmk juyjxy eqno mdwklum reg dgn cirh wmxfyj bnxlgo dlobk
oyv gshqyot jgcqe dsf gyohqst gqgeojo egoogjq dmqpyp
sypianq yss lmhu ulmh itilh ndkda lhiit
qbxxl bxxql ikald nfap qixwbqq
jtqhqty ljysnl nwoj toa bmmyj pal
ahktew sxody nkvsf pbxyt baws wgwfwej bevgzm jus hcvajfy kzrb jwgwewf
jzsb szbj ujngwf nfuuf lfiuxdu uufnf orsy
vgo hto isstyul gau wsmxoqw
uxw itwf epaw hec wape hemol rpwyosc xzxmrll eetz zui kagca
mjncux muv rygdeis rygdeis
qgkqjvf iprzibd fkvqqgj llcrl vbh vlf lllrc zwrunt
dslsa wvoex eqbwj tjem gbx ayn xcan fnacl xggxon gnwjlh
yzosv hcxjiz yvon gcgd
bixpny ecln sda eymt bjiwk
rlcad lrdca adoqfzs rgty mds pwb kmwj
wkai pmryffq rrdmodc wgyx taz yxwg nkap
auynzwc vzg uapdv qkrh
ldmuysp oyu kpn ejbl mfifa bzs hwyn brlw qpzqx uyilao ysdumpl
czoxoj pwnultl wezolbw lyk aonesgb
nqy nhb nle yycp lgtbo ojf dytwyh ufa
rwr eph obg peh pejret prjtee ovgz
vdqf vdqf ycjrg ovzl lelbe vdqf
gvagdqm gvdgqam dmb zaxe nepzwn
emwh bkkbgec qwdgk mhvfsrf wmdfpp ekzuua
mbqw lgkyazt ckyhvnq uladwo owt
qwiwd pbo tkjoqda zapo dygqopv zzdlwfn
qty dhb iinncba ytq kvh idgoevt chx waq
ulffsvk vplsz ulffsvk uxsh cpwgxd ikgcacx nrirke uowcjvn
gknmxr grkxnm fco dilyyj grmxkn
saqxkh uhue nvu fef xsuxq ekyyoc bcaavd
qltwqa vrmpv vhra nof yprauc vkreojm eaq igiy mec
wvheiyg uthy gpvcs nhnjrne mqaejr tfnsly zfbhn entcc nystfl cpq
zxv jzk dwsjgrd gqqxhp xqxu naunwc yeh qzpkz awcnnu aoosa icadax
vpmqmg qmvpgm tqs mvpqmg
inehzu zwxeoy jxia fcyzxc hwikd
bzwnp kamsen ajpn kdls bzh xqcb bzwnp cmjnfa wmgx
hbuhc qgvhxy smzkxh zzebox hbcuh net wyrdppc yvgxqh
oeum oemu iyags xaipdi euom
tqljgoq ghtdhw xhnni lux qltojqg lki zxztda pcqjif acpzvwy
ydijaq kbyjxpu onyd hsfgz geqvbg
rwoih xog dtbzyr ryzbdt tdbyzr
vcdxf zosw pardxfz bmb mscmain lwfc jvq hbszcqh fxomsmm ahnugx
zutsemg pqzil ddv nsstz gmeuzst bedvy xkzzjpw xlqbd
xxf ltnnu yeb hbml agj meovtjr qrul kexerkw xxf
tqrpd hhcx bmdv nlmr pnu pajdtc rpatqi yekedx oeiuew epsshog
ttbfpv plairk toh jagfsg njnqpa tmwh vwqp irtxv
vdky uwc tkkkztp vdky vdky qlcw lza
rzie yundymy pwgx wtwtbg kpiw mewnb liveysj uvsbn
jgfvyny hacg pzra arpz uowswu puzsfu hoe heo vrq naup
hqv vrl uko qgpikho lligvxa wdld qgpikho
whvby yomxwj dieffc jkprinh dsaqy yfrnba woyq yexeb mjn cbszn xeswvvo
wowtgu rciyg rlas bra quyfec ihe thuu asxhscu bsbdpbi ogxosu
vydsaet tvnkjq piedkzj foeiqz zqivt iatsju tjnqvk drauaf vqitz invoz
cppn jqzw zmxr qksuas iifmjg xtkgf cppn cppn jpsd
nkifpsq cxdx bokxhm ebww kghagrp bofhrl grc cheuzyj
ibgrlvm hrcx jjuoh ipmt
hcoqkh fzt rgravb cimauj jxjq blct qhc vjxw pqpg qzp
jycxz xcv czxjy vxc
liljaur cgmg neldxb xfummcq yfhiukd dnqhl iolxn cmewhb
hpvoihj fkwokod txy uuktw vmqqb dpldzh yxmcay cyaxmy xycaym wekr
ccnaf wuxc ecadb vbgpt ccntf sezo skjdkbf fnctc
hqdtwho kdhyman bjtcjvr bwllva ncyffyr
xprn jrrvmj pdw yvexm ewbflbe eapml rvrmjj xmevy rxyzhf
wjcbpy qdgtcp cfjh muww fhg sgfdleo nelpte yucqa aavev
rci vqypsqt xmg rzii
gramh wwprtc ampdhw dajr
ovrm mdyhpbl mdylbph aykz
cbmo fxs nuugu guunu upt ljjuhjw nituh utp kxqc
rhabal rhabal rhabal vah lfrs
nrq qway ftzp rtjcks mbygdtd hsiqbh wypqb rtjcks cllp hsiqbh
ywa anhcf nvd puqkwg molrwck wsctx xvd molrwck
wox jzq jfen wcvus cswvu oxw irg lmu tpj viahm jesic
qenad neqad smlgi ydwzq ppdemvs ucyuf qtunm eoqx jlgv
sucpl nrdwbl ltvetok npbw ozzw hafyay sjmui sjmui jkqlq pyn pbuopx
nxgaiu ybyl meo kgh saqjaz xhbqr otelcyp vkwc
iqrl ldjlwvl ajhrl dnhutr gkknyqs mcvluet fgyu ogiz cxo aiunl orb
psd cyq xpoyqny yqc kozqh vonfd uhozwz pds hcpw
tvaxder tulwmw qiw avddbmh irog vynjzcc refx efxr emnvk
myjx npqk whm egw kpy igrrohg ukglx ldnuqw caqg ynx fckhnsh
dafv bkdoqg zcqvbco xgikoac cvbqczo
rtzhpwk ukuyp bayhzp upkuy ahbpyz
oarcuv pnlkxvw fqdkj hwzsz nauwl lpufibz vzfbgc unkluxy rwh xuknuyl
vxhsaj ppdxw qrswqtu ulwv uqtqwsr ppxwd
cww cww cww scu
wiiikwa bfpewt zbgxfkl iqpk tpbwfe aazdcxj ipqk icggn fwn fjr
net ovxuwpz yvzmzd yvzmzd
xgar czuhp vuhisaq fgrqxy evvrtf mnmar lsk
hld mxuedug itswju vmmejqx snzslqj toe bbmugph mgubhpb mowj nrjnzu
qbz ouhye hsldmp lcf hyhlrb ewvle zko
cke mupaq quapm eck
owu zdt lales tzd apjjo fhpx bmuktbw dvehpz
libvl zxypk azazc vtsom ohdzycb
kiowxnc scxygrf ckxnwio ycxsrgf
vcjj fqz lfawfx mps zhv qykch vhz psu zud spu fnpvkx
scfvum fuktgk tua ieosetl wwmjtt exnsw wwmttj plvd pfb kku pdbom
wkfw snukd wkfw gyaojdf bjw htagy cdsp
beh gatqxcu ibrooxr ssww orrioxb eenkqz
jlv affah mtbemf tylh aafhf
zqfajd uwzrw csouuip qzadjf
gsnlrw tcel hha tfbzrp ild aenqa
iirfxef kdux yvj vbzgj
ibx pfll rgkp nancij llpf xib gbkfy
uvw kkbavj pznsnk okigtxl ogitxkl eobbs xhaz wroabn ltogxki
bivdf lotvmoh vrb kpaeeue tdab qhukcb qmy kuqf kesu
egs hbsfeu esg twxko uib
ocraimu qilp ijmx eco nhevqp juxf ksejr bcqqau uhpt
pyx jmpglf juokd dxszjw cml vcjge pfg
gxwrt btmimse dkpbha idmz mtignka ngakmti
dpjhm jyalra hukf imocr lkgt rqywn quhe fukh
nbau xyc bdh yni xaawxm cyx xwaaxm akx gyodqe htbifc
bywdxe bfrp rvb rndl onal jghiwb nuta aint qlciwcx
fpic yrqce land soxhci qzc zoebsq hcdohcc fzhcl iyxb dqinum hchdcoc
zok ghgp zok lmk
ozfz zofz dkdekzb sqc
gfti zuqvg cexmtyl qwuqnj stepb erduqhy cuoizcs qudyreh kqvfdd guzqv
jrugz jzugr lmqu jgihgo hjfbz duxkn unxkd
ckiys dbqmi ckiys ckiys
iylp uvvdp pluifaa djo
esxec rwvel djxppqf jymwt ilm aiz upn aiz wrfefwi rwvel
nitgjr pokxuy puhdwg qtxpb veylp zqvzkbd lrvpcgu zuy rnigjt ibci
jboyzq ogcldr hlon ywav jqqtz qjzqt vyaw cok
aqdw jxn hqknh azbylg
jya qpxtmsj hqrtsgg qjtpxsm
pofcs sxw dlvru dlvur swx
yphvvb qqyyfsp sjkbff dqyerxe jxzes oof
pwbya txk bbwsj ywgimd kmdpc bawpy lbnt
bkbazff ldmaq tyfl acqurpy ndnrp
asw ctiv mnxzyc weeuwb gsn bzk irbyhxl cgqomj izy zbk
yrxcrbt bcrryxt pofe wwzl
vuaqez kbtuyai vuaqez dxqud uvo gmhtg dxqud
tpzs gqdxpxo zzpgta uurjx xpqxodg
cil lsv vznqw vro zqzvjhm jhgauzw uxnwk lci zpgpu frjvyzo tsv
zfvcuim gwn gnw dxfppok
btb goof iwadca aac tbb jha uvzi
qah ned ipmure kyta ffhrwe njz paq kaag xmlui
rkmw vrblwyy gpax hxsf zpbza gypuwf jbib ypcjwd vrlybyw
yfjljn uxpvg huik jsnah nkhsg yfjljn lqzsz
hagjlqx agnax jqalxgh rvjgtc mjrmph azznzcq gxajlqh
ipki bhoabp rmiyl dmjyxl zzsmap aju
tyjrr rigrf ciq qic avmwu jtr wpq
vuf cosgytm toycgms ufv qzpcbrs
epzgxr lydrsj ezxrpg expzgr
ecm prj kmak makk jpr
ccwyq txy okj matxa socoa
zrjphq gigayv ywkfmru yrwukmf fxjjrha gqkxx zhjy tisutx kufrywm izjfj igg
lfhgsro gsroflh wrpo lofhgsr
kgkgj wkhnab ubrjaoa ubrjaoa ubrjaoa ggdgh
hztutpn epnqmz ffcroq mnqpez niibpn kdloak xjui ozttj lyzsc pzgq inpnib
kruz sjqp mmd hhdxjgc mauouma asevvo upjwqi hxcgjhd etqzagp
zylf qime cho oraid svytv gqrjufv mker cho vnkyiin tjms
dotjul qyv hnh cibtg gdpauyx wzp
fabtira ejxoeor cqyethv ndjrq hnxn joq otng lrr csytrub
txhgepd fwdaanm nawdamf pxine qqrn pronw exnip qwkimt rvy
kuxzhi jln urzxtw rzu ebsuylm tscru qwlhfgq nnu nuchvz vuht
cqgu camlr umkltcf stx izp rtdwxff wkfvs
jhje cxix lefcrsu nebv idfzhic xqri xkft
utzxb znb ietupd uqgbhje aobip oawjwm hetyan uqtqv hpwzyri kwxyu
jvzvbt xuyvp aegdkb srbw bzabpf lyfriez cruyfu
nhi nih aeb ihn
hcf zypt djcm pkjx pvhh
rhvxcfk exydvk ids hybme hnk yfchvs mjbo meocn
rpboxr rxoprb hdzje zhedj
ziildbo apzvatr vsv isndq ebxyy ntm tdttg wkvdh qnids vkdhw xxolip
ywu uyw ipcjz pjzci xjn kvgk vsocprw
euzo njlpv ndrlhi drlnhi ivmjkb fjrtxta skvgmrd
gbyvj dkck gevpfvb lhadhx rgjcdn yraxh bdk oen vqryd bkr
vgkp hncttxb wgxh gdyjo bbdfzvc xhgw rznzgda yxrrlo gxhw
ifjlb fpecyic svhjp ilmj oxgr svhaf
vbqky lhccj xtmm xzjyykn oqmdq qywir bswly
euxxziv totzer vsxfx leo djho uoeaz edaig fbu lumbi
ooqtwq pvo kid vpo jxin bod btqc fbyuz
jhabi mronu htqqyz umjcbv sgnbp wyn cetmt pcjf
tnrkcyl dduuhxh rylkctn pwj rtynkcl mzzfomr
rxx ldqffi ulappk nltawbn tplhb kyb cqyi
vzkw gviooah vxh xeae ohvcad oaiwcj dkx
sdofdjt hcifv dqws sia mlwm vfich kavh myzue roops mzuye
uxs nlbmjp nlbmjp tlaxa tlaxa
ynnisp twx xtw jgkc yinpns
kumorsm wav xhx bpvz clqc ffmadzl ndny ymslo lobv
ljzabj tqhves mezh pwn wue dwfqq lynvtt boeknvi xqbd pkud tzlanis
lgq qiikzl oihnsr pivtjmu qhic yvmeebg rxu qgl yuxnqse dvu faxqez
ldk mlwja vmdqr yzlxiua amlubt ejmzfx nonm zhkxbn gaqbnqq
ttc ctt kneknx smtnaft abljip tct
uybhbiw zwojzlm cfxoopp abulenj znz zzn opllzmm yufk witwxzp
qvkybwi rdbxb qiuizmo fqgne jgot jxz dqhapn
vzinf ehaley amnk laheye invfz
pedakl ivld agzyhr wmzba tzzzg bazwm wjwgux thrnxkn
cmyhae nwfs nfsw kmh pxkaffq
vdf szupev tyunp qiiu deevxmy wozvtt nelnr kgdexy gparqj hajavz biizn
pwspk skpwp ontbjee pkspw cfbj
ihsmh djxtak wkzllao oyr djxtak prc
uhvihqq jrgf hdfek pdrfpt tghz gthz awae wcygi wujti svq fhedk
gnfhsj odqlt netmsul rviio nkzw nkzw
xyvc clxw cyxv lxcw
duegck pkviu npwsp zdx wpvn dmxgnv ixv fybs xteru
vih kgk hads boaddu daiwo hozoufv nef vtcplc isiw
tzqoo dqlgvno jzlay sywx ecej addt ecej addt mnfcu
ymgmby zegudpx ipsjai ger wcwjw brzebb
eqekxlx itra xekelxq exqkexl
rciu ojaa ircu nxjga puvmwou remgu
sltth pprimb slnxopq avtir hvpv ppww fhfap wisn kzs jcuuuuf
xbppc ydpbq zhjh oym iljzvk vsb
ueye shtps uccehi ccheiu dqm yeeu
gwywf lcpv qza qza gzuovj jfzffyh oybfxqv
aawi ynsvdco azdoz cqr tnyquq xlyvbx eca kcalpes
zumgzhy rou kguqa vubw bwgd qprxcg etnbev nqmi
fyd tuoz uwclqn cgl lrpkf irz dizv nxze clg jghx jbpt
kwuanos eorjr tcahp kwuanos cyrpfji zxayggd kwuanos jkqt qqvbork lizk
vtu ovje vhg ovje vtu zcy hrhtr puawfgv
bliz exp wot svxv epx
jiqgxwj yips hjsatc jgsrno msfp vxvbt bba bqmw xjgpgog
vpvypp ggwp wggp gndp hedpse afji hcqgof
hxueubt hiynoa qqzaj ohb qway
akq nfnes sdrlza nfnes weq
udxpdpx gctuv llhxuow rqtetm hdbnpte oebapv civy oeobu ftgivd pykj
pbgbvn jgmr xrz dfn gosjobw ndf
gnf dtbsnc fwcmml tscdnb fgn qgadusl eifpk
vmnv yuxrup qcphi tanc tnca kjrv cphqi
hclggs sghglc fgplp odn pfglp emkrulf whwtmbs qnuyg
wcxtr ani ain sha hsa zxbkf bzxokat qezo ljqxi xqcwfmd dxo
waiq smpbu dbyka uibxjrg nze wiqa rfpts ddjsjv jqqjez bpusm
lpcxf vsbj owjwc tuqj vkrgrh jsjdepv oil lxrjox frsxsi clr
vzunp prwk nnd rfs vpuzn
pqpqv lvsk sqxf nhobsm hakbn ywj
xxu uxx szqnmi lnwtmx
akq nmlw fupwsth jduvhva
nac wwlxqck hpbce vxxqa fyp xvxqa kxwclqw yvlmv bfwi
pzxjbj nvwv mdooiez vvftp enjrsck iypu uhru fpx omtd
llxgp qwf pwaj cuhb scloot hbcu jgp vjw ooclst
sisd akawvzd wvdzkaa gyoij ikt eeeosb jiwiup
tche vxj sbctqv jvx gosur usgor ibo yqxo qqgd zspl
cidd welisl fxblxqk qxbklfx fbdoqcz glhq iylodvz zvds ghlq
cnsa hrxst mrnkqtj bptq jmi cpbcofs kveyeur uzmga modphm rtx kntqjrm
dvyup usfaq rtghoec bvcos fqsua zohwwg
onf vncybi dlaxni oqyqqkn
okfwa qyyx ebnv llql nphq etdt ytgivlo jwgwz kiob
ann vqnqvpx wth lpwid bjvzw xpwqxcj azg ioeyzzp onwf
smy epzomx xep yid zctvrfj astdj cfg fgc eriuxt
rljqgin wzobzrh cuwtx vcsbx tmg tuysq vxipgho
ewp rsrnsj wgeyin lrji ddgt utol xxwut fjiwopa
upu ftvqbk tfkvbq fdwga rmu puu hbiasjw
cfl lmqkb lfc wbtlfi uqsjs ejgmphi tbliwf nzcela gzb
zop unwmiu acull mkwh hvruknw rfk mmhaz iqmenq fifino
iczua bjut tlgf zicau jtbu
mtka ipd mdifj kps
irqkysw xfsjl tedx yckkbx iktxb sqxn pbfvubv uudzppz
mdrn cihat wcext kufs awwtjok pfjg
wdevt tyo zzbp pqlqq wdevt
yhatqkv ntuhw tdfd buxazh xbcsv bas gkv rbzi tddf jbj bsa
malip hiiy qezz yhii wlfojre
zqnfll bssveq lprwbep bhqml tztbt
npnxotu yupdytb jptqo klfydfe fpucmfq svxcqr unopxnt
gdpz gwj iytiohu efk ctjzf asade abhotq brmhu tbtdur zzksbh
kxft klzslf tjdzciy lzslkf
ejei ezmemvg xlt zte tbwhz dgnfpao zotck wus uaz gbwbb
dgednf vypmbs eiytot empfmny
uopmui uehue wdvzt adpfcif mutl ifaztka vydi xumtz orstno
dleero olxiq gxnlfm nfmxlg wloeavr olhrwg hrjd yicj ymyeex qav gxyjgfq
hevj rqcne zycgb qgqtn rqcne ptfvu yyyu zlm hevj
zrkhuh sttnkt hkuzhr vqtu
ppsfm kcao qjq dgadglx cxaawjn pbucfu fed qgioarc dfe ricoaqg
vmawf oktunea zraoir gkt zraoir jcvkqoq
mqgml ecawug ugwace szwul iwbmooj owmiojb
auggaw cypcuw npci vuyxijd pofswjx vdkrgx xylk rom ksj
qmwx jgsrdj ikva xzxw avik
zzhcqu rbg pywjdn wyndpj zchuqz
wzd wqycftu yldezp zovuy oydia hovewe
kfid qkkk thak qhbf rvzlzvu uuxh pbj hkat gow oeqcw knqqzha
sua itv hfpg bdqye bznlrk hfpg bdqye kvir kaai ggtz jqn
ulggl guitamm tkpckso fupacz otxtqpd jxnqc
ueesb ndyik vjftz jgqqv nrcf
krh dqpmsw fybzynl zhjbvkw exefc rhs neq ldprb bhhvxm pjwirun
ymavl qwxr yavml wagwc ekokrpq zewppw iumcgin cxdvwx
wwdukav kuawvwd kowv dkwvuwa
eazot bil tzu vdwwbm fvauwrq
esq tixokph yspf ztoxfut lgzush pwv swh pwv auqhuu tixokph
pdbeyxi poio mugfkb brwbbx aao uszw fokjeb uswz
sbs ryjr ptispi tvnhu htunv vthnu
czjmg hbdjhvi jrkoy fpgwc syafy aar kvnq eaecsb wqzpx
twtp dvl uvyje qtlzj dsvyr qpjnj eyoigx bhgpccy gwn dtuf
mxit xunctu vbyks wmqc jriuupl ybvks uncutx nsoxwrb ykt prc
yye mgf uhc irowpc dsdv iwaxod ftavlj dxzp tcch tcch mefz
rxe xwrrgl xwrrgl duu rxe xbbgoe
ucsz akswcd ojrmqq cox hgfh lxwu ltnnf cenikcp
opjhdp svwezr svwezr opjhdp
qojlkl ircxqnt utfmdg fcvr vehkcvt ufmzcpv xwlh ddavv xel bwlz fii
rzkayeh iursm zhily hdnq fqydfvt uwoy hptpiqu tdqy bgr xdr
ymruz umzry hbltwya jhwhzk flh tahylbw bdbaimb qscbp ntkuf
uxpato owsqyao vaog oenomkc usrmnc epua vzkppls
qxqczbk qyguz alawj xgjawtw wxtjgwa snfcdmz
fjfgos rmpd mgs vbk dlls jkljao eoovdfb ucdvaoq qmjmqku ney porr
nmcrqz zcoxpk dlnzksd ymh zyg spxss ruyk bychq gsgv eusiuid mnrqcz
jbzadnx lzl sdamer okoico frqisrm lxet agriw
xceoqr qai vahc jjzifsn exg
igjpn wfy ukn aag quro wklsq cjq bgtjrdz gmub wyhh
fzlwnm mygfn vkzwvw zvhsex gfki
ijvzgai ebmeq wssfmbq uguh sfuutm nwkgmex dxael liakdxs rnf sky yowpxc
bjzkyjh fced nji esowk qxsubsk qgtts
nkdgo bbjfq fgnxnhd gfjchl jetdb xubsgj eiju ldlm oxsx znft bbqfj
xovcnob pxfe pmstes yzkdm iqlvha nmcziix fexp ivqalh rxecqps
xpyew xudfud wwqe qhfjlcu epv fnrbgyv ihli qngtx yjlfg ozqbzn esp
timl gcohx vqzic gzm shwlkkv icqzv urchuc
xpqq gaqzwo cci dowahsr gaqzwo
jjsagdl umbpxre kyre zvaryft tmw pxpnjy
aqovcz nunq nnuq xjrvvh autjmit jiatumt
elg lps lge zjjot hwz tmqrup xaxxmo zlbzp uftd fukdad kvpymsm
iokwzal ywti zbdmzbu lprywe wbgbwza ypogbga kzliwao wstqi eqm keaeaj gbabwwz
lwfpk mhufe eddzgd ljxyqy vhzkct uemhf
lwqil fzugdo faq feppo usl llwqi
nje hthr ropq qvcepu bexszfj avmzjvv zajmvvv fhcd xnc cnx qnuaux
kvksn dphbyz nsx wrcc ccrw
nzpa pzzunfv ygzjy gxrrtcj hrt trh pwxpg yifgjmo fnupzzv wbzx
aepti rbojui ypvhe ubojri tcema aan dntkw qjx bfvmyos tcm hvoqytn
qpwq exu jvsiwj gsw avr vbemldy
xsbzpf xbzyvx xax sxh vpxt gccy xxa zhgbwoa hwwxoky fhvdxfc pvtx
pnsa ovtjolz tyutl eyjjzt jvtoolz owbypvr tytlu ewtzgec
cyg dwwk eihsp aeuk bbnay aluwyz hdmv uaek mwt ihpse wjhnkeg
fhzx vjetz vjub tejvz
ewwyb jidhu pyvyenn igtnyd tiwr akwkkbi myz xxjwb jjrdeg
jbkuw kwir rkiw ubwkj
bltffuw lftwufb hhsh wfbtulf nrxaa rlszi toijxnz czlci
bqrm pga zgblgcw pgwhhn lcgzwbg bcgzlgw yqb
mhjj vjoa gnjlc kclcr ito ofksy giavy fpqeioj
bkiqmif izidbui sttxxi bswhkxp sduuw
mjgnvw mjgwnv ojzyuv gvj
qxn kkhc whd fgwk auzugg augzgu kqfov wfgk
spdxbnu xpfofsb bpfsoxf ahjywql spbxoff
bwqxhlm wbqlxmh kqgpl fyzgf guhkvgx ovk qhmp gnrmu wvd wedj
vvwf hcnc vvwsngj qedzoxm hcnc qedzoxm kjthdi cbwqep qtvu
gio iqklmro noqablo bab jiqc rwebyg rqkloim wzmgs uunl amqs iwj
snxj szobqt zcgvwv wiyqknu
uto jteikwd cew gqsks hmvjtcy sach
zpgl qnkoex amhufmr figns upv xezrl rjleak nwrna
pzkvrdz dtonazj gtr gfxucuf lstjl lsjtl rgkope kzpdzrv lyptn zfxjys ttk
ddxgm lumlgki jhv doft kok swy ckds swy ddxgm lbfbdv
qfs rcufzgz iaiqw qfs qfs
nvkbo sgv mquwb ritpye nbkov poex hraorm qrrr qdt qefl
irxannd fiud ehyb ggx plqg pvvn uuptop tcvbm abuf bcfnmw
qwya ukblz epmbfr vmlon yqwa
hlo mmv vmm mvm
svzpxun yugbbe sbbpxs dmy xspbbs zhpovyf fyovhzp cpbt pke
zgk gft zybs zrgcoo ypu bue htgo
xnesq srsx pkzaoh cfqzugh
lntd nvxetbv clykjpd svmibpx evxtvnb yldkpjc
jsqq tzwak hephg eqwczd ioisa yim tmdifn mceip
kuwqz wzkqu zwchmj lfec uexne iztp llityt
kvamkpc pvbryqh ion cwizjde gln kcpvmak pzzlw gnl
ydeqf bfaab sydqhbp smsxdjr pynrs cqymt
onb eiab bno nob
mqslq scnelxv hyllrf scnelxv mqslq wmnbk
pttu kubby lgop bbyuk gsk skg ikktlbb inbyvz
xznvl zwtdj vbxdyd clhw
hgy zudelp ickc drfjgn iyws xhc
zzv wik iorhat qkb kjb lykdz vrce yjsjwj
gyw xzgbi efus uuy
hwcy ujdun bjjuvd jbdvju onnk xeyy mmp onkn qyzl
jwfm ptjwrbl hhuv uolz adyweh qpj wxyogp igvnojq jmfw pqs fsnirby`
	input08 = input07
)

func main() {
	fmt.Println(advent2017.SumSkipDigits(input01, 1))
	fmt.Println(advent2017.SumSkipDigits(input02, len(input02)/2))
	fmt.Println(advent2017.RowMaxDiffChecksum(input03))
	fmt.Println(advent2017.RowDivChecksum(input04))
	fmt.Println(advent2017.MovesToCenter(input05))
	fmt.Println(advent2017.LeastCumulativeSpiral(input06))
	fmt.Println(advent2017.ValidPassphrases(input07))
	fmt.Println(advent2017.ValidAnagramPassphrases(input08))
}
