package station

import (
	"fmt"
)

const (
	AALSTRASSE_STUTTGART                                   = "Aalstraße Stuttgart"
	ABELSTRASSE_LUDWIGSBURG                                = "Abelstraße Ludwigsburg"
	ABZW_HOHENEGARTEN_MAINHARDT                            = "Abzw. Hohenegarten Mainhardt"
	ABZW_HOHENEGARTEN_B14_MAINHARDT                        = "Abzw. Hohenegarten B14 Mainhardt"
	ABZW_REHNENHAUWETZGAU_SCHWABISCH_GMUND                 = "Abzw. Rehnenhau/Wetzgau Schwäbisch Gmünd"
	ABZWEIG_ROHRBRONN                                      = "Abzweig Rohrbronn"
	ABZWEIG_HOLZGERLINGEN                                  = "Abzweig Holzgerlingen"
	ABZWEIG_SCHLATTSTALL                                   = "Abzweig Schlattstall"
	ABZWEIG_NEUWEILER                                      = "Abzweig Neuweiler"
	ABZWEIG_MONCHBERG                                      = "Abzweig Mönchberg"
	ABZWEIG_HEBSACK                                        = "Abzweig Hebsack"
	ABZWEIG_AURICH                                         = "Abzweig Aurich"
	ABZWEIG_DATZINGEN                                      = "Abzweig Dätzingen"
	ABZWEIG_ALTERSBERG_KAISERSBACH                         = "Abzweig Altersberg Kaisersbach"
	ABZWEIG_B466_REICHENBACH_IM_TALE_DEGGINGEN             = "Abzweig B466 Reichenbach im Täle (Deggingen)"
	ABZWEIG_BAHNHOF_MARKGRONINGEN                          = "Abzweig Bahnhof Markgröningen"
	ABZWEIG_BEIM_KLOSTER_ADELBERG                          = "Abzweig beim Kloster Adelberg"
	ABZWEIG_DAFERN_LIPPOLDSWEILER                          = "Abzweig Däfern Lippoldsweiler"
	ABZWEIG_DEGENFELD_LAUTERSTEIN                          = "Abzweig Degenfeld Lauterstein"
	ABZWEIG_DRACKENSTEIN_GOSBACH                           = "Abzweig Drackenstein Gosbach"
	ABZWEIG_EISENBACHSEE_PFAHLBRONN                        = "Abzweig Eisenbachsee Pfahlbronn"
	ABZWEIG_GREUTHOF_VORDERSTEINENBERG                     = "Abzweig Greuthof Vordersteinenberg"
	ABZWEIG_HOHREIN_GOPPINGEN                              = "Abzweig Hohrein Göppingen"
	ABZWEIG_KORNBERGHALLE_DURNAU                           = "Abzweig Kornberghalle Dürnau"
	ABZWEIG_NEUFURSTENHUTTE_GROSSERLACH                    = "Abzweig Neufürstenhütte Großerlach"
	ABZWEIG_OBERWALDEN_WANGEN_GP                           = "Abzweig Oberwälden Wangen (GP)"
	ABZWEIG_RAUHER_KAPF_BOBLINGEN                          = "Abzweig Rauher Kapf Böblingen"
	ABZWEIG_REICHENBACH_DONZDORF                           = "Abzweig Reichenbach Donzdorf"
	ABZWEIG_SCHMALENBERG_RUDERSBERG                        = "Abzweig Schmalenberg Rudersberg"
	ABZWEIG_STEINBACH_RUDERSBERG                           = "Abzweig Steinbach Rudersberg"
	ABZWEIG_WALTERTAL_HOHENSTADT_GP                        = "Abzweig Waltertal Hohenstadt (GP)"
	ABZWEIG_WINZINGEN_DONZDORF                             = "Abzweig Winzingen Donzdorf"
	ABZWEIGUNG_BAIERECK_THOMASHARDT                        = "Abzweigung Baiereck Thomashardt"
	ABZWEIGUNG_HEPSISAU_WEILHEIM_T                         = "Abzweigung Hepsisau Weilheim (T)"
	ACHALMSTRASSE_HOLZGERLINGEN                            = "Achalmstraße Holzgerlingen"
	ACHALMSTRASSE_NURTINGEN                                = "Achalmstraße Nürtingen"
	ADELSBACH_WINNENDEN                                    = "Adelsbach Winnenden"
	ADELSTETTEN_PFAHLBRONN                                 = "Adelstetten Pfahlbronn"
	ADENAUERBRUCKE_OBEN_OBERESSLINGEN                      = "Adenauerbrücke (oben) Oberesslingen"
	ADLER_HEININGEN_GP                                     = "Adler Heiningen (GP)"
	ADLER_ALTENSTADT_GEISLINGEN_STEIGE                     = "Adler Altenstadt Geislingen (Steige)"
	ADLERAPOTHEKE_STRASSDORF                               = "Adler/Apotheke Straßdorf"
	ADLERBRUCKE_PLUDERHAUSEN                               = "Adlerbrücke Plüderhausen"
	ADLERPLATZ_HOCHBERG                                    = "Adlerplatz Hochberg"
	ADLERSTRASSE_FELLBACH                                  = "Adlerstraße Fellbach"
	ADLERSTRASSE_WERNAU_N                                  = "Adlerstraße Wernau (N)"
	ADLERSTRASSE_HERRENBERG                                = "Adlerstraße Herrenberg"
	ADLERSTRASSE_OTTENBACH                                 = "Adlerstraße Ottenbach"
	ADLERWEG_GEMMRIGHEIM                                   = "Adlerweg Gemmrigheim"
	ADMIRALWEG_STUTTGART                                   = "Admiralweg Stuttgart"
	ADOLFKOLPINGSTRASSE_DONZDORF                           = "Adolf-Kolping-Straße Donzdorf"
	AFFSTATTER_TAL_HERRENBERG                              = "Affstätter Tal Herrenberg"
	AGENTUR_FUR_ARBEIT_BIETIGHEIM                          = "Agentur für Arbeit Bietigheim"
	AGENTUR_FUR_ARBEIT_LUDWIGSBURG                         = "Agentur für Arbeit Ludwigsburg"
	AGNESWEG_HOLZGERLINGEN                                 = "Agnesweg Holzgerlingen"
	AHORNSTRASSE_PLUDERHAUSEN                              = "Ahornstraße Plüderhausen"
	AHRENWEG_STUTTGART                                     = "Ährenweg Stuttgart"
	AICHELBACH_OPPENWEILER                                 = "Aichelbach Oppenweiler"
	AICHENBACHHOF_PLUDERHAUSEN                             = "Aichenbachhof Plüderhausen"
	AICHHOLZHOF_UNTERWEISSACH                              = "Aichholzhof Unterweissach"
	AICHSTRUT_WELZHEIM                                     = "Aichstrut Welzheim"
	AIDLINGER_STRASSE_DAGERSHEIM                           = "Aidlinger Straße Dagersheim"
	AIDLINGER_STRASSE_DEUFRINGEN                           = "Aidlinger Straße Deufringen"
	AKADEMIEGARTEN_NEUHAUSEN_F                             = "Akademiegärten Neuhausen (F)"
	AKAZIENWEG_HERRENBERG                                  = "Akazienweg Herrenberg"
	ALBBLICK_ZELL_ESSLINGEN                                = "Albblick Zell (Esslingen)"
	ALBBLICK_STUTTGART                                     = "Albblick Stuttgart"
	ALBERTEINSTEINGYMNASIUM_BOBLINGEN                      = "Albert-Einstein-Gymnasium Böblingen"
	ALBERTLUTHULIPLATZ_STUTTGART                           = "Albert-Luthuli-Platz Stuttgart"
	ALBERTSCHAFFLESTRASSE_STUTTGART                        = "Albert-Schäffle-Straße Stuttgart"
	ALBERTSCHWEITZERGYMNASIUM_LEONBERG                     = "Albert-Schweitzer-Gymnasium Leonberg"
	ALBERTSCHWEITZERSCHULE_GOPPINGEN                       = "Albert-Schweitzer-Schule Göppingen"
	ALBERTSCHWEITZERSTRASSE_BESIGHEIM                      = "Albert-Schweitzer-Straße Besigheim"
	ALBERTSCHWEITZERSTRASSE_HOFINGEN                       = "Albert-Schweitzer-Straße Höfingen"
	ALBRECHTHAUTGASSE_KUPPINGEN                            = "Albrecht-Haut-Gasse Kuppingen"
	ALBRECHTSTRASSE_LUDWIGSBURG                            = "Albrechtstraße Ludwigsburg"
	ALBSCHULE_STUTTGART                                    = "Albschule Stuttgart"
	ALBSPORTHALLE_BOHMENKIRCH                              = "Albsporthalle Böhmenkirch"
	ALBSTRASSE_REICHENBACH_F                               = "Albstraße Reichenbach (F)"
	ALBSTRASSE_KORNWESTHEIM                                = "Albstraße Kornwestheim"
	ALBSTRASSE_GROSSBETTLINGEN                             = "Albstraße Großbettlingen"
	ALBSTRASSE_NEUFFEN                                     = "Albstraße Neuffen"
	ALBSTRASSE_SPARWIESEN                                  = "Albstraße Sparwiesen"
	ALBSTRASSE_SIEDLERSTUBE_NURTINGEN                      = "Albstraße (Siedlerstube) Nürtingen"
	ALDINGER_STRASSE_KORNWESTHEIM                          = "Aldinger Straße Kornwestheim"
	ALDINGER_STRASSE_LUDWIGSBURG                           = "Aldinger Straße Ludwigsburg"
	ALDINGER_STRASSE_HEGNACH                               = "Aldinger Straße Hegnach"
	ALEMANNENSTRASSE_SIELMINGEN                            = "Alemannenstraße Sielmingen"
	ALEMANNENSTRASSE_OEFFINGEN                             = "Alemannenstraße Oeffingen"
	ALEXANDRINENPLATZ_HOCHBERG                             = "Alexandrinenplatz Hochberg"
	ALFREDDIEBOLDWEG_WAIBLINGEN                            = "Alfred-Diebold-Weg Waiblingen"
	ALFREDLEIKAMGARTEN_KORB                                = "Alfred-Leikam-Garten Korb"
	ALLEEALTER_WEG_TAMM                                    = "Allee/Alter Weg Tamm"
	ALLEEBRACHTERSTRASSE_TAMM                              = "Allee/Brächterstraße Tamm"
	ALLEENFELD_FREUDENTAL                                  = "Alleenfeld Freudental"
	ALLEENSTRASSE_ZELL_ESSLINGEN                           = "Alleenstraße Zell (Esslingen)"
	ALLMANDKLINGE_HOHENHASLACH                             = "Allmandklinge Hohenhaslach"
	ALLMENDSTRASSE_MAICHINGEN                              = "Allmendstraße Maichingen"
	ALLMERSBACH_AM_WEINBERG_ALLMERSBACH_AM_WEINBERG_ASPACH = "Allmersbach am Weinberg Allmersbach am Weinberg (Aspach)"
	ALM_MURRHARDT                                          = "Alm Murrhardt"
	ALPSEEWEG_STUTTGART                                    = "Alpseeweg Stuttgart"
	ALTBACH_ALTBACH                                        = "Altbach Altbach"
	ALTBACHER_HOF_ALTBACH                                  = "Altbacher Hof Altbach"
	ALTBACHER_WEG_LIEBERSBRONN                             = "Altbacher Weg Liebersbronn"
	ALTDORFER_STRASSE_HILDRIZHAUSEN                        = "Altdorfer Straße Hildrizhausen"
	ALTE_BERGSTRASSE_DEIZISAU                              = "Alte Bergstraße Deizisau"
	ALTE_BUNDESSTRASSE_WAIBLINGEN                          = "Alte Bundesstraße Waiblingen"
	ALTE_HEUSTEIGE_ZELL_ESSLINGEN                          = "Alte Heusteige Zell (Esslingen)"
	ALTE_KELTER_FELLBACH                                   = "Alte Kelter Fellbach"
	ALTE_KIRCHE_JEBENHAUSEN                                = "Alte Kirche Jebenhausen"
	ALTE_KRONE_STUTTGART                                   = "Alte Krone Stuttgart"
	ALTE_MUHLE_HAFNERHASLACH                               = "Alte Mühle Häfnerhaslach"
	ALTE_PLOCHINGER_STEIGE_KIRCHHEIM_T                     = "Alte Plochinger Steige Kirchheim (T)"
	ALTE_POST_DETTENHAUSEN                                 = "Alte Post Dettenhausen"
	ALTE_RENNINGER_STRASSE_111_WEIL_DER_STADT              = "Alte Renninger Straße 111 Weil der Stadt"
	ALTE_RENNINGER_STRASSE_19_WEIL_DER_STADT               = "Alte Renninger Straße 19 Weil der Stadt"
	ALTE_RENNINGER_STRASSE_49_WEIL_DER_STADT               = "Alte Renninger Straße 49 Weil der Stadt"
	ALTE_SCHIFFFAHRT_METTINGEN                             = "Alte Schifffahrt Mettingen"
	ALTE_SCHULE_NEUSTADT                                   = "Alte Schule Neustadt"
	ALTE_STEIGE_GERLINGEN                                  = "Alte Steige Gerlingen"
	ALTE_STEIGE_LIEBERSBRONN                               = "Alte Steige Liebersbronn"
	ALTE_STEIGE_SCHORNDORF                                 = "Alte Steige Schorndorf"
	ALTE_STRASSE_FORNSBACH                                 = "Alte Straße Fornsbach"
	ALTE_TALSTRASSE_ST_BERNHARDT                           = "Alte Talstraße St. Bernhardt"
	ALTE_WAIBLINGER_STRASSE_NEUSTADT                       = "Alte Waiblinger Straße Neustadt"
	ALTE_WINNENDER_STEIGE_WAIBLINGEN                       = "Alte Winnender Steige Waiblingen"
	ALTENBERGSTAFFEL_STUTTGART                             = "Altenbergstaffel Stuttgart"
	ALTENBURG_STUTTGART                                    = "Altenburg Stuttgart"
	ALTENHEIM_PLATTENHARDT                                 = "Altenheim Plattenhardt"
	ALTENHEIM_FELLBACH                                     = "Altenheim Fellbach"
	ALTENHEIM_BURGHALDE_SINDELFINGEN                       = "Altenheim Burghalde Sindelfingen"
	ALTENSTEIGER_STRASSE_NAGOLD                            = "Altensteiger Straße Nagold"
	ALTENZENTRUM_HOLZGERLINGEN                             = "Altenzentrum Holzgerlingen"
	ALTER_BAHNHOF_HEININGEN_GP                             = "Alter Bahnhof Heiningen (GP)"
	ALTER_BAHNHOFWEG_HOFINGEN                              = "Alter Bahnhofweg Höfingen"
	ALTER_FRIEDHOF_DARMSHEIM                               = "Alter Friedhof Darmsheim"
	ALTER_FRIEDHOF_BOBLINGEN                               = "Alter Friedhof Böblingen"
	ALTER_GUTSHOF_STUTTGART                                = "Alter Gutshof Stuttgart"
	ALTER_OSSWEILER_WEG_LUDWIGSBURG                        = "Alter Oßweiler Weg Ludwigsburg"
	ALTER_SPORTPLATZ_WENDLINGEN_N                          = "Alter Sportplatz Wendlingen (N)"
	ALTER_WEG_TAMM                                         = "Alter Weg Tamm"
	ALTES_RATHAUS_MAICHINGEN                               = "Altes Rathaus Maichingen"
	ALTES_RATHAUS_HILDRIZHAUSEN                            = "Altes Rathaus Hildrizhausen"
	ALTES_RATHAUS_UNTERJETTINGEN                           = "Altes Rathaus Unterjettingen"
	ALTES_RATHAUS_KEMNAT                                   = "Altes Rathaus Kemnat"
	ALTES_RATHAUS_BITTENFELD                               = "Altes Rathaus Bittenfeld"
	ALTES_RATHAUS_WEISSENSTEIN                             = "Altes Rathaus Weißenstein"
	ALTES_RATHAUS_DETTENHAUSEN                             = "Altes Rathaus Dettenhausen"
	ALTES_SCHULHAUS_ENSINGEN                               = "Altes Schulhaus Ensingen"
	ALTFURSTENHUTTE_GROSSERLACH                            = "Altfürstenhütte Großerlach"
	ALTSTADT_LEONBERG                                      = "Altstadt Leonberg"
	ALTSTADT_WALDENBUCH                                    = "Altstadt Waldenbuch"
	ALTSTADTGARAGE_HERRENBERG                              = "Altstadtgarage Herrenberg"
	ALTVATERWEG_KIRCHHEIM_T                                = "Altvaterweg Kirchheim (T)"
	ALUP_KONGEN                                            = "Alup Köngen"
	AM_AUSSEREN_GRABEN_STUTTGART                           = "Am Äußeren Graben Stuttgart"
	AM_BERGWALD_STUTTGART                                  = "Am Bergwald Stuttgart"
	AM_BISMARCKTURM_STUTTGART                              = "Am Bismarckturm Stuttgart"
	AM_BRENNER_LIPPOLDSWEILER                              = "Am Brenner Lippoldsweiler"
	AM_BRONNBACH_WEILER_SCHORNDORF                         = "Am Bronnbach Weiler (Schorndorf)"
	AM_BRUNNELE_ALBERSHAUSEN                               = "Am Brünnele Albershausen"
	AM_DRESSELBACH_SACHSENWEILER                           = "Am Dresselbach Sachsenweiler"
	AM_ESCHELBACH_HOLZGERLINGEN                            = "Am Eschelbach Holzgerlingen"
	AM_FELDRAND_STUTTGART                                  = "Am Feldrand Stuttgart"
	AM_FRIEDHOF_PLEIDELSHEIM                               = "Am Friedhof Pleidelsheim"
	AM_FRIEDHOF_GROSSSACHSENHEIM                           = "Am Friedhof Großsachsenheim"
	AM_GALGENBERG_GOPPINGEN                                = "Am Galgenberg Göppingen"
	AM_KATZENBACH_WAIBLINGEN                               = "Am Kätzenbach Waiblingen"
	AM_KINDELBERG_RENNINGEN                                = "Am Kindelberg Renningen"
	AM_KIRCHPLATZ_STUTTGART                                = "Am Kirchplatz Stuttgart"
	AM_KREUZWIESLE_LORCH                                   = "Am Kreuzwiesle Lorch"
	AM_KRIEGSBERGTURM_STUTTGART                            = "Am Kriegsbergturm Stuttgart"
	AM_KUGELRAIN_BAIERECK                                  = "Am Kugelrain Baiereck"
	AM_LINDLE_BONLANDEN                                    = "Am Lindle Bonlanden"
	AM_MAURENER_BERG_DITZINGEN                             = "Am Maurener Berg Ditzingen"
	AM_MEERBACH_BARTENBACH_GOPPINGEN                       = "Am Meerbach Bartenbach (Göppingen)"
	AM_MITTELKAI_STUTTGART                                 = "Am Mittelkai Stuttgart"
	AM_MUHLBACH_SCHORNDORF                                 = "Am Mühlbach Schorndorf"
	AM_OCHSENWALD_STUTTGART                                = "Am Ochsenwald Stuttgart"
	AM_PFAFFENWALD_HOCHDORF_EBERDINGEN                     = "Am Pfaffenwald Hochdorf (Eberdingen)"
	AM_ROMERKASTELL_STUTTGART                              = "Am Römerkastell Stuttgart"
	AM_SCHATTWALD_STUTTGART                                = "Am Schattwald Stuttgart"
	AM_SCHLOSS_BONNIGHEIM                                  = "Am Schloss Bönnigheim"
	AM_SCHLOSSBERG_HOFINGEN                                = "Am Schlossberg Höfingen"
	AM_SCHONEN_RAIN_ESSLINGEN_N                            = "Am Schönen Rain Esslingen (N)"
	AM_SONNENBERG_LUDWIGSBURG                              = "Am Sonnenberg Ludwigsburg"
	AM_TURMLE_UNTERMBERG                                   = "Am Türmle Untermberg"
	AM_UNTEREN_SCHLOSSBERG_NECKARREMS                      = "Am unteren Schlossberg Neckarrems"
	AM_WARMEN_MUHLHAUSEN_IM_TALE                           = "Am Warmen Mühlhausen im Täle"
	AM_WASEN_OTLINGEN                                      = "Am Wasen Ötlingen"
	AM_WASSERTURM_MARKGRONINGEN                            = "Am Wasserturm Markgröningen"
	AM_WIESENBACH_HOLZHAUSEN_UHINGEN                       = "Am Wiesenbach Holzhausen (Uhingen)"
	AM_WIESENGRUND_BERKHEIM                                = "Am Wiesengrund Berkheim"
	AM_WIESENRAIN_KIRCHHEIM_T                              = "Am Wiesenrain Kirchheim (T)"
	AM_ZIPFELBACH_BITTENFELD                               = "Am Zipfelbach Bittenfeld"
	AMSELWEG_WELZHEIM                                      = "Amselweg Welzheim"
	AMSELWEG_GEMMRIGHEIM                                   = "Amselweg Gemmrigheim"
	AMSELWEG_HERRENBERG                                    = "Amselweg Herrenberg"
	AMSTERDAMER_STRASSE_BOBLINGEN                          = "Amsterdamer Straße Böblingen"
	AMSTETTER_STRASSE_STUTTGART                            = "Amstetter Straße Stuttgart"
	AMTSGERICHT_NURTINGEN                                  = "Amtsgericht Nürtingen"
	AMUNDSENSTRASSE_SINDELFINGEN                           = "Amundsenstraße Sindelfingen"
	AN_DER_WETTE_KORNWESTHEIM                              = "An der Wette Kornwestheim"
	ANHAUSERSTRASSE_HEGENSBERG                             = "Anhäuserstraße Hegensberg"
	ANNASCHIEBERWEG_ESSLINGEN_N                            = "Anna-Schieber-Weg Esslingen (N)"
	ANNEFRANKSCHULE_WENDLINGEN_N                           = "Anne-Frank-Schule Wendlingen (N)"
	ANNEFRANKSTRASSE_HOLZGERLINGEN                         = "Anne-Frank-Straße Holzgerlingen"
	ANTONSCHMIDTSTRASSE_WAIBLINGEN                         = "Anton-Schmidt-Straße Waiblingen"
	ANTONIAVISCONTISTRASSE_BIETIGHEIM                      = "Antonia-Visconti-Straße Bietigheim"
	ANTWERPENER_STRASSE_STUTTGART                          = "Antwerpener Straße Stuttgart"
	AOKWALDHORN_BACKNANG                                   = "AOK/Waldhorn Backnang"
	APOTHEKE_OTLINGEN                                      = "Apotheke Ötlingen"
	APOTHEKE_BEILSTEIN                                     = "Apotheke Beilstein"
	ARBEITERZENTRUM_BOBLINGEN                              = "Arbeiterzentrum Böblingen"
	ARBEITSAGENTURPOST_WAIBLINGEN                          = "Arbeitsagentur/Post Waiblingen"
	ARNDTSPITTASTRASSE_STUTTGART                           = "Arndt-/Spittastraße Stuttgart"
	ARSENALPLATZ_LUDWIGSBURG                               = "Arsenalplatz Ludwigsburg"
	ARZTEHAUS_MURRHARDT                                    = "Ärztehaus Murrhardt"
	ASANG_STUTTGART                                        = "Asang Stuttgart"
	ASEMWALD_STUTTGART                                     = "Asemwald Stuttgart"
	ASPACHER_BRUCKE_BACKNANG                               = "Aspacher Brücke Backnang"
	ASPEN_STUTTGART                                        = "Aspen Stuttgart"
	ASPEN_HOCHDORF_ES                                      = "Aspen Hochdorf (ES)"
	ASPENWALDSTRASSE_STUTTGART                             = "Aspenwaldstraße Stuttgart"
	ASPERG_ASPERG                                          = "Asperg Asperg"
	ASPERGER_STRASSE_STUTTGART                             = "Asperger Straße Stuttgart"
	ASPERGLEN_ASPERGLEN                                    = "Asperglen Asperglen"
	AUBERLENWEG_STUTTGART                                  = "Auberlenweg Stuttgart"
	AUCHTERTSTRASSE_SCHLIERBACH                            = "Auchtertstraße Schlierbach"
	AUENWALDHALLE_UNTERBRUDEN                              = "Auenwaldhalle Unterbrüden"
	AUERHAHNSTRASSE_SCHORNDORF                             = "Auerhahnstraße Schorndorf"
	AUF_DEM_BUHL_LORCH                                     = "Auf dem Bühl Lorch"
	AUF_DEM_KIES_BESIGHEIM                                 = "Auf dem Kies Besigheim"
	AUF_DEM_LERCHENBERG_NURTINGEN                          = "Auf dem Lerchenberg Nürtingen"
	AUF_DEM_WASEN_LUDWIGSBURG                              = "Auf dem Wasen Ludwigsburg"
	AUF_DEN_RUDERN_BURGSTALL_BURGSTETTEN                   = "Auf den Rüdern Burgstall (Burgstetten)"
	AUF_DER_GANS_STUTTGART                                 = "Auf der Gans Stuttgart"
	AUFHAUSEN_AUFHAUSEN_GEISLINGEN                         = "Aufhausen Aufhausen (Geislingen)"
	AUFHAUSER_STR_NELLINGEN_UL                             = "Aufhauser Str. Nellingen (UL)"
	AUGSBURGER_PLATZ_STUTTGART                             = "Augsburger Platz Stuttgart"
	AUGUSTBEBELSTRASSE_EGLOSHEIM                           = "August-Bebel-Straße Eglosheim"
	AUGUSTBRANDLESTRASSE_FELLBACH                          = "August-Brändle-Straße Fellbach"
	AUGUSTINUM_STUTTGART                                   = "Augustinum Stuttgart"
	AUSSERE_HAUPTSTRASSE_GEISLINGEN_STEIGE                 = "Äußere Hauptstraße Geislingen (Steige)"
	AUSSIEDLERHOFE_MOGLINGEN                               = "Aussiedlerhöfe Möglingen"
	AUSSIEDLERHOFE_GROSSBOTTWAR                            = "Aussiedlerhöfe Großbottwar"
	AUSSIEDLERHOFE_HOFEN_BONNIGHEIM                        = "Aussiedlerhöfe Hofen (Bönnigheim)"
	AUSSIEDLERHOFE_HASLACH                                 = "Aussiedlerhöfe Haslach"
	AUSSIEDLERHOFE_STOTTEN                                 = "Aussiedlerhöfe Stötten"
	AUSSIEDLERHOFE_HEININGEN_GP                            = "Aussiedlerhöfe Heiningen (GP)"
	AUSTRASSE_GRUNBACH                                     = "Austraße Grunbach"
	AUSTRASSE_KIRCHHEIM_T                                  = "Austraße Kirchheim (T)"
	AUSTRASSE_VAIHINGEN_E                                  = "Austraße Vaihingen (E)"
	AUSTRASSE_WENDLINGEN_N                                 = "Austraße Wendlingen (N)"
	AUTALHALLE_BAD_UBERKINGEN                              = "Autalhalle Bad Überkingen"
	AUTENSTRASSE_DITZINGEN                                 = "Autenstraße Ditzingen"
	AUTMUT_TISCHARDT                                       = "Autmut Tischardt"
	AUWIESEN_STUTTGART                                     = "Auwiesen Stuttgart"
	AUWIESENBRUCKE_BIETIGHEIM                              = "Auwiesenbrücke Bietigheim"
	AUWIESENSCHULE_NECKARTENZLINGEN                        = "Auwiesenschule Neckartenzlingen"
	AVE_MARIA_DEGGINGEN                                    = "Ave Maria Deggingen"
	B_297HOHE_LINDE_KIRNECK                                = "B 297/Hohe Linde Kirneck"
	B_297KORNSTR_OBERKIRNECK                               = "B 297/Kornstr. Oberkirneck"
	B14_MAINHARDT                                          = "B14 Mainhardt"
	B28_BERNECK                                            = "B28 Berneck"
	B312_NECKARTAILFINGEN                                  = "B312 Neckartailfingen"
	B466_HAUSEN_BAD_UBERKINGEN                             = "B466 Hausen (Bad Überkingen)"
	BAACH_SCHNAIT                                          = "Baach Schnait"
	BAACHER_STRASSE_HOFEN                                  = "Baacher Straße Höfen"
	BAACHER_WEG_HERTMANNSWEILER                            = "Baacher Weg Hertmannsweiler"
	BACHBRUCKE_SCHLIERBACH                                 = "Bachbrücke Schlierbach"
	BACHHALDE_STUTTGART                                    = "Bachhalde Stuttgart"
	BACHLENWEG_STUTTGART                                   = "Bächlenweg Stuttgart"
	BACHSTRASSE_SINDELFINGEN                               = "Bachstraße Sindelfingen"
	BACHSTRASSE_ALTDORF_BB                                 = "Bachstraße Altdorf (BB)"
	BACHSTRASSE_GROSSBOTTWAR                               = "Bachstraße Großbottwar"
	BACHSTRASSE_WEISSACH_BB                                = "Bachstraße Weissach (BB)"
	BACHSTRASSE_SCHLAT                                     = "Bachstraße Schlat"
	BACKHAUS_DACHTEL                                       = "Backhaus Dachtel"
	BACKHAUS_OBERBERKEN                                    = "Backhaus Oberberken"
	BACKHAUS_GROSSHEPPACH                                  = "Backhaus Großheppach"
	BACKHAUS_ESCHENBACH                                    = "Backhaus Eschenbach"
	BACKHAUS_OBERWALDEN                                    = "Backhaus Oberwälden"
	BACKHAUSLE_HANWEILER                                   = "Backhäusle Hanweiler"
	BACKNANG_BACKNANG                                      = "Backnang Backnang"
	BACKNANGER_STRASSE_WINNENDEN                           = "Backnanger Straße Winnenden"
	BACKNANGER_STRASSE_RUDERSBERG                          = "Backnanger Straße Rudersberg"
	BAD_LEINFELDEN                                         = "Bad Leinfelden"
	BAD_BAD_UBERKINGEN                                     = "Bad Bad Überkingen"
	BAD_CANNSTATT_STUTTGART                                = "Bad Cannstatt Stuttgart"
	BAD_CANNSTATT_WILHELMSPLATZ_STUTTGART                  = "Bad Cannstatt Wilhelmsplatz Stuttgart"
	BADBERKAWEG_MARBACH_N                                  = "Bad-Berka-Weg Marbach (N)"
	BADEPARK_BIETIGHEIM                                    = "Badepark Bietigheim"
	BADEZENTRUM_SINDELFINGEN                               = "Badezentrum Sindelfingen"
	BADSTRASSE_ALTBACH                                     = "Badstraße Altbach"
	BADSTRASSE_LIPPOLDSWEILER                              = "Badstraße Lippoldsweiler"
	BADWIESEN_KIRCHHEIM_T                                  = "Badwiesen Kirchheim (T)"
	BAHNBRUCKE_BONDORF                                     = "Bahnbrücke Bondorf"
	BAHNHOF_BERNHAUSEN                                     = "Bahnhof Bernhausen"
	BAHNHOF_HOLZMADEN                                      = "Bahnhof Holzmaden"
	BAHNHOF_STEINHEIM_M                                    = "Bahnhof Steinheim (M)"
	BAHNHOF_DONZDORF                                       = "Bahnhof Donzdorf"
	BAHNHOF_ROHRDORF_CW                                    = "Bahnhof Rohrdorf (CW)"
	BAHNHOF_ARENA_LUDWIGSBURG                              = "Bahnhof (Arena) Ludwigsburg"
	BAHNHOF_DAMMSTRASSE_WAIBLINGEN                         = "Bahnhof (Dammstraße) Waiblingen"
	BAHNHOF_FLUGFELD_BOBLINGEN                             = "Bahnhof (Flugfeld) Böblingen"
	BAHNHOF_INDUSTRIESTR_RENNINGEN                         = "Bahnhof (Industriestr.) Renningen"
	BAHNHOF_PLOCHINGER_STRASSE_NURTINGEN                   = "Bahnhof (Plochinger Straße) Nürtingen"
	BAHNHOF_POST_NURTINGEN                                 = "Bahnhof (Post) Nürtingen"
	BAHNHOF_SCHLUSSELACKER_MAICHINGEN                      = "Bahnhof (Schlüsseläcker) Maichingen"
	BAHNHOF_SUD_OBERESSLINGEN                              = "Bahnhof (Süd) Oberesslingen"
	BAHNHOF_TROLLINGERWEG_METTINGEN                        = "Bahnhof (Trollingerweg) Mettingen"
	BAHNHOF_UHLANDSTRASSE_FRICKENHAUSEN                    = "Bahnhof (Uhlandstraße) Frickenhausen"
	BAHNHOF_WARTHSTRASSE_KORNTAL                           = "Bahnhof (Warthstraße) Korntal"
	BAHNHOFSALLEE_BAD_BOLL                                 = "Bahnhofsallee Bad Boll"
	BAHNHOFSSTEG_GOPPINGEN                                 = "Bahnhofssteg Göppingen"
	BAHNHOFSTRASSE_NEUHAUSEN_F                             = "Bahnhofstraße Neuhausen (F)"
	BAHNHOFSTRASSE_NURTINGEN                               = "Bahnhofstraße Nürtingen"
	BAHNHOFSTRASSE_MOGLINGEN                               = "Bahnhofstraße Möglingen"
	BAHNHOFSTRASSE_ERDMANNHAUSEN                           = "Bahnhofstraße Erdmannhausen"
	BAHNHOFSTRASSE_WENDLINGEN_N                            = "Bahnhofstraße Wendlingen (N)"
	BAHNHOFSTRASSE_UNTERLENNINGEN                          = "Bahnhofstraße Unterlenningen"
	BAHNHOFSTRASSE_WEILHEIM_T                              = "Bahnhofstraße Weilheim (T)"
	BAHNHOFSTRASSE_WEISSACH_BB                             = "Bahnhofstraße Weissach (BB)"
	BAHNHOFSTRASSE_HEMMINGEN                               = "Bahnhofstraße Hemmingen"
	BAHNHOFSTRASSE_SCHONAICH                               = "Bahnhofstraße Schönaich"
	BAHNHOFSTRASSE_PLUDERHAUSEN                            = "Bahnhofstraße Plüderhausen"
	BAHNHOFSTRASSE_NELLMERSBACH                            = "Bahnhofstraße Nellmersbach"
	BAHNHOFSTRASSE_ENDERSBACH                              = "Bahnhofstraße Endersbach"
	BAHNHOFSTRASSE_SERSHEIM                                = "Bahnhofstraße Sersheim"
	BAHNHOFSTRASSE_GINGEN_F                                = "Bahnhofstraße Gingen (F)"
	BAHNHOFSTRASSE_MUHLHAUSEN_IM_TALE                      = "Bahnhofstraße Mühlhausen im Täle"
	BAHNHOFSTRASSE_BIRENBACH                               = "Bahnhofstraße Birenbach"
	BAHNHOFSTRASSE_WASCHENBEUREN                           = "Bahnhofstraße Wäschenbeuren"
	BAHNUBERFUHRUNG_GERADSTETTEN                           = "Bahnüberführung Geradstetten"
	BAHNUNTERFUHRUNG_WINTERBACH                            = "Bahnunterführung Winterbach"
	BAJASTRASSE_WAIBLINGEN                                 = "Bajastraße Waiblingen"
	BALZHOLZ_RATHAUS_BEUREN                                = "Balzholz Rathaus Beuren"
	BALZHOLZER_STRASSE_BEUREN                              = "Balzholzer Straße Beuren"
	BANK_NECKARGRONINGEN                                   = "Bank Neckargröningen"
	BANNHALDE_GROSSSACHSENHEIM                             = "Bannhalde Großsachsenheim"
	BANNMUHLE_REICHENBACH_F                                = "Bannmühle Reichenbach (F)"
	BANRAIN_URBACH                                         = "Banrain Urbach"
	BARBAROSSATHERMEN_GOPPINGEN                            = "Barbarossa-Thermen Göppingen"
	BARBAROSSASEE_GOPPINGEN                                = "Barbarossasee Göppingen"
	BARENWIESEN_SERACH                                     = "Bärenwiesen Serach"
	BARTENBACH_SULZBACH_M                                  = "Bartenbach Sulzbach (M)"
	BARTENWEG_SINDELFINGEN                                 = "Bartenweg Sindelfingen"
	BAUERNBRESLAUER_STRASSE_DITZINGEN                      = "Bauern-/Breslauer Straße Ditzingen"
	BAUHOF_NURTINGEN                                       = "Bauhof Nürtingen"
	BAUHOF_GOPPINGEN                                       = "Bauhof Göppingen"
	BAUMGARTEN_SALACH                                      = "Baumgarten Salach"
	BAUMGARTENWEG_STUTTGART                                = "Baumgartenweg Stuttgart"
	BAUMWASENSTRASSE_SCHORNDORF                            = "Baumwasenstraße Schorndorf"
	BAUSCHE_WELZHEIM                                       = "Bausche Welzheim"
	BBW_SCHELMENHOLZ                                       = "BBW Schelmenholz"
	BEBELSTRASSE_GEISLINGEN_STEIGE                         = "Bebelstraße Geislingen (Steige)"
	BEERHALDENSTRASSE_ENZWEIHINGEN                         = "Beerhaldenstraße Enzweihingen"
	BEETHOVENSTRASSE_STUTTGART                             = "Beethovenstraße Stuttgart"
	BEETHOVENSTRASSE_PLOCHINGEN                            = "Beethovenstraße Plochingen"
	BEETHOVENSTRASSE_MERKLINGEN                            = "Beethovenstraße Merklingen"
	BEETHOVENSTRASSE_NEUWEILER                             = "Beethovenstraße Neuweiler"
	BEETHOVENSTRASSE_UNTERWEISSACH                         = "Beethovenstraße Unterweissach"
	BEETHOVENSTRASSE_HERRENBERG                            = "Beethovenstraße Herrenberg"
	BEETHOVENSTRASSE_BITTENFELD                            = "Beethovenstraße Bittenfeld"
	BEETHOVENSTRASSE_LORCH                                 = "Beethovenstraße Lorch"
	BEETHOVENSTRASSE_35_LORCH                              = "Beethovenstraße 35 Lorch"
	BEETHOVENSTRASSE_ABZW_LORCH                            = "Beethovenstraße Abzw. Lorch"
	BEIHINGER_PLATZ_FREIBERG_N                             = "Beihinger Platz Freiberg (N)"
	BEIHINGER_STRASSE_HOHENECK                             = "Beihinger Straße Hoheneck"
	BEIHINGER_STRASSE_PLEIDELSHEIM                         = "Beihinger Straße Pleidelsheim"
	BEILSTEINER_STRASSE_OBERSTENFELD                       = "Beilsteiner Straße Oberstenfeld"
	BEILSTEINER_STRASSE_STUTTGART                          = "Beilsteiner Straße Stuttgart"
	BEIM_KREISEL_ENDERSBACH                                = "Beim Kreisel Endersbach"
	BEIM_WASSERTURM_WAIBLINGEN                             = "Beim Wasserturm Waiblingen"
	BEIM_WASSERTURM_BACKNANG                               = "Beim Wasserturm Backnang"
	BELFORTER_PLATZ_LEONBERG                               = "Belforter Platz Leonberg"
	BEMPFLINGEN_BEMPFLINGEN                                = "Bempflingen Bempflingen"
	BENNINGEN_N_BENNINGEN_N                                = "Benningen (N) Benningen (N)"
	BENZACH_BEUTELSBACH                                    = "Benzach Beutelsbach"
	BENZACH_BILDUNGSZENTRUM_BEUTELSBACH                    = "Benzach Bildungszentrum Beutelsbach"
	BENZENHOFWEG_KIRCHHEIM_T                               = "Benzenhofweg Kirchheim (T)"
	BENZSTRASSE_URBACH                                     = "Benzstraße Urbach"
	BENZSTRASSE_HERRENBERG                                 = "Benzstraße Herrenberg"
	BERGFRIEDHOF_STUTTGART                                 = "Bergfriedhof Stuttgart"
	BERGHAUSTRASSE_STUTTGART                               = "Berghaustraße Stuttgart"
	BERGHEIMER_HOF_STUTTGART                               = "Bergheimer Hof Stuttgart"
	BERGSIEDLUNG_DACHTEL                                   = "Bergsiedlung Dachtel"
	BERGSIEDLUNG_KEPLERSTRASSE_DACHTEL                     = "Bergsiedlung (Keplerstraße) Dachtel"
	BERGSTAFFELSTRASSE_STUTTGART                           = "Bergstaffelstraße Stuttgart"
	BERGSTRASSE_KORNTAL                                    = "Bergstraße Korntal"
	BERGSTRASSE_SULZGRIES                                  = "Bergstraße Sulzgries"
	BERGSTRASSE_WERNAU_N                                   = "Bergstraße Wernau (N)"
	BERGSTRASSE_RECHBERGHAUSEN                             = "Bergstraße Rechberghausen"
	BERGSTRASSE_LORCH                                      = "Bergstraße Lorch"
	BERGWALD_GECHINGEN                                     = "Bergwald Gechingen"
	BERLINER_PLATZ_LUDWIGSBURG                             = "Berliner Platz Ludwigsburg"
	BERLINER_PLATZ_HOHE_STRASSE_STUTTGART                  = "Berliner Platz (Hohe Straße) Stuttgart"
	BERLINER_PLATZ_LIEDERHALLE_STUTTGART                   = "Berliner Platz (Liederhalle) Stuttgart"
	BERLINER_RING_BACKNANG                                 = "Berliner Ring Backnang"
	BERLINER_STRASSE_MAICHINGEN                            = "Berliner Straße Maichingen"
	BERLINER_STRASSE_HERRENBERG                            = "Berliner Straße Herrenberg"
	BERLINER_STRASSE_BOBLINGEN                             = "Berliner Straße Böblingen"
	BERLINER_STRASSE_ALDINGEN                              = "Berliner Straße Aldingen"
	BERLINER_STRASSE_ASPERG                                = "Berliner Straße Asperg"
	BERNHALDEN_OPPENWEILER                                 = "Bernhalden Oppenweiler"
	BERNHAUSER_STRASSE_STUTTGART                           = "Bernhäuser Straße Stuttgart"
	BERNHAUSER_WEG_NEUHAUSEN_F                             = "Bernhäuser Weg Neuhausen (F)"
	BERNSTEINSTRASSE_STUTTGART                             = "Bernsteinstraße Stuttgart"
	BERTHABENZSTRASSE_ENSINGEN                             = "Bertha-Benz-Straße Ensingen"
	BERUFLICHES_SCHULZENTRUM_ZELL_ESSLINGEN                = "Berufliches Schulzentrum Zell (Esslingen)"
	BERUFSBILDUNGSZENTRUM_STUTTGART                        = "Berufsbildungszentrum Stuttgart"
	BERUFSSCHULE_PATTONVILLE                               = "Berufsschule Pattonville"
	BERUFSSCHULZENTRUM_WAIBLINGEN                          = "Berufsschulzentrum Waiblingen"
	BERUFSSCHULZENTRUM_GEISLINGEN_STEIGE                   = "Berufsschulzentrum Geislingen (Steige)"
	BERUFSSCHULZENTRUM_ODE_GOPPINGEN                       = "Berufsschulzentrum Öde Göppingen"
	BERWINKEL_SULZBACH_M                                   = "Berwinkel Sulzbach (M)"
	BESIGHEIM_BESIGHEIM                                    = "Besigheim Besigheim"
	BESIGHEIMER_STRASSE_EGLOSHEIM                          = "Besigheimer Straße Eglosheim"
	BESIGHEIMER_STRASSE_LOCHGAU                            = "Besigheimer Straße Löchgau"
	BESIGHEIMER_STRASSE_FREUDENTAL                         = "Besigheimer Straße Freudental"
	BESIGHEIMER_STRASSE_OTTMARSHEIM                        = "Besigheimer Straße Ottmarsheim"
	BESKIDENSTRASSE_STUTTGART                              = "Beskidenstraße Stuttgart"
	BETHEL_WELZHEIM                                        = "Bethel Welzheim"
	BETRIEBSHOF_WARMBRONN                                  = "Betriebshof Warmbronn"
	BETTACKER_NEBRINGEN                                    = "Bettäcker Nebringen"
	BETTLINGER_FORUM_GROSSBETTLINGEN                       = "Bettlinger Forum Großbettlingen"
	BETZ_KONGEN                                            = "Betz Köngen"
	BETZSTRASSE_GOPPINGEN                                  = "Betzstraße Göppingen"
	BEUNDSTRASSE_EISLINGEN_F                               = "Beundstraße Eislingen (F)"
	BEURENER_STEIGE_BEUREN                                 = "Beurener Steige Beuren"
	BEUTELSBACH_BEUTELSBACH                                = "Beutelsbach Beutelsbach"
	BEZIRKSAMT_BEZGENRIET                                  = "Bezirksamt Bezgenriet"
	BEZIRKSAMT_WISSGOLDINGEN                               = "Bezirksamt Wißgoldingen"
	BIEGEL_BACKNANG                                        = "Biegel Backnang"
	BIELER_WEG_STUTTGART                                   = "Bieler Weg Stuttgart"
	BIETIGHEIM_BIETIGHEIM                                  = "Bietigheim Bietigheim"
	BIETIGHEIMER_STRASSE_LUDWIGSBURG                       = "Bietigheimer Straße Ludwigsburg"
	BIETIGHEIMER_STRASSE_GROSSINGERSHEIM                   = "Bietigheimer Straße Großingersheim"
	BIETIGHEIMER_STRASSE_KLEINSACHSENHEIM                  = "Bietigheimer Straße Kleinsachsenheim"
	BIETIGHEIMER_STRASSE_METTERZIMMERN                     = "Bietigheimer Straße Metterzimmern"
	BIETIGHEIMER_STRASSE_MURR                              = "Bietigheimer Straße Murr"
	BIHLPLATZ_STUTTGART                                    = "Bihlplatz Stuttgart"
	BILDACKERSTRASSE_HOHENACKER                            = "Bildäckerstraße Hohenacker"
	BILDSTOCKLE_LEONBERG                                   = "Bildstöckle Leonberg"
	BILDUNGSCAMPUS_GROSSSACHSENHEIM                        = "Bildungscampus Großsachsenheim"
	BILDUNGSZENTRUM_WEILHEIM_T                             = "Bildungszentrum Weilheim (T)"
	BILDUNGSZENTRUM_MARKGRONINGEN                          = "Bildungszentrum Markgröningen"
	BILDUNGSZENTRUM_COTTENWEILER                           = "Bildungszentrum Cottenweiler"
	BILDUNGSZENTRUM_II_WINNENDEN                           = "Bildungszentrum II Winnenden"
	BILDUNGSZENTRUM_WEST_LUDWIGSBURG                       = "Bildungszentrum West Ludwigsburg"
	BILFINGER_STRASSE_FREIBERG_N                           = "Bilfinger Straße Freiberg (N)"
	BINSENSTRASSE_SCHALKSTETTEN                            = "Binsenstraße Schalkstetten"
	BIRKACH_FRIEDHOF_STUTTGART                             = "Birkach Friedhof Stuttgart"
	BIRKACH_WEST_STUTTGART                                 = "Birkach West Stuttgart"
	BIRKACHER_STRASSE_STUTTGART                            = "Birkacher Straße Stuttgart"
	BIRKENALLEE_BAHNUNTERFUHRUNG_PLUDERHAUSEN              = "Birkenallee (Bahnunterführung) Plüderhausen"
	BIRKENKOPF_STUTTGART                                   = "Birkenkopf Stuttgart"
	BIRKENSTRASSE_HOLZGERLINGEN                            = "Birkenstraße Holzgerlingen"
	BIRKENWEG_WAIBLINGEN                                   = "Birkenweg Waiblingen"
	BIRKENWEG_ALTBACH                                      = "Birkenweg Altbach"
	BIRKENWEG_BISSINGEN_LB                                 = "Birkenweg Bissingen (LB)"
	BIRKENWEG_GEISLINGEN_STEIGE                            = "Birkenweg Geislingen (Steige)"
	BIRKENWEISSBUCH_VORDERWEISSBUCH                        = "Birkenweißbuch Vorderweißbuch"
	BIRKHAU_AFFALTERBACH                                   = "Birkhau Affalterbach"
	BIRKHECKENSTRASSE_STUTTGART                            = "Birkheckenstraße Stuttgart"
	BIRKHOF_KAISERSBACH                                    = "Birkhof Kaisersbach"
	BIRKHOF_DEGGINGEN                                      = "Birkhof Deggingen"
	BIRKHOFE_ROSSWALDEN                                    = "Birkhöfe Roßwälden"
	BISMARCKPLATZ_STUTTGART_WEST                           = "Bismarckplatz Stuttgart West"
	BISMARCKSTAFFEL_STUTTGART                              = "Bismarckstaffel Stuttgart"
	BISMARCKSTRASSE_ESSLINGEN_N                            = "Bismarckstraße Esslingen (N)"
	BISMARCKSTRASSE_KIRCHHEIM_T                            = "Bismarckstraße Kirchheim (T)"
	BISSINGER_STRASSE_EGLOSHEIM                            = "Bissinger Straße Eglosheim"
	BISSINGER_STRASSE_UNTERMBERG                           = "Bissinger Straße Untermberg"
	BISSINGER_STRASSE_OCHSENWANG                           = "Bissinger Straße Ochsenwang"
	BITTENFELDER_STRASSE_HOHENACKER                        = "Bittenfelder Straße Hohenacker"
	BLAMMERBERGSTRASSE_WEIL_DER_STADT                      = "Blammerbergstraße Weil der Stadt"
	BLATTERT_MURR                                          = "Blattert Murr"
	BLEICHEREI_UHINGEN                                     = "Bleicherei Uhingen"
	BLEICHSTRASSE_GOPPINGEN                                = "Bleichstraße Göppingen"
	BLICK_STUTTGART                                        = "Blick Stuttgart"
	BLOSENBERGKIRCHE_LEONBERG                              = "Blosenbergkirche Leonberg"
	BLUHENDES_BAROCK_LUDWIGSBURG                           = "Blühendes Barock Ludwigsburg"
	BLUMENWILHELMSTRASSE_REICHENBACH_F                     = "Blumen-/Wilhelmstraße Reichenbach (F)"
	BLUMENSTRASSE_BACKNANG                                 = "Blumenstraße Backnang"
	BLUMENSTRASSE_UNTERENSINGEN                            = "Blumenstraße Unterensingen"
	BLUMENSTRASSE_REICHENBACH_F                            = "Blumenstraße Reichenbach (F)"
	BLUMENSTRASSE_WAIBLINGEN                               = "Blumenstraße Waiblingen"
	BLUMENSTRASSE_KLEININGERSHEIM                          = "Blumenstraße Kleiningersheim"
	BOBLINGEN_BOBLINGEN                                    = "Böblingen Böblingen"
	BODELSCHWINGHSCHULE_MURRHARDT                          = "Bodelschwinghschule Murrhardt"
	BOEHRINGER_AREAL_GOPPINGEN                             = "Boehringer Areal Göppingen"
	BOHNAU_KIRCHHEIM_T                                     = "Bohnau Kirchheim (T)"
	BOHNAUHAUS_KIRCHHEIM_T                                 = "Bohnauhaus Kirchheim (T)"
	BOHNREUTE_KIRCHENKIRNBERG                              = "Bohnreute Kirchenkirnberg"
	BOHRINGSWEILER_GROSSERLACH                             = "Böhringsweiler Großerlach"
	BOLLER_STRASSE_DURNAU                                  = "Boller Straße Dürnau"
	BOLZPLATZ_HOHENGEHREN                                  = "Bolzplatz Hohengehren"
	BOLZSTRASSE_KORNWESTHEIM                               = "Bolzstraße Kornwestheim"
	BOLZSTRASSE_BIETIGHEIM                                 = "Bolzstraße Bietigheim"
	BOMBACH_AICH                                           = "Bombach Aich"
	BONDORF_BONDORF                                        = "Bondorf Bondorf"
	BONHOLZ_WALDENBUCH                                     = "Bonholz Waldenbuch"
	BONHOLZ_ALFDORF                                        = "Bonholz Alfdorf"
	BONHOLZ_ABZWEIG_ALFDORF                                = "Bonholz Abzweig Alfdorf"
	BONLANDENER_STRASSE_BONLANDEN                          = "Bonlandener Straße Bonlanden"
	BONLANDER_STRASSE_ECHTERDINGEN                         = "Bonländer Straße Echterdingen"
	BOPSER_STUTTGART                                       = "Bopser Stuttgart"
	BOPSERWALDSTRASSE_GERLINGEN                            = "Bopserwaldstraße Gerlingen"
	BORKUMSTRASSE_NEUWSIEDL_STUTTGART                      = "Borkumstraße (Neuw.-siedl.) Stuttgart"
	BORSENPLATZ_STUTTGART                                  = "Börsenplatz Stuttgart"
	BORSIGSTRASSE_STUTTGART                                = "Borsigstraße Stuttgart"
	BORSIGSTRASSE_BISSINGEN_LB                             = "Borsigstraße Bissingen (LB)"
	BOSCH_RUTESHEIM                                        = "Bosch Rutesheim"
	BOSCH_TOR_III_SCHWIEBERDINGEN                          = "Bosch Tor III Schwieberdingen"
	BOSCH_TOR_IV_SCHWIEBERDINGEN                           = "Bosch Tor IV Schwieberdingen"
	BOSCHHOF_TURKHEIM                                      = "Boschhof Türkheim"
	BOSKOPWEG_WAIBLINGEN                                   = "Boskopweg Waiblingen"
	BOTNANG_STUTTGART                                      = "Botnang Stuttgart"
	BOTNANG_FREIBAD_STUTTGART                              = "Botnang Freibad Stuttgart"
	BOTNANGER_SATTEL_STUTTGART                             = "Botnanger Sattel Stuttgart"
	BOTTROPER_STRASSE_STUTTGART                            = "Bottroper Straße Stuttgart"
	BOTTWARTALSTRASSE_HOHENECK                             = "Bottwartalstraße Hoheneck"
	BRAHMSSTRASSE_KIRCHHEIM_T                              = "Brahmsstraße Kirchheim (T)"
	BRANDENBURGER_STRASSE_OBERESSLINGEN                    = "Brandenburger Straße Oberesslingen"
	BRANDHOF_L_1080_GSCHWEND                               = "Brandhof L 1080 Gschwend"
	BRAUEREI_SCHWIEBERDINGEN                               = "Brauerei Schwieberdingen"
	BRAUEREIPLATZ_MAGSTADT                                 = "Brauereiplatz Magstadt"
	BRAUNGARTWEG_ZOLLBERG                                  = "Braungartweg Zollberg"
	BRECH_PFAHLBRONN                                       = "Brech Pfahlbronn"
	BRECH_KAISERLINDE_PFAHLBRONN                           = "Brech Kaiserlinde Pfahlbronn"
	BREECH_KRZG_RATTENHARZ_BORTLINGEN                      = "Breech Krzg. Rattenharz Börtlingen"
	BREECH_WENDEPLATTE_BORTLINGEN                          = "Breech Wendeplatte Börtlingen"
	BREITENF_REINHOLDMAIERPLATZ_WELZHEIM                   = "Breitenf. Reinhold-Maier-Platz Welzheim"
	BREITENFURST_LORCHER_STRASSE_WELZHEIM                  = "Breitenfürst Lorcher Straße Welzheim"
	BREITENFURST_SCHULHAUS_WELZHEIM                        = "Breitenfürst Schulhaus Welzheim"
	BREITENSTEINER_STRASSE_BOBLINGEN                       = "Breitensteiner Straße Böblingen"
	BREITER_WEG_BOHMENKIRCH                                = "Breiter Weg Böhmenkirch"
	BREITER_WEG_EBHAUSEN                                   = "Breiter Weg Ebhausen"
	BREITESTRASSE_HEININGEN_GP                             = "Breitestraße Heiningen (GP)"
	BREITFELDERHOF_OTTENBACH                               = "Breitfelderhof Ottenbach"
	BREITWIESEN_GERLINGEN                                  = "Breitwiesen Gerlingen"
	BREITWIESENHAUS_GERLINGEN                              = "Breitwiesenhaus Gerlingen"
	BREND_PFAHLBRONN                                       = "Brend Pfahlbronn"
	BRENDLE_GROSSMARKT_STUTTGART                           = "Brendle (Großmarkt) Stuttgart"
	BRENNERSTRASSE_ELTINGEN                                = "Brennerstraße Eltingen"
	BRESLAUER_STRASSE_BOBLINGEN                            = "Breslauer Straße Böblingen"
	BRESLAUER_STRASSE_SCHWIEBERDINGEN                      = "Breslauer Straße Schwieberdingen"
	BRESLAUER_STRASSE_BEILSTEIN                            = "Breslauer Straße Beilstein"
	BRESLAUER_STRASSEKINDERGARTEN_DITZINGEN                = "Breslauer Straße/Kindergarten Ditzingen"
	BREUNINGERLAND_LUDWIGSBURG                             = "Breuningerland Ludwigsburg"
	BREUNINGERLAND_SINDELFINGEN                            = "Breuningerland Sindelfingen"
	BRIEFZENTRUM_SALACH                                    = "Briefzentrum Salach"
	BROMBERGERHOFE_HOHENHASLACH                            = "Brombergerhöfe Hohenhaslach"
	BRONNTOR_HERRENBERG                                    = "Bronntor Herrenberg"
	BRONNWIESENWEG_RUDERSBERG                              = "Bronnwiesenweg Rudersberg"
	BRUCH_BRUCH_WEISSACH                                   = "Bruch Bruch (Weissach)"
	BRUCH_KAISERSBACH                                      = "Bruch Kaisersbach"
	BRUCK_LORCH                                            = "Bruck Lorch"
	BRUCKEN_UNTERLENNINGEN                                 = "Brucken Unterlenningen"
	BRUCKEN_UNTERE_STRASSE_UNTERLENNINGEN                  = "Brucken Untere Straße Unterlenningen"
	BRUCKENSTRASSE_WERNAU_N                                = "Bruckenstraße Wernau (N)"
	BRUCKENSTRASSE_HESSIGHEIM                              = "Brückenstraße Hessigheim"
	BRUCKENSTRASSE_ALDINGEN                                = "Brückenstraße Aldingen"
	BRUCKENSTRASSE_WINNENDEN                               = "Brückenstraße Winnenden"
	BRUCKENSTRASSE_BIRENBACH                               = "Brückenstraße Birenbach"
	BRUCKENSTRASSE_GINGEN_F                                = "Brückenstraße Gingen (F)"
	BRUCKENSTRASSE_GOPPINGEN                               = "Brückenstraße Göppingen"
	BRUCKENWEG_GRUIBINGEN                                  = "Brückenweg Gruibingen"
	BRUCKLE_WALDENBRONN                                    = "Brückle Wäldenbronn"
	BRUCKWIESEN_GEISLINGEN_STEIGE                          = "Bruckwiesen Geislingen (Steige)"
	BRUCKWIESEN_SPIELPLATZ_HATTENHOFEN                     = "Bruckwiesen Spielplatz Hattenhofen"
	BRUCKWIESENSTRASSE_HATTENHOFEN                         = "Bruckwiesenstraße Hattenhofen"
	BRUDERHAUS_STUTTGART                                   = "Bruderhaus Stuttgart"
	BRUHL_WERNAU_N                                         = "Brühl Wernau (N)"
	BRUHL_LINDORF                                          = "Brühl Lindorf"
	BRUHL_FREILICHTMUSEUM_BEUREN                           = "Brühl (Freilichtmuseum) Beuren"
	BRUHLSIEDLUNG_NEUHAUSEN_F                              = "Brühlsiedlung Neuhausen (F)"
	BRUHLSTRASSE_MOGLINGEN                                 = "Brühlstraße Möglingen"
	BRUHLSTRASSE_GROSSINGERSHEIM                           = "Brühlstraße Großingersheim"
	BRUHLSTRASSE_DECKENPFRONN                              = "Brühlstraße Deckenpfronn"
	BRUHLSTRASSE_MOTZINGEN                                 = "Brühlstraße Mötzingen"
	BRUHLSTRASSE_22_MAGSTADT                               = "Brühlstraße 22 Magstadt"
	BRUHLWEG_SALACH                                        = "Brühlweg Salach"
	BRUNNEN_WALDENBRONN                                    = "Brunnen Wäldenbronn"
	BRUNNENGARTEN_WIESENSTEIG                              = "Brunnengarten Wiesensteig"
	BRUNNENSTRASSE_WELZHEIM                                = "Brunnenstraße Welzheim"
	BRUNNENSTRASSE_GINGEN_F                                = "Brunnenstraße Gingen (F)"
	BRUNNENSTRASSE_BARTENBACH_GOPPINGEN                    = "Brunnenstraße Bartenbach (Göppingen)"
	BRUNNENWIESEN_SPIELBERG                                = "Brunnenwiesen Spielberg"
	BRUNNER_STRASSE_LUDWIGSBURG                            = "Brünner Straße Ludwigsburg"
	BUBENBAD_STUTTGART                                     = "Bubenbad Stuttgart"
	BUCH_HOLZGERLINGEN                                     = "Buch Holzgerlingen"
	BUCH_BERLINER_STRASSE_BIETIGHEIM                       = "Buch Berliner Straße Bietigheim"
	BUCH_BRESLAUER_STRASSE_BIETIGHEIM                      = "Buch Breslauer Straße Bietigheim"
	BUCH_BUCHSCHULE_BIETIGHEIM                             = "Buch Buchschule Bietigheim"
	BUCH_BUCHZENTRUM_BIETIGHEIM                            = "Buch Buchzentrum Bietigheim"
	BUCH_DRESDNER_STRASSE_BIETIGHEIM                       = "Buch Dresdner Straße Bietigheim"
	BUCH_GRONINGER_WEG_BIETIGHEIM                          = "Buch Gröninger Weg Bietigheim"
	BUCH_POSENER_STRASSE_BIETIGHEIM                        = "Buch Posener Straße Bietigheim"
	BUCH_SUCYSTRASSE_BIETIGHEIM                            = "Buch Sucystraße Bietigheim"
	BUCHENBRONN_FORSTHAUS_EBERSBACH_F                      = "Büchenbronn Forsthaus Ebersbach (F)"
	BUCHENBRONN_GEMEINDEHAUS_EBERSBACH_F                   = "Büchenbronn Gemeindehaus Ebersbach (F)"
	BUCHENBRONN_KONIGSEICHE_EBERSBACH_F                    = "Büchenbronn Königseiche Ebersbach (F)"
	BUCHENBRONN_SIEDLUNG_EBERSBACH_F                       = "Büchenbronn Siedlung Ebersbach (F)"
	BUCHENBRONNER_STRASSE_EBERSBACH_F                      = "Büchenbronner Straße Ebersbach (F)"
	BUCHENGEHREN_PFAHLBRONN                                = "Buchengehren Pfahlbronn"
	BUCHENGEHREN_SAGMUHLE_PFAHLBRONN                       = "Buchengehren Sägmühle Pfahlbronn"
	BUCHENRAIN_OCHSENBACH                                  = "Buchenrain Ochsenbach"
	BUCHENSTRASSE_STEINACH_BERGLEN                         = "Buchenstraße Steinach (Berglen)"
	BUCHENSTRASSE_FAURNDAU                                 = "Buchenstraße Faurndau"
	BUCHENWEG_WAIBLINGEN                                   = "Buchenweg Waiblingen"
	BUCHENWEG_SCHORNDORF                                   = "Buchenweg Schorndorf"
	BUCHHALDENSCHULE_AIDLINGEN                             = "Buchhaldenschule Aidlingen"
	BUCHHALDENSTRASSE_SCHNAIT                              = "Buchhaldenstraße Schnait"
	BUCHRAIN_HOLZHEIM_GOPPINGEN                            = "Buchrain Holzheim (Göppingen)"
	BUCHRAIN_FRIEDHOF_STUTTGART                            = "Buchrain Friedhof Stuttgart"
	BUCHSENSTRASSE_STUTTGART                               = "Büchsenstraße Stuttgart"
	BUCHSTEINER_GINGEN_F                                   = "Buchsteiner Gingen (F)"
	BUCHWALD_STUTTGART                                     = "Buchwald Stuttgart"
	BUDAPESTER_PLATZ_STUTTGART                             = "Budapester Platz Stuttgart"
	BUHL_KORB                                              = "Bühl Korb"
	BUHLACKER_SCHONAICH                                    = "Bühläcker Schönaich"
	BUHLACKERSTRASSE_STETTEN_I_R                           = "Bühläckerstraße Stetten (i. R.)"
	BUHLBRONN_BUHLBRONN                                    = "Buhlbronn Buhlbronn"
	BUHLENSTRASSE_HOLZGERLINGEN                            = "Bühlenstraße Holzgerlingen"
	BUHLER_STRASSE_BOBLINGEN                               = "Bühler Straße Böblingen"
	BUHLFELD_OPPENWEILER                                   = "Bühlfeld Oppenweiler"
	BUHNERSTRASSE_SCHMIDEN                                 = "Bühnerstraße Schmiden"
	BULKESWEG_KIRCHHEIM_T                                  = "Bulkesweg Kirchheim (T)"
	BUNSENSTR_OBERENSINGEN                                 = "Bunsenstr. Oberensingen"
	BUNZWANGER_STRASSE_ALBERSHAUSEN                        = "Bünzwanger Straße Albershausen"
	BUNZWANGER_STRASSE_UHINGEN                             = "Bünzwanger Straße Uhingen"
	BUOCHER_STRASSE_BREUNINGSWEILER                        = "Buocher Straße Breuningsweiler"
	BUOLWEG_ENZWEIHINGEN                                   = "Buolweg Enzweihingen"
	BURG_ESSLINGEN_N                                       = "Burg Esslingen (N)"
	BURG_BURG                                              = "Bürg Bürg"
	BURGERHAUS_BEUREN                                      = "Bürgerhaus Beuren"
	BURGERHAUS_MAICHINGEN                                  = "Bürgerhaus Maichingen"
	BURGERHAUS_AICHELBERG_GP                               = "Bürgerhaus Aichelberg (GP)"
	BURGERHAUS_RSKN_RUDERN                                 = "Bürgerhaus RSKN Rüdern"
	BURGERMUHLE_BONNIGHEIM                                 = "Burgermühle Bönnigheim"
	BURGERSEE_KIRCHHEIM_T                                  = "Bürgersee Kirchheim (T)"
	BURGERZENTRUM_HEGENLOHE                                = "Bürgerzentrum Hegenlohe"
	BURGERZENTRUM_HALLENBAD_WAIBLINGEN                     = "Bürgerzentrum (Hallenbad) Waiblingen"
	BURGERZENTRUM_REMSBRUCKE_WAIBLINGEN                    = "Bürgerzentrum (Remsbrücke) Waiblingen"
	BURGGYMNASIUM_SCHORNDORF                               = "Burggymnasium Schorndorf"
	BURGHALDENFRIEDHOF_SINDELFINGEN                        = "Burghaldenfriedhof Sindelfingen"
	BURGHALDENSTRASSE_POPPENWEILER                         = "Burghaldenstraße Poppenweiler"
	BURGHOLZ_PFAHLBRONN                                    = "Burgholz Pfahlbronn"
	BURGHOLZHOF_STUTTGART                                  = "Burgholzhof Stuttgart"
	BURGKLINGE_GERLINGEN                                   = "Burgklinge Gerlingen"
	BURGPLATZ_BACKNANG                                     = "Burgplatz Backnang"
	BURGPLATZ_BONNIGHEIM                                   = "Burgplatz Bönnigheim"
	BURGSTALL_UHINGEN                                      = "Burgstall Uhingen"
	BURGSTALL_M_BURGSTALL_BURGSTETTEN                      = "Burgstall (M) Burgstall (Burgstetten)"
	BURGSTALLER_STRASSE_ERBSTETTEN                         = "Burgstaller Straße Erbstetten"
	BURGSTALLER_STRASSE_KIRCHBERG_M                        = "Burgstaller Straße Kirchberg (M)"
	BURGSTRASSE_ECHTERDINGEN                               = "Burgstraße Echterdingen"
	BURGSTRASSE_URBACH                                     = "Burgstraße Urbach"
	BURGUNDERSTRASSE_BEUTELSBACH                           = "Burgunderstraße Beutelsbach"
	BURGUNDERSTRASSE_METTINGEN                             = "Burgunderstraße Mettingen"
	BURGUNDERWEG_HANWEILER                                 = "Burgunderweg Hanweiler"
	BURKHARDSHOF_BIRKMANNSWEILER                           = "Burkhardshof Birkmannsweiler"
	BURKHARDTSMUHLE_WALDENBUCH                             = "Burkhardtsmühle Waldenbuch"
	BUSBAHNHOF_WELZHEIM                                    = "Busbahnhof Welzheim"
	BUSHOF_ABZWEIG_SULZBACH_M                              = "Bushof Abzweig Sulzbach (M)"
	BUSNAUER_PLATZ_STUTTGART                               = "Büsnauer Platz Stuttgart"
	BUSNAUER_STRASSE_41_WARMBRONN                          = "Büsnauer Straße 41 Warmbronn"
	BUSNAUER_STRASSE_6_WARMBRONN                           = "Büsnauer Straße 6 Warmbronn"
	BUSSBACHSTRASSE_STUTTGART                              = "Bußbachstraße Stuttgart"
	BUTZBERG_GRAB                                          = "Butzberg Grab"
	CAFE_ROMMEL_SCHWAIKHEIM                                = "Cafe Rommel Schwaikheim"
	CAFE_STOLL_OBERENSINGEN                                = "Cafe Stoll Oberensingen"
	CALWER_BRUCKE_SINDELFINGEN                             = "Calwer Brücke Sindelfingen"
	CALWER_STRASSE_SINDELFINGEN                            = "Calwer Straße Sindelfingen"
	CALWER_STRASSE_BOBLINGEN                               = "Calwer Straße Böblingen"
	CALWER_STRASSE_OBERJESINGEN                            = "Calwer Straße Oberjesingen"
	CALWER_STRASSE_HULB_BOBLINGEN                          = "Calwer Straße (Hulb) Böblingen"
	CANNSTATTER_PLATZ_FELLBACH                             = "Cannstatter Platz Fellbach"
	CANNSTATTER_STRASSE_METTINGEN                          = "Cannstatter Straße Mettingen"
	CANNSTATTER_STRASSE_ALDINGEN                           = "Cannstatter Straße Aldingen"
	CANNSTATTER_WASEN_STUTTGART                            = "Cannstatter Wasen Stuttgart"
	CARLBENZSTRASSE_OTTMARSHEIM                            = "Carl-Benz-Straße Ottmarsheim"
	CARLBENZSTRASSE_LAICHINGEN                             = "Carl-Benz-Straße Laichingen"
	CARLKAELBLESTRASSE_BACKNANG                            = "Carl-Kaelble-Straße Backnang"
	CARLSCHMINCKESTRASSE_ELTINGEN                          = "Carl-Schmincke-Straße Eltingen"
	CARLZEISSSTRASSE_HARTHAUSEN                            = "Carl-Zeiss-Straße Harthausen"
	CARLZEISSSTRASSE_GEBERSHEIM                            = "Carl-Zeiss-Straße Gebersheim"
	CARLZEISSSTRASSE_OTTMARSHEIM                           = "Carl-Zeiss-Straße Ottmarsheim"
	CHARLOTTENPLATZ_ESSLINGEN_N                            = "Charlottenplatz Esslingen (N)"
	CHARLOTTENPLATZ_STUTTGART                              = "Charlottenplatz Stuttgart"
	CHAUSSEEBERGSTRASSE_BESIGHEIM                          = "Chausseebergstraße Besigheim"
	CHRISTKONIGSKIRCHE_GOPPINGEN                           = "Christkönigskirche Göppingen"
	CHRISTOFSHOF_WALDHAUSEN_GEISLINGEN                     = "Christofshof Waldhausen (Geislingen)"
	CHRISTOFSTRASSE_WAIBLINGEN                             = "Christofstraße Waiblingen"
	CHRISTOPHSBAD_GOPPINGEN                                = "Christophsbad Göppingen"
	CHRISTOPHSTRASSE_BACKNANG                              = "Christophstraße Backnang"
	CHRISTOPHSTRASSE_GOPPINGEN                             = "Christophstraße Göppingen"
	COMBURGSTRASSE_LUDWIGSBURG                             = "Comburgstraße Ludwigsburg"
	CRONHUTTE_KAISERSBACH                                  = "Cronhütte Kaisersbach"
	DACHSWALD_STUTTGART                                    = "Dachswald Stuttgart"
	DACHSWALDWEG_STUTTGART                                 = "Dachswaldweg Stuttgart"
	DAFERN_LIPPOLDSWEILER                                  = "Däfern Lippoldsweiler"
	DAIMLERPLATZ_STUTTGART                                 = "Daimlerplatz Stuttgart"
	DAIMLERSTEG_SINDELFINGEN                               = "Daimlersteg Sindelfingen"
	DAIMLERSTRASSE_OEFFINGEN                               = "Daimlerstraße Oeffingen"
	DAIMLERSTRASSE_SCHONAICH                               = "Daimlerstraße Schönaich"
	DAIMLERSTRASSE_MOGLINGEN                               = "Daimlerstraße Möglingen"
	DAIMLERSTRASSE_LUDWIGSBURG                             = "Daimlerstraße Ludwigsburg"
	DAIMLERSTRASSE_HOLZGERLINGEN                           = "Daimlerstraße Holzgerlingen"
	DAIMLERSTRASSE_HERRENBERG                              = "Daimlerstraße Herrenberg"
	DAMMSTRASSE_NURTINGEN                                  = "Dammstraße Nürtingen"
	DANZIGER_PLATZ_WAIBLINGEN                              = "Danziger Platz Waiblingen"
	DANZIGER_STRASSE_MUNCHINGEN                            = "Danziger Straße Münchingen"
	DANZIGER_STRASSE_BOBLINGEN                             = "Danziger Straße Böblingen"
	DANZIGER_STRASSE_LUDWIGSBURG                           = "Danziger Straße Ludwigsburg"
	DANZIGER_STRASSE_KIRCHBERG_M                           = "Danziger Straße Kirchberg (M)"
	DANZIGER_STRASSE_EISLINGEN_F                           = "Danziger Straße Eislingen (F)"
	DANZIGER_WEG_SERSHEIM                                  = "Danziger Weg Sersheim"
	DATZINGER_STRASSE_SCHAFHAUSEN                          = "Dätzinger Straße Schafhausen"
	DAUERNBERG_SPIEGELBERG                                 = "Dauernberg Spiegelberg"
	DAUERNBERG_ABZWEIG_SPIEGELBERG                         = "Dauernberg Abzweig Spiegelberg"
	DEGERLOCH_STUTTGART                                    = "Degerloch Stuttgart"
	DEGERLOCH_ALBSTRASSE_STUTTGART                         = "Degerloch Albstraße Stuttgart"
	DEGERLOCH_EPPLESTRASSE_STUTTGART                       = "Degerloch Epplestraße Stuttgart"
	DEGERLOCH_ZOB_STUTTGART                                = "Degerloch ZOB Stuttgart"
	DENKENDORFER_STRASSE_NELLINGEN                         = "Denkendorfer Straße Nellingen"
	DENTELTAL_SPIEGELBERG                                  = "Denteltal Spiegelberg"
	DESSAUER_STRASSE_STUTTGART                             = "Dessauer Straße Stuttgart"
	DETTENHAUSEN_DETTENHAUSEN                              = "Dettenhausen Dettenhausen"
	DETTINGEN_T_DETTINGEN_T                                = "Dettingen (T) Dettingen (T)"
	DEUTSCHES_ROTES_KREUZ_SECHSELBERG                      = "Deutsches Rotes Kreuz Sechselberg"
	DEVIZESSTRASSE_WAIBLINGEN                              = "Devizesstraße Waiblingen"
	DIAKONIE_STETTEN_I_R                                   = "Diakonie Stetten (i. R.)"
	DIAKONISSENMUTTERHAUS_AIDLINGEN                        = "Diakonissenmutterhaus Aidlingen"
	DIEB_HEIMSHEIM                                         = "Dieb Heimsheim"
	DIEGELSBERG_UHINGEN                                    = "Diegelsberg Uhingen"
	DIEPOLDSTRASSE_BERNHAUSEN                              = "Diepoldstraße Bernhausen"
	DIESELSTRASSE_OEFFINGEN                                = "Dieselstraße Oeffingen"
	DIESELSTRASSE_RUTESHEIM                                = "Dieselstraße Rutesheim"
	DIESELSTRASSE_OPPENWEILER                              = "Dieselstraße Oppenweiler"
	DIESELSTRASSE_GOPPINGEN                                = "Dieselstraße Göppingen"
	DIETRICHBONHOEFFERSTRASSE_GERADSTETTEN                 = "Dietrich-Bonhoeffer-Straße Geradstetten"
	DIEZENHALDE_ZENTRUM_BOBLINGEN                          = "Diezenhalde (Zentrum) Böblingen"
	DIEZENHALDE_ZENTRUM_SUD_BOBLINGEN                      = "Diezenhalde (Zentrum) Süd Böblingen"
	DIEZENHALDENWEG_BOBLINGEN                              = "Diezenhaldenweg Böblingen"
	DILLMANNSTRASSE_STUTTGART                              = "Dillmannstraße Stuttgart"
	DINGLESMAD_GSCHWEND                                    = "Dinglesmad Gschwend"
	DINKELSTRASSE_KARLSHOF_STUTTGART                       = "Dinkelstraße (Karlshof) Stuttgart"
	DITZENBACHER_STR_AUENDORF                              = "Ditzenbacher Str. Auendorf"
	DITZENBRUNNER_STRASSE_DITZINGEN                        = "Ditzenbrunner Straße Ditzingen"
	DITZENBRUNNERKNIELSTRASSE_DITZINGEN                    = "Ditzenbrunner-/Knielstraße Ditzingen"
	DITZINGEN_DITZINGEN                                    = "Ditzingen Ditzingen"
	DITZINGER_STRASSE_STUTTGART                            = "Ditzinger Straße Stuttgart"
	DLW_BIETIGHEIM                                         = "DLW Bietigheim"
	DOBELSTRASSE_STUTTGART                                 = "Dobelstraße Stuttgart"
	DOFFINGER_STRASSE_DARMSHEIM                            = "Döffinger Straße Darmsheim"
	DOGGENBURG_STUTTGART                                   = "Doggenburg Stuttgart"
	DOLLENSTRASSE_ALFDORF                                  = "Döllenstraße Alfdorf"
	DOMO_SINDELFINGEN                                      = "Domo Sindelfingen"
	DORFBRUNNEN_HAFNERHASLACH                              = "Dorfbrunnen Häfnerhaslach"
	DORFGEMEINSCHAFT_TENNENTAL_DECKENPFRONN                = "Dorfgemeinschaft Tennental Deckenpfronn"
	DORFHALLE_STEINBACH                                    = "Dorfhalle Steinbach"
	DORFPLATZ_RIET                                         = "Dorfplatz Riet"
	DORFPLATZ_GUSSENSTADT                                  = "Dorfplatz Gussenstadt"
	DORFSTRASSE_NECKARREMS                                 = "Dorfstraße Neckarrems"
	DORFSTRASSE_PFLUGFELDEN                                = "Dorfstraße Pflugfelden"
	DORFSTRASSE_ROSSWALDEN                                 = "Dorfstraße Roßwälden"
	DORFWIESEN_SCHMIDEN                                    = "Dorfwiesen Schmiden"
	DORNHALDENSTRASSE_STUTTGART                            = "Dornhaldenstraße Stuttgart"
	DORNIERSTRASSE_DITZINGEN                               = "Dornierstraße Ditzingen"
	DORNIERSTRASSE_BOBLINGEN                               = "Dornierstraße Böblingen"
	DORNIERSTRASSE_SIRNAU                                  = "Dornierstraße Sirnau"
	DORNROSCHENWEG_STUTTGART                               = "Dornröschenweg Stuttgart"
	DOROTHEENSTRASSE_STUTTGART                             = "Dorotheenstraße Stuttgart"
	DOROTHEENWEG_WERNAU_N                                  = "Dorotheenweg Wernau (N)"
	DRHERBERTKONIGPLATZ_GOPPINGEN                          = "Dr.-Herbert-König-Platz Göppingen"
	DRPFEIFFERSTRASSE_GOPPINGEN                            = "Dr.-Pfeiffer-Straße Göppingen"
	DREIECK_WEILER_EBERSBACH                               = "Dreieck Weiler (Ebersbach)"
	DRESCHERSTRASSE_RUTESHEIM                              = "Drescherstraße Rutesheim"
	DRESCHHALLE_GUNDELBACH                                 = "Dreschhalle Gündelbach"
	DRESDENER_RING_BACKNANG                                = "Dresdener Ring Backnang"
	DRIEFBRUNNEN_PLATTENHARDT                              = "Driefbrunnen Plattenhardt"
	DRK_SINDELFINGEN                                       = "DRK Sindelfingen"
	DROSSELWEG_SINDELFINGEN                                = "Drosselweg Sindelfingen"
	DROSSELWEG_FELLBACH                                    = "Drosselweg Fellbach"
	DROSSELWEG_DOFFINGEN                                   = "Drosselweg Döffingen"
	DUGENDORF_SALACH                                       = "Dugendorf Salach"
	DULKHAUSLE_WIFLINGSHAUSEN                              = "Dulkhäusle Wiflingshausen"
	DURERSTRASSE_GOPPINGEN                                 = "Dürerstraße Göppingen"
	DURERWEG_LORCH                                         = "Dürerweg Lorch"
	DURNAUER_STRASSE_BAD_BOLL                              = "Dürnauer Straße Bad Boll"
	DURNAUER_WEG_STUTTGART                                 = "Dürnauer Weg Stuttgart"
	DURRMARBACHER_WEG_BISSINGEN_LB                         = "Dürr/Marbacher Weg Bissingen (LB)"
	DURRROSENSTRASSE_BISSINGEN_LB                          = "Dürr/Rosenstraße Bissingen (LB)"
	DURRBACH_WEILER_SCHORNDORF                             = "Dürrbach Weiler (Schorndorf)"
	DURRBACHSTRASSE_STUTTGART                              = "Dürrbachstraße Stuttgart"
	DURRLEWANG_STUTTGART                                   = "Dürrlewang Stuttgart"
	EBELSTRASSE_HOHENECK                                   = "Ebelstraße Hoheneck"
	EBENE_ESSLINGEN_N                                      = "Ebene Esslingen (N)"
	EBENESCHULE_RECHBERGHAUSEN                             = "Ebene/Schule Rechberghausen"
	EBERDINGER_STRASSE_HOCHDORF_EBERDINGEN                 = "Eberdinger Straße Hochdorf (Eberdingen)"
	EBERHARDBAUERSTRASSE_PLIENSAUVORSTADT                  = "Eberhard-Bauer-Straße Pliensauvorstadt"
	EBERSBACH_F_EBERSBACH_F                                = "Ebersbach (F) Ebersbach (F)"
	EBERSBACHER_STRASSE_SCHLIERBACH                        = "Ebersbacher Straße Schlierbach"
	EBERSHALDENFRIEDHOF_ESSLINGEN_N                        = "Ebershaldenfriedhof Esslingen (N)"
	EBITZWEG_STUTTGART                                     = "Ebitzweg Stuttgart"
	EBNI_KAISERSBACH                                       = "Ebni Kaisersbach"
	EBNISEE_KAISERSBACH                                    = "Ebnisee Kaisersbach"
	EBNISEESTRASSE_STUTTGART                               = "Ebniseestraße Stuttgart"
	EBNISEESTRASSE_ALTHUTTE                                = "Ebniseestraße Althütte"
	EBNISEESTRASSE_OBERWEISSACH                            = "Ebniseestraße Oberweissach"
	ECHTERDINGEN_ECHTERDINGEN                              = "Echterdingen Echterdingen"
	ECKARTSWEILER_WELZHEIM                                 = "Eckartsweiler Welzheim"
	ECKENERSTRASSE_SIRNAU                                  = "Eckenerstraße Sirnau"
	ECKWALDEN_ORTSEINGANG_BAD_BOLL                         = "Eckwälden Ortseingang Bad Boll"
	EDEKA_UNTERWEISSACH                                    = "Edeka Unterweissach"
	EDELMANNSHOF_RUDERSBERG                                = "Edelmannshof Rudersberg"
	EDELWEISSWEG_STUTTGART                                 = "Edelweißweg Stuttgart"
	EDENHOF_LORCH                                          = "Edenhof Lorch"
	EGELSBERG_EGELSBERGSTRASSE_WEILHEIM_T                  = "Egelsberg Egelsbergstraße Weilheim (T)"
	EGELSBERG_HOCHHAUS_WEILHEIM_T                          = "Egelsberg Hochhaus Weilheim (T)"
	EGELSBERG_TECKSTRASSE_WEILHEIM_T                       = "Egelsberg Teckstraße Weilheim (T)"
	EGELSEE_RIELINGSHAUSEN                                 = "Egelsee Rielingshausen"
	EGELSEE_HEIMSHEIM                                      = "Egelsee Heimsheim"
	EGERT_ZELL_ESSLINGEN                                   = "Egert Zell (Esslingen)"
	EHNINGEN_EHNINGEN                                      = "Ehningen Ehningen"
	EHRENHALDE_STUTTGART                                   = "Ehrenhalde Stuttgart"
	EIBENWEG_NURTINGEN                                     = "Eibenweg Nürtingen"
	EICHBUHL_WANGEN_GP                                     = "Eichbühl Wangen (GP)"
	EICHE_SCHONAICH                                        = "Eiche Schönaich"
	EICHENDORFFSCHULE_BOBLINGEN                            = "Eichendorff-Schule Böblingen"
	EICHENDORFFSCHULE_ZOLLBERG                             = "Eichendorffschule Zollberg"
	EICHENDORFFSTRASSE_NECKARTAILFINGEN                    = "Eichendorffstraße Neckartailfingen"
	EICHENDORFFSTRASSE_ZOLLBERG                            = "Eichendorffstraße Zollberg"
	EICHENDORFFSTRASSE_KIRCHHEIM_T                         = "Eichendorffstraße Kirchheim (T)"
	EICHENDORFFSTRASSE_NURTINGEN                           = "Eichendorffstraße Nürtingen"
	EICHENHOF_EISLINGEN_F                                  = "Eichenhof Eislingen (F)"
	EICHENPFADLE_DAGERSHEIM                                = "Eichenpfädle Dagersheim"
	EICHENWEG_KORNWESTHEIM                                 = "Eichenweg Kornwestheim"
	EICHENWEG_HERRENBERG                                   = "Eichenweg Herrenberg"
	EICHENWEG_BISSINGEN_LB                                 = "Eichenweg Bissingen (LB)"
	EICHGRABEN_MARBACH_N                                   = "Eichgraben Marbach (N)"
	EICHHALDE_BAD_BOLL                                     = "Eichhalde Bad Boll"
	EICHHOLZ_SINDELFINGEN                                  = "Eichholz Sindelfingen"
	EICHWALD_BACKNANG                                      = "Eichwald Backnang"
	EICHWASEN_NECKARTENZLINGEN                             = "Eichwasen Neckartenzlingen"
	EIERHAUSLE_SCHWIEBERDINGEN                             = "Eierhäusle Schwieberdingen"
	EIERHOF_WELZHEIM                                       = "Eierhof Welzheim"
	EIERNEST_STUTTGART                                     = "Eiernest Stuttgart"
	EINFELDSTRASSE_HOLZHAUSEN_UHINGEN                      = "Einfeldstraße Holzhausen (Uhingen)"
	EINKAUFSZENTRUM_GERLINGEN                              = "Einkaufszentrum Gerlingen"
	EINKAUFSZENTRUM_OBERJETTINGEN                          = "Einkaufszentrum Oberjettingen"
	EINKAUFSZENTRUM_GOSBACH                                = "Einkaufszentrum Gosbach"
	EINKAUFSZENTRUM_NAGOLD                                 = "Einkaufszentrum Nagold"
	EINSTEINSTRASSE_GULTSTEIN                              = "Einsteinstraße Gültstein"
	EINSTEINSTRASSE_KIRCHHEIM_T                            = "Einsteinstraße Kirchheim (T)"
	EINSTEINSTRASSE_SERSHEIM                               = "Einsteinstraße Sersheim"
	EINTRACHT_BACKNANG                                     = "Eintracht Backnang"
	EISBERG_L362_NAGOLD                                    = "Eisberg L362 Nagold"
	EISENBAHNBRUCKE_GEISLINGEN_STEIGE                      = "Eisenbahnbrücke Geislingen (Steige)"
	EISENBAHNSTRASSE_PLOCHINGEN                            = "Eisenbahnstraße Plochingen"
	EISENSCHMIEDMUHLE_MURRHARDT                            = "Eisenschmiedmühle Murrhardt"
	EISENTALSTRASSE_WAIBLINGEN                             = "Eisentalstraße Waiblingen"
	EISGASSE_HEMMINGEN                                     = "Eisgasse Hemmingen"
	EISLINGEN_F_EISLINGEN_F                                = "Eislingen (F) Eislingen (F)"
	EISLINGER_STRASSE_SALACH                               = "Eislinger Straße Salach"
	ELBENPLATZ_BOBLINGEN                                   = "Elbenplatz Böblingen"
	ELBESTRASSE_STUTTGART                                  = "Elbestraße Stuttgart"
	ELISABETHKRANZSTRASSE_LUDWIGSBURG                      = "Elisabeth-Kranz-Straße Ludwigsburg"
	ELISABETHSELBERTGYMNASIUM_BERNHAUSEN                   = "Elisabeth-Selbert-Gymnasium Bernhausen"
	ELLENTAL_BIETIGHEIM                                    = "Ellental Bietigheim"
	ELLENTALGYMNASIEN_BIETIGHEIM                           = "Ellentalgymnasien Bietigheim"
	ELLENWEILER_OPPENWEILER                                = "Ellenweiler Oppenweiler"
	ELSABRANDSTROMSTRASSE_BOBLINGEN                        = "Elsa-Brändström-Straße Böblingen"
	ELSABRANDSTROMSTRASSE_HOFINGEN                         = "Elsa-Brändström-Straße Höfingen"
	ELSENHALDE_SCHONAICH                                   = "Elsenhalde Schönaich"
	ELSERRING_BESIGHEIM                                    = "Elserring Besigheim"
	ELSTERSTAFFEL_STUTTGART                                = "Elsterstaffel Stuttgart"
	ELSTERWEG_DOFFINGEN                                    = "Elsterweg Döffingen"
	ELTINGER_STRASSE_SINDELFINGEN                          = "Eltinger Straße Sindelfingen"
	ELTINGER_STRASSE_STUTTGART                             = "Eltinger Straße Stuttgart"
	ELWERTSTRASSE_STUTTGART                                = "Elwertstraße Stuttgart"
	EMERHOLZ_STUTTGART                                     = "Emerholz Stuttgart"
	EMILMUNZSTRASSE_WAIBLINGEN                             = "Emil-Münz-Straße Waiblingen"
	ENBW_CITY_STUTTGART                                    = "EnBW City Stuttgart"
	ENDERSBACH_ENDERSBACH                                  = "Endersbach Endersbach"
	ENDERSBACHER_STRASSE_STETTEN_I_R                       = "Endersbacher Straße Stetten (i. R.)"
	ENDERSBACHER_STRASSE_BEINSTEIN                         = "Endersbacher Straße Beinstein"
	ENDWIESENSTRASSE_HOPFIGHEIM                            = "Endwiesenstraße Höpfigheim"
	ENGELBERG_LEONBERG                                     = "Engelberg Leonberg"
	ENGELBERG_WINTERBACH                                   = "Engelberg Winterbach"
	ENGELBERG_WALDORFSCHULE_WINTERBACH                     = "Engelberg Waldorfschule Winterbach"
	ENGELBOLDSTRASSE_STUTTGART                             = "Engelboldstraße Stuttgart"
	ENSINGER_STRASSE_KLEINGLATTBACH                        = "Ensinger Straße Kleinglattbach"
	ENTENACKER_BISSINGEN_LB                                = "Entenäcker Bissingen (LB)"
	ENZBACHWEG_SCHLAT                                      = "Enzbachweg Schlat"
	ENZBRUCKE_BIETIGHEIM                                   = "Enzbrücke Bietigheim"
	ENZBRUCKE_VAIHINGEN_E                                  = "Enzbrücke Vaihingen (E)"
	ENZENHARDTPLATZ_NURTINGEN                              = "Enzenhardtplatz Nürtingen"
	ENZPLATZ_BESIGHEIM                                     = "Enzplatz Besigheim"
	ENZSTRASSE_KORNWESTHEIM                                = "Enzstraße Kornwestheim"
	ENZSTRASSE_UNTERMBERG                                  = "Enzstraße Untermberg"
	ENZWEIHINGEN_B10_ENZWEIHINGEN                          = "Enzweihingen B10 Enzweihingen"
	EPPLERINWEG_SCHORNDORF                                 = "Epplerinweg Schorndorf"
	ERDBEERWEG_STUTTGART                                   = "Erdbeerweg Stuttgart"
	ERDMANNHAUSEN_ERDMANNHAUSEN                            = "Erdmannhausen Erdmannhausen"
	ERDMANNHAUSER_STRASSE_MARBACH_N                        = "Erdmannhäuser Straße Marbach (N)"
	ERGENZINGEN_ERGENZINGEN                                = "Ergenzingen Ergenzingen"
	ERHOLUNGSHEIM_GULTSTEIN                                = "Erholungsheim Gültstein"
	ERICHSCHUMMSTRASSE_MURRHARDT                           = "Erich-Schumm-Straße Murrhardt"
	ERLACH_STETTEN_LEINFELDENECHT                          = "Erlach Stetten (Leinfelden-Echt.)"
	ERLACH_GROSSERLACH                                     = "Erlach Großerlach"
	ERLACH_ABZWEIG_GROSSERLACH                             = "Erlach Abzweig Großerlach"
	ERLENHOF_ODERNHARDT                                    = "Erlenhof Ödernhardt"
	ERLENSTRASSE_SCHORNDORF                                = "Erlenstraße Schorndorf"
	ERLIGHEIM_ERLIGHEIM                                    = "Erligheim Erligheim"
	ERLIGHEIMER_STRASSE_LOCHGAU                            = "Erligheimer Straße Löchgau"
	ERNSTBAUERPLATZ_RENNINGEN                              = "Ernst-Bauer-Platz Renningen"
	ERNSTREUTERPLATZ_STUTTGART                             = "Ernst-Reuter-Platz Stuttgart"
	ERNSTSIGLEGYMNASIUM_KORNWESTHEIM                       = "Ernst-Sigle-Gymnasium Kornwestheim"
	ERWINSCHOETTLEPLATZ_STUTTGART                          = "Erwin-Schoettle-Platz Stuttgart"
	ESCHENBACHER_STR_GAMMELSHAUSEN                         = "Eschenbacher Str. Gammelshausen"
	ESCHENBRUNNLESTRASSE_SINDELFINGEN                      = "Eschenbrünnlestraße Sindelfingen"
	ESCHENRIED_SINDELFINGEN                                = "Eschenried Sindelfingen"
	ESCHENSTRUET_SULZBACH_M                                = "Eschenstruet Sulzbach (M)"
	ESCHENWEG_SCHELMENHOLZ                                 = "Eschenweg Schelmenholz"
	ESELSHALDEN_WELZHEIM                                   = "Eselshalden Welzheim"
	ESELSMUHLE_MUSBERG                                     = "Eselsmühle Musberg"
	ESELWEG_STUTTGART                                      = "Eselweg Stuttgart"
	ESSEGGER_STRASSE_SINDELFINGEN                          = "Essegger Straße Sindelfingen"
	ESSLINGEN_N_ESSLINGEN_N                                = "Esslingen (N) Esslingen (N)"
	ESSLINGER_STRASSE_FELLBACH                             = "Esslinger Straße Fellbach"
	ESSLINGER_STRASSE_WOLFSCHLUGEN                         = "Esslinger Straße Wolfschlugen"
	ESSLINGER_STRASSE_DEIZISAU                             = "Esslinger Straße Deizisau"
	ESSLINGER_STRASSE_DENKENDORF                           = "Esslinger Straße Denkendorf"
	ESSLINGER_STRASSE_PLOCHINGEN                           = "Esslinger Straße Plochingen"
	ESSLINGER_WEG_MAGSTADT                                 = "Esslinger Weg Magstadt"
	ESZET_STUTTGART                                        = "Eszet Stuttgart"
	ETTERSTRASSE_SCHLATTSTALL                              = "Etterstraße Schlattstall"
	ETZELWEG_KOHLBERG                                      = "Etzelweg Kohlberg"
	ETZWIESENBRUCKE_BACKNANG                               = "Etzwiesenbrücke Backnang"
	EUGENBOLZSTRASSE_ST_BERNHARDT                          = "Eugen-Bolz-Straße St. Bernhardt"
	EUGENBOLZSTRASSE_OST_BOBLINGEN                         = "Eugen-Bolz-Straße Ost Böblingen"
	EUGENSPLATZ_STUTTGART                                  = "Eugensplatz Stuttgart"
	EUGENSTRASSE_FELLBACH                                  = "Eugenstraße Fellbach"
	EUGENSTRASSE_GOPPINGEN                                 = "Eugenstraße Göppingen"
	EULENBERG_MUSBERG                                      = "Eulenberg Musberg"
	EULENHOF_KAISERSBACH                                   = "Eulenhof Kaisersbach"
	EUROPAPLATZ_STUTTGART                                  = "Europaplatz Stuttgart"
	EV_AKADEMIEREHAKLINIK_BAD_BOLL                         = "Ev. Akademie/Reha-Klinik Bad Boll"
	EV_KINDERGARTEN_GINGEN_F                               = "Ev. Kindergarten Gingen (F)"
	EV_KINDERGARTEN_LORCH                                  = "Ev. Kindergarten Lorch"
	EV_KIRCHE_SACHSENWEILER                                = "Ev. Kirche Sachsenweiler"
	EVANGELISCHE_KIRCHE_SCHONAICH                          = "Evangelische Kirche Schönaich"
	EVANGELISCHE_KIRCHE_UNTERJETTINGEN                     = "Evangelische Kirche Unterjettingen"
	EVANGELISCHE_KIRCHE_STETTEN_I_R                        = "Evangelische Kirche Stetten (i. R.)"
	EWSARENALORCHER_STR_GOPPINGEN                          = "EWS-Arena/Lorcher Str. Göppingen"
	EWSARENANORDL_RINGSTR_GOPPINGEN                        = "EWS-Arena/Nördl. Ringstr. Göppingen"
	EYBTALHALLE_EYBACH                                     = "Eybtalhalle Eybach"
	FA_ALLGAIER_UHINGEN                                    = "Fa. Allgaier Uhingen"
	FABRIKSTRASSE_WENDLINGEN_N                             = "Fabrikstraße Wendlingen (N)"
	FABRIKSTRASSE_OTLINGEN                                 = "Fabrikstraße Ötlingen"
	FABRIKSTRASSE_NEIDLINGEN                               = "Fabrikstraße Neidlingen"
	FABRIKSTRASSE_BISSINGEN_AN_DER_TECK_T                  = "Fabrikstraße Bissingen an der Teck (T)"
	FABRIKSTRASSE_UHINGEN                                  = "Fabrikstraße Uhingen"
	FAHRGERGASSE_GEMMRIGHEIM                               = "Fährgergasse Gemmrigheim"
	FAHRZEUGWERKE_EISLINGEN_F                              = "Fahrzeugwerke Eislingen (F)"
	FALBENHENNENSTRASSE_STUTTGART                          = "Falbenhennenstraße Stuttgart"
	FALKENSTRASSE_JEBENHAUSEN                              = "Falkenstraße Jebenhausen"
	FALKENWEG_EGLOSHEIM                                    = "Falkenweg Eglosheim"
	FARBERSTRASSE_KUCHEN_F                                 = "Färberstraße Kuchen (F)"
	FASANENHOF_STUTTGART                                   = "Fasanenhof Stuttgart"
	FASANENHOF_ABZWEIG_STUTTGART                           = "Fasanenhof Abzweig Stuttgart"
	FASANENSTRASSE_JEBENHAUSEN                             = "Fasanenstraße Jebenhausen"
	FAULHABER_SCHONAICH                                    = "Faulhaber Schönaich"
	FAURNDAU_FAURNDAU                                      = "Faurndau Faurndau"
	FAUSLERSTRASSE_JESINGEN                                = "Fauslerstraße Jesingen"
	FAUSTSTRASSE_STUTTGART                                 = "Fauststraße Stuttgart"
	FAUTSPACH_SECHSELBERG                                  = "Fautspach Sechselberg"
	FAVORITEPARK_EGLOSHEIM                                 = "Favoritepark Eglosheim"
	FEHRBELLINER_STR_STUTTGART                             = "Fehrbelliner Str. Stuttgart"
	FELDBERGSTRASSE_SINDELFINGEN                           = "Feldbergstraße Sindelfingen"
	FELDBERGSTRASSE_BOBLINGEN                              = "Feldbergstraße Böblingen"
	FELDWIESEN_BUNZWANGEN                                  = "Feldwiesen Bünzwangen"
	FELLBACH_FELLBACH                                      = "Fellbach Fellbach"
	FELLBACHER_STRASSE_ROMMELSHAUSEN                       = "Fellbacher Straße Rommelshausen"
	FELLBACHER_STRASSE_OSSWEIL                             = "Fellbacher Straße Oßweil"
	FERDINANDPORSCHESTR_NAGOLD                             = "Ferdinand-Porsche-Str. Nagold"
	FERNSEHTURM_STUTTGART                                  = "Fernsehturm Stuttgart"
	FESSLER_MUHLE_SERSHEIM                                 = "Fessler Mühle Sersheim"
	FESTHALLE_RUTESHEIM                                    = "Festhalle Rutesheim"
	FESTHALLE_HOLZHAUSEN_UHINGEN                           = "Festhalle Holzhausen (Uhingen)"
	FESTPLATZ_MARKGRONINGEN                                = "Festplatz Markgröningen"
	FESTPLATZ_BALTMANNSWEILER                              = "Festplatz Baltmannsweiler"
	FEUERBACH_STUTTGART                                    = "Feuerbach Stuttgart"
	FEUERBACH_BOSCH_STUTTGART                              = "Feuerbach Bosch Stuttgart"
	FEUERBACH_FRIEDHOF_STUTTGART                           = "Feuerbach Friedhof Stuttgart"
	FEUERBACH_PFOSTENWALDLE_STUTTGART                      = "Feuerbach Pfostenwäldle Stuttgart"
	FEUERBACHER_STRASSE_LEONBERG                           = "Feuerbacher Straße Leonberg"
	FEUERBACHER_WEG_STUTTGART                              = "Feuerbacher Weg Stuttgart"
	FEUERSEE_KAPPISHAUSERN                                 = "Feuersee Kappishäusern"
	FEUERSEE_STUTTGART                                     = "Feuersee Stuttgart"
	FEUERWACHE_KUCHEN_F                                    = "Feuerwache Kuchen (F)"
	FEUERWACHE_GOPPINGEN                                   = "Feuerwache Göppingen"
	FEUERWEHR_GERLINGEN                                    = "Feuerwehr Gerlingen"
	FEUERWEHR_ALTBACH                                      = "Feuerwehr Altbach"
	FEUERWEHR_BACKNANG                                     = "Feuerwehr Backnang"
	FEUERWEHR_URBACH                                       = "Feuerwehr Urbach"
	FEUERWEHR_KIRCHBERG_M                                  = "Feuerwehr Kirchberg (M)"
	FEUERWEHR_RECHBERGHAUSEN                               = "Feuerwehr Rechberghausen"
	FEUERWEHRGERATEHAUS_GARTRINGEN                         = "Feuerwehrgerätehaus Gärtringen"
	FEUERWEHRHAUS_HORRHEIM                                 = "Feuerwehrhaus Horrheim"
	FEUERWEHRHAUS_ESCHENBACH                               = "Feuerwehrhaus Eschenbach"
	FEUERWEHRMAGAZIN_ENSINGEN                              = "Feuerwehrmagazin Ensingen"
	FICHTESTRASSE_STUTTGART                                = "Fichtestraße Stuttgart"
	FILDERAIRPORTAREAL_BERNHAUSEN                          = "Filder-Airport-Areal Bernhausen"
	FILDERBAHNSTRASSE_STUTTGART                            = "Filderbahnstraße Stuttgart"
	FILDERBLICKWEG_STUTTGART                               = "Filderblickweg Stuttgart"
	FILDERKLINIK_BONLANDEN                                 = "Filderklinik Bonlanden"
	FILDERSTADT_BERNHAUSEN                                 = "Filderstadt Bernhausen"
	FILDORADO_BONLANDEN                                    = "Fildorado Bonlanden"
	FILHARMONIE_BERNHAUSEN                                 = "Filharmonie Bernhausen"
	FILSBRUCKE_EBERSBACH_F                                 = "Filsbrücke Ebersbach (F)"
	FILSECKSTRASSE_FAURNDAU                                = "Filseckstraße Faurndau"
	FINANZAMT_WAIBLINGEN                                   = "Finanzamt Waiblingen"
	FINANZAMT_ESSLINGEN_N                                  = "Finanzamt Esslingen (N)"
	FINKENBERG_WAIBLINGEN                                  = "Finkenberg Waiblingen"
	FINKENSTRASSE_PLATTENHARDT                             = "Finkenstraße Plattenhardt"
	FINKENWEG_SIRNAU                                       = "Finkenweg Sirnau"
	FINKENWEG_ALBERSHAUSEN                                 = "Finkenweg Albershausen"
	FIRMA_BINZ_LORCH                                       = "Firma Binz Lorch"
	FIRMA_FESTO_BERKHEIM                                   = "Firma Festo Berkheim"
	FIRMA_ZUGEL_WUSTENROT                                  = "Firma Zügel Wüstenrot"
	FISCHERPFAD_BIETIGHEIM                                 = "Fischerpfad Bietigheim"
	FISCHERWORTH_GROSSINGERSHEIM                           = "Fischerwörth Großingersheim"
	FISCHSTRASSE_GOPPINGEN                                 = "Fischstraße Göppingen"
	FLACHTER_STRASSE_STUTTGART                             = "Flachter Straße Stuttgart"
	FLACHTER_STRASSE_WEISSACH_BB                           = "Flachter Straße Weissach (BB)"
	FLANDERNSTRASSE_ST_BERNHARDT                           = "Flandernstraße St. Bernhardt"
	FLATTICHSCHULE_MUNCHINGEN                              = "Flattichschule Münchingen"
	FLEINSBACH_BERNHAUSEN                                  = "Fleinsbach Bernhausen"
	FLICKENWIESEN_GROSSBOTTWAR                             = "Flickenwiesen Großbottwar"
	FLUGFELDALLEE_BOBLINGEN                                = "Flugfeld-Allee Böblingen"
	FLUGHAFENMESSE_ECHTERDINGEN                            = "Flughafen/Messe Echterdingen"
	FOHRENBUHL_BACKNANG                                    = "Föhrenbühl Backnang"
	FOHRICH_STUTTGART                                      = "Föhrich Stuttgart"
	FORCHENRAINSTRASSE_GERLINGEN                           = "Forchenrainstraße Gerlingen"
	FORCHENWALDSTRASSE_SCHELMENHOLZ                        = "Forchenwaldstraße Schelmenholz"
	FORELLENWEG_STUTTGART                                  = "Forellenweg Stuttgart"
	FORNSBACH_FORNSBACH                                    = "Fornsbach Fornsbach"
	FORNSBACHER_WEG_BACKNANG                               = "Fornsbacher Weg Backnang"
	FORSCHUNG_DENKENDORF                                   = "Forschung Denkendorf"
	FORSTBODEN_GROSSASPACH                                 = "Forstboden Großaspach"
	FORSTHAUS_I_STUTTGART                                  = "Forsthaus I Stuttgart"
	FORSTHAUS_II_STUTTGART                                 = "Forsthaus II Stuttgart"
	FORSTHAUS_PARKPLATZ_STUTTGART                          = "Forsthaus Parkplatz Stuttgart"
	FORSTSTRASSE_STETTEN_LEINFELDENECHT                    = "Forststraße Stetten (Leinfelden-Echt.)"
	FORSTSTRASSE_KAISERSBACH                               = "Forststraße Kaisersbach"
	FORUM_AM_SCHLOSSPARK_LUDWIGSBURG                       = "Forum am Schlosspark Ludwigsburg"
	FRANK_LEINFELDEN                                       = "Frank Leinfelden"
	FRANKENSTRASSE_STUTTGART                               = "Frankenstraße Stuttgart"
	FRANKENSTRASSE_SINDELFINGEN                            = "Frankenstraße Sindelfingen"
	FRANKENWEILER_GRAB                                     = "Frankenweiler Grab"
	FRANKFURTER_STRASSE_EGLOSHEIM                          = "Frankfurter Straße Eglosheim"
	FRAUBRONNSTRASSE_STUTTGART                             = "Fraubronnstraße Stuttgart"
	FRAUENKOPF_STUTTGART                                   = "Frauenkopf Stuttgart"
	FRAUENKREUZ_LEONBERG                                   = "Frauenkreuz Leonberg"
	FRAUENLANDERSTRASSE_STETTEN_I_R                        = "Frauenländerstraße Stetten (i. R.)"
	FRAUENSTRASSE_GEISLINGEN_STEIGE                        = "Frauenstraße Geislingen (Steige)"
	FREIBAD_KIRCHHEIM_T                                    = "Freibad Kirchheim (T)"
	FREIBAD_BOBLINGEN                                      = "Freibad Böblingen"
	FREIBAD_STETTEN_I_R                                    = "Freibad Stetten (i. R.)"
	FREIBAD_WAIBLINGEN                                     = "Freibad Waiblingen"
	FREIBAD_BONNIGHEIM                                     = "Freibad Bönnigheim"
	FREIBAD_GOPPINGEN                                      = "Freibad Göppingen"
	FREIBAD_UHINGEN                                        = "Freibad Uhingen"
	FREIBADFITKOM_BESIGHEIM                                = "Freibad/FitKom Besigheim"
	FREIBERG_STUTTGART                                     = "Freiberg Stuttgart"
	FREIBERG_N_FREIBERG_N                                  = "Freiberg (N) Freiberg (N)"
	FREIBERG_SCHULZENTRUM_STUTTGART                        = "Freiberg Schulzentrum Stuttgart"
	FREIBERGER_STRASSE_BIETIGHEIM                          = "Freiberger Straße Bietigheim"
	FREIBERGSTRASSE_STUTTGART                              = "Freibergstraße Stuttgart"
	FREIBURGER_ALLEE_BOBLINGEN                             = "Freiburger Allee Böblingen"
	FREIHOFPLATZ_STUTTGART                                 = "Freihofplatz Stuttgart"
	FREILICHTMUSEUM_BEUREN                                 = "Freilichtmuseum Beuren"
	FREITAGSHOF_WERNAU_N                                   = "Freitagshof Wernau (N)"
	FREIWALDAUSTRASSE_KIRCHHEIM_T                          = "Freiwaldaustraße Kirchheim (T)"
	FREIZEITPARK_KORNWESTHEIM                              = "Freizeitpark Kornwestheim"
	FREIZEITZENTRUM_BUOCH                                  = "Freizeitzentrum Buoch"
	FREUDENSTADTER_STRASSE_BOBLINGEN                       = "Freudenstädter Straße Böblingen"
	FREUDENTALER_STRASSE_LOCHGAU                           = "Freudentaler Straße Löchgau"
	FREYTAGWEG_STUTTGART                                   = "Freytagweg Stuttgart"
	FRICKENHAUSEN_FRICKENHAUSEN                            = "Frickenhausen Frickenhausen"
	FRICKENHAUSER_STRASSE_TISCHARDT                        = "Frickenhäuser Straße Tischardt"
	FRIDINGER_STRASSE_STUTTGART                            = "Fridinger Straße Stuttgart"
	FRIEDENSKIRCHE_URBACH                                  = "Friedenskirche Urbach"
	FRIEDENSSCHULE_NEUSTADT                                = "Friedensschule Neustadt"
	FRIEDENSTRASSE_STUTTGART                               = "Friedenstraße Stuttgart"
	FRIEDENSTRASSE_5_LUDWIGSBURG                           = "Friedenstraße 5 Ludwigsburg"
	FRIEDENSTRASSE_52_LUDWIGSBURG                          = "Friedenstraße 52 Ludwigsburg"
	FRIEDHOF_MUSBERG                                       = "Friedhof Musberg"
	FRIEDHOF_SIELMINGEN                                    = "Friedhof Sielmingen"
	FRIEDHOF_ECHTERDINGEN                                  = "Friedhof Echterdingen"
	FRIEDHOF_RENNINGEN                                     = "Friedhof Renningen"
	FRIEDHOF_BACKNANG                                      = "Friedhof Backnang"
	FRIEDHOF_ROMMELSHAUSEN                                 = "Friedhof Rommelshausen"
	FRIEDHOF_ZELL_ESSLINGEN                                = "Friedhof Zell (Esslingen)"
	FRIEDHOF_OHMDEN                                        = "Friedhof Ohmden"
	FRIEDHOF_ZIZISHAUSEN                                   = "Friedhof Zizishausen"
	FRIEDHOF_MUNCHINGEN                                    = "Friedhof Münchingen"
	FRIEDHOF_WEIL_DER_STADT                                = "Friedhof Weil der Stadt"
	FRIEDHOF_ALTDORF_BB                                    = "Friedhof Altdorf (BB)"
	FRIEDHOF_WELZHEIM                                      = "Friedhof Welzheim"
	FRIEDHOF_WINNENDEN                                     = "Friedhof Winnenden"
	FRIEDHOF_UNTERBRUDEN                                   = "Friedhof Unterbrüden"
	FRIEDHOF_KORNWESTHEIM                                  = "Friedhof Kornwestheim"
	FRIEDHOF_POPPENWEILER                                  = "Friedhof Poppenweiler"
	FRIEDHOF_BESIGHEIM                                     = "Friedhof Besigheim"
	FRIEDHOF_GROSSINGERSHEIM                               = "Friedhof Großingersheim"
	FRIEDHOF_WAIBLINGEN                                    = "Friedhof Waiblingen"
	FRIEDHOF_NUFRINGEN                                     = "Friedhof Nufringen"
	FRIEDHOF_THOMASHARDT                                   = "Friedhof Thomashardt"
	FRIEDHOF_HOLZGERLINGEN                                 = "Friedhof Holzgerlingen"
	FRIEDHOF_NECKARREMS                                    = "Friedhof Neckarrems"
	FRIEDHOF_WENDLINGEN_N                                  = "Friedhof Wendlingen (N)"
	FRIEDHOF_GRAB                                          = "Friedhof Grab"
	FRIEDHOF_BITTENFELD                                    = "Friedhof Bittenfeld"
	FRIEDHOF_ROSSWAG                                       = "Friedhof Roßwag"
	FRIEDHOF_DEGGINGEN                                     = "Friedhof Deggingen"
	FRIEDHOF_GEISLINGEN_STEIGE                             = "Friedhof Geislingen (Steige)"
	FRIEDHOF_BOHMENKIRCH                                   = "Friedhof Böhmenkirch"
	FRIEDHOF_WASCHENBEUREN                                 = "Friedhof Wäschenbeuren"
	FRIEDHOF_HATTENHOFEN                                   = "Friedhof Hattenhofen"
	FRIEDHOF_HEININGEN_GP                                  = "Friedhof Heiningen (GP)"
	FRIEDHOF_BORTLINGEN                                    = "Friedhof Börtlingen"
	FRIEDHOF_NORD_EISLINGEN_F                              = "Friedhof Nord Eislingen (F)"
	FRIEDHOF_NORD_GOPPINGEN                                = "Friedhof Nord Göppingen"
	FRIEDHOF_SANKT_PETER_BIETIGHEIM                        = "Friedhof Sankt Peter Bietigheim"
	FRIEDHOF_SUD_EISLINGEN_F                               = "Friedhof Süd Eislingen (F)"
	FRIEDHOF_SUD_GOPPINGEN                                 = "Friedhof Süd Göppingen"
	FRIEDHOFSTRASSE_STUTTGART                              = "Friedhofstraße Stuttgart"
	FRIEDRICHBREININGSTRASSE_BESIGHEIM                     = "Friedrich-Breining-Straße Besigheim"
	FRIEDRICHEBERTSTRASSE_SINDELFINGEN                     = "Friedrich-Ebert-Straße Sindelfingen"
	FRIEDRICHEBERTSTRASSE_BIETIGHEIM                       = "Friedrich-Ebert-Straße Bietigheim"
	FRIEDRICHLISTPLATZ_BOBLINGEN                           = "Friedrich-List-Platz Böblingen"
	FRIEDRICHLISTSTRASSE_FELLBACH                          = "Friedrich-List-Straße Fellbach"
	FRIEDRICHLISTSTRASSE_GROSSASPACH                       = "Friedrich-List-Straße Großaspach"
	FRIEDRICHLISTSTRASSE_WERNAU_N                          = "Friedrich-List-Straße Wernau (N)"
	FRIEDRICHSCHILLERGYMNASIUM_FELLBACH                    = "Friedrich-Schiller-Gymnasium Fellbach"
	FRIEDRICHSCHILLERSCHULE_NEUHAUSEN_F                    = "Friedrich-Schiller-Schule Neuhausen (F)"
	FRIEDRICHSTRASSE_SIELMINGEN                            = "Friedrichstraße Sielmingen"
	FRIEDRICHSTRASSE_SCHMIDEN                              = "Friedrichstraße Schmiden"
	FRIEDRICHSTRASSE_LUDWIGSBURG                           = "Friedrichstraße Ludwigsburg"
	FRIEDRICHSTRASSE_GOPPINGEN                             = "Friedrichstraße Göppingen"
	FRIEDRICHSWAHL_STUTTGART                               = "Friedrichswahl Stuttgart"
	FRIESENSTRASSE_OSSWEIL                                 = "Friesenstraße Oßweil"
	FRITZWALTERWEG_STUTTGART                               = "Fritz-Walter-Weg Stuttgart"
	FROBELSCHULE_HERRENBERG                                = "Fröbelschule Herrenberg"
	FROBELSTRASSE_BERNHAUSEN                               = "Fröbelstraße Bernhausen"
	FROBELSTRASSE_WAIBLINGEN                               = "Fröbelstraße Waiblingen"
	FRONACKERSTRASSE_WAIBLINGEN                            = "Fronackerstraße Waiblingen"
	FRONACKERSTRASSE_SINDELFINGEN                          = "Fronäckerstraße Sindelfingen"
	FRONTALSTRASSE_DECKENPFRONN                            = "Frontalstraße Deckenpfronn"
	FROSCHEGERT_GROTZINGEN                                 = "Froschegert Grötzingen"
	FRUHLINGSTRASSE_ALBERSHAUSEN                           = "Frühlingstraße Albershausen"
	FRUHLINGSTRASSE_ESCHENBACH                             = "Frühlingstraße Eschenbach"
	FRUHMESSHOF_KIRCHBERG_M                                = "Frühmeßhof Kirchberg (M)"
	FRUWIRTHSTRASSE_STUTTGART                              = "Fruwirthstraße Stuttgart"
	FUCHSBAUEINKAUFSMARKT_DITZINGEN                        = "Fuchsbau/Einkaufsmarkt Ditzingen"
	FUCHSECKSATTEL_AUENDORF                                = "Fuchsecksattel Auendorf"
	FUCHSECKSTRASSE_GOPPINGEN                              = "Fuchseckstraße Göppingen"
	FUCHSECKSTRASSE_SUSSEN                                 = "Fuchseckstraße Süßen"
	FUCHSGRUBE_WAIBLINGEN                                  = "Fuchsgrube Waiblingen"
	FUCHSLOCHER_DONNSTETTEN                                = "Fuchslöcher Donnstetten"
	FUCHSRAIN_STUTTGART                                    = "Fuchsrain Stuttgart"
	FUCHSRAIN_HARDT                                        = "Fuchsrain Hardt"
	FULLERSTRASSE_GERLINGEN                                = "Füllerstraße Gerlingen"
	FURFELDER_STRASSE_STUTTGART                            = "Fürfelder Straße Stuttgart"
	FURSTENHOF_GROSSASPACH                                 = "Fürstenhof Großaspach"
	FURTBACHSCHULE_MOGLINGEN                               = "Furtbachschule Möglingen"
	FURTHMUHLE_AIDLINGEN                                   = "Furthmühle Aidlingen"
	GABLENBERG_STUTTGART                                   = "Gablenberg Stuttgart"
	GAILDORFER_STRASSE_BACKNANG                            = "Gaildorfer Straße Backnang"
	GAISBURG_STUTTGART                                     = "Gaisburg Stuttgart"
	GAISERPLATZ_KIRCHHEIM_T                                = "Gaiserplatz Kirchheim (T)"
	GALERIE_WAIBLINGEN                                     = "Galerie Waiblingen"
	GALGENBERG_WAIBLINGEN                                  = "Galgenberg Waiblingen"
	GALGENBERGBRUCKE_WAIBLINGEN                            = "Galgenbergbrücke Waiblingen"
	GAMMELSHAUSER_STR_DURNAU                               = "Gammelshauser Str. Dürnau"
	GAMP_ROHRDORF_CW                                       = "Gamp Rohrdorf (CW)"
	GANSBUHL_SCHORNDORF                                    = "Gänsbühl Schorndorf"
	GANSFUSSALLEE_LUDWIGSBURG                              = "Gänsfußallee Ludwigsburg"
	GANSSEE_BOBLINGEN                                      = "Ganssee Böblingen"
	GANSWASEN_PLUDERHAUSEN                                 = "Gänswasen Plüderhausen"
	GARTENFREUNDE_GOPPINGEN                                = "Gartenfreunde Göppingen"
	GARTENSTADT_LEONBERG                                   = "Gartenstadt Leonberg"
	GARTENSTADT_OBERESSLINGEN                              = "Gartenstadt Oberesslingen"
	GARTENSTRASSE_FELLBACH                                 = "Gartenstraße Fellbach"
	GARTENSTRASSE_BACKNANG                                 = "Gartenstraße Backnang"
	GARTENSTRASSE_NABERN                                   = "Gartenstraße Nabern"
	GARTENSTRASSE_BEUREN                                   = "Gartenstraße Beuren"
	GARTENSTRASSE_HESSIGHEIM                               = "Gartenstraße Hessigheim"
	GARTENSTRASSE_HOCHDORF_EBERDINGEN                      = "Gartenstraße Hochdorf (Eberdingen)"
	GARTENSTRASSE_LUDWIGSBURG                              = "Gartenstraße Ludwigsburg"
	GARTENSTRASSE_GECHINGEN                                = "Gartenstraße Gechingen"
	GARTRINGEN_GARTRINGEN                                  = "Gärtringen Gärtringen"
	GARTRINGER_STRASSE_NUFRINGEN                           = "Gärtringer Straße Nufringen"
	GASTHAUS_KRONE_OTTENBACH                               = "Gasthaus Krone Ottenbach"
	GAUFELDEN_GAUFELDEN                                    = "Gäufelden Gäufelden"
	GAUSMANNSWEILER_WELZHEIM                               = "Gausmannsweiler Welzheim"
	GAUSSSTRASSE_STUTTGART                                 = "Gaußstraße Stuttgart"
	GAUSSWEG_GOPPINGEN                                     = "Gaussweg Göppingen"
	GEBENWEILER_KAISERSBACH                                = "Gebenweiler Kaisersbach"
	GEBERSHEIMER_STRASSE_RUTESHEIM                         = "Gebersheimer Straße Rutesheim"
	GEBERSHEIMERROSEGGERSTRASSE_ELTINGEN                   = "Gebersheimer-/Roseggerstraße Eltingen"
	GEGEN_EICH_OSSWEIL                                     = "Gegen Eich Oßweil"
	GEHRENWALD_STUTTGART                                   = "Gehrenwald Stuttgart"
	GEIBELSTRASSE_STUTTGART                                = "Geibelstraße Stuttgart"
	GEISINGER_DORFPLATZ_FREIBERG_N                         = "Geisinger Dorfplatz Freiberg (N)"
	GEISLINGEN_STEIGE_GEISLINGEN_STEIGE                    = "Geislingen (Steige) Geislingen (Steige)"
	GEISLINGEN_WEST_GEISLINGEN_STEIGE                      = "Geislingen West Geislingen (Steige)"
	GEISLINGER_STRASSE_LAICHINGEN                          = "Geislinger Straße Laichingen"
	GEISLINGER_STRASSE_21_ELTINGEN                         = "Geislinger Straße 21 Eltingen"
	GEISLINGER_STRASSE_3_ELTINGEN                          = "Geislinger Straße 3 Eltingen"
	GEISLINGER_STRASSE_51_ELTINGEN                         = "Geislinger Straße 51 Eltingen"
	GEISSBERG_BIRENBACH                                    = "Geißberg Birenbach"
	GEISSHALDE_DEUFRINGEN                                  = "Geißhalde Deufringen"
	GELEENER_STRASSE_BOBLINGEN                             = "Geleener Straße Böblingen"
	GELEENER_STRASSEKINDERGARTEN_BOBLINGEN                 = "Geleener Straße/Kindergarten Böblingen"
	GEMEINDEHALLE_KIRCHHEIM_N                              = "Gemeindehalle Kirchheim (N)"
	GEMEINDEHALLE_ALTBACH                                  = "Gemeindehalle Altbach"
	GEMEINDEHALLE_KIRCHENKIRNBERG                          = "Gemeindehalle Kirchenkirnberg"
	GEMEINDEHALLE_HOFEN                                    = "Gemeindehalle Höfen"
	GEMEINDEHAUS_BOBLINGEN                                 = "Gemeindehaus Böblingen"
	GEMEINDEHAUS_BUOCH                                     = "Gemeindehaus Buoch"
	GEMEINDEHAUS_ESCHENBACH                                = "Gemeindehaus Eschenbach"
	GEMEINDELANDERWEG_ALBERSHAUSEN                         = "Gemeindeländerweg Albershausen"
	GEMEINDEZENTRUM_OEFFINGEN                              = "Gemeindezentrum Oeffingen"
	GEMEINDEZENTRUM_BERKHEIM                               = "Gemeindezentrum Berkheim"
	GEMEINDEZENTRUM_OBERJETTINGEN                          = "Gemeindezentrum Oberjettingen"
	GEMEINDEZENTRUM_KUCHEN_F                               = "Gemeindezentrum Kuchen (F)"
	GEMEINSCHAFTSSCHULE_BAD_BOLL                           = "Gemeinschaftsschule Bad Boll"
	GENERATIONENHAUS_STUTTGART                             = "Generationenhaus Stuttgart"
	GENERATORSTRASSE_STUTTGART                             = "Generatorstraße Stuttgart"
	GEORGELSERWEG_SUSSEN                                   = "Georg-Elser-Weg Süßen"
	GEORGIIWALDSTADION_LIEBERSBRONN                        = "Georgii-Waldstadion Liebersbronn"
	GERADSTETTEN_GERADSTETTEN                              = "Geradstetten Geradstetten"
	GERBERSTRASSE_BACKNANG                                 = "Gerberstraße Backnang"
	GERIATRISCHES_ZENTRUM_KENNENBURG                       = "Geriatrisches Zentrum Kennenburg"
	GERLINGEN_GERLINGEN                                    = "Gerlingen Gerlingen"
	GERLINGER_STRASSE_RAMTEL                               = "Gerlinger Straße Ramtel"
	GERLINGER_STRASSE_STUTTGART                            = "Gerlinger Straße Stuttgart"
	GERLINGER_TOR_GERLINGEN                                = "Gerlinger Tor Gerlingen"
	GERMANENSTRASSE_HOLZGERLINGEN                          = "Germanenstraße Holzgerlingen"
	GEROKSRUHE_STUTTGART                                   = "Geroksruhe Stuttgart"
	GEROKSTRASSE_BISSINGEN_LB                              = "Gerokstraße Bissingen (LB)"
	GEROKWEG_KUCHEN_F                                      = "Gerokweg Kuchen (F)"
	GERTRUDBAUMERWEG_BACKNANG                              = "Gertrud-Bäumer-Weg Backnang"
	GESCHICHTSHAUS_OWEN                                    = "Geschichtshaus Owen"
	GESCHWISTERSCHOLLGYMNASIUM_STUTTGART                   = "Geschwister-Scholl-Gymnasium Stuttgart"
	GESTUTSHOF_SCHARNHAUSEN                                = "Gestütshof Scharnhausen"
	GESTUTSWEG_WEIL                                        = "Gestütsweg Weil"
	GESUNDHEITSZENTRUM_BACKNANG                            = "Gesundheitszentrum Backnang"
	GEWERBE_TURKHEIM                                       = "Gewerbe Türkheim"
	GEWERBEGEBIET_STETTEN_LEINFELDENECHT                   = "Gewerbegebiet Stetten (Leinfelden-Echt.)"
	GEWERBEGEBIET_ALFDORF                                  = "Gewerbegebiet Alfdorf"
	GEWERBEGEBIET_FREUDENTAL                               = "Gewerbegebiet Freudental"
	GEWERBEGEBIET_NUFRINGEN                                = "Gewerbegebiet Nufringen"
	GEWERBEGEBIET_THOMASHARDT                              = "Gewerbegebiet Thomashardt"
	GEWERBEGEBIET_NELLMERSBACH                             = "Gewerbegebiet Nellmersbach"
	GEWERBEGEBIET_SPARWIESEN                               = "Gewerbegebiet Sparwiesen"
	GEWERBEGEBIET_MERKLINGEN                               = "Gewerbegebiet Merklingen"
	GEWERBEGEBIET_STEIGE_RUTESHEIM                         = "Gewerbegebiet Steige Rutesheim"
	GEWERBEGEBIET_SUD_HEIMERDINGEN                         = "Gewerbegebiet Süd Heimerdingen"
	GEWERBEGEBIET_SUD_DITZINGEN                            = "Gewerbegebiet Süd Ditzingen"
	GEWERBEPARK_ESCHENBACH                                 = "Gewerbepark Eschenbach"
	GEWERBESTRASSE_STUTTGART                               = "Gewerbestraße Stuttgart"
	GEWERBESTRASSE_AKADEMIE_WAIBLINGEN                     = "Gewerbestraße (Akademie) Waiblingen"
	GIEBEL_STUTTGART                                       = "Giebel Stuttgart"
	GIESHOF_SPIEGELBERG                                    = "Gieshof Spiegelberg"
	GINGEN_F_GINGEN_F                                      = "Gingen (F) Gingen (F)"
	GISELASTRASSE_WAIBLINGEN                               = "Giselastraße Waiblingen"
	GLASHUTTE_WALDENBUCH                                   = "Glashütte Waldenbuch"
	GLASPALAST_SINDELFINGEN                                = "Glaspalast Sindelfingen"
	GLAUNERWEG_STUTTGART                                   = "Glaunerweg Stuttgart"
	GLEMSAUE_DITZINGEN                                     = "Glemsaue Ditzingen"
	GLEMSBRUCKE_UNTERRIEXINGEN                             = "Glemsbrücke Unterriexingen"
	GLEMSECK_LEONBERG                                      = "Glemseck Leonberg"
	GLEMSECKSTRASSE_ELTINGEN                               = "Glemseckstraße Eltingen"
	GLEMSTAL_GERLINGEN                                     = "Glemstal Gerlingen"
	GLEMSTAL_SCHWIEBERDINGEN                               = "Glemstal Schwieberdingen"
	GLOCKENSTRASSE_STUTTGART                               = "Glockenstraße Stuttgart"
	GMEINWEILER_KAISERSBACH                                = "Gmeinweiler Kaisersbach"
	GNESENER_STRASSE_STUTTGART                             = "Gnesener Straße Stuttgart"
	GOCKELHOF_KIRCHENKIRNBERG                              = "Göckelhof Kirchenkirnberg"
	GOCKELHOF_KREUZUNG_FORNSBACH                           = "Göckelhof Kreuzung Fornsbach"
	GOERDELERSTRASSE_KORNWESTHEIM                          = "Goerdelerstraße Kornwestheim"
	GOETHEHOLDERLINSTR_MOTZINGEN                           = "Goethe-/Hölderlinstr. Mötzingen"
	GOETHESTR_FAURNDAU                                     = "Goethestr. Faurndau"
	GOETHESTRASSE_GERLINGEN                                = "Goethestraße Gerlingen"
	GOETHESTRASSE_WENDLINGEN_N                             = "Goethestraße Wendlingen (N)"
	GOETHESTRASSE_DAGERSHEIM                               = "Goethestraße Dagersheim"
	GOETHESTRASSE_KOHLBERG                                 = "Goethestraße Kohlberg"
	GOETHESTRASSE_ALDINGEN                                 = "Goethestraße Aldingen"
	GOETHESTRASSE_MOTZINGEN                                = "Goethestraße Mötzingen"
	GOLDACKER_STEINENBRONN                                 = "Goldäcker Steinenbronn"
	GOLDBERG_BOBLINGEN                                     = "Goldberg Böblingen"
	GOLDBERG_REALSCHULE_SINDELFINGEN                       = "Goldberg Realschule Sindelfingen"
	GOLDBERG_WASSERTURM_SINDELFINGEN                       = "Goldberg Wasserturm Sindelfingen"
	GOLDBERGGYMNASIUM_SINDELFINGEN                         = "Goldberg-Gymnasium Sindelfingen"
	GOLDBODEN_WINTERBACH                                   = "Goldboden Winterbach"
	GOLDMORGEN_DETTINGEN_T                                 = "Goldmorgen Dettingen (T)"
	GOLDMUHLESTRASSE_SINDELFINGEN                          = "Goldmühlestraße Sindelfingen"
	GOLFPARK_GOPPINGEN                                     = "Golfpark Göppingen"
	GOLFPLATZ_LEONBERG                                     = "Golfplatz Leonberg"
	GOLFPLATZ_EGLOSHEIM                                    = "Golfplatz Eglosheim"
	GOLLENHOFER_STRASSE_WEILER_ZUM_STEIN                   = "Gollenhofer Straße Weiler zum Stein"
	GOPPINGEN_GOPPINGEN                                    = "Göppingen Göppingen"
	GOPPINGER_STRASSE_RAMTEL                               = "Göppinger Straße Ramtel"
	GOPPINGER_STRASSE_ZELL_A                               = "Göppinger Straße Zell (A)"
	GOTTLIEBDAIMLERSCHULE_SINDELFINGEN                     = "Gottlieb-Daimler-Schule Sindelfingen"
	GOTTLIEBDAIMLERSTRASSE_NAGOLD                          = "Gottlieb-Daimler-Straße Nagold"
	GOTTLOBBAUKNECHTPLATZ_WELZHEIM                         = "Gottlob-Bauknecht-Platz Welzheim"
	GOTTLOBKAMMPLATZ_SCHORNDORF                            = "Gottlob-Kamm-Platz Schorndorf"
	GOTTSCHEER_STRASSE_SINDELFINGEN                        = "Gottscheer Straße Sindelfingen"
	GRABENSTR_STADTM_VAIHINGEN_E                           = "Grabenstr. (Stadtm.) Vaihingen (E)"
	GRABENSTRASSE_MARBACH_N                                = "Grabenstraße Marbach (N)"
	GRABENSTRASSE_WEIL_DER_STADT                           = "Grabenstraße Weil der Stadt"
	GRABENSTRASSE_PLUDERHAUSEN                             = "Grabenstraße Plüderhausen"
	GRABENSTRASSE_GARTRINGEN                               = "Grabenstraße Gärtringen"
	GRABENSTRASSE_BAHNHOF_SCHORNDORF                       = "Grabenstraße (Bahnhof) Schorndorf"
	GRAFDEGENFELDSTRASSE_RECHBERGHAUSEN                    = "Graf-Degenfeld-Straße Rechberghausen"
	GRAFENBERGER_STRASSE_KLEINBETTLINGEN                   = "Grafenberger Straße Kleinbettlingen"
	GRAFENBERGWEG_SCHORNDORF                               = "Grafenbergweg Schorndorf"
	GRAFENWEG_HERRENBERG                                   = "Grafenweg Herrenberg"
	GRAIRICH_KAISERSBACH                                   = "Grairich Kaisersbach"
	GRAUBACHTAL_HATTENHOFEN                                = "Graubachtal Hattenhofen"
	GRAUHALDE_SCHORNDORF                                   = "Grauhalde Schorndorf"
	GRAUHALDENHOF_RUDERSBERG                               = "Grauhaldenhof Rudersberg"
	GREUT_STUTTGART                                        = "Greut Stuttgart"
	GREUTACKERSTRASSE_MONCHBERG                            = "Greutäckerstraße Mönchberg"
	GREUTHOFLE_VORDERSTEINENBERG                           = "Greuthöfle Vordersteinenberg"
	GREUTTERSTRASSE_STUTTGART                              = "Greutterstraße Stuttgart"
	GRIEBENACKER_PLATTENHARDT                              = "Griebenäcker Plattenhardt"
	GROENERSTRASSE_LUDWIGSBURG                             = "Groenerstraße Ludwigsburg"
	GRONINGER_STRASSEBAUMHALDE_DITZINGEN                   = "Gröninger Straße/Baumhalde Ditzingen"
	GRONINGERHALDENSTRASSE_DITZINGEN                       = "Gröninger-/Haldenstraße Ditzingen"
	GRORACH_NEUENHAUS                                      = "Grörach Neuenhaus"
	GROSSE_GASSE_GINGEN_F                                  = "Große Gasse Gingen (F)"
	GROSSER_FORST_NURTINGEN                                = "Großer Forst Nürtingen"
	GROSSGLOCKNERSTRASSE_STUTTGART                         = "Großglocknerstraße Stuttgart"
	GROSSHOCHBERG_SPIEGELBERG                              = "Großhöchberg Spiegelberg"
	GROSSHOCHBERGER_STRASSE_SPIEGELBERG                    = "Großhöchberger Straße Spiegelberg"
	GROSSSACHSENHEIMER_STRASSE_KLEINSACHSENHEIM            = "Großsachsenheimer Straße Kleinsachsenheim"
	GROTZINGER_STRASSE_OBERENSINGEN                        = "Grötzinger Straße Oberensingen"
	GROTZSTRASSE_BISSINGEN_LB                              = "Grotzstraße Bissingen (LB)"
	GROTZSTRASSE_FRIEDHOF_BISSINGEN_LB                     = "Grotzstraße (Friedhof) Bissingen (LB)"
	GRUIBINGER_STRASSE_BAD_BOLL                            = "Gruibinger Straße Bad Boll"
	GRUNACKERSTRASSE_SINDELFINGEN                          = "Grünäckerstraße Sindelfingen"
	GRUNBACH_GRUNBACH                                      = "Grunbach Grunbach"
	GRUNBACH_DONZDORF                                      = "Grünbach Donzdorf"
	GRUND_UND_HAUPTSCH_GOLDBERG_SINDELFINGEN               = "Grund- und Hauptsch. Goldberg Sindelfingen"
	GRUNDGENSSTRASSE_STUTTGART                             = "Gründgensstraße Stuttgart"
	GRUNDSCHULE_GRUNBACH                                   = "Grundschule Grunbach"
	GRUNDSCHULE_KLEINGLATTBACH                             = "Grundschule Kleinglattbach"
	GRUNDSCHULE_SPARWIESEN                                 = "Grundschule Sparwiesen"
	GRUNDSCHULE_GRUIBINGEN                                 = "Grundschule Gruibingen"
	GRUNDSCHULE_ESCHENBACH                                 = "Grundschule Eschenbach"
	GRUNDSCHULE_BAD_DITZENBACH                             = "Grundschule Bad Ditzenbach"
	GRUNDSCHULE_RECHBERGHAUSEN                             = "Grundschule Rechberghausen"
	GRUNENBERGSTRASSE_GUTENBERG                            = "Grünenbergstraße Gutenberg"
	GRUNLINGWEG_STUTTGART                                  = "Grünlingweg Stuttgart"
	GRUNWIESENSTRASSE_BIETIGHEIM                           = "Grünwiesenstraße Bietigheim"
	GSCHWENDER_STRASSE_WELZHEIM                            = "Gschwender Straße Welzheim"
	GULTSTEIN_GULTSTEIN                                    = "Gültstein Gültstein"
	GUNTTERSTRASSE_MARBACH_N                               = "Güntterstraße Marbach (N)"
	GUSTAVRAUSTRASSE_BIETIGHEIM                            = "Gustav-Rau-Straße Bietigheim"
	GUSTAVWERNERSCHULE_WALDDORF                            = "Gustav-Werner-Schule Walddorf"
	GUTENBERGER_STRASSE_DONNSTETTEN                        = "Gutenberger Straße Donnstetten"
	GUTENBERGER_STRASSE_17_OBERLENNINGEN                   = "Gutenberger Straße 17 Oberlenningen"
	GUTENBERGER_STRASSE_47_OBERLENNINGEN                   = "Gutenberger Straße 47 Oberlenningen"
	GUTENBERGER_STRASSE_66_OBERLENNINGEN                   = "Gutenberger Straße 66 Oberlenningen"
	GUTENBERGSTRASSE_SCHMIDEN                              = "Gutenbergstraße Schmiden"
	GUTENHALDE_BONLANDEN                                   = "Gutenhalde Bonlanden"
	GUTTENBRUNNSTRASSE_SINDELFINGEN                        = "Guttenbrunnstraße Sindelfingen"
	GYMNASIUM_ASPERG                                       = "Gymnasium Asperg"
	GYMNASIUM_MURRHARDT                                    = "Gymnasium Murrhardt"
	GYMNASIUM_WEIL_DER_STADT                               = "Gymnasium Weil der Stadt"
	GYMNASIUM_HOLZGERLINGEN                                = "Gymnasium Holzgerlingen"
	GYMNASIUM_ECHTERDINGEN                                 = "Gymnasium Echterdingen"
	GYMNASIUM_KORNTAL                                      = "Gymnasium Korntal"
	HABICHTWEG_KIRCHHEIM_T                                 = "Habichtweg Kirchheim (T)"
	HAFENBAHNSTRASSE_STUTTGART                             = "Hafenbahnstraße Stuttgart"
	HAFNERSTRASSE_NEUENHAUS                                = "Häfnerstraße Neuenhaus"
	HAGENBUCH_DONZDORF                                     = "Hagenbuch Donzdorf"
	HAGER_SULZBACH_M                                       = "Hager Sulzbach (M)"
	HAGHOF_PFAHLBRONN                                      = "Haghof Pfahlbronn"
	HAGWEG_UNTERKIRNECK                                    = "Hagweg Unterkirneck"
	HAHNWEIDSTRASSE_KIRCHHEIM_T                            = "Hahnweidstraße Kirchheim (T)"
	HAIER_HAIERSCHULE_FAURNDAU                             = "Haier Haierschule Faurndau"
	HAIER_WALDORFSCHULE_FAURNDAU                           = "Haier Waldorfschule Faurndau"
	HAIGST_STUTTGART                                       = "Haigst Stuttgart"
	HAILFINGER_STRASSE_BONDORF                             = "Hailfinger Straße Bondorf"
	HAILINGSTRASSE_GOPPINGEN                               = "Hailingstraße Göppingen"
	HAINBACH_WALDENBRONN                                   = "Hainbach Wäldenbronn"
	HALDE_OTLINGEN                                         = "Halde Ötlingen"
	HALDE_NECKARHAUSEN                                     = "Halde Neckarhausen"
	HALDE_HEMMINGEN                                        = "Hälde Hemmingen"
	HALDENRAINWEG_ALTBACH                                  = "Haldenrainweg Altbach"
	HALDENSTRASSE_OTLINGEN                                 = "Haldenstraße Ötlingen"
	HALLENBAD_WALDENBUCH                                   = "Hallenbad Waldenbuch"
	HALLENBAD_MAICHINGEN                                   = "Hallenbad Maichingen"
	HALLENBAD_HERRENBERG                                   = "Hallenbad Herrenberg"
	HALLENBAD_BOBLINGEN                                    = "Hallenbad Böblingen"
	HALLENBAD_KORNWESTHEIM                                 = "Hallenbad Kornwestheim"
	HALLENBAD_BISSINGEN_LB                                 = "Hallenbad Bissingen (LB)"
	HALLENBAD_EISLINGEN_F                                  = "Hallenbad Eislingen (F)"
	HALLSCHLAG_STUTTGART                                   = "Hallschlag Stuttgart"
	HAMMERSCHLAG_SCHORNDORF                                = "Hammerschlag Schorndorf"
	HAMMERSCHMIEDE_MURRHARDT                               = "Hammerschmiede Murrhardt"
	HANDELSTRASSE_STUTTGART                                = "Händelstraße Stuttgart"
	HANDWERKSTRASSE_STUTTGART                              = "Handwerkstraße Stuttgart"
	HANFSTRASSE_DETTINGEN_T                                = "Hanfstraße Dettingen (T)"
	HANGELSTEIN_KIRCHE_ZELL_ESSLINGEN                      = "Hangelstein Kirche Zell (Esslingen)"
	HANGLICH_WANGEN_GP                                     = "Hanglich Wangen (GP)"
	HANGWEIDE_ROMMELSHAUSEN                                = "Hangweide Rommelshausen"
	HANNSKLEMMSTRASSE_OST_BOBLINGEN                        = "Hanns-Klemm-Straße Ost Böblingen"
	HANNSKLEMMSTRASSE_WEST_BOBLINGEN                       = "Hanns-Klemm-Straße West Böblingen"
	HANSBRUMMERPLATZ_LEINFELDEN                            = "Hans-Brümmer-Platz Leinfelden"
	HANSREHNSTIFT_STUTTGART                                = "Hans-Rehn-Stift Stuttgart"
	HANSSEYFFSTRASSE_GOPPINGEN                             = "Hans-Seyff-Straße Göppingen"
	HANSTHOMAWEG_KIRCHHEIM_T                               = "Hans-Thoma-Weg Kirchheim (T)"
	HANSVONHUTTENPLATZ_HARDT                               = "Hans-von-Hutten-Platz Hardt"
	HARBACH_MURRHARDT                                      = "Harbach Murrhardt"
	HARDT_HARDT                                            = "Hardt Hardt"
	HARDTHOF_SCHWIEBERDINGEN                               = "Hardthof Schwieberdingen"
	HARDTLINDE_MURR                                        = "Hardtlinde Murr"
	HARDTSCHULE_EBERSBACH_F                                = "Hardtschule Ebersbach (F)"
	HARDTWALDHALLE_KLEINASPACH                             = "Hardtwaldhalle Kleinaspach"
	HARFENBRUCKE_SINDELFINGEN                              = "Harfenbrücke Sindelfingen"
	HARTENECKSTRASSE_FREIBERG_N                            = "Harteneckstraße Freiberg (N)"
	HARTENECKSTRASSE_LUDWIGSBURG                           = "Harteneckstraße Ludwigsburg"
	HARTHAUSER_WEG_SIELMINGEN                              = "Harthäuser Weg Sielmingen"
	HARTWEG_BESIGHEIM                                      = "Hartweg Besigheim"
	HASELBACHMUHLE_ABZWEIG_SULZBACH_M                      = "Haselbachmühle Abzweig Sulzbach (M)"
	HASELSTEINSTRASSE_BREUNINGSWEILER                      = "Haselsteinstraße Breuningsweiler"
	HASENHALDE_BACKNANG                                    = "Hasenhälde Backnang"
	HASENHEIM_HARTHAUSEN                                   = "Hasenheim Harthausen"
	HASENHOF_WALDENBUCH                                    = "Hasenhof Waldenbuch"
	HASLACH_DARMSHEIM                                      = "Häslach Darmsheim"
	HATTENHOFER_STRASSE_SCHLIERBACH                        = "Hattenhofer Straße Schlierbach"
	HAUBENWASEN_PFAHLBRONN                                 = "Haubenwasen Pfahlbronn"
	HAUBERSBRONN_HAUBERSBRONN                              = "Haubersbronn Haubersbronn"
	HAUFFSTRASSE_KORNTAL                                   = "Hauffstraße Korntal"
	HAUFFSTRASSE_OBERENSINGEN                              = "Hauffstraße Oberensingen"
	HAUFFSTRASSE_KORNWESTHEIM                              = "Hauffstraße Kornwestheim"
	HAUFFSTRASSE_RECHBERGHAUSEN                            = "Hauffstraße Rechberghausen"
	HAUGENNEST_OBERBOIHINGEN                               = "Haugennest Oberboihingen"
	HAUPTBAHNHOF_OBEN_STUTTGART                            = "Hauptbahnhof (oben) Stuttgart"
	HAUPTBAHNHOF_TIEF_STUTTGART                            = "Hauptbahnhof (tief) Stuttgart"
	HAUPTBF_ARNULFKLETTPLATZ_STUTTGART                     = "Hauptbf (Arnulf-Klett-Platz) Stuttgart"
	HAUPTFRIEDHOF_STUTTGART                                = "Hauptfriedhof Stuttgart"
	HAUPTSTRBAHNHOF_EISLINGEN_F                            = "Hauptstr./Bahnhof Eislingen (F)"
	HAUPTSTRASSE_BERNHAUSEN                                = "Hauptstraße Bernhausen"
	HAUPTSTRASSE_BONLANDEN                                 = "Hauptstraße Bonlanden"
	HAUPTSTRASSE_GROSSBOTTWAR                              = "Hauptstraße Großbottwar"
	HAUPTSTRASSE_MERKLINGEN                                = "Hauptstraße Merklingen"
	HAUPTSTRASSE_DAGERSHEIM                                = "Hauptstraße Dagersheim"
	HAUPTSTRASSE_HEGNACH                                   = "Hauptstraße Hegnach"
	HAUPTSTRASSE_KAISERSBACH                               = "Hauptstraße Kaisersbach"
	HAUPTSTRASSE_URBACH                                    = "Hauptstraße Urbach"
	HAUPTSTRASSE_LEUTENBACH                                = "Hauptstraße Leutenbach"
	HAUPTSTRASSE_RIELINGSHAUSEN                            = "Hauptstraße Rielingshausen"
	HAUPTSTRASSE_PLEIDELSHEIM                              = "Hauptstraße Pleidelsheim"
	HAUPTSTRASSE_HOPFIGHEIM                                = "Hauptstraße Höpfigheim"
	HAUPTSTRASSE_UNTERRIEXINGEN                            = "Hauptstraße Unterriexingen"
	HAUPTSTRASSE_GROSSSACHSENHEIM                          = "Hauptstraße Großsachsenheim"
	HAUPTSTRASSE_PEROUSE                                   = "Hauptstraße Perouse"
	HAUPTSTRASSE_WEIL_IM_SCHONBUCH                         = "Hauptstraße Weil im Schönbuch"
	HAUPTSTRASSE_HATTENHOFEN                               = "Hauptstraße Hattenhofen"
	HAUPTSTRASSE_MACHTOLSHEIM                              = "Hauptstraße Machtolsheim"
	HAUPTSTRASSE_77_WARMBRONN                              = "Hauptstraße 77 Warmbronn"
	HAUPTSTRASSE_99_NECKARWEIHINGEN                        = "Hauptstraße 99 Neckarweihingen"
	HAUS_AM_REMSUFER_NECKARREMS                            = "Haus am Remsufer Neckarrems"
	HAUS_AM_WEINBERG_STUTTGART                             = "Haus am Weinberg Stuttgart"
	HAUS_AN_DER_METTER_BIETIGHEIM                          = "Haus an der Metter Bietigheim"
	HAUS_DER_BURGER_ALDINGEN                               = "Haus der Bürger Aldingen"
	HAUS_DER_FEUERWEHR_ALDINGEN                            = "Haus der Feuerwehr Aldingen"
	HAUS_DER_JUGEND_NECKARGRONINGEN                        = "Haus der Jugend Neckargröningen"
	HAUS_DER_VEREINE_DAGERSHEIM                            = "Haus der Vereine Dagersheim"
	HAUS_FRANK_MARKGRONINGEN                               = "Haus Frank Markgröningen"
	HAUSACKER_SIELMINGEN                                   = "Hausäcker Sielmingen"
	HAUSEN_MURRHARDT                                       = "Hausen Murrhardt"
	HAUSENER_STRASSE_MERKLINGEN                            = "Hausener Straße Merklingen"
	HAUSENRING_STUTTGART                                   = "Hausenring Stuttgart"
	HAUSGARTEN_WAIBLINGEN                                  = "Hausgärten Waiblingen"
	HAUSWEINBERG_BEINSTEIN                                 = "Hausweinberg Beinstein"
	HAYDNSTRASSE_HOFINGEN                                  = "Haydnstraße Höfingen"
	HAYDNSTRASSEBADER_SCHORNDORF                           = "Haydnstraße/Bäder Schorndorf"
	HEBSACKER_STRASSE_STUTTGART                            = "Hebsacker Straße Stuttgart"
	HECKBACHSTRASSE_KLEINHEPPACH                           = "Heckbachstraße Kleinheppach"
	HECKENGAUSTRASSE_NAGOLD                                = "Heckengäustraße Nagold"
	HECKENROSENWEG_NAGOLD                                  = "Heckenrosenweg Nagold"
	HECKENWEG_BESIGHEIM                                    = "Heckenweg Besigheim"
	HECKENWEG_HERTMANNSWEILER                              = "Heckenweg Hertmannsweiler"
	HEDELFINGEN_STUTTGART                                  = "Hedelfingen Stuttgart"
	HEDELFINGER_STRASSE_STUTTGART                          = "Hedelfinger Straße Stuttgart"
	HEERSTRASSE_STUTTGART                                  = "Heerstraße Stuttgart"
	HEERSTRASSE_GULTSTEIN                                  = "Heerstraße Gültstein"
	HEERWEG_GROSSBETTLINGEN                                = "Heerweg Großbettlingen"
	HEGELSEIDENSTRASSE_STUTTGART                           = "Hegel-/Seidenstraße Stuttgart"
	HEGELGYMNASIUM_STUTTGART                               = "Hegel-Gymnasium Stuttgart"
	HEGELSTRASSE_KIRCHHEIM_T                               = "Hegelstraße Kirchheim (T)"
	HEGENLOHER_STRASSE_THOMASHARDT                         = "Hegenloher Straße Thomashardt"
	HEGNACHER_HOHE_WAIBLINGEN                              = "Hegnacher Höhe Waiblingen"
	HEGNACHER_STRASSE_HOHENACKER                           = "Hegnacher Straße Hohenacker"
	HEIDEHOFSTRASSE_STUTTGART                              = "Heidehofstraße Stuttgart"
	HEIDESTRASSE_NECKARHALDE                               = "Heidestraße Neckarhalde"
	HEIDHOFE_ALLEENSTRASSE_BOHMENKIRCH                     = "Heidhöfe Alleenstraße Böhmenkirch"
	HEIDHOFE_SCHAFHAUS_BOHMENKIRCH                         = "Heidhöfe Schafhaus Böhmenkirch"
	HEIGELINSTRASSE_STUTTGART                              = "Heigelinstraße Stuttgart"
	HEILBAD_HOHENECK                                       = "Heilbad Hoheneck"
	HEILBRONNER_STRASSE_LUDWIGSBURG                        = "Heilbronner Straße Ludwigsburg"
	HEILIGENACKER_GEISLINGEN_STEIGE                        = "Heiligenäcker Geislingen (Steige)"
	HEIMATMUSEUM_UNTERWEISSACH                             = "Heimatmuseum Unterweissach"
	HEIMBACH_JEBENHAUSEN                                   = "Heimbach Jebenhausen"
	HEIMBERG_STUTTGART                                     = "Heimberg Stuttgart"
	HEIMERDINGEN_HEIMERDINGEN                              = "Heimerdingen Heimerdingen"
	HEININGERJAHNSTRASSE_GOPPINGEN                         = "Heininger-/Jahnstraße Göppingen"
	HEINKELSTRASSE_KIRCHHEIM_T                             = "Heinkelstraße Kirchheim (T)"
	HEINKELSTRASSE_BOBLINGEN                               = "Heinkelstraße Böblingen"
	HEINKELSTRASSE_LOCHGAU                                 = "Heinkelstraße Löchgau"
	HEINLESMUHLE_VORDERSTEINENBERG                         = "Heinlesmühle Vordersteinenberg"
	HEINRICHESSIGSTRASSE_LEONBERG                          = "Heinrich-Essig-Straße Leonberg"
	HELENELANGESCHULE_MARKGRONINGEN                        = "Helene-Lange-Schule Markgröningen"
	HELENENBURGWEG_BIETIGHEIM                              = "Helenenburgweg Bietigheim"
	HELFFERICHSTRASSE_STUTTGART                            = "Helfferichstraße Stuttgart"
	HELLERSHOF_VORDERSTEINENBERG                           = "Hellershof Vordersteinenberg"
	HELLERSHOF_KREUZUNG_VORDERSTEINENBERG                  = "Hellershof Kreuzung Vordersteinenberg"
	HEMMINGEN_HEMMINGEN                                    = "Hemmingen Hemmingen"
	HEMMINGER_STRASSE_STUTTGART                            = "Hemminger Straße Stuttgart"
	HEMMINGER_STRASSE_HEIMERDINGEN                         = "Hemminger Straße Heimerdingen"
	HENNERSDORFER_STRASSE_NEUWEILER                        = "Hennersdorfer Straße Neuweiler"
	HENRIETTENSTRASSE_KIRCHHEIM_T                          = "Henriettenstraße Kirchheim (T)"
	HERBSTHALDE_STUTTGART                                  = "Herbsthalde Stuttgart"
	HERDERPLATZ_STUTTGART                                  = "Herderplatz Stuttgart"
	HERDERSTRASSE_OBERESSLINGEN                            = "Herderstraße Oberesslingen"
	HERDWEG_DITZINGEN                                      = "Herdweg Ditzingen"
	HERDWEG_JEBENHAUSEN                                    = "Herdweg Jebenhausen"
	HERDWEG_POSTAMT_BOBLINGEN                              = "Herdweg (Postamt) Böblingen"
	HERMANNESSIGSTRASSE_SCHWIEBERDINGEN                    = "Hermann-Essig-Straße Schwieberdingen"
	HERMANNHESSEREALSCHULE_GOPPINGEN                       = "Hermann-Hesse-Realschule Göppingen"
	HERMANNHESSESTRASSE_AIDLINGEN                          = "Hermann-Hesse-Straße Aidlingen"
	HERMANNHESSESTRASSE_NECKARWEIHINGEN                    = "Hermann-Hesse-Straße Neckarweihingen"
	HERMANNLONSSTRASSE_NURTINGEN                           = "Hermann-Löns-Straße Nürtingen"
	HERMANNLONSSTRASSE_GERLINGEN                           = "Hermann-Löns-Straße Gerlingen"
	HERMANNLONSWEG_WINNENDEN                               = "Hermann-Löns-Weg Winnenden"
	HERMANNSCHWABHALLE_WINNENDEN                           = "Hermann-Schwab-Halle Winnenden"
	HERRENBACHSTRASSE_BAIERECK                             = "Herrenbachstraße Baiereck"
	HERRENBERG_HERRENBERG                                  = "Herrenberg Herrenberg"
	HERRENBERGER_STRASSE_MAICHINGEN                        = "Herrenberger Straße Maichingen"
	HERRENBERGER_STRASSE_BOBLINGEN                         = "Herrenberger Straße Böblingen"
	HERRENBERGER_STRASSE_KAYH                              = "Herrenberger Straße Kayh"
	HERRENBERGER_STRASSE_OBERJETTINGEN                     = "Herrenberger Straße Oberjettingen"
	HERRENWIESENWEG_NECKARHALDE                            = "Herrenwiesenweg Neckarhalde"
	HERRSCHAFTSWEG_PFLUGFELDEN                             = "Herrschaftsweg Pflugfelden"
	HERTICHSTRASSE_33_ELTINGEN                             = "Hertichstraße 33 Eltingen"
	HERTICHSTRASSE_73_ELTINGEN                             = "Hertichstraße 73 Eltingen"
	HERWEGHSTRASSE_STUTTGART                               = "Herweghstraße Stuttgart"
	HERZOGPHILIPPPLATZ_PARKSIEDLUNG                        = "Herzog-Philipp-Platz Parksiedlung"
	HERZOGWEG_HERRENBERG                                   = "Herzogweg Herrenberg"
	HESLACH_VOGELRAIN_STUTTGART                            = "Heslach Vogelrain Stuttgart"
	HESSIGHEIMER_STRASSE_MUNDELSHEIM                       = "Hessigheimer Straße Mundelsheim"
	HEUBACHSTRASSE_GOPPINGEN                               = "Heubachstraße Göppingen"
	HEUBERGSTRASSE_KORNWESTHEIM                            = "Heubergstraße Kornwestheim"
	HEUMADEN_STUTTGART                                     = "Heumaden Stuttgart"
	HEUMADEN_KAISERSBACH                                   = "Heumaden Kaisersbach"
	HEUMADEN_BOCKELSTRASSE_STUTTGART                       = "Heumaden Bockelstraße Stuttgart"
	HEUMADEN_ROSE_STUTTGART                                = "Heumaden Rose Stuttgart"
	HEUMADEN_SCHULE_STUTTGART                              = "Heumaden Schule Stuttgart"
	HEUSEE_PLUDERHAUSEN                                    = "Heusee Plüderhausen"
	HEUSTEIGSTRASSE_BOBLINGEN                              = "Heusteigstraße Böblingen"
	HEUTINGSHEIMER_STRASSE_STUTTGART                       = "Heutingsheimer Straße Stuttgart"
	HEUWEG_RUTESHEIM                                       = "Heuweg Rutesheim"
	HEYDSCHULE_MARKGRONINGEN                               = "Heyd-Schule Markgröningen"
	HIEBERSCHULE_UHINGEN                                   = "Hieberschule Uhingen"
	HILDENSTRASSE_GUNDELBACH                               = "Hildenstraße Gündelbach"
	HILDRIZHAUSER_STRASSE_HERRENBERG                       = "Hildrizhauser Straße Herrenberg"
	HILLERLOCHGAUER_STRASSE_BIETIGHEIM                     = "Hiller-/Löchgauer Straße Bietigheim"
	HILLERPLATZ_BIETIGHEIM                                 = "Hillerplatz Bietigheim"
	HIMMELSLEITER_STUTTGART                                = "Himmelsleiter Stuttgart"
	HINDENBURGSTRASSE_NELLINGEN                            = "Hindenburgstraße Nellingen"
	HINDENBURGSTRASSE_MOGLINGEN                            = "Hindenburgstraße Möglingen"
	HINDENBURGSTRASSE_WEISSACH_BB                          = "Hindenburgstraße Weissach (BB)"
	HINDENBURGSTRASSE_HERRENBERG                           = "Hindenburgstraße Herrenberg"
	HINDENBURGSTRASSE_DENKENDORF                           = "Hindenburgstraße Denkendorf"
	HINTERBUCHELBERG_MURRHARDT                             = "Hinterbüchelberg Murrhardt"
	HINTERE_SCHMIEDGASSE_SCHWABISCH_GMUND                  = "Hintere Schmiedgasse Schwäbisch Gmünd"
	HINTERE_TOS_BACKNANG                                   = "Hintere Tos Backnang"
	HINTERER_HOLZWEG_RUDERN                                = "Hinterer Holzweg Rüdern"
	HINTERSTEINENBERG_SUD_VORDERSTEINENBERG                = "Hintersteinenberg Süd Vordersteinenberg"
	HINTERWESTERMURR_FORNSBACH                             = "Hinterwestermurr Fornsbach"
	HIRSCH_GOSBACH                                         = "Hirsch Gosbach"
	HIRSCH_HAUSEN_BAD_UBERKINGEN                           = "Hirsch Hausen (Bad Überkingen)"
	HIRSCH_SCHNITTLINGEN                                   = "Hirsch Schnittlingen"
	HIRSCHBERGER_STRASSE_RAMTEL                            = "Hirschberger Straße Ramtel"
	HIRSCHBERGSTRASSE_GERLINGEN                            = "Hirschbergstraße Gerlingen"
	HIRSCHBERGSTRASSE_EGLOSHEIM                            = "Hirschbergstraße Eglosheim"
	HIRSCHGARTENWEG_MOGLINGEN                              = "Hirschgartenweg Möglingen"
	HIRSCHHOF_LENGLINGEN                                   = "Hirschhof Lenglingen"
	HIRSCHLANDER_STRASSE_DITZINGEN                         = "Hirschlander Straße Ditzingen"
	HIRSCHLANDER_STRASSE_HOFINGEN                          = "Hirschlander Straße Höfingen"
	HIRSCHLANDHOF_OBERESSLINGEN                            = "Hirschlandhof Oberesslingen"
	HIRSCHLANDSTRASSE_OBERESSLINGEN                        = "Hirschlandstraße Oberesslingen"
	HIRSCHPLATZ_FAURNDAU                                   = "Hirschplatz Faurndau"
	HIRSCHSTRASSE_OSSWEIL                                  = "Hirschstraße Oßweil"
	HIRSCHSTRASSE_ECHTERDINGEN                             = "Hirschstraße Echterdingen"
	HIRSCHSTRASSE_FAURNDAU                                 = "Hirschstraße Faurndau"
	HOCHBERGER_STRASSE_HOCHDORF_REMSECK                    = "Hochberger Straße Hochdorf (Remseck)"
	HOCHBERGER_STRASSE_BITTENFELD                          = "Hochberger Straße Bittenfeld"
	HOCHBERGER_STRASSE_POPPENWEILER                        = "Hochberger Straße Poppenweiler"
	HOCHBERGER_WALD_HOCHBERG                               = "Hochberger Wald Hochberg"
	HOCHDORFER_STRASSE_NOTZINGEN                           = "Hochdorfer Straße Notzingen"
	HOCHDORFER_STRASSE_BITTENFELD                          = "Hochdorfer Straße Bittenfeld"
	HOCHDORFER_STRASSE_EBERDINGEN                          = "Hochdorfer Straße Eberdingen"
	HOCHHAUS_REICHENBACH_F                                 = "Hochhaus Reichenbach (F)"
	HOCHHAUS_LORCH                                         = "Hochhaus Lorch"
	HOCHSCHULE_ESSLINGEN_N                                 = "Hochschule Esslingen (N)"
	HOCHSCHULE_GOPPINGEN                                   = "Hochschule Göppingen"
	HOCHSCHULZENTRUM_ESSLINGEN_N                           = "Hochschulzentrum Esslingen (N)"
	HOCHSTETTER_WEG_ZOLLBERG                               = "Hochstetter Weg Zollberg"
	HOCHWACHTTURM_WAIBLINGEN                               = "Hochwachtturm Waiblingen"
	HOCHWANG_RATHAUS_OBERLENNINGEN                         = "Hochwang Rathaus Oberlenningen"
	HOCHWIESENSTRASSE_BONDORF                              = "Hochwiesenstraße Bondorf"
	HOF_STETTEN_LEINFELDENECHT                             = "Hof Stetten (Leinfelden-Echt.)"
	HOF_SACHSENWEILER                                      = "Hof Sachsenweiler"
	HOFEN_STUTTGART                                        = "Hofen Stuttgart"
	HOFEN_HOFEN_BONNIGHEIM                                 = "Hofen Hofen (Bönnigheim)"
	HOFENER_STRASSE_BONNIGHEIM                             = "Hofener Straße Bönnigheim"
	HOFERSTRASSE_LUDWIGSBURG                               = "Hoferstraße Ludwigsburg"
	HOFFELD_STUTTGART                                      = "Hoffeld Stuttgart"
	HOFGARTENSTR_ROMMELSHAUSEN                             = "Hofgartenstr. Rommelshausen"
	HOFHALDE_BARTENBACH_GOPPINGEN                          = "Hofhalde Bartenbach (Göppingen)"
	HOFINGEN_HOFINGEN                                      = "Höfingen Höfingen"
	HOFMEISTERSCHAUWERK_SINDELFINGEN                       = "Hofmeister/Schauwerk Sindelfingen"
	HOFSTRASSE_HEGENSBERG                                  = "Hofstraße Hegensberg"
	HOFSTRASSE_DARMSHEIM                                   = "Hofstraße Darmsheim"
	HOFWIESENSTRASSE_GERLINGEN                             = "Hofwiesenstraße Gerlingen"
	HOGY_GOPPINGEN                                         = "Hogy Göppingen"
	HOHBAUM_PLUDERHAUSEN                                   = "Hohbaum Plüderhausen"
	HOHBERGSCHULE_PLUDERHAUSEN                             = "Hohbergschule Plüderhausen"
	HOHBERGSTRASSE_WEITMARS                                = "Hohbergstraße Weitmars"
	HOHE_EICHE_STUTTGART                                   = "Hohe Eiche Stuttgart"
	HOHEBILDSTRASSE_WALDDORF                               = "Hohebildstraße Walddorf"
	HOHENACKERSTRASSE_SCHMIDEN                             = "Hohenackerstraße Schmiden"
	HOHENBRACH_GRAB                                        = "Hohenbrach Grab"
	HOHENFREIBAD_STUTTGART                                 = "Höhenfreibad Stuttgart"
	HOHENKREUZ_HOHENKREUZ                                  = "Hohenkreuz Hohenkreuz"
	HOHENRAINSTRASSE_NECKARWEIHINGEN                       = "Hohenrainstraße Neckarweihingen"
	HOHENSTADTER_WEG_GOPPINGEN                             = "Hohenstädter Weg Göppingen"
	HOHENSTANGE_FRANKFURTER_STRASSE_TAMM                   = "Hohenstange Frankfurter Straße Tamm"
	HOHENSTANGE_STUTTGARTER_STRASSE_TAMM                   = "Hohenstange Stuttgarter Straße Tamm"
	HOHENSTANGE_TUBINGER_STRASSE_TAMM                      = "Hohenstange Tübinger Straße Tamm"
	HOHENSTANGE_ULMER_STRASSE_TAMM                         = "Hohenstange Ulmer Straße Tamm"
	HOHENSTAUFENSTRASSE_ASPERG                             = "Hohenstaufenstraße Asperg"
	HOHENSTAUFENSTRASSE_SCHORNDORF                         = "Hohenstaufenstraße Schorndorf"
	HOHENSTEINSTRASSE_STUTTGART                            = "Hohensteinstraße Stuttgart"
	HOHENSTEINSTRASSE_GOPPINGEN                            = "Hohensteinstraße Göppingen"
	HOHENSTRASSE_FELLBACH                                  = "Höhenstraße Fellbach"
	HOHENWEG_OBERBERKEN                                    = "Höhenweg Oberberken"
	HOHENZOLLERNPLATZ_LUDWIGSBURG                          = "Hohenzollernplatz Ludwigsburg"
	HOHENZOLLERNSTRASSE_PLEIDELSHEIM                       = "Hohenzollernstraße Pleidelsheim"
	HOHENZOLLERNSTRASSE_HOLZGERLINGEN                      = "Hohenzollernstraße Holzgerlingen"
	HOHENZOLLERNSTRASSE_SCHAFHAUSEN                        = "Hohenzollernstraße Schafhausen"
	HOHERBAUMWEG_NAGOLD                                    = "Hoher-Baum-Weg Nagold"
	HOHES_GESTADE_NURTINGEN                                = "Hohes Gestade Nürtingen"
	HOHEWARTSTRASSE_STUTTGART                              = "Hohewartstraße Stuttgart"
	HOHLWEG_STETTEN_LEINFELDENECHT                         = "Hohlweg Stetten (Leinfelden-Echt.)"
	HOHNWEILER_FORSTSTRASSE_LIPPOLDSWEILER                 = "Hohnweiler Forststraße Lippoldsweiler"
	HOHNWEILER_RATHAUS_LIPPOLDSWEILER                      = "Hohnweiler Rathaus Lippoldsweiler"
	HOHNWEILER_RATHAUSSTRASSE_LIPPOLDSWEILER               = "Hohnweiler Rathausstraße Lippoldsweiler"
	HOHREIN_ORTSMITTE_GOPPINGEN                            = "Hohrein Ortsmitte Göppingen"
	HOHROT_GROSSASPACH                                     = "Hohrot Großaspach"
	HOHWEG_FLACHT                                          = "Hohweg Flacht"
	HOLBEINWEG_KIRCHHEIM_T                                 = "Holbeinweg Kirchheim (T)"
	HOLDERACKER_STUTTGART                                  = "Holderäcker Stuttgart"
	HOLDERBUSCHLE_GROSSSACHSENHEIM                         = "Holderbüschle Großsachsenheim"
	HOLDERLINPLATZ_STUTTGART                               = "Hölderlinplatz Stuttgart"
	HOLDERLINSTRASSE_STUTTGART                             = "Hölderlinstraße Stuttgart"
	HOLDERLINSTRASSE_MARBACH_N                             = "Hölderlinstraße Marbach (N)"
	HOLDERLINWEG_URBACH                                    = "Hölderlinweg Urbach"
	HOLDERWEG_STETTEN_LEINFELDENECHT                       = "Holderweg Stetten (Leinfelden-Echt.)"
	HOLDIS_PFAHLBRONN                                      = "Höldis Pfahlbronn"
	HOLL_AICHELBERG_AICHWALD                               = "Holl Aichelberg (Aichwald)"
	HOLZBERGWEG_SCHORNDORF                                 = "Holzbergweg Schorndorf"
	HOLZGERLINGEN_HOLZGERLINGEN                            = "Holzgerlingen Holzgerlingen"
	HOLZGERLINGER_STRASSE_ALTDORF_BB                       = "Holzgerlinger Straße Altdorf (BB)"
	HOLZHAUSER_STRASSE_WANGEN_GP                           = "Holzhäuser Straße Wangen (GP)"
	HOLZHEIMER_STRASSE_EISLINGEN_F                         = "Holzheimer Straße Eislingen (F)"
	HOLZLACHSTRASSE_FRICKENHAUSEN                          = "Hölzlachstraße Frickenhausen"
	HOLZSTEIG_GULTSTEIN                                    = "Holzsteig Gültstein"
	HOLZWEILERHOF_WINZERHAUSEN                             = "Holzweilerhof Winzerhausen"
	HONOLDWEG_STUTTGART                                    = "Honoldweg Stuttgart"
	HOPFIGHEIMER_STRASSE_STEINHEIM_M                       = "Höpfigheimer Straße Steinheim (M)"
	HOPFIGHEIMER_STRASSE_BIETIGHEIM                        = "Höpfigheimer Straße Bietigheim"
	HORBSTRASSE_RUIT                                       = "Horbstraße Ruit"
	HORDTHOF_MURRHARDT                                     = "Hördthof Murrhardt"
	HORNBACH_ALDINGEN                                      = "Hornbach Aldingen"
	HORNBERGSTRLACHEBRUCKLE_DITZINGEN                      = "Hornbergstr./Lachebrückle Ditzingen"
	HORNBERGSTRASSE_KORNWESTHEIM                           = "Hornbergstraße Kornwestheim"
	HORNLE_DREIBRONNENSTRASSE_MARBACH_N                    = "Hörnle Dreibronnenstraße Marbach (N)"
	HORNLE_HOCHHAUS_MARBACH_N                              = "Hörnle Hochhaus Marbach (N)"
	HORNLE_WIESBADENER_PLATZ_MARBACH_N                     = "Hörnle Wiesbadener Platz Marbach (N)"
	HORNLESWEG_KOHLBERG                                    = "Hörnlesweg Kohlberg"
	HORRENWINKEL_STEINHEIM_M                               = "Horrenwinkel Steinheim (M)"
	HORSCHBACHSCHULE_MURRHARDT                             = "Hörschbachschule Murrhardt"
	HOSSLINSWART_HOSSLINSWART                              = "Hößlinswart Hößlinswart"
	HUBERTUSWEG_WIFLINGSHAUSEN                             = "Hubertusweg Wiflingshausen"
	HUGELSTRASSE_UNTERBRUDEN                               = "Hügelstraße Unterbrüden"
	HUGOWOLFWEG_KIRCHHEIM_T                                = "Hugo-Wolf-Weg Kirchheim (T)"
	HUGSTRASSE_GOPPINGEN                                   = "Hugstraße Göppingen"
	HULB_BOBLINGEN                                         = "Hulb Böblingen"
	HULBE_SCHWIEBERDINGEN                                  = "Hülbe Schwieberdingen"
	HULBEN_HOLZGERLINGEN                                   = "Hülben Holzgerlingen"
	HULBEN_NORD_HOLZGERLINGEN                              = "Hülben Nord Holzgerlingen"
	HUMBOLDTSTRASSE_LINSENHOFEN                            = "Humboldtstraße Linsenhofen"
	HUMMELBAUM_RENNINGEN                                   = "Hummelbaum Renningen"
	HUMMELBUHL_SULZBACH_M                                  = "Hummelbühl Sulzbach (M)"
	HUMPFENTALSTRASSE_NURTINGEN                            = "Humpfentalstraße Nürtingen"
	HUNDSACKER_STRUMPFELBACH_WEINSTADT                     = "Hundsäcker Strümpfelbach (Weinstadt)"
	HUNDSBERG_VORDERSTEINENBERG                            = "Hundsberg Vordersteinenberg"
	HUNGERBUHLSTRASSE_SCHORNDORF                           = "Hungerbühlstraße Schorndorf"
	HUOBER_ERDMANNHAUSEN                                   = "Huober Erdmannhausen"
	HUSARENHOF_BESIGHEIM                                   = "Husarenhof Besigheim"
	HUTTENBUHL_VORDERSTEINENBERG                           = "Hüttenbühl Vordersteinenberg"
	HUTTENBUHLSEE_ABZW_VORDERSTEINENBERG                   = "Hüttenbühlsee Abzw. Vordersteinenberg"
	HUTTLEN_SPIEGELBERG                                    = "Hüttlen Spiegelberg"
	HUTWIESENSTRASSE_MAGSTADT                              = "Hutwiesenstraße Magstadt"
	IBM_EHNINGEN                                           = "IBM Ehningen"
	IHINGER_HOF_RENNINGEN                                  = "Ihinger Hof Renningen"
	IKEA_LUDWIGSBURG                                       = "IKEA Ludwigsburg"
	IKEA_SINDELFINGEN                                      = "IKEA Sindelfingen"
	ILSFELDER_STRASSE_OTTMARSHEIM                          = "Ilsfelder Straße Ottmarsheim"
	ILTISWEG_FAURNDAU                                      = "Iltisweg Faurndau"
	IM_ALTEN_SEE_STEINENBRONN                              = "Im Alten See Steinenbronn"
	IM_BACKENSCHLAG_BONDORF                                = "Im Backenschlag Bondorf"
	IM_BAUMSTUCKLE_WAIBLINGEN                              = "Im Baumstückle Waiblingen"
	IM_BORNRAIN_MOGLINGEN                                  = "Im Bornrain Möglingen"
	IM_BRAUNKIEL_ALTBACH                                   = "Im Braunkiel Altbach"
	IM_CHAUSSEEFELD_STUTTGART                              = "Im Chausseefeld Stuttgart"
	IM_DEGEN_STUTTGART                                     = "Im Degen Stuttgart"
	IM_ELSENTAL_STUTTGART                                  = "Im Elsental Stuttgart"
	IM_ENGENLAUCH_BARTENBACH_GOPPINGEN                     = "Im Engenlauch Bartenbach (Göppingen)"
	IM_FELDLE_BIETIGHEIM                                   = "Im Feldle Bietigheim"
	IM_GAILER_GECHINGEN                                    = "Im Gailer Gechingen"
	IM_GANSLESGRUND_NURTINGEN                              = "Im Gänslesgrund Nürtingen"
	IM_GRUND_BAD_UBERKINGEN                                = "Im Grund Bad Überkingen"
	IM_HAG_STUTTGART                                       = "Im Hag Stuttgart"
	IM_HEGNACH_EBERSBACH_F                                 = "Im Hegnach Ebersbach (F)"
	IM_HILLER_MIEDELSBACH                                  = "Im Hiller Miedelsbach"
	IM_HIMMERREICH_STUTTGART                               = "Im Himmerreich Stuttgart"
	IM_HOLDER_PLIENSAUVORSTADT                             = "Im Holder Pliensauvorstadt"
	IM_HUMMERHOLZ_WEILER_ZUM_STEIN                         = "Im Hummerholz Weiler zum Stein"
	IM_KAISEMER_STUTTGART                                  = "Im Kaisemer Stuttgart"
	IM_KAPF_RUIT                                           = "Im Kapf Ruit"
	IM_KOLLER_SIELMINGEN                                   = "Im Köller Sielmingen"
	IM_KOPFERT_STUTTGART                                   = "Im Köpfert Stuttgart"
	IM_KUSTERFELD_BACKNANG                                 = "Im Kusterfeld Backnang"
	IM_LANGEN_FELD_MARKGRONINGEN                           = "Im langen Feld Markgröningen"
	IM_LAUCH_STUTTGART                                     = "Im Lauch Stuttgart"
	IM_LETTEN_EHNINGEN                                     = "Im Letten Ehningen"
	IM_MADER_STUTTGART                                     = "Im Mäder Stuttgart"
	IM_MOLDENGRABEN_KORNWESTHEIM                           = "Im Moldengraben Kornwestheim"
	IM_MORGEN_ALBERSHAUSEN                                 = "Im Morgen Albershausen"
	IM_MUNZEN_KIRCHHEIM_T                                  = "Im Münzen Kirchheim (T)"
	IM_RAISER_STUTTGART                                    = "Im Raiser Stuttgart"
	IM_SAMANN_WAIBLINGEN                                   = "Im Sämann Waiblingen"
	IM_SEE_KLEINGLATTBACH                                  = "Im See Kleinglattbach"
	IM_SEELE_HERRENBERG                                    = "Im Seele Herrenberg"
	IM_STEINGRABEN_HERRENBERG                              = "Im Steingraben Herrenberg"
	IM_STOCKACH_WEILHEIM_T                                 = "Im Stockach Weilheim (T)"
	IM_TAL_MUNKLINGEN                                      = "Im Tal Münklingen"
	IM_UNTEREN_ZEHEN_WIFLINGSHAUSEN                        = "Im Unteren Zehen Wiflingshausen"
	IM_VOGELSANG_NEUENHAUS                                 = "Im Vogelsang Neuenhaus"
	IM_VOGELSANG_HERRENBERG                                = "Im Vogelsang Herrenberg"
	IM_WALDECK_ASPERG                                      = "Im Waldeck Asperg"
	IM_WEIZEN_KORNWESTHEIM                                 = "Im Weizen Kornwestheim"
	IM_WIESENGRUND_BACKNANG                                = "Im Wiesengrund Backnang"
	IM_WIESENGRUND_KORNWESTHEIM                            = "Im Wiesengrund Kornwestheim"
	IMENTAL_UNTERJETTINGEN                                 = "Imental Unterjettingen"
	IMMANUELDORNFELDSTRASSE_NECKARWEIHINGEN                = "Immanuel-Dornfeld-Straße Neckarweihingen"
	IMMELMANNZENTRUM_GOPPINGEN                             = "Immelmannzentrum Göppingen"
	IMMO_NAGOLD                                            = "IMMO Nagold"
	IN_DEN_AULESWIESEN_ALLMERSBACH_IM_TAL                  = "In den Äuleswiesen Allmersbach im Tal"
	IN_DEN_BEETEN_GROSSINGERSHEIM                          = "In den Beeten Großingersheim"
	IN_DEN_ENTENACKERN_STUTTGART                           = "In den Entenäckern Stuttgart"
	IN_DEN_FRAUENGARTEN_GROSSBOTTWAR                       = "In den Frauengärten Großbottwar"
	IN_DEN_HOFACKERN_WEILER_ZUM_STEIN                      = "In den Hofäckern Weiler zum Stein"
	IN_DEN_MESSENWIESEN_ROSSWALDEN                         = "In den Messenwiesen Roßwälden"
	IN_DEN_STEGWIESEN_STUTTGART                            = "In den Stegwiesen Stuttgart"
	IN_DEN_WEINGARTEN_EISLINGEN_F                          = "In den Weingärten Eislingen (F)"
	IN_DER_MISSE_EBHAUSEN                                  = "In der Misse Ebhausen"
	IN_DER_WERRE_STUTTGART                                 = "In der Werre Stuttgart"
	INDEXSTRASSEBAHNHOF_OBERESSLINGEN                      = "Indexstraße/Bahnhof Oberesslingen"
	INDUSTRIEGEBIET_PLATTENHARDT                           = "Industriegebiet Plattenhardt"
	INDUSTRIEGEBIET_AICHSCHIESS                            = "Industriegebiet Aichschieß"
	INDUSTRIEGEBIET_WENDLINGEN_N                           = "Industriegebiet Wendlingen (N)"
	INDUSTRIEGEBIET_NABERN                                 = "Industriegebiet Nabern"
	INDUSTRIEGEBIET_UNTERRIEXINGEN                         = "Industriegebiet Unterriexingen"
	INDUSTRIEGEBIET_BIRKMANNSWEILER                        = "Industriegebiet Birkmannsweiler"
	INDUSTRIEGEBIET_ECHTERDINGEN                           = "Industriegebiet Echterdingen"
	INDUSTRIEGEBIET_ALBERSHAUSEN                           = "Industriegebiet Albershausen"
	INDUSTRIEGEBIET_BORTLINGEN                             = "Industriegebiet Börtlingen"
	INDUSTRIEGEBIET_OST_NEUHAUSEN_F                        = "Industriegebiet Ost Neuhausen (F)"
	INDUSTRIEGEBIET_OST_DITZINGEN                          = "Industriegebiet Ost Ditzingen"
	INDUSTRIEGEBIET_OST_GOPPINGEN                          = "Industriegebiet Ost Göppingen"
	INDUSTRIEGEBIET_SUD_JEBENHAUSEN                        = "Industriegebiet Süd Jebenhausen"
	INDUSTRIEGEBIET_WEST_NEUHAUSEN_F                       = "Industriegebiet West Neuhausen (F)"
	INDUSTRIEGEBIETFA_KNECHT_LORCH                         = "Industriegebiet/Fa. Knecht Lorch"
	INDUSTRIESTRASSE_STUTTGART                             = "Industriestraße Stuttgart"
	INDUSTRIESTRASSE_BACKNANG                              = "Industriestraße Backnang"
	INDUSTRIESTRASSE_MERKLINGEN                            = "Industriestraße Merklingen"
	INDUSTRIESTRASSE_SERSHEIM                              = "Industriestraße Sersheim"
	INGPARK_NORD_NAGOLD                                    = "INGpark Nord Nagold"
	INGPARK_SUD_NAGOLD                                     = "INGpark Süd Nagold"
	INSELBAD_ZIZISHAUSEN                                   = "Inselbad Zizishausen"
	INSELSTRASSE_STUTTGART                                 = "Inselstraße Stuttgart"
	ISARSTRASSE_WALDREMS                                   = "Isarstraße Waldrems"
	ISEGRIMWEG_STUTTGART                                   = "Isegrimweg Stuttgart"
	ISELSHAUSER_STRASSE_MOTZINGEN                          = "Iselshauser Straße Mötzingen"
	JAGERHAUS_LIEBERSBRONN                                 = "Jägerhaus Liebersbronn"
	JAHNHALLE_ENDERSBACH                                   = "Jahnhalle Endersbach"
	JAHNSTRASSE_KORNTAL                                    = "Jahnstraße Korntal"
	JAHNSTRASSE_WEILHEIM_T                                 = "Jahnstraße Weilheim (T)"
	JAHNSTRASSE_BESIGHEIM                                  = "Jahnstraße Besigheim"
	JAHNSTRASSE_POPPENWEILER                               = "Jahnstraße Poppenweiler"
	JAHNSTRASSE_WALDDORF                                   = "Jahnstraße Walddorf"
	JAKOBBUTTERSTRASSE_SCHMIDEN                            = "Jakob-/Butterstraße Schmiden"
	JAKOBRAIBLEANLAGE_SCHORNDORF                           = "Jakob-Raible-Anlage Schorndorf"
	JAKOBSTRASSE_BERKHEIM                                  = "Jakobstraße Berkheim"
	JELINSTRASSE_MOHRINGEN                                 = "Jelinstraße Möhringen"
	JENISCHSTRASSE_LUDWIGSBURG                             = "Jenischstraße Ludwigsburg"
	JENNERSTRASSE_KUPPINGEN                                = "Jennerstraße Kuppingen"
	JESISTRASSE_WAIBLINGEN                                 = "Jesistraße Waiblingen"
	JETTINGER_STRASSE_KUPPINGEN                            = "Jettinger Straße Kuppingen"
	JOHANNESKEPLERGYMNASIUM_LEONBERG                       = "Johannes-Kepler-Gymnasium Leonberg"
	JOHANNESKEPLERSCHULE_MAGSTADT                          = "Johannes-Kepler-Schule Magstadt"
	JOHANNESTAUFERKIRCHE_MAGSTADT                          = "Johannes-Täufer-Kirche Magstadt"
	JOHANNESGRABEN_STUTTGART                               = "Johannesgraben Stuttgart"
	JOHANNESKIRCHE_BERNHAUSEN                              = "Johanneskirche Bernhausen"
	JOHANNESKIRCHE_BACKNANG                                = "Johanneskirche Backnang"
	JOHANNESKIRCHE_STUTTGART                               = "Johanneskirche Stuttgart"
	JOHANNESSTRASSE_ZELL_ESSLINGEN                         = "Johannesstraße Zell (Esslingen)"
	JOHANNESSTRASSE_SCHORNDORF                             = "Johannesstraße Schorndorf"
	JOHANNESSTRASSE_KORNWESTHEIM                           = "Johannesstraße Kornwestheim"
	JOHANNESSTRASSE_HERRENBERG                             = "Johannesstraße Herrenberg"
	JOHNFKENNEDYSTR_GOPPINGEN                              = "John-F.-Kennedy-Str. Göppingen"
	JOSEPHHAYDNSTRASSE_WEIL_DER_STADT                      = "Joseph-Haydn-Straße Weil der Stadt"
	JUGENDDORF_VAIHINGEN_E                                 = "Jugenddorf Vaihingen (E)"
	JUGENDHERBERGE_MURRHARDT                               = "Jugendherberge Murrhardt"
	JUGENDHERBERGE_HOHENSTAUFEN                            = "Jugendherberge Hohenstaufen"
	JUNGBORN_NURTINGEN                                     = "Jungborn Nürtingen"
	JUNGE_WEINBERGE_WAIBLINGEN                             = "Junge Weinberge Waiblingen"
	JUNKERSSTRASSE_BOBLINGEN                               = "Junkersstraße Böblingen"
	JURASTRASSE_STUTTGART                                  = "Jurastraße Stuttgart"
	JUSIWEG_ZOLLBERG                                       = "Jusiweg Zollberg"
	JUX_JUX                                                = "Jux Jux"
	K_1206_BAHNHOF_REICHENBACH_F                           = "K 1206 (Bahnhof) Reichenbach (F)"
	KFERDBRAUNSTRASSE_BACKNANG                             = "K.-Ferd.-Braun-Straße Backnang"
	KHORNSCHUCHSTRASSE_URBACH                              = "K.-Hornschuch-Straße Urbach"
	KAISERLUDWIGSTRASSE_MURRHARDT                          = "Kaiser-Ludwig-Straße Murrhardt"
	KAISERBAU_AGENTUR_FUR_ARBEIT_GOPPINGEN                 = "Kaiserbau (Agentur für Arbeit) Göppingen"
	KAISEREICHE_SCHLICHTEN                                 = "Kaisereiche Schlichten"
	KAISERSTRASSE_NELLINGEN                                = "Kaiserstraße Nellingen"
	KAISERSTRASSE_BEUTELSBACH                              = "Kaiserstraße Beutelsbach"
	KALBERSTELLE_WEIL_IM_SCHONBUCH                         = "Kälberstelle Weil im Schönbuch"
	KALKOFEN_ENDERSBACH                                    = "Kalkofen Endersbach"
	KALKOFENSTRASSE_HERRENBERG                             = "Kalkofenstraße Herrenberg"
	KALLENBERG_ALTHUTTE                                    = "Kallenberg Althütte"
	KALLENBERG_LACKFABRIK_MUNCHINGEN                       = "Kallenberg Lackfabrik Münchingen"
	KALLENBERG_RASTHAUS_MUNCHINGEN                         = "Kallenberg Rasthaus Münchingen"
	KALTENTAL_STUTTGART                                    = "Kaltental Stuttgart"
	KALTENTHALSTRASSE_ALDINGEN                             = "Kaltenthalstraße Aldingen"
	KAMERALAMTSSTRASSE_STUTTGART                           = "Kameralamtsstraße Stuttgart"
	KAMMGARNSPINNEREI_BIETIGHEIM                           = "Kammgarnspinnerei Bietigheim"
	KAMMGARNSPINNEREI_B27_BIETIGHEIM                       = "Kammgarnspinnerei B27 Bietigheim"
	KANTSTRASSE_SIELMINGEN                                 = "Kantstraße Sielmingen"
	KANTSTRASSE_GEISLINGEN_STEIGE                          = "Kantstraße Geislingen (Steige)"
	KANTSTRASSE_GOPPINGEN                                  = "Kantstraße Göppingen"
	KAPELLE_PLOCHINGEN                                     = "Kapelle Plochingen"
	KAPELLENBERG_DOFFINGEN                                 = "Kapellenberg Döffingen"
	KAPELLENSTRASSE_WENDLINGEN_N                           = "Kapellenstraße Wendlingen (N)"
	KAPELLENSTRASSE_SUSSEN                                 = "Kapellenstraße Süßen"
	KAPF_KAPFHOFWEG_VORDERSTEINENBERG                      = "Kapf Kapfhofweg Vordersteinenberg"
	KAPF_MOHNWIESENWEG_VORDERSTEINENBERG                   = "Kapf Mohnwiesenweg Vordersteinenberg"
	KAPFFSTRASSE_URBACH                                    = "Kapffstraße Urbach"
	KAPPELBERGSTRASSE_FELLBACH                             = "Kappelbergstraße Fellbach"
	KAPPELGASSE_SCHWABISCH_GMUND                           = "Kappelgasse Schwäbisch Gmünd"
	KARLGEORGPFLEIDERERSTRASSE_HERTMANNSWEILER             = "Karl-Georg-Pfleiderer-Straße Hertmannsweiler"
	KARLJOOSSTRASSE_KORNWESTHEIM                           = "Karl-Joos-Straße Kornwestheim"
	KARLMAIALLEE_BIETIGHEIM                                = "Karl-Mai-Allee Bietigheim"
	KARLOLGAKRANKENHAUS_STUTTGART                          = "Karl-Olga-Krankenhaus Stuttgart"
	KARLPFAFFSTRASSE_STUTTGART                             = "Karl-Pfaff-Straße Stuttgart"
	KARLPFAFFSTRASSE_PLIENSAUVORSTADT                      = "Karl-Pfaff-Straße Pliensauvorstadt"
	KARLSHOF_KLEINASPACH                                   = "Karlshof Kleinaspach"
	KARLSHOHE_LUDWIGSBURG                                  = "Karlshöhe Ludwigsburg"
	KARLSPLATZ_LUDWIGSBURG                                 = "Karlsplatz Ludwigsburg"
	KARLSRUHER_ALLEE_PFLUGFELDEN                           = "Karlsruher Allee Pflugfelden"
	KARLSTRASSE_MUSBERG                                    = "Karlstraße Musberg"
	KARLSTRASSE_BERNHAUSEN                                 = "Karlstraße Bernhausen"
	KARLSTRASSE_BOBLINGEN                                  = "Karlstraße Böblingen"
	KARLSTRASSE_BISSINGEN_LB                               = "Karlstraße Bissingen (LB)"
	KARLSTRASSE_ROMMELSHAUSEN                              = "Karlstraße Rommelshausen"
	KARNSBERG_MURRHARDT                                    = "Karnsberg Murrhardt"
	KAROLINGERSTRASSE_SCHMIDEN                             = "Karolingerstraße Schmiden"
	KARPATENWEG_OTLINGEN                                   = "Karpatenweg Ötlingen"
	KASBACH_MURRHARDT                                      = "Käsbach Murrhardt"
	KASERNE_MOHRINGEN_STUTTGART                            = "Kaserne Möhringen Stuttgart"
	KASPARSWALD_STETTEN_LEINFELDENECHT                     = "Kasparswald Stetten (Leinfelden-Echt.)"
	KATH_KINDERGARTEN_LORCH                                = "Kath. Kindergarten Lorch"
	KATH_KIRCHE_GINGEN_F                                   = "Kath. Kirche Gingen (F)"
	KATHARINENHOF_STRUMPFELBACH_BACKNANG                   = "Katharinenhof Strümpfelbach (Backnang)"
	KATHARINENHOSPITAL_STUTTGART                           = "Katharinenhospital Stuttgart"
	KATHARINENSTAFFEL_ESSLINGEN_N                          = "Katharinenstaffel Esslingen (N)"
	KATHARINENSTRASSE_EGLOSHEIM                            = "Katharinenstraße Eglosheim"
	KATHARINENSTRASSE_GUNDELBACH                           = "Katharinenstraße Gündelbach"
	KATHEKOLLWITZSTRASSE_HERRENBERG                        = "Käthe-Kollwitz-Straße Herrenberg"
	KATHOLISCHE_KIRCHE_DETTINGEN_T                         = "Katholische Kirche Dettingen (T)"
	KATHOLISCHE_KIRCHE_BEMPFLINGEN                         = "Katholische Kirche Bempflingen"
	KATHOLISCHE_KIRCHE_AIDLINGEN                           = "Katholische Kirche Aidlingen"
	KATHOLISCHE_KIRCHE_GROSSHEPPACH                        = "Katholische Kirche Großheppach"
	KATHOLISCHE_KIRCHE_HEGNACH                             = "Katholische Kirche Hegnach"
	KATHOLISCHE_KIRCHE_OBERSTENFELD                        = "Katholische Kirche Oberstenfeld"
	KATHOLISCHE_KIRCHE_UNTERJETTINGEN                      = "Katholische Kirche Unterjettingen"
	KATHOLISCHE_KIRCHE_BESIGHEIM                           = "Katholische Kirche Besigheim"
	KATHOLISCHE_KIRCHE_ZELL_A                              = "Katholische Kirche Zell (A)"
	KATHOLISCHE_KIRCHE_GRUIBINGEN                          = "Katholische Kirche Gruibingen"
	KATZENBACHSTRASSE_STUTTGART                            = "Katzenbachstraße Stuttgart"
	KATZENKOPF_WALDENBRONN                                 = "Katzenkopf Wäldenbronn"
	KATZENLOCH_GEISLINGEN_STEIGE                           = "Katzenloch Geislingen (Steige)"
	KEGELPLATZ_STETTEN_I_R                                 = "Kegelplatz Stetten (i. R.)"
	KEHLSTRASSE_VAIHINGEN_E                                = "Kehlstraße Vaihingen (E)"
	KEILBERGSTRASSE_BOBLINGEN                              = "Keilbergstraße Böblingen"
	KELTENBURGSTRASSE_BOBLINGEN                            = "Keltenburgstraße Böblingen"
	KELTENMUSEUM_HOCHDORF_EBERDINGEN                       = "Keltenmuseum Hochdorf (Eberdingen)"
	KELTENSTRASSE_RENNINGEN                                = "Keltenstraße Renningen"
	KELTER_HOF_UND_LEMBACH                                 = "Kelter Hof und Lembach"
	KELTER_FRICKENHAUSEN                                   = "Kelter Frickenhausen"
	KELTER_STETTEN_I_R                                     = "Kelter Stetten (i. R.)"
	KELTER_BRUCH_WEISSACH                                  = "Kelter Bruch (Weissach)"
	KELTER_STRUMPFELBACH_WEINSTADT                         = "Kelter Strümpfelbach (Weinstadt)"
	KELTER_WINNENDEN                                       = "Kelter Winnenden"
	KELTER_SCHNAIT                                         = "Kelter Schnait"
	KELTER_STEINHEIM_M                                     = "Kelter Steinheim (M)"
	KELTER_GROSSBOTTWAR                                    = "Kelter Großbottwar"
	KELTERSTADION_ROMMELSHAUSEN                            = "Kelter/Stadion Rommelshausen"
	KELTERPLATZ_BESIGHEIM                                  = "Kelterplatz Besigheim"
	KELTERPLATZ_OBERBRUDEN                                 = "Kelterplatz Oberbrüden"
	KELTERRAIN_ECHTERDINGEN                                = "Kelterrain Echterdingen"
	KELTERSCHULE_NECKARREMS                                = "Kelterschule Neckarrems"
	KELTERSIEDLUNG_SCHLECHTBACH                            = "Keltersiedlung Schlechtbach"
	KELTERSTRASSE_STUTTGART                                = "Kelterstraße Stuttgart"
	KELTERSTRASSE_ROMMELSHAUSEN                            = "Kelterstraße Rommelshausen"
	KELTERSTRASSE_GRUNBACH                                 = "Kelterstraße Grunbach"
	KELTERSTRASSE_FRICKENHAUSEN                            = "Kelterstraße Frickenhausen"
	KELTERSTRASSE_BISSINGEN_LB                             = "Kelterstraße Bissingen (LB)"
	KELTERWEG_BURGSTALL_BURGSTETTEN                        = "Kelterweg Burgstall (Burgstetten)"
	KELTERWEG_SCHONAICH                                    = "Kelterweg Schönaich"
	KELTERWIESENBACH_GERADSTETTEN                          = "Kelterwiesenbach Geradstetten"
	KEMNATER_STRASSE_STUTTGART                             = "Kemnater Straße Stuttgart"
	KENNENBURG_KENNENBURG                                  = "Kennenburg Kennenburg"
	KEPLERSTRASSE_OBERESSLINGEN                            = "Keplerstraße Oberesslingen"
	KEPLERSTRASSE_RUTESHEIM                                = "Keplerstraße Rutesheim"
	KEPLERSTRASSE_DEIZISAU                                 = "Keplerstraße Deizisau"
	KEPLERSTRASSE_GOPPINGEN                                = "Keplerstraße Göppingen"
	KERNERSTEG_ALDINGEN                                    = "Kernersteg Aldingen"
	KERSCHENSTEINERSCHULE_STUTTGART                        = "Kerschensteinerschule Stuttgart"
	KESSELTOBELSTRASSE_FAURNDAU                            = "Kesseltobelstraße Faurndau"
	KIEBITZWEG_KIRCHHEIM_T                                 = "Kiebitzweg Kirchheim (T)"
	KIENBACHSTRASSE_STUTTGART                              = "Kienbachstraße Stuttgart"
	KIENBACHSTRASSE_FELLBACH                               = "Kienbachstraße Fellbach"
	KIESACKERFA_WOLLIN_WALDHAUSEN_B_SCHORND_LORCH          = "Kiesäcker/Fa. Wollin Waldhausen b Schornd (Lorch)"
	KIESELHOF_MURRHARDT                                    = "Kieselhof Murrhardt"
	KILLESBERG_STUTTGART                                   = "Killesberg Stuttgart"
	KIMMICHSWEILER_WEG_LIEBERSBRONN                        = "Kimmichsweiler Weg Liebersbronn"
	KINDERGARTEN_HOHENGEHREN                               = "Kindergarten Hohengehren"
	KINDERGARTEN_FORNSBACH                                 = "Kindergarten Fornsbach"
	KINDERGARTEN_METTERZIMMERN                             = "Kindergarten Metterzimmern"
	KINDERGARTEN_JUX                                       = "Kindergarten Jux"
	KINDERGARTEN_OCHSENBACH                                = "Kindergarten Ochsenbach"
	KINDERGARTEN_ROSSWALDEN                                = "Kindergarten Roßwälden"
	KINDERGARTEN_BAIERECK                                  = "Kindergarten Baiereck"
	KINDERGARTEN_NAGOLD                                    = "Kindergarten Nagold"
	KINDERGARTEN_ALBSTRASSE_ALDINGEN                       = "Kindergarten Albstraße Aldingen"
	KINDERGARTEN_LEONBERGER_STRASSE_ALDINGEN               = "Kindergarten Leonberger Straße Aldingen"
	KIRBACHHOF_OCHSENBACH                                  = "Kirbachhof Ochsenbach"
	KIRCHACKERSTRASSE_ST_BERNHARDT                         = "Kirchackerstraße St. Bernhardt"
	KIRCHBERG_GECHINGEN                                    = "Kirchberg Gechingen"
	KIRCHBERG_M_KIRCHBERG_M                                = "Kirchberg (M) Kirchberg (M)"
	KIRCHE_HARTHAUSEN                                      = "Kirche Harthausen"
	KIRCHE_MUSBERG                                         = "Kirche Musberg"
	KIRCHE_MUNCHINGEN                                      = "Kirche Münchingen"
	KIRCHE_RUIT                                            = "Kirche Ruit"
	KIRCHE_WEILER_SCHORNDORF                               = "Kirche Weiler (Schorndorf)"
	KIRCHE_LIEBERSBRONN                                    = "Kirche Liebersbronn"
	KIRCHE_OTLINGEN                                        = "Kirche Ötlingen"
	KIRCHE_SCHLICHTEN                                      = "Kirche Schlichten"
	KIRCHE_HEGENLOHE                                       = "Kirche Hegenlohe"
	KIRCHE_DAGERSHEIM                                      = "Kirche Dagersheim"
	KIRCHE_HASLACH                                         = "Kirche Haslach"
	KIRCHE_GULTSTEIN                                       = "Kirche Gültstein"
	KIRCHE_AFFSTATT                                        = "Kirche Affstätt"
	KIRCHE_NEUENHAUS                                       = "Kirche Neuenhaus"
	KIRCHE_OCHSENBACH                                      = "Kirche Ochsenbach"
	KIRCHE_STEINENBRONN                                    = "Kirche Steinenbronn"
	KIRCHE_MUHLHAUSEN_IM_TALE                              = "Kirche Mühlhausen im Täle"
	KIRCHE_SCHLAT                                          = "Kirche Schlat"
	KIRCHE_GAMMELSHAUSEN                                   = "Kirche Gammelshausen"
	KIRCHE_HOHENSTADT_GP                                   = "Kirche Hohenstadt (GP)"
	KIRCHE_ROSSWALDEN                                      = "Kirche Roßwälden"
	KIRCHE_NENNINGEN                                       = "Kirche Nenningen"
	KIRCHE_SCHLIERBACH                                     = "Kirche Schlierbach"
	KIRCHE_WEITMARS                                        = "Kirche Weitmars"
	KIRCHE_MACHTOLSHEIM                                    = "Kirche Machtolsheim"
	KIRCHE_SANKT_MAGNUS_WERNAU_N                           = "Kirche Sankt Magnus Wernau (N)"
	KIRCHHEIM_N_KIRCHHEIM_N                                = "Kirchheim (N) Kirchheim (N)"
	KIRCHHEIM_T_KIRCHHEIM_T                                = "Kirchheim (T) Kirchheim (T)"
	KIRCHHEIMER_STRHEGELSTRASSE_WEILHEIM_T                 = "Kirchheimer Str./Hegelstraße Weilheim (T)"
	KIRCHHEIMER_STRASSE_STUTTGART                          = "Kirchheimer Straße Stuttgart"
	KIRCHHEIMER_STRASSE_WENDLINGEN_N                       = "Kirchheimer Straße Wendlingen (N)"
	KIRCHHEIMER_STRASSE_NOTZINGEN                          = "Kirchheimer Straße Notzingen"
	KIRCHHEIMER_STRASSE_NURTINGEN                          = "Kirchheimer Straße Nürtingen"
	KIRCHHEIMER_STRASSE_BONNIGHEIM                         = "Kirchheimer Straße Bönnigheim"
	KIRCHHEIMER_STRASSE_KONGEN                             = "Kirchheimer Straße Köngen"
	KIRCHHEIMER_STRASSE_ZELL_A                             = "Kirchheimer Straße Zell (A)"
	KIRCHPLATZ_WELZHEIM                                    = "Kirchplatz Welzheim"
	KIRCHPLATZ_OBERBRUDEN                                  = "Kirchplatz Oberbrüden"
	KIRCHPLATZ_GRUNBACH                                    = "Kirchplatz Grunbach"
	KIRCHPLATZ_BAD_BOLL                                    = "Kirchplatz Bad Boll"
	KIRCHPLATZRATHAUS_RECHBERGHAUSEN                       = "Kirchplatz/Rathaus Rechberghausen"
	KIRCHSTRASSE_BOHMENKIRCH                               = "Kirchstraße Böhmenkirch"
	KIRCHSTRASSE_HEININGEN_GP                              = "Kirchstraße Heiningen (GP)"
	KIRCHSTRASSENELKENWEG_DEIZISAU                         = "Kirchstraße/Nelkenweg Deizisau"
	KIRCHTALSTRASSE_STUTTGART                              = "Kirchtalstraße Stuttgart"
	KIRCHTALSTRASSE_KORNWESTHEIM                           = "Kirchtalstraße Kornwestheim"
	KIRSCHENHARDTHOF_BURGSTALL_BURGSTETTEN                 = "Kirschenhardthof Burgstall (Burgstetten)"
	KIRSCHHALDENWEG_BESIGHEIM                              = "Kirschhaldenweg Besigheim"
	KITTENESHALDE_KIRCHHEIM_T                              = "Kitteneshalde Kirchheim (T)"
	KITZENER_STRASSE_OTTENBACH                             = "Kitzener Straße Ottenbach"
	KLAFFENBACH_RUDERSBERG                                 = "Klaffenbach Rudersberg"
	KLAFFENSTEINSTRASSE_BOBLINGEN                          = "Klaffensteinstraße Böblingen"
	KLAIBERSTRASSE_ENZWEIHINGEN                            = "Klaiberstraße Enzweihingen"
	KLEIDERWERK_ADLER_NECKARTENZLINGEN                     = "Kleiderwerk Adler Neckartenzlingen"
	KLEINBOTTWAR_KLEINBOTTWAR                              = "Kleinbottwar Kleinbottwar"
	KLEINBOTTWARER_STRASSE_STEINHEIM_M                     = "Kleinbottwarer Straße Steinheim (M)"
	KLEINER_OSTRING_STUTTGART                              = "Kleiner Ostring Stuttgart"
	KLEINER_SCHLOSSPLATZ_STUTTGART                         = "Kleiner Schlossplatz Stuttgart"
	KLEINERLACH_GROSSERLACH                                = "Kleinerlach Großerlach"
	KLEINFELDFRIEDHOF_FELLBACH                             = "Kleinfeldfriedhof Fellbach"
	KLEINGARTEN_SCHORNDORF                                 = "Kleingärten Schorndorf"
	KLEINGARTENANLAGE_JEBENHAUSEN                          = "Kleingartenanlage Jebenhausen"
	KLEINGLATTBACHER_STRASSE_ENSINGEN                      = "Kleinglattbacher Straße Ensingen"
	KLEINHEPPACHER_STRASSE_GROSSHEPPACH                    = "Kleinheppacher Straße Großheppach"
	KLEINHEPPACHER_STRASSE_BEINSTEIN                       = "Kleinheppacher Straße Beinstein"
	KLEINHOHENHEIM_STUTTGART                               = "Kleinhohenheim Stuttgart"
	KLEINSACHSENHEIMER_STRASSE_METTERZIMMERN               = "Kleinsachsenheimer Straße Metterzimmern"
	KLEINSACHSENHEIMER_STRASSE_GROSSSACHSENHEIM            = "Kleinsachsenheimer Straße Großsachsenheim"
	KLEISTSTRASSE_STUTTGART                                = "Kleiststraße Stuttgart"
	KLINGEN_MURRHARDT                                      = "Klingen Murrhardt"
	KLINGENMUHLE_WELZHEIM                                  = "Klingenmühle Welzheim"
	KLINGENSTRASSE_MUSBERG                                 = "Klingenstraße Musberg"
	KLINGENSTRASSE_AFFALTERBACH                            = "Klingenstraße Affalterbach"
	KLINGENSTRASSE_BITTENFELD                              = "Klingenstraße Bittenfeld"
	KLINIK_AM_EICHERT_GOPPINGEN                            = "Klinik am Eichert Göppingen"
	KLINIK_SCHILLERHOHE_GERLINGEN                          = "Klinik Schillerhöhe Gerlingen"
	KLINIKSCHULE_MARKGRONINGEN                             = "Klinik/Schule Markgröningen"
	KLINIKUM_LUDWIGSBURG                                   = "Klinikum Ludwigsburg"
	KLINIKUM_ESSLINGEN_OBERESSLINGEN                       = "Klinikum Esslingen Oberesslingen"
	KLINKERNFELD_BERNHAUSEN                                = "Klinkernfeld Bernhausen"
	KLOPFERBACH_GROSSASPACH                                = "Klöpferbach Großaspach"
	KLOSTER_LORCH                                          = "Kloster Lorch"
	KLOSTERGARTEN_SINDELFINGEN                             = "Klostergarten Sindelfingen"
	KLOSTERPLATZ_OEFFINGEN                                 = "Klosterplatz Oeffingen"
	KLOSTERSTRASSE_STETTEN_I_R                             = "Klosterstraße Stetten (i. R.)"
	KLOSTERWEG_KONGEN                                      = "Klosterweg Köngen"
	KNAPPENWEG_STUTTGART                                   = "Knappenweg Stuttgart"
	KNIEBISSTRASSE_MAGSTADT                                = "Kniebisstraße Magstadt"
	KOHLENWEG_ALDINGEN                                     = "Kohlenweg Aldingen"
	KOHLERWEG_HEGENSBERG                                   = "Kohlerweg Hegensberg"
	KOHLPLATTE_UNTERJETTINGEN                              = "Kohlplatte Unterjettingen"
	KOHLSTRASSE_LUDWIGSBURG                                = "Köhlstraße Ludwigsburg"
	KOLPINGSIEDLUNG_STUTTGART                              = "Kolpingsiedlung Stuttgart"
	KOLPINGSTRASSE_WENDLINGEN_N                            = "Kolpingstraße Wendlingen (N)"
	KONIGWILHELMPLATZ_MARBACH_N                            = "König-Wilhelm-Platz Marbach (N)"
	KONIGSALLEE_WEIL                                       = "Königsallee Weil"
	KONIGSBAU_PLUDERHAUSEN                                 = "Königsbau Plüderhausen"
	KONIGSBERGER_STRASSE_EHNINGEN                          = "Königsberger Straße Ehningen"
	KONIGSBERGER_STRASSE_WERNAU_N                          = "Königsberger Straße Wernau (N)"
	KONIGSBERGER_STRASSE_HERRENBERG                        = "Königsberger Straße Herrenberg"
	KONIGSBRONNHOF_RUDERSBERG                              = "Königsbronnhof Rudersberg"
	KONIGSKNOLL_SINDELFINGEN                               = "Königsknoll Sindelfingen"
	KONIGSTRASSLE_STUTTGART                                = "Königsträßle Stuttgart"
	KONRADHAUSSMANNWEG_SCHORNDORF                          = "Konrad-Haussmann-Weg Schorndorf"
	KONRADWIDERHOLTHALLE_KIRCHHEIM_T                       = "Konrad-Widerholt-Halle Kirchheim (T)"
	KORBER_STEIGE_WAIBLINGEN                               = "Korber Steige Waiblingen"
	KORBER_STRASSE_KLEINHEPPACH                            = "Korber Straße Kleinheppach"
	KORNBECKSTRASSE_LUDWIGSBURG                            = "Kornbeckstraße Ludwigsburg"
	KORNBERGSATTEL_GRUIBINGEN                              = "Kornbergsattel Gruibingen"
	KORNBERGWEG_PLOCHINGEN                                 = "Kornbergweg Plochingen"
	KORNHAUSPLATZ_GOPPINGEN                                = "Kornhausplatz Göppingen"
	KORNTAL_KORNTAL                                        = "Korntal Korntal"
	KORNTALER_STRASSE_STUTTGART                            = "Korntaler Straße Stuttgart"
	KORNTALER_STRASSE_LEONBERG                             = "Korntaler Straße Leonberg"
	KORNTALERSCHUBARTSTRASSE_DITZINGEN                     = "Korntaler-/Schubartstraße Ditzingen"
	KORNWESTHEIM_KORNWESTHEIM                              = "Kornwestheim Kornwestheim"
	KORNWESTHEIMER_STRASSE_STUTTGART                       = "Kornwestheimer Straße Stuttgart"
	KORNWESTHEIMER_STRASSE_MUNCHINGEN                      = "Kornwestheimer Straße Münchingen"
	KORSCHSTRASSE_ZELL_ESSLINGEN                           = "Körschstraße Zell (Esslingen)"
	KORSCHTAL_DEIZISAU                                     = "Körschtal Deizisau"
	KOTTWEIL_STEINACH_BERGLEN                              = "Kottweil Steinach (Berglen)"
	KRAFTWERK_MUNSTER_STUTTGART                            = "Kraftwerk Münster Stuttgart"
	KRAHERWALD_STUTTGART                                   = "Kräherwald Stuttgart"
	KRANKENHAUS_LEONBERG                                   = "Krankenhaus Leonberg"
	KRANKENHAUS_MARBACH_N                                  = "Krankenhaus Marbach (N)"
	KRANKENHAUS_SINDELFINGEN                               = "Krankenhaus Sindelfingen"
	KRANKENHAUS_BIETIGHEIM                                 = "Krankenhaus Bietigheim"
	KRANKENHAUS_VAIHINGEN_E                                = "Krankenhaus Vaihingen (E)"
	KRANKENHAUS_GEISLINGEN_STEIGE                          = "Krankenhaus Geislingen (Steige)"
	KRANKENHAUS_SEEBACH_GEISLINGEN_STEIGE                  = "Krankenhaus (Seebach) Geislingen (Steige)"
	KRANSTRASSE_STUTTGART                                  = "Kranstraße Stuttgart"
	KRAPFENREUT_EBERSBACH_F                                = "Krapfenreut Ebersbach (F)"
	KREHLSTRASSE_STUTTGART                                 = "Krehlstraße Stuttgart"
	KREHWINKEL_ASPERGLEN                                   = "Krehwinkel Asperglen"
	KREISBERUFSSCHULE_LUDWIGSBURG                          = "Kreisberufsschule Ludwigsburg"
	KREISBERUFSSCHULE_BIETIGHEIM                           = "Kreisberufsschule Bietigheim"
	KREISBERUFSSCHULZENTRUM_BACKNANG                       = "Kreisberufsschulzentrum Backnang"
	KREISKRANKENHAUS_BOBLINGEN                             = "Kreiskrankenhaus Böblingen"
	KREISKRANKENHAUS_B28_HERRENBERG                        = "Kreiskrankenhaus (B28) Herrenberg"
	KREISLERSTRASSE_SCHOPFLOCH_LENNINGEN                   = "Kreislerstraße Schopfloch (Lenningen)"
	KREISSPARKASSE_BALTMANNSWEILER                         = "Kreissparkasse Baltmannsweiler"
	KREISSPARKASSE_PLOCHINGEN                              = "Kreissparkasse Plochingen"
	KREISSPARKASSE_OBERBOIHINGEN                           = "Kreissparkasse Oberboihingen"
	KREISSPARKASSE_BORTLINGEN                              = "Kreissparkasse Börtlingen"
	KREISSTRASSE_UNTERKIRNECK                              = "Kreisstraße Unterkirneck"
	KREISTIERHEIM_BOBLINGEN                                = "Kreistierheim Böblingen"
	KREISVERKEHR_SCHANBACH                                 = "Kreisverkehr Schanbach"
	KREISVERKEHR_HIRSCHLANDEN                              = "Kreisverkehr Hirschlanden"
	KREISVERKEHR_GAMMELSHAUSEN                             = "Kreisverkehr Gammelshausen"
	KREMSER_STRASSE_BOBLINGEN                              = "Kremser Straße Böblingen"
	KRETTENHOF_GOPPINGEN                                   = "Krettenhof Göppingen"
	KREUZACKER_BONLANDEN                                   = "Kreuzäcker Bonlanden"
	KREUZBRUNNEN_SCHARNHAUSER_PARK                         = "Kreuzbrunnen Scharnhauser Park"
	KREUZSTRASSE_STUTTGART                                 = "Kreuzstraße Stuttgart"
	KREUZSTRASSE_OBERESSLINGEN                             = "Kreuzstraße Oberesslingen"
	KREUZSTRASSE_DATZINGEN                                 = "Kreuzstraße Dätzingen"
	KREUZSTRASSE_KAISERSBACH                               = "Kreuzstraße Kaisersbach"
	KREUZSTRASSE_OBERSTENFELD                              = "Kreuzstraße Oberstenfeld"
	KREUZUNG_AICHSCHIESS                                   = "Kreuzung Aichschieß"
	KREUZUNG_HEININGEN_BACKNANG                            = "Kreuzung Heiningen (Backnang)"
	KREUZUNG_AICHELBERG_GP                                 = "Kreuzung Aichelberg (GP)"
	KREUZUNG_ADELBERG_RECHBERGHAUSEN                       = "Kreuzung Adelberg Rechberghausen"
	KREUZUNG_ESCHENBACHE_EISLINGEN_F                       = "Kreuzung Eschenbäche Eislingen (F)"
	KREUZUNG_GRUIBINGEN_AUENDORF_GAMMELSHAUSEN             = "Kreuzung Gruibingen Auendorf Gammelshausen"
	KREUZWEGACKER_STEINHEIM_M                              = "Kreuzwegäcker Steinheim (M)"
	KRIEGERDENKMAL_HEGENSBERG                              = "Kriegerdenkmal Hegensberg"
	KRIEGERDENKMAL_GROTZINGEN                              = "Kriegerdenkmal Grötzingen"
	KRIEGSBERG_HOHENACKER                                  = "Kriegsberg Hohenacker"
	KRING_STEINENBRONN                                     = "Kring Steinenbronn"
	KROATENHOF_NURTINGEN                                   = "Kroatenhof Nürtingen"
	KRONE_GRONAU                                           = "Krone Gronau"
	KRONE_BAIERECK                                         = "Krone Baiereck"
	KRONEALTENSTADT_GEISLINGEN_STEIGE                      = "Krone/Altenstadt Geislingen (Steige)"
	KRONENBRUNNEN_PLATTENHARDT                             = "Kronenbrunnen Plattenhardt"
	KRONENPLATZ_WINNENDEN                                  = "Kronenplatz Winnenden"
	KRONENSTRASSE_BERKHEIM                                 = "Kronenstraße Berkheim"
	KRONENSTRASSE_AICHELBERG_AICHWALD                      = "Kronenstraße Aichelberg (Aichwald)"
	KRONENZENTRUM_BIETIGHEIM                               = "Kronenzentrum Bietigheim"
	KRONERHOF_AUFHAUSEN_GEISLINGEN                         = "Krönerhof Aufhausen (Geislingen)"
	KRUICHLING_KIRCHHEIM_T                                 = "Kruichling Kirchheim (T)"
	KRUMMBACHTAL_GERLINGEN                                 = "Krummbachtal Gerlingen"
	KRUMMENACKER_KRUMMENACKER                              = "Krummenacker Krummenacker"
	KRUMMHARDT_ABZWEIG_AICHSCHIESS                         = "Krummhardt Abzweig Aichschieß"
	KRUMMHARDT_ORT_KRUMMHARDT                              = "Krummhardt Ort Krummhardt"
	KRUMMW_SCHWARZENGASSE_EISLINGEN_F                      = "Krummw. Schwarzengasse Eislingen (F)"
	KRUMMWALDEN_BRUCKENSTR_EISLINGEN_F                     = "Krummwälden Brückenstr. Eislingen (F)"
	KUCHEN_KUCHEN_F                                        = "Kuchen Kuchen (F)"
	KUCHENGRUNDOVR_BACKNANG                                = "Kuchengrund/OVR Backnang"
	KUHWASEN_STUTTGART                                     = "Kühwasen Stuttgart"
	KULLENWIESEN_HARDT                                     = "Kullenwiesen Hardt"
	KUNKELINHALLE_SCHORNDORF                               = "Künkelinhalle Schorndorf"
	KUNKELINSCHULE_SCHORNDORF                              = "Künkelinschule Schorndorf"
	KUNSTAKADEMIE_STUTTGART                                = "Kunstakademie Stuttgart"
	KUNSTHALLE_GOPPINGEN                                   = "Kunsthalle Göppingen"
	KUPPINGER_STRASSE_NUFRINGEN                            = "Kuppinger Straße Nufringen"
	KURKLINIK_BAD_DITZENBACH                               = "Kurklinik Bad Ditzenbach"
	KURMARKER_KASERNE_STUTTGART                            = "Kurmärker Kaserne Stuttgart"
	KURSAAL_STUTTGART                                      = "Kursaal Stuttgart"
	KURTSCHUMACHERSTRASSE_STUTTGART                        = "Kurt-Schumacher-Straße Stuttgart"
	KURZACH_NASSACH                                        = "Kurzach Nassach"
	KURZE_STRASSE_GEMMRIGHEIM                              = "Kurze Straße Gemmrigheim"
	LACHENSTRASSE_SCHONAICH                                = "Lachenstraße Schönaich"
	LACHENTORSTRASSE_HOFINGEN                              = "Lachentorstraße Höfingen"
	LAIBLINGER_WEG_SCHWIEBERDINGEN                         = "Laiblinger Weg Schwieberdingen"
	LAICHLESTRASSE_GERLINGEN                               = "Laichlestraße Gerlingen"
	LAIHLE_STUTTGART                                       = "Laihle Stuttgart"
	LAILENSACKER_PLATTENHARDT                              = "Lailensäcker Plattenhardt"
	LAMBERTWEG_STUTTGART                                   = "Lambertweg Stuttgart"
	LAMM_NECKARGRONINGEN                                   = "Lamm Neckargröningen"
	LAMM_GOSBACH                                           = "Lamm Gosbach"
	LAMM_TREFFELHAUSEN                                     = "Lamm Treffelhausen"
	LAMMTAL_GARTRINGEN                                     = "Lammtal Gärtringen"
	LANDAUER_STRASSE_STUTTGART                             = "Landauer Straße Stuttgart"
	LANDHAUS_STUTTGART                                     = "Landhaus Stuttgart"
	LANDHAUSKREUZUNG_STUTTGART                             = "Landhauskreuzung Stuttgart"
	LANDHAUSSIEDLUNG_MAICHINGEN                            = "Landhaussiedlung Maichingen"
	LANDRATSAMT_LUDWIGSBURG                                = "Landratsamt Ludwigsburg"
	LANDRATSAMT_GOPPINGEN                                  = "Landratsamt Göppingen"
	LANGE_ANWANDEN_SINDELFINGEN                            = "Lange Anwanden Sindelfingen"
	LANGE_MORGEN_WEILHEIM_T                                = "Lange Morgen Weilheim (T)"
	LANGE_STRASSE_SIELMINGEN                               = "Lange Straße Sielmingen"
	LANGE_STRASSE_MUNDELSHEIM                              = "Lange Straße Mundelsheim"
	LANGE_WEIDEN_WINNENDEN                                 = "Lange Weiden Winnenden"
	LANGENBERG_WELZHEIM                                    = "Langenberg Welzheim"
	LANGENBUHL_ELTINGEN                                    = "Längenbühl Eltingen"
	LANGER_WEG_WALDENBRONN                                 = "Langer Weg Wäldenbronn"
	LANGHANS_BEILSTEIN                                     = "Langhans Beilstein"
	LANGMANTEL_HOHENHASLACH                                = "Langmantel Hohenhaslach"
	LANNERSTRASSE_OPPELSBOHM                               = "Lannerstraße Oppelsbohm"
	LAPP_KABEL_STUTTGART                                   = "Lapp Kabel Stuttgart"
	LAUBERT_ERBSTETTEN                                     = "Laubert Erbstetten"
	LAUBSANGERWEG_FELLBACH                                 = "Laubsängerweg Fellbach"
	LAUCHHAU_STUTTGART                                     = "Lauchhau Stuttgart"
	LAUERHALDE_WARMBRONN                                   = "Lauerhalde Warmbronn"
	LAUFENMUHLE_WELZHEIM                                   = "Laufenmühle Welzheim"
	LAUSTRASSE_STUTTGART                                   = "Laustraße Stuttgart"
	LAUTERGARTEN_DONZDORF                                  = "Lautergarten Donzdorf"
	LAUTERN_SULZBACH_M                                     = "Lautern Sulzbach (M)"
	LAUTERTAL_SULZBACH_M                                   = "Lautertal Sulzbach (M)"
	LAUTERTAL_WUSTENROT                                    = "Lautertal Wüstenrot"
	LECHTSTRASSE_NECKARWEIHINGEN                           = "Lechtstraße Neckarweihingen"
	LEDERBERG_STUTTGART                                    = "Lederberg Stuttgart"
	LEDERGASSE_SCHWABISCH_GMUND                            = "Ledergasse Schwäbisch Gmünd"
	LEDERSTRASSE_SCHORNDORF                                = "Lederstraße Schorndorf"
	LEERWASEN_NEIDLINGEN                                   = "Leerwasen Neidlingen"
	LEHENBRUCKE_ASPERG                                     = "Lehenbrücke Asperg"
	LEHENSTRASSE_STUTTGART                                 = "Lehenstraße Stuttgart"
	LEHENWEILER_AIDLINGEN                                  = "Lehenweiler Aidlingen"
	LEHENWEILER_ABZWEIG_AIDLINGEN                          = "Lehenweiler Abzweig Aidlingen"
	LEHLESTRASSE_FAURNDAU                                  = "Lehlestraße Faurndau"
	LEHNENBERG_REICHENBACH_BERGLEN                         = "Lehnenberg Reichenbach (Berglen)"
	LEHNENBERG_KREUZUNG_REICHENBACH_BERGLEN                = "Lehnenberg Kreuzung Reichenbach (Berglen)"
	LEIBNIZSTRASSE_STUTTGART                               = "Leibnizstraße Stuttgart"
	LEIBNIZSTRASSE_KORNWESTHEIM                            = "Leibnizstraße Kornwestheim"
	LEINFELDEN_LEINFELDEN                                  = "Leinfelden Leinfelden"
	LEINFELDENER_STRASSE_STUTTGART                         = "Leinfeldener Straße Stuttgart"
	LEINFELDER_HAUS_LEINFELDEN                             = "Leinfelder Haus Leinfelden"
	LEINTEL_REICHENBACH_F                                  = "Leintel Reichenbach (F)"
	LEINTELSTRASSE_EBERSBACH_F                             = "Leintelstraße Ebersbach (F)"
	LEIPZIGER_PLATZ_STUTTGART                              = "Leipziger Platz Stuttgart"
	LEIPZIGER_STRASSE_SINDELFINGEN                         = "Leipziger Straße Sindelfingen"
	LEMBERGSCHULE_NAGOLD                                   = "Lembergschule Nagold"
	LENAUDENKMAL_ESSLINGEN_N                               = "Lenaudenkmal Esslingen (N)"
	LENAUSTRASSE_ALBERSHAUSEN                              = "Lenaustraße Albershausen"
	LENBACHSTRASSE_GOPPINGEN                               = "Lenbachstraße Göppingen"
	LENZHALDE_ESSLINGEN_N                                  = "Lenzhalde Esslingen (N)"
	LENZHALDE_SCHARNHAUSEN                                 = "Lenzhalde Scharnhausen"
	LENZHALDE_KORNWESTHEIM                                 = "Lenzhalde Kornwestheim"
	LEOCENTER_ELTINGEN                                     = "Leo-Center Eltingen"
	LEOBAD_ELTINGEN                                        = "Leobad Eltingen"
	LEOBENER_STRASSE_STUTTGART                             = "Leobener Straße Stuttgart"
	LEONARDODAVINCIPLATZ_BOBLINGEN                         = "Leonardo-da-Vinci-Platz Böblingen"
	LEONBERG_LEONBERG                                      = "Leonberg Leonberg"
	LEONBERGER_DREIECK_ELTINGEN                            = "Leonberger Dreieck Eltingen"
	LEONBERGER_STRASSE_ELTINGEN                            = "Leonberger Straße Eltingen"
	LEONBERGER_STRASSE_GERLINGEN                           = "Leonberger Straße Gerlingen"
	LEONBERGER_STRASSE_HEIMSHEIM                           = "Leonberger Straße Heimsheim"
	LERCHENACKER_OBERESSLINGEN                             = "Lerchenäcker Oberesslingen"
	LERCHENACKER_BACKNANG                                  = "Lerchenäcker Backnang"
	LERCHENBERG_ORTSMITTE_GOPPINGEN                        = "Lerchenberg Ortsmitte Göppingen"
	LERCHENRAINSCHULE_STUTTGART                            = "Lerchenrainschule Stuttgart"
	LERCHENSTRASSE_OBERBOIHINGEN                           = "Lerchenstraße Oberboihingen"
	LERCHENSTRASSE_WEILHEIM_T                              = "Lerchenstraße Weilheim (T)"
	LERCHENSTRASSE_GROSSASPACH                             = "Lerchenstraße Großaspach"
	LERCHENWEG_WALDENBUCH                                  = "Lerchenweg Waldenbuch"
	LERCHENWIESEN_STUTTGART                                = "Lerchenwiesen Stuttgart"
	LESSINGSTRASSE_SCHONAICH                               = "Lessingstraße Schönaich"
	LESSINGSTRASSE_RUTESHEIM                               = "Lessingstraße Rutesheim"
	LESSINGSTRASSE_GOPPINGEN                               = "Lessingstraße Göppingen"
	LETTENACKER_PLOCHINGEN                                 = "Lettenäcker Plochingen"
	LETTENSTICH_WELZHEIM                                   = "Lettenstich Welzheim"
	LETTENSTRASSE_NEUHAUSEN_F                              = "Lettenstraße Neuhausen (F)"
	LEUTENBACHER_STRASSE_NELLMERSBACH                      = "Leutenbacher Straße Nellmersbach"
	LIBANONSTRASSE_STUTTGART                               = "Libanonstraße Stuttgart"
	LICHTENBERGER_STRASSE_OBERSTENFELD                     = "Lichtenberger Straße Oberstenfeld"
	LICHTENSTEINSTRASSE_BACKNANG                           = "Lichtensteinstraße Backnang"
	LICHTENSTERN_GYMNASIUM_GROSSSACHSENHEIM                = "Lichtenstern Gymnasium Großsachsenheim"
	LIEBENAU_WALDENBUCH                                    = "Liebenau Waldenbuch"
	LIEBFRAUENKIRCHE_EISLINGEN_F                           = "Liebfrauenkirche Eislingen (F)"
	LIEMANNSKLINGE_ABZWEIG_SULZBACH_M                      = "Liemannsklinge Abzweig Sulzbach (M)"
	LIEMERSBACH_GROSSERLACH                                = "Liemersbach Großerlach"
	LILIENSTRASSE_HOLZGERLINGEN                            = "Lilienstraße Holzgerlingen"
	LILIENSTRASSE_MALMSHEIM                                = "Lilienstraße Malmsheim"
	LILIENSTRASSE_ROMMELSHAUSEN                            = "Lilienstraße Rommelshausen"
	LILIENTHALSTRASSE_STUTTGART                            = "Lilienthalstraße Stuttgart"
	LIMBURGSCHULEBISSINGER_STRASSE_WEILHEIM_T              = "Limburgschule/Bissinger Straße Weilheim (T)"
	LIMBURGSCHULEKELTERSTRASSE_WEILHEIM_T                  = "Limburgschule/Kelterstraße Weilheim (T)"
	LIMESWEG_GRAB                                          = "Limesweg Grab"
	LIMESWEG_MAITIS                                        = "Limesweg Maitis"
	LINDACH_RECHBERGHAUSEN                                 = "Lindach Rechberghausen"
	LINDACHSCHULE_STETTEN_LEINFELDENECHT                   = "Lindachschule Stetten (Leinfelden-Echt.)"
	LINDE_HARTHAUSEN                                       = "Linde Harthausen"
	LINDE_KONGEN                                           = "Linde Köngen"
	LINDE_UNTERENSINGEN                                    = "Linde Unterensingen"
	LINDE_KAPPISHAUSERN                                    = "Linde Kappishäusern"
	LINDE_JEBENHAUSEN                                      = "Linde Jebenhausen"
	LINDENMUSEUM_STUTTGART                                 = "Linden-Museum Stuttgart"
	LINDENHOF_MOTZINGEN                                    = "Lindenhof Mötzingen"
	LINDENHOF_GEISLINGEN_STEIGE                            = "Lindenhof Geislingen (Steige)"
	LINDENHOF_BOHMENKIRCH                                  = "Lindenhof Böhmenkirch"
	LINDENPLATZ_OBERENSINGEN                               = "Lindenplatz Oberensingen"
	LINDENPLATZ_HEBSACK                                    = "Lindenplatz Hebsack"
	LINDENPLATZ_NEUFFEN                                    = "Lindenplatz Neuffen"
	LINDENPLATZ_UNTERWEISSACH                              = "Lindenplatz Unterweissach"
	LINDENPLATZ_STEINBACH                                  = "Lindenplatz Steinbach"
	LINDENSCHULE_GEISLINGEN_STEIGE                         = "Lindenschule Geislingen (Steige)"
	LINDENSTRASSE_BEMPFLINGEN                              = "Lindenstraße Bempflingen"
	LINDENSTRASSE_KORNWESTHEIM                             = "Lindenstraße Kornwestheim"
	LINDENSTRASSE_HOCHBERG                                 = "Lindenstraße Hochberg"
	LINDENSTRASSE_STEINENBRONN                             = "Lindenstraße Steinenbronn"
	LINDENSTRASSE_BISSINGEN_LB                             = "Lindenstraße Bissingen (LB)"
	LINDENSTRASSE_HEININGEN_GP                             = "Lindenstraße Heiningen (GP)"
	LINDENSTRASSE_WISSGOLDINGEN                            = "Lindenstraße Wißgoldingen"
	LINDENSTRASSE_MERKLINGEN                               = "Lindenstraße Merklingen"
	LINDENTAL_SCHLECHTBACH                                 = "Lindental Schlechtbach"
	LINDENTALER_STRASSE_SCHLECHTBACH                       = "Lindentaler Straße Schlechtbach"
	LINDENWEG_SCHORNDORF                                   = "Lindenweg Schorndorf"
	LINDORFER_WEG_KIRCHHEIM_T                              = "Lindorfer Weg Kirchheim (T)"
	LINDPAINTNERSTRASSE_STUTTGART                          = "Lindpaintnerstraße Stuttgart"
	LINGWIESEN_MUNCHINGEN                                  = "Lingwiesen Münchingen"
	LINSENHOFEN_LINSENHOFEN                                = "Linsenhofen Linsenhofen"
	LINSENHOFER_STRASSE_BEUREN                             = "Linsenhofer Straße Beuren"
	LISEMEITNERGYMNASIUM_ALDINGEN                          = "Lise-Meitner-Gymnasium Aldingen"
	LISEMEITNERSTR_OBERENSINGEN                            = "Lise-Meitner-Str. Oberensingen"
	LISEMEITNERSTRASSE_GOPPINGEN                           = "Lise-Meitner-Straße Göppingen"
	LISEMEITNERSTRASSE_NAGOLD                              = "Lise-Meitner-Straße Nagold"
	LISTSTRASSE_STUTTGART                                  = "Liststraße Stuttgart"
	LISZTSTRASSE_MURRHARDT                                 = "Lisztstraße Murrhardt"
	LOBENROT_ORT_LOBENROT                                  = "Lobenrot Ort Lobenrot"
	LOBENROTER_HOF_SCHANBACH                               = "Lobenroter Hof Schanbach"
	LOCHGAUER_STRASSE_BIETIGHEIM                           = "Löchgauer Straße Bietigheim"
	LOCHLESACKER_FREIBERG_N                                = "Löchlesäcker Freiberg (N)"
	LOCHLESBERG_DOFFINGEN                                  = "Löchlesberg Döffingen"
	LOHACKERSTRASSE_STUTTGART                              = "Lohäckerstraße Stuttgart"
	LOHMUHLE_WUSTENROT                                     = "Lohmühle Wüstenrot"
	LORCH_LORCH                                            = "Lorch Lorch"
	LORCHER_STRASSE_PFAHLBRONN                             = "Lorcher Straße Pfahlbronn"
	LORTZINGSTRASSE_ROMMELSHAUSEN                          = "Lortzingstraße Rommelshausen"
	LOSACKER_STUTTGART                                     = "Losäcker Stuttgart"
	LOSBURGSTRASSE_ALTBACH                                 = "Losburgstraße Altbach"
	LOSCHER_OST_MOGLINGEN                                  = "Löscher Ost Möglingen"
	LOSCHER_WEST_MOGLINGEN                                 = "Löscher West Möglingen"
	LOTTER_LUDWIGSBURG                                     = "Lotter Ludwigsburg"
	LOWEN_WOLFSCHLUGEN                                     = "Löwen Wolfschlugen"
	LOWEN_KUCHEN_F                                         = "Löwen Kuchen (F)"
	LOWEN_BORTLINGEN                                       = "Löwen Börtlingen"
	LOWENBRUNNEN_FELLBACH                                  = "Löwenbrunnen Fellbach"
	LOWENKELLER_SCHORNDORF                                 = "Löwenkeller Schorndorf"
	LOWENSTEINER_STRASSE_SPIEGELBERG                       = "Löwensteiner Straße Spiegelberg"
	LOWENTOR_STUTTGART                                     = "Löwentor Stuttgart"
	LOWENTORBRUCKE_STUTTGART                               = "Löwentorbrücke Stuttgart"
	LUDWIGFINCKHSTRASSE_WENDLINGEN_N                       = "Ludwig-Finckh-Straße Wendlingen (N)"
	LUDWIGHERRSTRASSE_KORNWESTHEIM                         = "Ludwig-Herr-Straße Kornwestheim"
	LUDWIGUHLANDSCHULE_HEIMSHEIM                           = "Ludwig-Uhland-Schule Heimsheim"
	LUDWIGSBURG_LUDWIGSBURG                                = "Ludwigsburg Ludwigsburg"
	LUDWIGSBURGER_STRASSE_OEFFINGEN                        = "Ludwigsburger Straße Oeffingen"
	LUDWIGSBURGER_STRASSE_SCHWIEBERDINGEN                  = "Ludwigsburger Straße Schwieberdingen"
	LUDWIGSBURGER_STRASSE_MOGLINGEN                        = "Ludwigsburger Straße Möglingen"
	LUDWIGSBURGER_STRASSE_TAMM                             = "Ludwigsburger Straße Tamm"
	LUDWIGSBURGER_STRASSE_WAIBLINGEN                       = "Ludwigsburger Straße Waiblingen"
	LUDWIGSBURGER_STRASSE_KORNWESTHEIM                     = "Ludwigsburger Straße Kornwestheim"
	LUDWIGSBURGER_STRASSE_HOHENECK                         = "Ludwigsburger Straße Hoheneck"
	LUDWIGSBURGER_STRASSE_BISSINGEN_LB                     = "Ludwigsburger Straße Bissingen (LB)"
	LUDWIGSBURGER_STRASSE_GROSSSACHSENHEIM                 = "Ludwigsburger Straße Großsachsenheim"
	LUDWIGSTRASSE_MUSBERG                                  = "Ludwigstraße Musberg"
	LUDWIGSTRASSE_FAURNDAU                                 = "Ludwigstraße Faurndau"
	LUFTFRACHTZENTRUM_WEST_BERNHAUSEN                      = "Luftfrachtzentrum West Bernhausen"
	LUGINSLAND_STUTTGART                                   = "Luginsland Stuttgart"
	LUGLENWEG_HERRENBERG                                   = "Lüglenweg Herrenberg"
	LUIKENWEG_STUTTGART                                    = "Luikenweg Stuttgart"
	LUTHERKIRCHE_FELLBACH                                  = "Lutherkirche Fellbach"
	LUTZENBERG_ALTHUTTE                                    = "Lutzenberg Althütte"
	LUTZESTRASSE_SCHNAIT                                   = "Lützestraße Schnait"
	MAGELLANSTRASSE_ECHTERDINGEN                           = "Magellanstraße Echterdingen"
	MAGSTADT_MAGSTADT                                      = "Magstadt Magstadt"
	MAGSTADTER_STRASSE_SCHAFHAUSEN                         = "Magstadter Straße Schafhausen"
	MAHDENTALSTRASSE_SINDELFINGEN                          = "Mahdentalstraße Sindelfingen"
	MAHDERKLINGE_STUTTGART                                 = "Mähderklinge Stuttgart"
	MAICHINGEN_MAICHINGEN                                  = "Maichingen Maichingen"
	MAICKLERSTRASSE_FELLBACH                               = "Maicklerstraße Fellbach"
	MAIENPLATZ_BOBLINGEN                                   = "Maienplatz Böblingen"
	MAIENWALTERSTRASSE_SULZGRIES                           = "Maienwalterstraße Sulzgries"
	MAIERHOFE_WEILHEIM_T                                   = "Maierhöfe Weilheim (T)"
	MAIERHOFSTRASSE_LORCH                                  = "Maierhofstraße Lorch"
	MAILLE_ESSLINGEN_N                                     = "Maille Esslingen (N)"
	MAINZER_ALLEE_LUDWIGSBURG                              = "Mainzer Allee Ludwigsburg"
	MALMSHEIM_MALMSHEIM                                    = "Malmsheim Malmsheim"
	MANFREDVONARDENNEALLEE_BACKNANG                        = "Manfred-von-Ardenne-Allee Backnang"
	MANFREDWORNERSTRASSE_GOPPINGEN                         = "Manfred-Wörner-Straße Göppingen"
	MANN__HUMMEL_LUDWIGSBURG                               = "Mann + Hummel Ludwigsburg"
	MANNENBERG_RUDERSBERG                                  = "Mannenberg Rudersberg"
	MANNENBERG_IM_HAU_RUDERSBERG                           = "Mannenberg Im Hau Rudersberg"
	MANNENWEILER_GRAB                                      = "Mannenweiler Grab"
	MANNHOLZ_PFAHLBRONN                                    = "Mannholz Pfahlbronn"
	MANOLZWEILER_WINTERBACH                                = "Manolzweiler Winterbach"
	MANOSQUER_STRASSE_LEINFELDEN                           = "Manosquer Straße Leinfelden"
	MANSLER_STRASSE_KORB                                   = "Mansler Straße Korb"
	MANZELLER_WEG_STUTTGART                                = "Manzeller Weg Stuttgart"
	MANZEN_MANZENSTR_GOPPINGEN                             = "Manzen Manzenstr. Göppingen"
	MANZEN_OBERE_SIEDLUNG_GOPPINGEN                        = "Manzen Obere Siedlung Göppingen"
	MANZEN_SPITZENBERGSTR_GOPPINGEN                        = "Manzen Spitzenbergstr. Göppingen"
	MANZEN_SPORTPLATZ_GOPPINGEN                            = "Manzen Sportplatz Göppingen"
	MARABUSTRASSE_STUTTGART                                = "Marabustraße Stuttgart"
	MARBACH_N_MARBACH_N                                    = "Marbach (N) Marbach (N)"
	MARBACHER_STRASSE_KIRCHBERG_M                          = "Marbacher Straße Kirchberg (M)"
	MARBACHER_STRASSE_AFFALTERBACH                         = "Marbacher Straße Affalterbach"
	MARBACHER_STRASSE_HEININGEN_BACKNANG                   = "Marbacher Straße Heiningen (Backnang)"
	MARBACHER_STRASSE_GROSSASPACH                          = "Marbacher Straße Großaspach"
	MARCOPOLOWEG_STUTTGART                                 = "Marco-Polo-Weg Stuttgart"
	MARCONIBRUCKE_STUTTGART                                = "Marconibrücke Stuttgart"
	MARCONISTRASSE_STUTTGART                               = "Marconistraße Stuttgart"
	MARGARETENSTRASSE_PLUDERHAUSEN                         = "Margaretenstraße Plüderhausen"
	MARGARETENWEG_GERLINGEN                                = "Margaretenweg Gerlingen"
	MARIECURIESCHULE_RAMTEL                                = "Marie-Curie-Schule Ramtel"
	MARIENSILBERBURGSTRASSE_STUTTGART                      = "Marien-/Silberburgstraße Stuttgart"
	MARIENHEIM_BACKNANG                                    = "Marienheim Backnang"
	MARIENHOSPITAL_STUTTGART                               = "Marienhospital Stuttgart"
	MARIENPLATZ_STUTTGART                                  = "Marienplatz Stuttgart"
	MARIENSTRASSE_STUTTGART                                = "Marienstraße Stuttgart"
	MARIENSTRASSE_WERNAU_N                                 = "Marienstraße Wernau (N)"
	MARIENSTRASSE_UHINGEN                                  = "Marienstraße Uhingen"
	MARIENSTRASSE_FAURNDAU                                 = "Marienstraße Faurndau"
	MARKGRABEN_HEBSACK                                     = "Markgraben Hebsack"
	MARKGRONINGER_STRASSE_EGLOSHEIM                        = "Markgröninger Straße Eglosheim"
	MARKGRONINGER_STRASSE_MOGLINGEN                        = "Markgröninger Straße Möglingen"
	MARKGRONINGER_STRASSE_ASPERG                           = "Markgröninger Straße Asperg"
	MARKLINEUM_GOPPINGEN                                   = "Märklineum Göppingen"
	MARKTBAUERNSTRASSE_DITZINGEN                           = "Markt-/Bauernstraße Ditzingen"
	MARKTGASSE_WAIBLINGEN                                  = "Marktgasse Waiblingen"
	MARKTPLATZ_WEILHEIM_T                                  = "Marktplatz Weilheim (T)"
	MARKTPLATZ_ESSLINGEN_N                                 = "Marktplatz Esslingen (N)"
	MARKTPLATZ_KIRCHHEIM_T                                 = "Marktplatz Kirchheim (T)"
	MARKTPLATZ_NECKARTENZLINGEN                            = "Marktplatz Neckartenzlingen"
	MARKTPLATZ_WEISSACH_BB                                 = "Marktplatz Weissach (BB)"
	MARKTPLATZ_SULZBACH_M                                  = "Marktplatz Sulzbach (M)"
	MARKTPLATZ_FORNSBACH                                   = "Marktplatz Fornsbach"
	MARKTPLATZ_ALFDORF                                     = "Marktplatz Alfdorf"
	MARKTPLATZ_RUDERSBERG                                  = "Marktplatz Rudersberg"
	MARKTPLATZ_HORRHEIM                                    = "Marktplatz Horrheim"
	MARKTPLATZ_WASCHENBEUREN                               = "Marktplatz Wäschenbeuren"
	MARKTPLATZ_GOPPINGEN                                   = "Marktplatz Göppingen"
	MARKTPLATZ_SALACH                                      = "Marktplatz Salach"
	MARKTPLATZ_SPARWIESEN                                  = "Marktplatz Sparwiesen"
	MARKTPLATZ_HEIMSHEIM                                   = "Marktplatz Heimsheim"
	MARKTPLATZ_GSCHWEND                                    = "Marktplatz Gschwend"
	MARKTPLATZ_BOBLINGER_STRASSE_SINDELFINGEN              = "Marktplatz (Böblinger Straße) Sindelfingen"
	MARKTPLATZ_RATHAUS_SINDELFINGEN                        = "Marktplatz (Rathaus) Sindelfingen"
	MARKTSTRASSE_BONLANDEN                                 = "Marktstraße Bonlanden"
	MARKTSTRASSE_GOPPINGEN                                 = "Marktstraße Göppingen"
	MARKUSKIRCHE_STUTTGART                                 = "Markuskirche Stuttgart"
	MARQUARDTSCHULE_PLOCHINGEN                             = "Marquardtschule Plochingen"
	MARRENSTRASSE_DONZDORF                                 = "Marrenstraße Donzdorf"
	MARTINLUTHERSTRASSE_ECHTERDINGEN                       = "Martin-Luther-Straße Echterdingen"
	MARTINLUTHERSTRASSE_LUDWIGSBURG                        = "Martin-Luther-Straße Ludwigsburg"
	MARTINSHAUS_BESIGHEIM                                  = "Martinshaus Besigheim"
	MARTINSKIRCHE_STUTTGART                                = "Martinskirche Stuttgart"
	MARTINSKIRCHE_KIRCHHEIM_T                              = "Martinskirche Kirchheim (T)"
	MARTINSTRASSE_NUSSDORF                                 = "Martinstraße Nussdorf"
	MARXENHOF_KIRCHENKIRNBERG                              = "Marxenhof Kirchenkirnberg"
	MASSSTABFABRIK_MAINHARDT                               = "Maßstabfabrik Mainhardt"
	MATHILDENSTRASSE_STUTTGART                             = "Mathildenstraße Stuttgart"
	MAUBACH_MAUBACH                                        = "Maubach Maubach"
	MAUBACHER_STRASSE_BACKNANG                             = "Maubacher Straße Backnang"
	MAUERACKER_TAILFINGEN                                  = "Maueräcker Tailfingen"
	MAULBRONNER_STRASSE_HORRHEIM                           = "Maulbronner Straße Horrheim"
	MAURENER_STRASSE_EHNINGEN                              = "Maurener Straße Ehningen"
	MAURENER_WEG_BOBLINGEN                                 = "Maurener Weg Böblingen"
	MAURER_STEINENBRONN                                    = "Maurer Steinenbronn"
	MAUSERSTRASSE_OSSWEIL                                  = "Mauserstraße Oßweil"
	MAXBORNGYMNASIUM_BACKNANG                              = "Max-Born-Gymnasium Backnang"
	MAXEYTHDAIMLERSTRASSE_WINNENDEN                        = "Max-Eyth-/Daimlerstraße Winnenden"
	MAXEYTHSEE_STUTTGART                                   = "Max-Eyth-See Stuttgart"
	MAXEYTHSTRASSE_WINNENDEN                               = "Max-Eyth-Straße Winnenden"
	MAXEYTHSTRASSE_WERNAU_N                                = "Max-Eyth-Straße Wernau (N)"
	MAXEYTHSTRASSE_KORNWESTHEIM                            = "Max-Eyth-Straße Kornwestheim"
	MAXPLANCKINSTITUTE_STUTTGART                           = "Max-Planck-Institute Stuttgart"
	MAXPLANCKSTRASSE_FELLBACH                              = "Max-Planck-Straße Fellbach"
	MAXPLANCKSTRASSE_OBERESSLINGEN                         = "Max-Planck-Straße Oberesslingen"
	MAXPLANCKSTRASSE_NEUFFEN                               = "Max-Planck-Straße Neuffen"
	MAXPLANCKSTRASSE_KORNWESTHEIM                          = "Max-Planck-Straße Kornwestheim"
	MAYBACHSTRASSE_STUTTGART                               = "Maybachstraße Stuttgart"
	MAYBACHSTRASSE_SCHONAICH                               = "Maybachstraße Schönaich"
	MAYBACHSTRASSE_KORNWESTHEIM                            = "Maybachstraße Kornwestheim"
	MAYBACHSTRASSE_FRICKENHAUSEN                           = "Maybachstraße Frickenhausen"
	MAYBACHSTRASSE_RECHBERGHAUSEN                          = "Maybachstraße Rechberghausen"
	MEDIUS_KLINIK_RUIT                                     = "medius Klinik Ruit"
	MEDIUS_KLINIK_KIRCHHEIM_T                              = "medius Klinik Kirchheim (T)"
	MEDIUS_KLINIK_NURTINGEN                                = "medius Klinik Nürtingen"
	MEHRZWECKHALLE_OSSWEIL                                 = "Mehrzweckhalle Oßweil"
	MEISENWEG_WAIBLINGEN                                   = "Meisenweg Waiblingen"
	MELLIBEESESTRASSE_SINDELFINGEN                         = "Melli-Beese-Straße Sindelfingen"
	MENDELSSOHNSTRASSE_WAIBLINGEN                          = "Mendelssohnstraße Waiblingen"
	MENZELSTRASSE_FREIBERG_N                               = "Menzelstraße Freiberg (N)"
	MERCEDESBENZ_STUTTGART                                 = "Mercedes-Benz Stuttgart"
	MERCEDESBENZ_HST_A_SINDELFINGEN                        = "Mercedes-Benz Hst. A Sindelfingen"
	MERCEDESBENZ_HST_B_SINDELFINGEN                        = "Mercedes-Benz Hst. B Sindelfingen"
	MERCEDESBENZ_HST_C_SINDELFINGEN                        = "Mercedes-Benz Hst. C Sindelfingen"
	MERCEDESBENZ_HST_D_SINDELFINGEN                        = "Mercedes-Benz Hst. D Sindelfingen"
	MERCEDESBENZ_HST_E_SINDELFINGEN                        = "Mercedes-Benz Hst. E Sindelfingen"
	MERCEDESBENZ_HST_F_SINDELFINGEN                        = "Mercedes-Benz Hst. F Sindelfingen"
	MERCEDESBENZ_HST_I_SINDELFINGEN                        = "Mercedes-Benz Hst. I Sindelfingen"
	MERCEDESBENZ_HST_K_SINDELFINGEN                        = "Mercedes-Benz Hst. K Sindelfingen"
	MERCEDESBENZ_HST_N_SINDELFINGEN                        = "Mercedes-Benz Hst. N Sindelfingen"
	MERCEDESBENZ_HST_O_SINDELFINGEN                        = "Mercedes-Benz Hst. O Sindelfingen"
	MERCEDESBENZ_HST_P_SINDELFINGEN                        = "Mercedes-Benz Hst. P Sindelfingen"
	MERCEDESBENZ_P305_SINDELFINGEN                         = "Mercedes-Benz P305 Sindelfingen"
	MERCEDESBENZ_P307_SINDELFINGEN                         = "Mercedes-Benz P307 Sindelfingen"
	MERCEDESBENZ_TOR_I_SINDELFINGEN                        = "Mercedes-Benz Tor I Sindelfingen"
	MERCEDESBENZ_TOR_III_SINDELFINGEN                      = "Mercedes-Benz Tor III Sindelfingen"
	MERCEDESBENZ_TOR_VII_SINDELFINGEN                      = "Mercedes-Benz Tor VII Sindelfingen"
	MERCEDESBENZ_WELT_STUTTGART                            = "Mercedes-Benz Welt Stuttgart"
	MERCEDESSTR_UHINGEN                                    = "Mercedesstr. Uhingen"
	MERCEDESSTRASSE_STUTTGART                              = "Mercedesstraße Stuttgart"
	MERCEDESSTRASSE_HIRSCHLANDEN                           = "Mercedesstraße Hirschlanden"
	MERKELSCHES_BAD_ESSLINGEN_N                            = "Merkel'sches Bad Esslingen (N)"
	MERKLINGEN_MERKLINGEN                                  = "Merklingen Merklingen"
	MERKLINGER_STRASSE_WEIL_DER_STADT                      = "Merklinger Straße Weil der Stadt"
	MERKLINGER_STRASSE_MACHTOLSHEIM                        = "Merklinger Straße Machtolsheim"
	MESSE_WEST_ECHTERDINGEN                                = "Messe West Echterdingen"
	MESSEHALLE_SINDELFINGEN                                = "Messehalle Sindelfingen"
	METLANGEN_ABZWEIG_SCHWABISCH_GMUND                     = "Metlangen Abzweig Schwäbisch Gmünd"
	METLANGEN_ORT_SCHWABISCH_GMUND                         = "Metlangen Ort Schwäbisch Gmünd"
	METTELBACH_KIRCHENKIRNBERG                             = "Mettelbach Kirchenkirnberg"
	METTELBERG_FORNSBACH                                   = "Mettelberg Fornsbach"
	METTERBRUCKE_GROSSSACHSENHEIM                          = "Metterbrücke Großsachsenheim"
	METTINGEN_METTINGEN                                    = "Mettingen Mettingen"
	METZELHOF_LORCH                                        = "Metzelhof Lorch"
	METZELWIESEN_MERKLINGEN                                = "Metzelwiesen Merklingen"
	METZGERHAU_STUTTGART                                   = "Metzgerhau Stuttgart"
	METZINGER_STRASSE_KOHLBERG                             = "Metzinger Straße Kohlberg"
	METZINGER_STRASSE_NURTINGEN                            = "Metzinger Straße Nürtingen"
	METZINGER_STRASSE_NECKARTENZLINGEN                     = "Metzinger Straße Neckartenzlingen"
	METZINGER_STRASSE_BONLANDEN                            = "Metzinger Straße Bonlanden"
	METZINGER_STRASSE_GROSSBETTLINGEN                      = "Metzinger Straße Großbettlingen"
	METZSTRASSE_STUTTGART                                  = "Metzstraße Stuttgart"
	MICHELAU_SCHLECHTBACH                                  = "Michelau Schlechtbach"
	MICHELAU_ORTSMITTE_SCHLECHTBACH                        = "Michelau Ortsmitte Schlechtbach"
	MICHELAUER_STRASSE_STEINENBERG                         = "Michelauer Straße Steinenberg"
	MICHELBERGGYMNASIUM_GEISLINGEN_STEIGE                  = "Michelberggymnasium Geislingen (Steige)"
	MIEDELSBACHSTEINENBERG_MIEDELSBACH                     = "Miedelsbach-Steinenberg Miedelsbach"
	MIEDELSBACHER_STRASSE_HAUBERSBRONN                     = "Miedelsbacher Straße Haubersbronn"
	MIKROZENTRUM_WAIBLINGEN                                = "Mikrozentrum Waiblingen"
	MILANWEG_KIRCHHEIM_T                                   = "Milanweg Kirchheim (T)"
	MILCHHOF_STUTTGART                                     = "Milchhof Stuttgart"
	MILCHZENTRALE_GEISLINGEN_STEIGE                        = "Milchzentrale Geislingen (Steige)"
	MILLOCKERSTRASSE_STUTTGART                             = "Millöckerstraße Stuttgart"
	MINERALBADER_STUTTGART                                 = "Mineralbäder Stuttgart"
	MINIGOLFPLATZ_RUTESHEIM                                = "Minigolfplatz Rutesheim"
	MITTE_SCHLAITDORF                                      = "Mitte Schlaitdorf"
	MITTE_SCHWIEBERDINGEN                                  = "Mitte Schwieberdingen"
	MITTE_GERADSTETTEN                                     = "Mitte Geradstetten"
	MITTE_REUDERN                                          = "Mitte Reudern"
	MITTE_NEIDLINGEN                                       = "Mitte Neidlingen"
	MITTE_NUFRINGEN                                        = "Mitte Nufringen"
	MITTE_HAUBERSBRONN                                     = "Mitte Haubersbronn"
	MITTE_PATTONVILLE                                      = "Mitte Pattonville"
	MITTEHELENENSTRASSE_MARKGRONINGEN                      = "Mitte/Helenenstraße Markgröningen"
	MITTEVOLKSBANK_MARKGRONINGEN                           = "Mitte/Volksbank Markgröningen"
	MITTEWERNERSTRASSE_MARKGRONINGEN                       = "Mitte/Wernerstraße Markgröningen"
	MITTELBRUDEN_OBERBRUDEN                                = "Mittelbrüden Oberbrüden"
	MITTELFISCHBACH_ABZWEIGUNG_GROSSERLACH                 = "Mittelfischbach Abzweigung Großerlach"
	MITTELSCHONTAL_BACKNANG                                = "Mittelschöntal Backnang"
	MITTELWEILER_PFAHLBRONN                                = "Mittelweiler Pfahlbronn"
	MITTENBUHL_DOFFINGEN                                   = "Mittenbühl Döffingen"
	MITTLERE_BRUCKE_SCHORNDORF                             = "Mittlere Brücke Schorndorf"
	MITTLERE_MUHLE_HOLZGERLINGEN                           = "Mittlere Mühle Holzgerlingen"
	MITTLERE_UFERSTRASSE_SCHORNDORF                        = "Mittlere Uferstraße Schorndorf"
	MITTNACHTSTRASSE_STUTTGART                             = "Mittnachtstraße Stuttgart"
	MOGLINGER_STRASSE_ASPERG                               = "Möglinger Straße Asperg"
	MOGLINGER_STRASSE_PFLUGFELDEN                          = "Möglinger Straße Pflugfelden"
	MOHRINGEN_BF_MOHRINGEN                                 = "Möhringen Bf Möhringen"
	MOHRINGEN_FREIBAD_STUTTGART                            = "Möhringen Freibad Stuttgart"
	MOHRINGEN_RATHAUS_STUTTGART                            = "Möhringen Rathaus Stuttgart"
	MOHRINGER_STRASSE_SIELMINGEN                           = "Möhringer Straße Sielmingen"
	MOLTKESTRASSE_KUCHEN_F                                 = "Moltkestraße Kuchen (F)"
	MONCHFELD_STUTTGART                                    = "Mönchfeld Stuttgart"
	MONCHHOF_KAISERSBACH                                   = "Mönchhof Kaisersbach"
	MONCHSBERGSTRASSE_FREIBERG_N                           = "Mönchsbergstraße Freiberg (N)"
	MONCHSBRUCKE_SCHORNDORF                                = "Mönchsbrücke Schorndorf"
	MONCHSBRUNNEN_SINDELFINGEN                             = "Mönchsbrunnen Sindelfingen"
	MONCHWEG_STUTTGART                                     = "Mönchweg Stuttgart"
	MONREPOS_EGLOSHEIM                                     = "Monrepos Eglosheim"
	MONSHEIMER_WEG_STUTTGART                               = "Mönsheimer Weg Stuttgart"
	MORBACH_GRAB                                           = "Morbach Grab"
	MORIKESCHULE_ELTINGEN                                  = "Mörike-Schule Eltingen"
	MORIKESCHULE_OTLINGEN                                  = "Mörikeschule Ötlingen"
	MORIKESTRASSE_STUTTGART                                = "Mörikestraße Stuttgart"
	MORIKESTRASSE_LUDWIGSBURG                              = "Mörikestraße Ludwigsburg"
	MORIKESTRASSE_OBERRIEXINGEN                            = "Mörikestraße Oberriexingen"
	MORIKESTRASSE_GROSSBETTLINGEN                          = "Mörikestraße Großbettlingen"
	MORIKESTRASSE_GOPPINGEN                                = "Mörikestraße Göppingen"
	MORIKESTRASSE_HEININGEN_GP                             = "Mörikestraße Heiningen (GP)"
	MORSESTRASSE_STUTTGART                                 = "Morsestraße Stuttgart"
	MOSELSTRASSE_STUTTGART                                 = "Moselstraße Stuttgart"
	MOTORSTRASSE_24_STUTTGART                              = "Motorstraße 24 Stuttgart"
	MOTZINGER_STRASSE_UNTERJETTINGEN                       = "Mötzinger Straße Unterjettingen"
	MOTZINGER_STRASSE_OSCHELBRONN_GAUFELDEN                = "Mötzinger Straße Öschelbronn (Gäufelden)"
	MOTZINGER_STRASSE_NAGOLD                               = "Mötzinger Straße Nagold"
	MOZARTSCHULE_NEUHAUSEN_F                               = "Mozartschule Neuhausen (F)"
	MOZARTSTRASSE_HOCHDORF_ES                              = "Mozartstraße Hochdorf (ES)"
	MOZARTSTRASSE_FELLBACH                                 = "Mozartstraße Fellbach"
	MOZARTSTRASSE_EISLINGEN_F                              = "Mozartstraße Eislingen (F)"
	MOZARTSTRASSE_GOPPINGEN                                = "Mozartstraße Göppingen"
	MUCKENSTURM_STUTTGART                                  = "Muckensturm Stuttgart"
	MUHLE_ALDINGEN                                         = "Mühle Aldingen"
	MUHLENBUCKEL_BERNHAUSEN                                = "Mühlenbuckel Bernhausen"
	MUHLHAUSEN_STUTTGART                                   = "Mühlhausen Stuttgart"
	MUHLHAUSEN_SCHLOSS_STUTTGART                           = "Mühlhausen Schloss Stuttgart"
	MUHLHAUSENER_STRASSE_LEHNINGEN                         = "Mühlhausener Straße Lehningen"
	MUHLHAUSER_STR_GRUIBINGEN                              = "Mühlhauser Str. Gruibingen"
	MUHLHAUSER_STRASSE_KORNWESTHEIM                        = "Mühlhäuser Straße Kornwestheim"
	MUHLHAUSER_STRASSE_ROSSWAG                             = "Mühlhäuser Straße Roßwag"
	MUHLHAUSER_STRASSE_OSSWEIL                             = "Mühlhäuser Straße Oßweil"
	MUHLSTEG_STUTTGART                                     = "Mühlsteg Stuttgart"
	MUHLSTRASSE_NURTINGEN                                  = "Mühlstraße Nürtingen"
	MUHLSTRASSE_OBERRIEXINGEN                              = "Mühlstraße Oberriexingen"
	MUHLSTRASSE_SUSSEN                                     = "Mühlstraße Süßen"
	MUHLSTRASSE_WALDHAUSEN_B_SCHORND_LORCH                 = "Mühlstraße Waldhausen b Schornd (Lorch)"
	MUHLWEG_HERRENBERG                                     = "Mühlweg Herrenberg"
	MUHLWEG_BEINSTEIN                                      = "Mühlweg Beinstein"
	MUHLWEG_ALDINGEN                                       = "Mühlweg Aldingen"
	MUHLWIESEN_BIETIGHEIM                                  = "Mühlwiesen Bietigheim"
	MULLERHEIM_MUNCHINGEN                                  = "Müllerheim Münchingen"
	MUNCHINGEN_MUNCHINGEN                                  = "Münchingen Münchingen"
	MUNCHINGER_STRASSE_DITZINGEN                           = "Münchinger Straße Ditzingen"
	MUNCHINGER_STRASSE_MARKGRONINGEN                       = "Münchinger Straße Markgröningen"
	MUNCHINGER_STRASSE_HEMMINGEN                           = "Münchinger Straße Hemmingen"
	MUNSTER_STUTTGART                                      = "Münster Stuttgart"
	MUNSTER_RATHAUS_STUTTGART                              = "Münster Rathaus Stuttgart"
	MUNSTER_VIADUKT_STUTTGART                              = "Münster Viadukt Stuttgart"
	MUNZENHALDE_PLUDERHAUSEN                               = "Münzenhalde Plüderhausen"
	MURKENBACHWEG_BOBLINGEN                                = "Murkenbachweg Böblingen"
	MURRBADER_BACKNANG                                     = "Murrbäder Backnang"
	MURRBRUCKE_BURGSTALL_BURGSTETTEN                       = "Murrbrücke Burgstall (Burgstetten)"
	MURRER_STRASSE_STEINHEIM_M                             = "Murrer Straße Steinheim (M)"
	MURRHARDT_MURRHARDT                                    = "Murrhardt Murrhardt"
	MURRHARDTER_STRASSE_SULZBACH_M                         = "Murrhardter Straße Sulzbach (M)"
	MURRHARLE_MURRHARDT                                    = "Murrhärle Murrhardt"
	MURRTALSCHULE_OPPENWEILER                              = "Murrtalschule Oppenweiler"
	MUTZENREIS_ZOLLBERG                                    = "Mutzenreis Zollberg"
	MUTZENREISSTRASSE_ZOLLBERG                             = "Mutzenreisstraße Zollberg"
	NACHBARSCHAFTSSCHULE_THOMASHARDT                       = "Nachbarschaftsschule Thomashardt"
	NACHBARSCHAFTSSCHULE_BRETZENACKER                      = "Nachbarschaftsschule Bretzenacker"
	NACHTIGALLENWEG_WAIBLINGEN                             = "Nachtigallenweg Waiblingen"
	NAGELESTAL_KIRCHHEIM_T                                 = "Nägelestal Kirchheim (T)"
	NAGELESTRASSE_STUTTGART                                = "Nägelestraße Stuttgart"
	NAGELSTRASSE_SCHARNHAUSEN                              = "Nagelstraße Scharnhausen"
	NAGOLDER_STEIG_MOTZINGEN                               = "Nagolder Steig Mötzingen"
	NASSACH_NASSACH                                        = "Nassach Nassach"
	NASSACH_BOLZPLATZ_UHINGEN                              = "Nassach Bolzplatz Uhingen"
	NASSACH_FRIEDHOF_UHINGEN                               = "Nassach Friedhof Uhingen"
	NASSACH_UNTERHUTT_UHINGEN                              = "Nassach Unterhütt Uhingen"
	NASSACHMUHLE_UHINGEN                                   = "Nassachmühle Uhingen"
	NASSACHMUHLE_SCHULE_UHINGEN                            = "Nassachmühle Schule Uhingen"
	NASTPLATZ_STUTTGART                                    = "Nastplatz Stuttgart"
	NASTSTRASSE_LUDWIGSBURG                                = "Naststraße Ludwigsburg"
	NATURFREUNDEHAUS_SECHSELBERG                           = "Naturfreundehaus Sechselberg"
	NATURFREUNDEHAUS_WELZHEIM                              = "Naturfreundehaus Welzheim"
	NATURSCHUTZZENTRUM_SCHOPFLOCH_LENNINGEN                = "Naturschutzzentrum Schopfloch (Lenningen)"
	NEBRINGER_STRASSE_BONDORF                              = "Nebringer Straße Bondorf"
	NECKAR_FORUM_ESSLINGEN_N                               = "Neckar Forum Esslingen (N)"
	NECKARBRUCKE_NURTINGEN                                 = "Neckarbrücke Nürtingen"
	NECKARBRUCKE_LUDWIGSBURG                               = "Neckarbrücke Ludwigsburg"
	NECKARBRUCKE_NECKARTAILFINGEN                          = "Neckarbrücke Neckartailfingen"
	NECKARFREIBAD_ESSLINGEN_N                              = "Neckarfreibad Esslingen (N)"
	NECKARHALDE_STUTTGART                                  = "Neckarhalde Stuttgart"
	NECKARHALDE_NECKARREMS                                 = "Neckarhalde Neckarrems"
	NECKARPARK_STUTTGART                                   = "Neckarpark Stuttgart"
	NECKARPARK_STADION_STUTTGART                           = "NeckarPark (Stadion) Stuttgart"
	NECKARSCHULE_ALDINGEN                                  = "Neckarschule Aldingen"
	NECKARSTRASSE_HEGNACH                                  = "Neckarstraße Hegnach"
	NECKARSTRASSE_KORNWESTHEIM                             = "Neckarstraße Kornwestheim"
	NECKARSTRASSE_ASPERG                                   = "Neckarstraße Asperg"
	NECKARSTRASSE_LUDWIGSBURG                              = "Neckarstraße Ludwigsburg"
	NECKARTAILFINGER_STRASSE_SCHLAITDORF                   = "Neckartailfinger Straße Schlaitdorf"
	NECKARTENZLINGER_STRASSE_BEMPFLINGEN                   = "Neckartenzlinger Straße Bempflingen"
	NECKARTOR_STUTTGART                                    = "Neckartor Stuttgart"
	NECKARWEIHINGER_STRASSE_OSSWEIL                        = "Neckarweihinger Straße Oßweil"
	NECKARZENTRUM_HOCHBERG                                 = "Neckarzentrum Hochberg"
	NECKLINSBERG_ASPERGLEN                                 = "Necklinsberg Asperglen"
	NECKLINSBERG_KREUZUNG_ASPERGLEN                        = "Necklinsberg Kreuzung Asperglen"
	NEIDLINGER_STRASSE_WEILHEIM_T                          = "Neidlinger Straße Weilheim (T)"
	NEL_MEZZO_GEISLINGEN_STEIGE                            = "Nel Mezzo Geislingen (Steige)"
	NELKENSTRASSE_AFFSTATT                                 = "Nelkenstraße Affstätt"
	NELLINGER_LINDE_NELLINGEN                              = "Nellinger Linde Nellingen"
	NELLINGER_STRASSE_BERKHEIM                             = "Nellinger Straße Berkheim"
	NELLMERSBACH_NELLMERSBACH                              = "Nellmersbach Nellmersbach"
	NETZ__IHK_NAGOLD                                       = "Netz / IHK Nagold"
	NETZESTRASSE_GRUNBUHL                                  = "Netzestraße Grünbühl"
	NEUE_BAHNHOFSTRASSE_VAIHINGEN_E                        = "Neue Bahnhofstraße Vaihingen (E)"
	NEUE_RAMTELSTRASSE_RAMTEL                              = "Neue Ramtelstraße Ramtel"
	NEUE_ROMMELSHAUSER_STRASSE_WAIBLINGEN                  = "Neue Rommelshauser Straße Waiblingen"
	NEUE_STRASSE_LIEBERSBRONN                              = "Neue Straße Liebersbronn"
	NEUE_STRASSE_STEINBACH                                 = "Neue Straße Steinbach"
	NEUE_WEINSTEIGE_61_STUTTGART                           = "Neue Weinsteige 61 Stuttgart"
	NEUENBUHL_FLACHT                                       = "Neuenbühl Flacht"
	NEUER_FRIEDHOF_LUDWIGSBURG                             = "Neuer Friedhof Ludwigsburg"
	NEUER_FRIEDHOF_SCHORNDORF                              = "Neuer Friedhof Schorndorf"
	NEUER_FRIEDHOF_BISSINGEN_LB                            = "Neuer Friedhof Bissingen (LB)"
	NEUER_MARKT_LEINFELDEN                                 = "Neuer Markt Leinfelden"
	NEUES_RATHAUS_SCHONAICH                                = "Neues Rathaus Schönaich"
	NEUFFEN_NEUFFEN                                        = "Neuffen Neuffen"
	NEUFFENER_STRASSE_LINSENHOFEN                          = "Neuffener Straße Linsenhofen"
	NEUFFENSTRASSE_WENDLINGEN_N                            = "Neuffenstraße Wendlingen (N)"
	NEUFFENSTRASSE_REICHENBACH_F                           = "Neuffenstraße Reichenbach (F)"
	NEUFURSTENHUTTE_GROSSERLACH                            = "Neufürstenhütte Großerlach"
	NEUGEREUT_STUTTGART                                    = "Neugereut Stuttgart"
	NEUHAUSER_STRASSE_MUNKLINGEN                           = "Neuhauser Straße Münklingen"
	NEUHAUSER_STRASSE_BERNHAUSEN                           = "Neuhäuser Straße Bernhausen"
	NEUHAUSER_STRASSE_DENKENDORF                           = "Neuhäuser Straße Denkendorf"
	NEUHOF_KIRCHBERG_M                                     = "Neuhof Kirchberg (M)"
	NEUHOF_WELZHEIM                                        = "Neuhof Welzheim"
	NEULAUTERN_WUSTENROT                                   = "Neulautern Wüstenrot"
	NEUSATZ_BESIGHEIM                                      = "Neusatz Besigheim"
	NEUSCHONTAL_BACKNANG                                   = "Neuschöntal Backnang"
	NEUSTADTHOHENACKER_NEUSTADT                            = "Neustadt-Hohenacker Neustadt"
	NEUSTADTER_STRASSE_WAIBLINGEN                          = "Neustädter Straße Waiblingen"
	NEUWEILER_STRASSE_KLEINSACHSENHEIM                     = "Neuweiler Straße Kleinsachsenheim"
	NEUWIESEN_B466_GEISLINGEN_STEIGE                       = "Neuwiesen (B466) Geislingen (Steige)"
	NEUWIESEN_GEWERBEGEBIET_GEISLINGEN_STEIGE              = "Neuwiesen Gewerbegebiet Geislingen (Steige)"
	NEUWIRTSH_PORSCHEP_STUTTGART                           = "Neuwirtsh. (Porschep.) Stuttgart"
	NEUWIRTSHAUSKREUZUNG_STUTTGART                         = "Neuwirtshauskreuzung Stuttgart"
	NIEBUHRWEG_STUTTGART                                   = "Niebuhrweg Stuttgart"
	NIEDERE_KLINGE_GEMMRIGHEIM                             = "Niedere Klinge Gemmrigheim"
	NIEDERHOFEN_ELTINGEN                                   = "Niederhofen Eltingen"
	NIEDERREUTIN_BONDORF                                   = "Niederreutin Bondorf"
	NIKOLAUSOTTOSTRASSE_ECHTERDINGEN                       = "Nikolaus-Otto-Straße Echterdingen"
	NIKOLAUSPFLEGE_STUTTGART                               = "Nikolauspflege Stuttgart"
	NIKOMODIESTRASSE_34_OBERLENNINGEN                      = "Nikomödiestraße 34 Oberlenningen"
	NIXENWEG_STUTTGART                                     = "Nixenweg Stuttgart"
	NOBELSTRASSE_STUTTGART                                 = "Nobelstraße Stuttgart"
	NOLDEWEG_LORCH                                         = "Noldeweg Lorch"
	NORD_BONLANDEN                                         = "Nord Bonlanden"
	NORD_WEIL_IM_SCHONBUCH                                 = "Nord Weil im Schönbuch"
	NORD_MAICHINGEN                                        = "Nord Maichingen"
	NORD_RUDERSBERG                                        = "Nord Rudersberg"
	NORD_HIRSCHLANDEN                                      = "Nord Hirschlanden"
	NORDWESTRING_BERNHAUSEN                                = "Nord-West-Ring Bernhausen"
	NORDWESTUMFAHRUNG_BERNHAUSEN                           = "Nord-West-Umfahrung Bernhausen"
	NORDBAHNHOF_KLEINGLATTBACH                             = "Nordbahnhof Kleinglattbach"
	NORDBAHNHOF_STUTTGART                                  = "Nordbahnhof Stuttgart"
	NORDHALDENSTRASSE_BEUTELSBACH                          = "Nordhaldenstraße Beutelsbach"
	NORDSEESCHWIEBERD_STRASSE_STUTTGART                    = "Nordsee-/Schwieberd. Straße Stuttgart"
	NORMANNENSTRASSE_WEIL_DER_STADT                        = "Normannenstraße Weil der Stadt"
	NOTZINGER_STRASSE_KIRCHHEIM_T                          = "Notzinger Straße Kirchheim (T)"
	NOVIZENWEGTHW_NEUHAUSEN_F                              = "Novizenweg/THW Neuhausen (F)"
	NUFRINGEN_NUFRINGEN                                    = "Nufringen Nufringen"
	NURNBERGER_STRASSE_STUTTGART                           = "Nürnberger Straße Stuttgart"
	NURTINGEN_NURTINGEN                                    = "Nürtingen Nürtingen"
	NURTINGER_STRASSE_GROSSBETTLINGEN                      = "Nürtinger Straße Großbettlingen"
	NURTINGER_STRASSE_GROTZINGEN                           = "Nürtinger Straße Grötzingen"
	NURTINGER_STRASSE_NECKARTAILFINGEN                     = "Nürtinger Straße Neckartailfingen"
	NURTINGER_STRASSE_BOBLINGEN                            = "Nürtinger Straße Böblingen"
	NUSSBAUMWEG_OBERSTENFELD                               = "Nussbaumweg Oberstenfeld"
	NUSSDORFER_STRASSE_AURICH                              = "Nußdorfer Straße Aurich"
	NUSSDORFER_STRASSE_EBERDINGEN                          = "Nußdorfer Straße Eberdingen"
	NUSSSTRASSE_SINDELFINGEN                               = "Nüßstraße Sindelfingen"
	OB_DEM_STEINBACH_STUTTGART                             = "Ob dem Steinbach Stuttgart"
	OBERAICHEN_LEINFELDEN                                  = "Oberaichen Leinfelden"
	OBERAICHEN_WALDHEIM_LEINFELDEN                         = "Oberaichen Waldheim Leinfelden"
	OBERAMTEIGASSE_BESIGHEIM                               = "Oberamteigasse Besigheim"
	OBERBOIHINGEN_OBERBOIHINGEN                            = "Oberboihingen Oberboihingen"
	OBERBOIHINGER_STRASSE_ZIZISHAUSEN                      = "Oberboihinger Straße Zizishausen"
	OBERBRUDENER_STRASSE_STEINBACH                         = "Oberbrüdener Straße Steinbach"
	OBERDORF_WEILER_OB_HELFENSTN                           = "Oberdorf Weiler ob Helfenstn."
	OBERE_ACKER_KAYH                                       = "Obere Äcker Kayh"
	OBERE_BAHNHOFSTRASSE_WAIBLINGEN                        = "Obere Bahnhofstraße Waiblingen"
	OBERE_GASSE_NELLINGEN_UL                               = "Obere Gasse Nellingen (UL)"
	OBERE_LIESACKER_ALTDORF_ES                             = "Obere Liesäcker Altdorf (ES)"
	OBERE_MARBACHER_STRASSE_LUDWIGSBURG                    = "Obere Marbacher Straße Ludwigsburg"
	OBERE_MONCHHALDE_STUTTGART                             = "Obere Mönchhalde Stuttgart"
	OBERE_SEEWIESEN_HOPFIGHEIM                             = "Obere Seewiesen Höpfigheim"
	OBERE_STRASSE_DETTINGEN_T                              = "Obere Straße Dettingen (T)"
	OBERE_ZIEGELEI_STUTTGART                               = "Obere Ziegelei Stuttgart"
	OBERER_EISBERGWEG_ZOLLBERG                             = "Oberer Eisbergweg Zollberg"
	OBERER_GRABEN_NEUFFEN                                  = "Oberer Graben Neuffen"
	OBERER_ROSBERG_WAIBLINGEN                              = "Oberer Rosberg Waiblingen"
	OBERER_SCHADBERG_KAISERSBACH                           = "Oberer Schadberg Kaisersbach"
	OBERER_WASEN_WELZHEIM                                  = "Oberer Wasen Welzheim"
	OBERESSLINGEN_OBERESSLINGEN                            = "Oberesslingen Oberesslingen"
	OBERFISCHBACH_GROSSERLACH                              = "Oberfischbach Großerlach"
	OBERHOF_OBERHOF                                        = "Oberhof Oberhof"
	OBERLENNINGEN_OBERLENNINGEN                            = "Oberlenningen Oberlenningen"
	OBERNDORF_RUDERSBERG                                   = "Oberndorf Rudersberg"
	OBERNDORF_BRUNNENPLATZ_RUDERSBERG                      = "Oberndorf Brunnenplatz Rudersberg"
	OBERNDORF_MANNENBERGER_STRASSE_RUDERSBERG              = "Oberndorf Mannenberger Straße Rudersberg"
	OBERNEUSTETTEN_KIRCHENKIRNBERG                         = "Oberneustetten Kirchenkirnberg"
	OBERRIEXINGER_STRASSE_GROSSSACHSENHEIM                 = "Oberriexinger Straße Großsachsenheim"
	OBERSCHONTAL_BACKNANG                                  = "Oberschöntal Backnang"
	OBERSCHONTAL_ABZWEIG_BACKNANG                          = "Oberschöntal Abzweig Backnang"
	OBERSTEINENBERG_WELZHEIM                               = "Obersteinenberg Welzheim"
	OBERTALWEG_WALDENBRONN                                 = "Obertalweg Wäldenbronn"
	OBERTURKH_BF_GOPPINGER_STR_STUTTGART                   = "Obertürkh. Bf (Göppinger Str.) Stuttgart"
	OBERTURKHEIM_STUTTGART                                 = "Obertürkheim Stuttgart"
	OBERWIESENSTRASSE_55_STUTTGART                         = "Oberwiesenstraße 55 Stuttgart"
	OBSTHALDE_BEINSTEIN                                    = "Obsthalde Beinstein"
	OCHSEN_SPIELBERG                                       = "Ochsen Spielberg"
	OCHSENTROG_BOBLINGEN                                   = "Ochsentrog Böblingen"
	ODERNHARDT_ODERNHARDT                                  = "Ödernhardt Ödernhardt"
	OESCHSCHULEN_EISLINGEN_F                               = "Oesch-Schulen Eislingen (F)"
	OFELE_OBERJETTINGEN                                    = "Öfele Oberjettingen"
	OHMSTRASSE_ZUFFENHAUSEN_BF_STUTTGART                   = "Ohmstraße (Zuffenhausen Bf) Stuttgart"
	OLGAKEPLERSTRASSE_DEIZISAU                             = "Olga-/Keplerstraße Deizisau"
	OLGAZEPPELINSTRASSE_DEIZISAU                           = "Olga-/Zeppelinstraße Deizisau"
	OLGAECK_STUTTGART                                      = "Olgaeck Stuttgart"
	OLMUHLE_WEISSACH_BB                                    = "Ölmühle Weissach (BB)"
	OPPENWEILER_OPPENWEILER                                = "Oppenweiler Oppenweiler"
	ORT_AICH                                               = "Ort Aich"
	ORTSEINGANG_HEGENSBERG                                 = "Ortseingang Hegensberg"
	ORTSMITTE_ERBSTETTEN                                   = "Ortsmitte Erbstetten"
	ORTSMITTE_KLEINASPACH                                  = "Ortsmitte Kleinaspach"
	ORTSMITTE_REICHENBACH_F                                = "Ortsmitte Reichenbach (F)"
	ORTSMITTE_OTLINGEN                                     = "Ortsmitte Ötlingen"
	ORTSMITTE_LINDORF                                      = "Ortsmitte Lindorf"
	ORTSMITTE_WARMBRONN                                    = "Ortsmitte Warmbronn"
	ORTSMITTE_FLACHT                                       = "Ortsmitte Flacht"
	ORTSMITTE_MERKLINGEN                                   = "Ortsmitte Merklingen"
	ORTSMITTE_OPPENWEILER                                  = "Ortsmitte Oppenweiler"
	ORTSMITTE_ALTHUTTE                                     = "Ortsmitte Althütte"
	ORTSMITTE_GROTZINGEN                                   = "Ortsmitte Grötzingen"
	ORTSMITTE_ALLMERSBACH_IM_TAL                           = "Ortsmitte Allmersbach im Tal"
	ORTSMITTE_SCHLECHTBACH                                 = "Ortsmitte Schlechtbach"
	ORTSMITTE_HOHENACKER                                   = "Ortsmitte Hohenacker"
	ORTSMITTE_BIRKMANNSWEILER                              = "Ortsmitte Birkmannsweiler"
	ORTSMITTE_OPPELSBOHM                                   = "Ortsmitte Oppelsbohm"
	ORTSMITTE_BAACH_WINNENDEN                              = "Ortsmitte Baach (Winnenden)"
	ORTSMITTE_HOFEN                                        = "Ortsmitte Höfen"
	ORTSMITTE_ENDERSBACH                                   = "Ortsmitte Endersbach"
	ORTSMITTE_MURR                                         = "Ortsmitte Murr"
	ORTSMITTE_AURICH                                       = "Ortsmitte Aurich"
	ORTSMITTE_GROSSINGERSHEIM                              = "Ortsmitte Großingersheim"
	ORTSMITTE_HOHENGEHREN                                  = "Ortsmitte Hohengehren"
	ORTSMITTE_EYBACH                                       = "Ortsmitte Eybach"
	ORTSMITTE_WALDHAUSEN_GEISLINGEN                        = "Ortsmitte Waldhausen (Geislingen)"
	ORTSMITTE_MAITIS                                       = "Ortsmitte Maitis"
	ORTSMITTE_LENGLINGEN                                   = "Ortsmitte Lenglingen"
	ORTSMITTE_ALBERSHAUSEN                                 = "Ortsmitte Albershausen"
	ORTSMITTE_BAD_BOLL                                     = "Ortsmitte Bad Boll"
	ORTSMITTE_DURNAU                                       = "Ortsmitte Dürnau"
	ORTSMITTE_MERKLINGEN_7                                 = "Ortsmitte Merklingen 7"
	ORTSMITTEKIRCHE_WALDHAUSEN_B_SCHORND_LORCH             = "Ortsmitte/Kirche Waldhausen b Schornd (Lorch)"
	ORTSMITTEKORNSTR_OBERKIRNECK                           = "Ortsmitte/Kornstr. Oberkirneck"
	OSCHELBRONN_OSCHELBRONN_BERGLEN                        = "Öschelbronn Öschelbronn (Berglen)"
	OSKAR_FRECH_SEEBAD_SCHORNDORF                          = "Oskar Frech Seebad Schorndorf"
	OSKARSCHLEMMERSTRASSE_STUTTGART                        = "Oskar-Schlemmer-Straße Stuttgart"
	OSSMANNSWEILER_EYBACH                                  = "Oßmannsweiler Eybach"
	OSSWEILER_WEG_ALDINGEN                                 = "Oßweiler Weg Aldingen"
	OST_REUDERN                                            = "Ost Reudern"
	OSTENDPLATZ_STUTTGART                                  = "Ostendplatz Stuttgart"
	OSTERBRONNSTRASSE_STUTTGART                            = "Osterbronnstraße Stuttgart"
	OSTERFELD_STUTTGART                                    = "Österfeld Stuttgart"
	OSTERHOLZ_LUDWIGSBURG                                  = "Osterholz Ludwigsburg"
	OSTERHOLZALLEE_LUDWIGSBURG                             = "Osterholzallee Ludwigsburg"
	OSTERREICHISCHER_PLATZ_STUTTGART                       = "Österreichischer Platz Stuttgart"
	OSTERWIESENWEG_KLEINGLATTBACH                          = "Osterwiesenweg Kleinglattbach"
	OSTFILDERN_NELLINGEN                                   = "Ostfildern Nellingen"
	OSTFRIEDHOF_OSSWEIL                                    = "Ostfriedhof Oßweil"
	OSTHEIM_LEOVETTERBAD_STUTTGART                         = "Ostheim Leo-Vetter-Bad Stuttgart"
	OSTPREUSSENWEG_NURTINGEN                               = "Ostpreußenweg Nürtingen"
	OSTRING_HOCHDORF_ES                                    = "Ostring Hochdorf (ES)"
	OSTSIEDLUNG_NURTINGEN                                  = "Ostsiedlung Nürtingen"
	OSTSTRASSE_LUDWIGSBURG                                 = "Oststraße Ludwigsburg"
	OSWALDSTRASSE_MAGSTADT                                 = "Oswaldstraße Magstadt"
	OTLINGEN_OTLINGEN                                      = "Ötlingen Ötlingen"
	OTLINGER_STRASSE_NOTZINGEN                             = "Ötlinger Straße Notzingen"
	OTTOHAHNSTRASSE_NURTINGEN                              = "Otto-Hahn-Straße Nürtingen"
	OTTOHAHNSTRASSE_KIRCHHEIM_T                            = "Otto-Hahn-Straße Kirchheim (T)"
	OTTOHAHNSTRASSE_NECKARWEIHINGEN                        = "Otto-Hahn-Straße Neckarweihingen"
	OTTOHIRSCHBRUCKEN_STUTTGART                            = "Otto-Hirsch-Brücken Stuttgart"
	OTTOLILIENTHALSTRASSE_BOBLINGEN                        = "Otto-Lilienthal-Straße Böblingen"
	OTTOMUHLSCHLEGELHAUS_ENDERSBACH                        = "Otto-Mühlschlegel-Haus Endersbach"
	OWEN_OWEN                                              = "Owen Owen"
	PADAGOGISCHE_HOCHSCHULE_LUDWIGSBURG                    = "Pädagogische Hochschule Ludwigsburg"
	PALAZZO_FREIBERG_N                                     = "Palazzo Freiberg (N)"
	PALMENWALDSTRASSE_BRUHL                                = "Palmenwaldstraße Brühl"
	PANORAMA_THERME_BEUREN                                 = "Panorama Therme Beuren"
	PANORAMAKLINIK_ESSLINGEN_N                             = "Panorama-Klinik Esslingen (N)"
	PANORAMASTR_RECHBERGHAUSEN                             = "Panoramastr. Rechberghausen"
	PANORAMASTRASSE_PLOCHINGEN                             = "Panoramastraße Plochingen"
	PANORAMASTRASSE_BISSINGEN_LB                           = "Panoramastraße Bissingen (LB)"
	PANORAMASTRASSE_LORCH                                  = "Panoramastraße Lorch"
	PANORAMASTRASSEBAHNHOF_FAURNDAU                        = "Panoramastraße/Bahnhof Faurndau"
	PANZERKASERNE_BOBLINGEN                                = "Panzerkaserne Böblingen"
	PAPIERFABRIK_GEMMRIGHEIM                               = "Papierfabrik Gemmrigheim"
	PAPPELWEG_MARKGRONINGEN                                = "Pappelweg Markgröningen"
	PARACELSUSSTRASSE_OBERESSLINGEN                        = "Paracelsusstraße Oberesslingen"
	PARADIESWEG_RUDERN                                     = "Paradiesweg Rüdern"
	PARKACKER_BIETIGHEIM                                   = "Parkäcker Bietigheim"
	PARKPLATZ_SCHLOSSGYMNASIUM_KIRCHHEIM_T                 = "Parkplatz Schlossgymnasium Kirchheim (T)"
	PARKPLATZ_ZOLLSTOCK_SPIEGELBERG                        = "Parkplatz Zollstock Spiegelberg"
	PARKSIEDLUNG_PARKSIEDLUNG                              = "Parksiedlung Parksiedlung"
	PARKSTRASSE_PARKSIEDLUNG                               = "Parkstraße Parksiedlung"
	PARKSTRASSE_PLIENSAUVORSTADT                           = "Parkstraße Pliensauvorstadt"
	PARKSTRASSE_LANDRATSAMT_BOBLINGEN                      = "Parkstraße (Landratsamt) Böblingen"
	PAULGERHARDTHAUS_BESIGHEIM                             = "Paul-Gerhardt-Haus Besigheim"
	PAULKOEPFFWEG_GOPPINGEN                                = "Paul-Koepff-Weg Göppingen"
	PAULLINCKESTRASSE_STUTTGART                            = "Paul-Lincke-Straße Stuttgart"
	PAULINENSTRASSE_NELLINGEN                              = "Paulinenstraße Nellingen"
	PAULUSKIRCHE_FELLBACH                                  = "Pauluskirche Fellbach"
	PAULUSKIRCHE_GEISLINGEN_STEIGE                         = "Pauluskirche Geislingen (Steige)"
	PAYERSTRASSE_STUTTGART                                 = "Payerstraße Stuttgart"
	PEREGRINASTRASSE_STUTTGART                             = "Peregrinastraße Stuttgart"
	PERONNASPLATZ_NEUHAUSEN_F                              = "Peronnas-Platz Neuhausen (F)"
	PEROUSER_STRASSE_MALMSHEIM                             = "Perouser Straße Malmsheim"
	PESTALOZZISCHULE_STUTTGART                             = "Pestalozzi-Schule Stuttgart"
	PESTALOZZISTRASSE_ENDERSBACH                           = "Pestalozzistraße Endersbach"
	PETERHEBELSTRASSE_NECKARWEIHINGEN                      = "Peter-Hebel-Straße Neckarweihingen"
	PETERSHALDE_PFAHLBRONN                                 = "Petershalde Pfahlbronn"
	PFAFFENWEG_STUTTGART                                   = "Pfaffenweg Stuttgart"
	PFAFFLESTRASSESCHILLERSCHULE_LORCH                     = "Pfäfflestraße/Schillerschule Lorch"
	PFARRACKER_NEUSTADT                                    = "Pfarräcker Neustadt"
	PFARRGARTENSTRASSE_KIRCHBERG_M                         = "Pfarrgartenstraße Kirchberg (M)"
	PFARRHAUS_GUNDELBACH                                   = "Pfarrhaus Gündelbach"
	PFARRSTRASSE_BIETIGHEIM                                = "Pfarrstraße Bietigheim"
	PFARRWIESENGYMNASIUM_SINDELFINGEN                      = "Pfarrwiesengymnasium Sindelfingen"
	PFEIFFERKLINGE_PLIENSAUVORSTADT                        = "Pfeifferklinge Pliensauvorstadt"
	PFINGSTHALDE_EYBACH                                    = "Pfingsthalde Eybach"
	PFINGSTWASEN_GOPPINGEN                                 = "Pfingstwasen Göppingen"
	PFITZERSTRASSE_SCHWABISCH_GMUND                        = "Pfitzerstraße Schwäbisch Gmünd"
	PFLEGEHEIM_WINTERBACH                                  = "Pflegeheim Winterbach"
	PFLEGEHEIM_SINDELFINGEN                                = "Pflegeheim Sindelfingen"
	PFLEGEHEIM_SCHLIERBACH                                 = "Pflegeheim Schlierbach"
	PFLUGFELDER_STRASSE_STUTTGART                          = "Pflugfelder Straße Stuttgart"
	PFLUGFELDER_STRASSE_KORNWESTHEIM                       = "Pflugfelder Straße Kornwestheim"
	PFORZHEIMER_STRASSE_HORRHEIM                           = "Pforzheimer Straße Horrheim"
	PFOSTENBERG_PLOCHINGEN                                 = "Pfostenberg Plochingen"
	PFULLINGER_STRASSE_STUTTGART                           = "Pfullinger Straße Stuttgart"
	PIEMONTESER_STRASSE_ERDMANNHAUSEN                      = "Piemonteser Straße Erdmannhausen"
	PLACKERTSTRASSE_BUNZWANGEN                             = "Pläckertstraße Bünzwangen"
	PLAISIRSCHULE_BACKNANG                                 = "Plaisirschule Backnang"
	PLANCKSTRASSE_STUTTGART                                = "Planckstraße Stuttgart"
	PLANCKWEG_GOPPINGEN                                    = "Planckweg Göppingen"
	PLATTENWALD_BACKNANG                                   = "Plattenwald Backnang"
	PLATZ_HOLZMADEN                                        = "Platz Holzmaden"
	PLEIDELSHEIMER_STRASSE_BIETIGHEIM                      = "Pleidelsheimer Straße Bietigheim"
	PLIENINGEN_STUTTGART                                   = "Plieningen Stuttgart"
	PLIENINGEN_GARBE_STUTTGART                             = "Plieningen Garbe Stuttgart"
	PLIENINGEN_POST_STUTTGART                              = "Plieningen Post Stuttgart"
	PLIENINGEN_SEEMUHLENWEG_STUTTGART                      = "Plieningen Seemühlenweg Stuttgart"
	PLIENINGER_STRASSE_STUTTGART                           = "Plieninger Straße Stuttgart"
	PLIENINGER_STRASSE_SCHARNHAUSEN                        = "Plieninger Straße Scharnhausen"
	PLIENINGER_STRASSE_NORD_SCHARNHAUSEN                   = "Plieninger Straße (Nord) Scharnhausen"
	PLIENSAUFRIEDHOF_ZOLLBERG                              = "Pliensaufriedhof Zollberg"
	PLIENSAUTURM_ESSLINGEN_N                               = "Pliensauturm Esslingen (N)"
	PLOCHINGEN_PLOCHINGEN                                  = "Plochingen Plochingen"
	PLOCHINGER_STR_DEIZISAU                                = "Plochinger Str. Deizisau"
	PLOCHINGER_STRASSE_KONGEN                              = "Plochinger Straße Köngen"
	PLUDERHAUSEN_PLUDERHAUSEN                              = "Plüderhausen Plüderhausen"
	POLIZEIHOCHSCHULE_HERRENBERG                           = "Polizeihochschule Herrenberg"
	POLIZEIPRASIDIUM_EINSATZ_GOPPINGEN                     = "Polizeipräsidium Einsatz Göppingen"
	POLIZEIWACHE_KIRCHHEIM_T                               = "Polizeiwache Kirchheim (T)"
	POPPENWEILER_STRASSE_HOCHDORF_REMSECK                  = "Poppenweiler Straße Hochdorf (Remseck)"
	PORSCHE_WEISSACH_BB                                    = "Porsche Weissach (BB)"
	PORSCHE_GROSSSACHSENHEIM                               = "Porsche Großsachsenheim"
	PORSCHESTRASSE_HERRENBERG                              = "Porschestraße Herrenberg"
	POST_RENNINGEN                                         = "Post Renningen"
	POST_DAGERSHEIM                                        = "Post Dagersheim"
	POST_GUTENBERG                                         = "Post Gutenberg"
	POST_OSCHELBRONN_GAUFELDEN                             = "Post Öschelbronn (Gäufelden)"
	POST_OPPENWEILER                                       = "Post Oppenweiler"
	POST_SCHORNBACH                                        = "Post Schornbach"
	POST_FREUDENTAL                                        = "Post Freudental"
	POST_HOHENSTEIN                                        = "Post Hohenstein"
	POST_BOHMENKIRCH                                       = "Post Böhmenkirch"
	POSTAMT_NEBRINGEN                                      = "Postamt Nebringen"
	POSTAMT_WIESENSTEIG                                    = "Postamt Wiesensteig"
	POSTDORFLE_STUTTGART                                   = "Postdörfle Stuttgart"
	POSTDORFLE_FELLBACH                                    = "Postdörfle Fellbach"
	POSTPLATZ_BOBLINGEN                                    = "Postplatz Böblingen"
	POSTSTRASSE_BEUTELSBACH                                = "Poststraße Beutelsbach"
	POSTSTRASSLE_BIETIGHEIM                                = "Poststräßle Bietigheim"
	POSTWEG_PLUDERHAUSEN                                   = "Postweg Plüderhausen"
	POTSDAMER_RING_BACKNANG                                = "Potsdamer Ring Backnang"
	PRAGFRIEDHOF_STUTTGART                                 = "Pragfriedhof Stuttgart"
	PRAGSATTEL_STUTTGART                                   = "Pragsattel Stuttgart"
	PREVORST_FEUERWEHR_GRONAU                              = "Prevorst Feuerwehr Gronau"
	PREVORST_LOWEN_GRONAU                                  = "Prevorst Löwen Gronau"
	PRINZEUGENPLATZ_GROSSHEPPACH                           = "Prinz-Eugen-Platz Großheppach"
	PULVERDINGEN_B10_ENZWEIHINGEN                          = "Pulverdingen B10 Enzweihingen"
	PULVERDINGEN_ORT_ENZWEIHINGEN                          = "Pulverdingen Ort Enzweihingen"
	PULVERDINGER_STRASSE_HOCHDORF_EBERDINGEN               = "Pulverdinger Straße Hochdorf (Eberdingen)"
	PUMPSTATION_HIRSCHLANDEN                               = "Pumpstation Hirschlanden"
	PUSTEBLUMENKREISEL_EBERSBACH_F                         = "Pusteblumenkreisel Ebersbach (F)"
	QUADRIUM_WERNAU_N                                      = "Quadrium Wernau (N)"
	QUELLENSTRASSE_BEINSTEIN                               = "Quellenstraße Beinstein"
	QUELLENWEG_URBACH                                      = "Quellenweg Urbach"
	QUERSPANGE_WAIBLINGEN                                  = "Querspange Waiblingen"
	QUERSTRASSE_SALACH                                     = "Querstraße Salach"
	RABENWIESEN_SUSSEN                                     = "Rabenwiesen Süßen"
	RAICHBERG_SCHULZENTRUM_EBERSBACH_F                     = "Raichberg Schulzentrum Ebersbach (F)"
	RAIFFEISENBANK_LIPPOLDSWEILER                          = "Raiffeisenbank Lippoldsweiler"
	RAIFFEISENSTRASSE_GROTZINGEN                           = "Raiffeisenstraße Grötzingen"
	RAIFFEISENSTRASSE_BONLANDEN                            = "Raiffeisenstraße Bonlanden"
	RAIFFEISENSTRASSE_WINZERHAUSEN                         = "Raiffeisenstraße Winzerhausen"
	RAIFFEISENSTRASSE_BARTENBACH_GOPPINGEN                 = "Raiffeisenstraße Bartenbach (Göppingen)"
	RAINACKERSTRASSE_BONLANDEN                             = "Rainäckerstraße Bonlanden"
	RAINWIESEN_NECKARGRONINGEN                             = "Rainwiesen Neckargröningen"
	RAITE_RENNINGEN                                        = "Raite Renningen"
	RAITELSBERG_STUTTGART                                  = "Raitelsberg Stuttgart"
	RAMTEL_GERLINGEN                                       = "Ramtel Gerlingen"
	RANDENSTRASSE_HASLACH                                  = "Randenstraße Haslach"
	RAPPACHSCHULE_STUTTGART                                = "Rappachschule Stuttgart"
	RAPPACHSTRASSE_STUTTGART                               = "Rappachstraße Stuttgart"
	RAPPENBERG_KIRCHBERG_M                                 = "Rappenberg Kirchberg (M)"
	RAPPENHOF_LEONBERG                                     = "Rappenhof Leonberg"
	RASTATTER_STRASSE_STUTTGART                            = "Rastatter Straße Stuttgart"
	RATHAUS_PLATTENHARDT                                   = "Rathaus Plattenhardt"
	RATHAUS_SIELMINGEN                                     = "Rathaus Sielmingen"
	RATHAUS_BERNHAUSEN                                     = "Rathaus Bernhausen"
	RATHAUS_ALTENRIET                                      = "Rathaus Altenriet"
	RATHAUS_OEFFINGEN                                      = "Rathaus Oeffingen"
	RATHAUS_SCHMIDEN                                       = "Rathaus Schmiden"
	RATHAUS_KAPPISHAUSERN                                  = "Rathaus Kappishäusern"
	RATHAUS_SCHARNHAUSEN                                   = "Rathaus Scharnhausen"
	RATHAUS_MAICHINGEN                                     = "Rathaus Maichingen"
	RATHAUS_LUDWIGSBURG                                    = "Rathaus Ludwigsburg"
	RATHAUS_ERDMANNHAUSEN                                  = "Rathaus Erdmannhausen"
	RATHAUS_KIRCHBERG_M                                    = "Rathaus Kirchberg (M)"
	RATHAUS_GUNDELBACH                                     = "Rathaus Gündelbach"
	RATHAUS_WALDREMS                                       = "Rathaus Waldrems"
	RATHAUS_MAUBACH                                        = "Rathaus Maubach"
	RATHAUS_GERADSTETTEN                                   = "Rathaus Geradstetten"
	RATHAUS_ROMMELSHAUSEN                                  = "Rathaus Rommelshausen"
	RATHAUS_ENDERSBACH                                     = "Rathaus Endersbach"
	RATHAUS_ALTBACH                                        = "Rathaus Altbach"
	RATHAUS_BALTMANNSWEILER                                = "Rathaus Baltmannsweiler"
	RATHAUS_DEIZISAU                                       = "Rathaus Deizisau"
	RATHAUS_KONGEN                                         = "Rathaus Köngen"
	RATHAUS_AICHELBERG_AICHWALD                            = "Rathaus Aichelberg (Aichwald)"
	RATHAUS_HOLZGERLINGEN                                  = "Rathaus Holzgerlingen"
	RATHAUS_NOTZINGEN                                      = "Rathaus Notzingen"
	RATHAUS_HOCHDORF_ES                                    = "Rathaus Hochdorf (ES)"
	RATHAUS_UNTERLENNINGEN                                 = "Rathaus Unterlenningen"
	RATHAUS_BEMPFLINGEN                                    = "Rathaus Bempflingen"
	RATHAUS_JESINGEN                                       = "Rathaus Jesingen"
	RATHAUS_OHMDEN                                         = "Rathaus Ohmden"
	RATHAUS_HEPSISAU                                       = "Rathaus Hepsisau"
	RATHAUS_BISSINGEN_AN_DER_TECK_T                        = "Rathaus Bissingen an der Teck (T)"
	RATHAUS_OCHSENWANG                                     = "Rathaus Ochsenwang"
	RATHAUS_SCHOPFLOCH_LENNINGEN                           = "Rathaus Schopfloch (Lenningen)"
	RATHAUS_KLEINBETTLINGEN                                = "Rathaus Kleinbettlingen"
	RATHAUS_TISCHARDT                                      = "Rathaus Tischardt"
	RATHAUS_NECKARHAUSEN                                   = "Rathaus Neckarhausen"
	RATHAUS_NECKARTAILFINGEN                               = "Rathaus Neckartailfingen"
	RATHAUS_NECKARTENZLINGEN                               = "Rathaus Neckartenzlingen"
	RATHAUS_MOTZINGEN                                      = "Rathaus Mötzingen"
	RATHAUS_HEMMINGEN                                      = "Rathaus Hemmingen"
	RATHAUS_GEBERSHEIM                                     = "Rathaus Gebersheim"
	RATHAUS_DATZINGEN                                      = "Rathaus Dätzingen"
	RATHAUS_MUNKLINGEN                                     = "Rathaus Münklingen"
	RATHAUS_GARTRINGEN                                     = "Rathaus Gärtringen"
	RATHAUS_ROHRAU                                         = "Rathaus Rohrau"
	RATHAUS_OSCHELBRONN_GAUFELDEN                          = "Rathaus Öschelbronn (Gäufelden)"
	RATHAUS_DOFFINGEN                                      = "Rathaus Döffingen"
	RATHAUS_AIDLINGEN                                      = "Rathaus Aidlingen"
	RATHAUS_DEUFRINGEN                                     = "Rathaus Deufringen"
	RATHAUS_BREITENSTEIN                                   = "Rathaus Breitenstein"
	RATHAUS_NEUWEILER                                      = "Rathaus Neuweiler"
	RATHAUS_KAYH                                           = "Rathaus Kayh"
	RATHAUS_DECKENPFRONN                                   = "Rathaus Deckenpfronn"
	RATHAUS_AFFSTATT                                       = "Rathaus Affstätt"
	RATHAUS_TAILFINGEN                                     = "Rathaus Tailfingen"
	RATHAUS_HEGNACH                                        = "Rathaus Hegnach"
	RATHAUS_GROSSERLACH                                    = "Rathaus Großerlach"
	RATHAUS_UNTERBRUDEN                                    = "Rathaus Unterbrüden"
	RATHAUS_ALTHUTTE                                       = "Rathaus Althütte"
	RATHAUS_DENKENDORF                                     = "Rathaus Denkendorf"
	RATHAUS_HEININGEN_BACKNANG                             = "Rathaus Heiningen (Backnang)"
	RATHAUS_HEUTENSBACH                                    = "Rathaus Heutensbach"
	RATHAUS_HAUBERSBRONN                                   = "Rathaus Haubersbronn"
	RATHAUS_NEUSTADT                                       = "Rathaus Neustadt"
	RATHAUS_BEINSTEIN                                      = "Rathaus Beinstein"
	RATHAUS_STRUMPFELBACH_WEINSTADT                        = "Rathaus Strümpfelbach (Weinstadt)"
	RATHAUS_BEUTELSBACH                                    = "Rathaus Beutelsbach"
	RATHAUS_WEILER_SCHORNDORF                              = "Rathaus Weiler (Schorndorf)"
	RATHAUS_LEUTENBACH                                     = "Rathaus Leutenbach"
	RATHAUS_BACKNANG                                       = "Rathaus Backnang"
	RATHAUS_UNTERWEISSACH                                  = "Rathaus Unterweissach"
	RATHAUS_WAIBLINGEN                                     = "Rathaus Waiblingen"
	RATHAUS_SCHWAIKHEIM                                    = "Rathaus Schwaikheim"
	RATHAUS_KORNWESTHEIM                                   = "Rathaus Kornwestheim"
	RATHAUS_OTTMARSHEIM                                    = "Rathaus Ottmarsheim"
	RATHAUS_OBERSTENFELD                                   = "Rathaus Oberstenfeld"
	RATHAUS_HOCHDORF_EBERDINGEN                            = "Rathaus Hochdorf (Eberdingen)"
	RATHAUS_BISSINGEN_LB                                   = "Rathaus Bissingen (LB)"
	RATHAUS_RUTESHEIM                                      = "Rathaus Rutesheim"
	RATHAUS_MALMSHEIM                                      = "Rathaus Malmsheim"
	RATHAUS_KOHLBERG                                       = "Rathaus Kohlberg"
	RATHAUS_HOFINGEN                                       = "Rathaus Höfingen"
	RATHAUS_STUTTGART                                      = "Rathaus Stuttgart"
	RATHAUS_OBERLENNINGEN                                  = "Rathaus Oberlenningen"
	RATHAUS_WEILER_ZUM_STEIN                               = "Rathaus Weiler zum Stein"
	RATHAUS_BRETZENACKER                                   = "Rathaus Bretzenacker"
	RATHAUS_OBERRIEXINGEN                                  = "Rathaus Oberriexingen"
	RATHAUS_GERLINGEN                                      = "Rathaus Gerlingen"
	RATHAUS_GROSSBETTLINGEN                                = "Rathaus Großbettlingen"
	RATHAUS_GEISLINGEN_STEIGE                              = "Rathaus Geislingen (Steige)"
	RATHAUS_GINGEN_F                                       = "Rathaus Gingen (F)"
	RATHAUS_KUCHEN_F                                       = "Rathaus Kuchen (F)"
	RATHAUS_REICHENBACH_IM_TALE_DEGGINGEN                  = "Rathaus Reichenbach im Täle (Deggingen)"
	RATHAUS_MUHLHAUSEN_IM_TALE                             = "Rathaus Mühlhausen im Täle"
	RATHAUS_WIESENSTEIG                                    = "Rathaus Wiesensteig"
	RATHAUS_BOHMENKIRCH                                    = "Rathaus Böhmenkirch"
	RATHAUS_NENNINGEN                                      = "Rathaus Nenningen"
	RATHAUS_BIRENBACH                                      = "Rathaus Birenbach"
	RATHAUS_HATTENHOFEN                                    = "Rathaus Hattenhofen"
	RATHAUS_UHINGEN                                        = "Rathaus Uhingen"
	RATHAUS_GRUIBINGEN                                     = "Rathaus Gruibingen"
	RATHAUS_OBERDRACKENSTEIN                               = "Rathaus Oberdrackenstein"
	RATHAUS_EBERSBACH_F                                    = "Rathaus Ebersbach (F)"
	RATHAUS_SUSSEN                                         = "Rathaus Süßen"
	RATHAUS_REICHENBACH_UNTER_RECHBERG_DONZDORF            = "Rathaus Reichenbach unter Rechberg (Donzdorf)"
	RATHAUS_WINZINGEN                                      = "Rathaus Winzingen"
	RATHAUS_STOTTEN                                        = "Rathaus Stötten"
	RATHAUS_GOPPINGEN                                      = "Rathaus Göppingen"
	RATHAUS_HASLACH                                        = "Rathaus Häslach"
	RATHAUS_DONNSTETTEN                                    = "Rathaus Donnstetten"
	RATHAUS_GECHINGEN                                      = "Rathaus Gechingen"
	RATHAUS_EBHAUSEN                                       = "Rathaus Ebhausen"
	RATHAUS_ALTENSTEIG                                     = "Rathaus Altensteig"
	RATHAUS_LEHNINGEN                                      = "Rathaus Lehningen"
	RATHAUS_NELLINGEN_UL                                   = "Rathaus Nellingen (UL)"
	RATHAUS_NOTARIAT_WALDDORF                              = "Rathaus (Notariat) Walddorf"
	RATHAUSDORFPLATZ_HOHENSTAUFEN                          = "Rathaus/Dorfplatz Hohenstaufen"
	RATHAUSSTADTHALLE_NECKARREMS                           = "Rathaus/Stadthalle Neckarrems"
	RATHAUSPLATZ_HIRSCHLANDEN                              = "Rathausplatz Hirschlanden"
	RATHAUSSTRASSE_EBERDINGEN                              = "Rathausstraße Eberdingen"
	RATTENHARZ_KAISERSTRKIRCHE_LORCH                       = "Rattenharz Kaiserstr./Kirche Lorch"
	RATTENHARZ_PULZHOFWEG_LORCH                            = "Rattenharz Pulzhofweg Lorch"
	RAUHER_KAPF_BOBLINGEN                                  = "Rauher Kapf Böblingen"
	RAUHWIESENSTRASSE_WINZINGEN                            = "Rauhwiesenstraße Winzingen"
	REALSCHULE_TAMM                                        = "Realschule Tamm"
	REALSCHULE_ZOLLBERG                                    = "Realschule Zollberg"
	REALSCHULE_WERNAU_N                                    = "Realschule Wernau (N)"
	REALSCHULE_NURTINGEN                                   = "Realschule Nürtingen"
	REALSCHULE_KORNWESTHEIM                                = "Realschule Kornwestheim"
	REALSCHULE_GROSSSACHSENHEIM                            = "Realschule Großsachsenheim"
	REALSCHULE_GERADSTETTEN                                = "Realschule Geradstetten"
	REALSCHULE_IM_AURAIN_BIETIGHEIM                        = "Realschule Im Aurain Bietigheim"
	REALSCHULE_REMSECK_PATTONVILLE                         = "Realschule Remseck Pattonville"
	REBENWEG_OWEN                                          = "Rebenweg Owen"
	REBHUHNWEG_LORCH                                       = "Rebhuhnweg Lorch"
	REBSTEIGE_MARKGRONINGEN                                = "Rebsteige Markgröningen"
	RECHBERGSTRASSE_KORNWESTHEIM                           = "Rechbergstraße Kornwestheim"
	RECHBERGSTRASSE_WEITMARS                               = "Rechbergstraße Weitmars"
	REESENBANKLE_PLUDERHAUSEN                              = "Reesenbänkle Plüderhausen"
	REICHENBACH_REICHENBACH_BERGLEN                        = "Reichenbach Reichenbach (Berglen)"
	REICHENBACH_F_REICHENBACH_F                            = "Reichenbach (F) Reichenbach (F)"
	REICHENBACHER_STRASSE_BALTMANNSWEILER                  = "Reichenbacher Straße Baltmannsweiler"
	REICHENBACHSTRASSE_REICHENBACH_F                       = "Reichenbachstraße Reichenbach (F)"
	REICHENBERG_OPPENWEILER                                = "Reichenberg Oppenweiler"
	REICHENHARDTSTRASSE_RECHBERGHAUSEN                     = "Reichenhardtstraße Rechberghausen"
	REICHERTSHALDE_LUDWIGSBURG                             = "Reichertshalde Ludwigsburg"
	REICHSDORFSTRASSE_HOHENSTAUFEN                         = "Reichsdorfstraße Hohenstaufen"
	REINHARDTSTRASSE_WOLFSCHLUGEN                          = "Reinhardtstraße Wolfschlugen"
	REINHOLDMAIERPLATZ_SCHORNDORF                          = "Reinhold-Maier-Platz Schorndorf"
	REINHOLDMAIERSTRASSE_GRUNBACH                          = "Reinhold-Maier-Straße Grunbach"
	REINHOLDSCHICKPLATZ_HERRENBERG                         = "Reinhold-Schick-Platz Herrenberg"
	REITERHOF_SINDELFINGEN                                 = "Reiterhof Sindelfingen"
	REITERHOF_BOBLINGEN                                    = "Reiterhof Böblingen"
	REITHALLE_SCHORNDORF                                   = "Reithalle Schorndorf"
	REITPRECHTS_AM_URSPRING_SCHWABISCH_GMUND               = "Reitprechts Am Urspring Schwäbisch Gmünd"
	REITPRECHTS_NEUBRUNNENGASSE_SCHWABISCH_GMUND           = "Reitprechts Neubrunnengasse Schwäbisch Gmünd"
	REITSTEIGE_AURICH                                      = "Reitsteige Aurich"
	REIZENWIESEN_WELZHEIM                                  = "Reizenwiesen Welzheim"
	REMBRANDTSTRASSE_STUTTGART                             = "Rembrandtstraße Stuttgart"
	REMBRANDTWEG_KIRCHHEIM_T                               = "Rembrandtweg Kirchheim (T)"
	REMSMURRCENTER_FELLBACH                                = "Rems-Murr-Center Fellbach"
	REMSMURRKLINIK_SCHORNDORF                              = "Rems-Murr-Klinik Schorndorf"
	REMSMURRKLINIKUM_WINNENDEN                             = "Rems-Murr-Klinikum Winnenden"
	REMSBRUCKE_WINTERBACH                                  = "Remsbrücke Winterbach"
	REMSECK_NECKARGRONINGEN                                = "Remseck Neckargröningen"
	REMSPARK_WAIBLINGEN                                    = "Remspark Waiblingen"
	REMSSTRASSE_LORCH                                      = "Remsstraße Lorch"
	REMSTALHALLE_WALDHAUSEN_B_SCHORND_LORCH                = "Remstalhalle Waldhausen b Schornd (Lorch)"
	REMSTALSTRASSE_NECKARREMS                              = "Remstalstraße Neckarrems"
	RENNINGEN_RENNINGEN                                    = "Renningen Renningen"
	RENNINGER_STRASSE_MAGSTADT                             = "Renninger Straße Magstadt"
	RESIDENZSCHLOSS_LUDWIGSBURG                            = "Residenzschloss Ludwigsburg"
	RESSESTRASSE_STUTTGART                                 = "Ressestraße Stuttgart"
	RETTERSBURG_RETTERSBURG                                = "Rettersburg Rettersburg"
	REUSCHKIRCHE_GOPPINGEN                                 = "Reuschkirche Göppingen"
	REUSSENSTEINSTRASSE_SCHOPFLOCH_LENNINGEN               = "Reußensteinstraße Schopfloch (Lenningen)"
	REUSSENSTEINSTRASSE_BOBLINGEN                          = "Reußensteinstraße Böblingen"
	REUSSENSTEINWEG_BONLANDEN                              = "Reußensteinweg Bonlanden"
	REUSSENSTEINWEG_HOCHDORF_ES                            = "Reußensteinweg Hochdorf (ES)"
	REUSTADT_HATTENHOFEN                                   = "Reustadt Hattenhofen"
	REUTLINGER_STRASSE_SIELMINGEN                          = "Reutlinger Straße Sielmingen"
	REUTLINGER_STRASSE_ESSLINGEN_N                         = "Reutlinger Straße Esslingen (N)"
	REUTLINGER_STRASSE_KEMNAT                              = "Reutlinger Straße Kemnat"
	REUTLINGER_STRASSE_STUTTGART                           = "Reutlinger Straße Stuttgart"
	REUTLINGER_STRASSE_GOPPINGEN                           = "Reutlinger Straße Göppingen"
	REUTTE_FRICKENHAUSEN                                   = "Reutte Frickenhausen"
	RICHARDWAGNERSTRASSE_BACKNANG                          = "Richard-Wagner-Straße Backnang"
	RICHTHOFENSTRASSE_GERLINGEN                            = "Richthofenstraße Gerlingen"
	RIEDBRUNNEN_GARTRINGEN                                 = "Riedbrunnen Gärtringen"
	RIEDENBERG_STUTTGART                                   = "Riedenberg Stuttgart"
	RIEDSEE_STUTTGART                                      = "Riedsee Stuttgart"
	RIEDWEG_LEINFELDEN                                     = "Riedweg Leinfelden"
	RIEDWIESEN_AICH                                        = "Riedwiesen Aich"
	RIEGELSTRASSE_NELLINGEN                                = "Riegelstraße Nellingen"
	RIELINGSHAUSER_STRASSE_MARBACH_N                       = "Rielingshäuser Straße Marbach (N)"
	RIEMENMUHLE_MERKLINGEN                                 = "Riemenmühle Merklingen"
	RIENHARZ_PFAHLBRONN                                    = "Rienharz Pfahlbronn"
	RIENHARZER_STRSCHULE_PFAHLBRONN                        = "Rienharzer Str./Schule Pfahlbronn"
	RIENZHOFER_MUHLE_BITTENFELD                            = "Rienzhofer Mühle Bittenfeld"
	RIESENGEBIRGSTRASSE_FAURNDAU                           = "Riesengebirgstraße Faurndau"
	RIETENAU_RIETENAU                                      = "Rietenau Rietenau"
	RIETENAUER_WEG_BACKNANG                                = "Rietenauer Weg Backnang"
	RIETER_STRASSE_HOCHDORF_EBERDINGEN                     = "Rieter Straße Hochdorf (Eberdingen)"
	RIETER_TAL_RIET                                        = "Rieter Tal Riet"
	RIETHMULLERHAUS_STUTTGART                              = "Riethmüllerhaus Stuttgart"
	RIGIPARK_HOLZHEIM_GOPPINGEN                            = "Rigipark Holzheim (Göppingen)"
	RILKEREALSCHULE_STUTTGART                              = "Rilke-Realschule Stuttgart"
	RINGSTRASSE_MAGSTADT                                   = "Ringstraße Magstadt"
	RINGSTRASSE_WINNENDEN                                  = "Ringstraße Winnenden"
	RINGSTRASSE_GERLINGEN                                  = "Ringstraße Gerlingen"
	RINGSTRASSE_SALACH                                     = "Ringstraße Salach"
	RISSHALDE_REICHENBACH_F                                = "Risshalde Reichenbach (F)"
	RISSLERINSTRASSE_SCHORNDORF                            = "Risslerinstraße Schorndorf"
	RITTERSTRASSE_SCHOCKINGEN                              = "Ritterstraße Schöckingen"
	ROBERTBOSCHCAMPUS_MALMSHEIM                            = "Robert-Bosch-Campus Malmsheim"
	ROBERTBOSCHKRANKENHAUS_STUTTGART                       = "Robert-Bosch-Krankenhaus Stuttgart"
	ROBERTBOSCHPLATZ_GERLINGEN                             = "Robert-Bosch-Platz Gerlingen"
	ROBERTBOSCHSTRASSE_SCHWIEBERDINGEN                     = "Robert-Bosch-Straße Schwieberdingen"
	ROBERTBOSCHSTRASSE_TAMM                                = "Robert-Bosch-Straße Tamm"
	ROBERTBOSCHSTRASSE_NECKARTENZLINGEN                    = "Robert-Bosch-Straße Neckartenzlingen"
	ROBERTBOSCHSTRASSE_DETTINGEN_T                         = "Robert-Bosch-Straße Dettingen (T)"
	ROBERTBOSCHSTRASSE_DARMSHEIM                           = "Robert-Bosch-Straße Darmsheim"
	ROBERTBOSCHSTRASSE_PFLUGFELDEN                         = "Robert-Bosch-Straße Pflugfelden"
	ROBERTBOSCHSTRASSE_PLEIDELSHEIM                        = "Robert-Bosch-Straße Pleidelsheim"
	ROBERTBOSCHSTRASSE_HOLZGERLINGEN                       = "Robert-Bosch-Straße Holzgerlingen"
	ROBERTBOSCHSTRASSE_SCHORNDORF                          = "Robert-Bosch-Straße Schorndorf"
	ROBERTFRANCKALLEE_LUDWIGSBURG                          = "Robert-Franck-Allee Ludwigsburg"
	ROBERTFRANCKSTRASSE_MURRHARDT                          = "Robert-Franck-Straße Murrhardt"
	ROBERTKOCHSTRASSE_PARKSIEDLUNG                         = "Robert-Koch-Straße Parksiedlung"
	ROBERTKOCHSTRASSE_LUDWIGSBURG                          = "Robert-Koch-Straße Ludwigsburg"
	ROGGENMUHLE_EYBACH                                     = "Roggenmühle Eybach"
	ROHR_SCHORNBACH                                        = "Rohr Schornbach"
	ROHR_STUTTGART                                         = "Rohr Stuttgart"
	ROHR_MITTE_STUTTGART                                   = "Rohr Mitte Stuttgart"
	ROHRACH_ALLMERSBACH_AM_WEINBERG_ASPACH                 = "Röhrach Allmersbach am Weinberg (Aspach)"
	ROHRACHWEG_SCHORNDORF                                  = "Röhrachweg Schorndorf"
	ROHRACKER_STUTTGART                                    = "Rohracker Stuttgart"
	ROHRACKER_KELTER_STUTTGART                             = "Rohracker Kelter Stuttgart"
	ROHRBRONN_ROHRBRONN                                    = "Rohrbronn Rohrbronn"
	ROHRER_HOHE_STUTTGART                                  = "Rohrer Höhe Stuttgart"
	ROHRER_STRASSE_STEINENBRONN                            = "Rohrer Straße Steinenbronn"
	ROHRER_WEG_STUTTGART                                   = "Rohrer Weg Stuttgart"
	ROMERSTRASSE_LEONBERG                                  = "Römerstraße Leonberg"
	ROMERSTRASSE_DETTINGEN_T                               = "Römerstraße Dettingen (T)"
	ROMERWEG_KIRCHBERG_M                                   = "Römerweg Kirchberg (M)"
	ROMMELSHAUSEN_ROMMELSHAUSEN                            = "Rommelshausen Rommelshausen"
	ROMMELSHAUSER_STRASSE_FELLBACH                         = "Rommelshauser Straße Fellbach"
	ROMMENTALER_STRASSE_SCHLAT                             = "Rommentaler Straße Schlat"
	RONTGENSTRASSE_OBERESSLINGEN                           = "Röntgenstraße Oberesslingen"
	RONTGENWEG_GOPPINGEN                                   = "Röntgenweg Göppingen"
	ROSE_SCHOCKINGEN                                       = "Rose Schöckingen"
	ROSEGGERWEG_GOPPINGEN                                  = "Roseggerweg Göppingen"
	ROSENACKERWEG_EGLOSHEIM                                = "Rosenackerweg Eglosheim"
	ROSENBERGJOHANNESSTRASSE_STUTTGART                     = "Rosenberg-/Johannesstraße Stuttgart"
	ROSENBERGSEIDENSTRASSE_STUTTGART                       = "Rosenberg-/Seidenstraße Stuttgart"
	ROSENBERGPLATZ_STUTTGART                               = "Rosenbergplatz Stuttgart"
	ROSENPLATZ_GROSSBOTTWAR                                = "Rosenplatz Großbottwar"
	ROSENSTEINBRUCKE_STUTTGART                             = "Rosensteinbrücke Stuttgart"
	ROSENSTEINPARK_STUTTGART                               = "Rosensteinpark Stuttgart"
	ROSENSTEINSTRASSE_STUTTGART                            = "Rosensteinstraße Stuttgart"
	ROSENSTRASSE_BERNHAUSEN                                = "Rosenstraße Bernhausen"
	ROSENSTRASSE_KEMNAT                                    = "Rosenstraße Kemnat"
	ROSENSTRASSE_OBERBRUDEN                                = "Rosenstraße Oberbrüden"
	ROSENSTRASSE_GERLINGEN                                 = "Rosenstraße Gerlingen"
	ROSSACKER_BEINSTEIN                                    = "Rossäcker Beinstein"
	ROSSBACHSTRASSE_STUTTGART                              = "Roßbachstraße Stuttgart"
	ROSSBACHSTRASSE_GOPPINGEN                              = "Roßbachstraße Göppingen"
	ROSSBAUM_GERLINGEN                                     = "Rossbaum Gerlingen"
	ROSSBERGSTAFFEL_BACKNANG                               = "Rossbergstaffel Backnang"
	ROSSBERGSTRASSE_HARTHAUSEN                             = "Roßbergstraße Harthausen"
	ROSSBERGWEG_HOLZGERLINGEN                              = "Roßbergweg Holzgerlingen"
	ROSSDORF_NURTINGEN                                     = "Roßdorf Nürtingen"
	ROSSDORF_BERLINER_STRASSE_NURTINGEN                    = "Roßdorf Berliner Straße Nürtingen"
	ROSSDORF_DURERPLATZ_NURTINGEN                          = "Roßdorf Dürerplatz Nürtingen"
	ROSSDORF_HOLBEINSTRASSE_NURTINGEN                      = "Roßdorf Holbeinstraße Nürtingen"
	ROSSDORF_KLEEWEG_NURTINGEN                             = "Roßdorf Kleeweg Nürtingen"
	ROSSDORF_LIEBERMANNSTRASSE_NURTINGEN                   = "Roßdorf Liebermannstraße Nürtingen"
	ROSSDORF_SCHULE_NURTINGEN                              = "Roßdorf Schule Nürtingen"
	ROSSLE_SCHANBACH                                       = "Rössle Schanbach"
	ROSSMARKT_KIRCHHEIM_T                                  = "Rossmarkt Kirchheim (T)"
	ROSSWAG_SPORTHALLE_ROSSWAG                             = "Roßwag Sporthalle Roßwag"
	ROSSWALDER_STRASSE_HOCHDORF_ES                         = "Roßwälder Straße Hochdorf (ES)"
	ROTBAUMLESFELD_LUDWIGSBURG                             = "Rotbäumlesfeld Ludwigsburg"
	ROTBUHL_SINDELFINGEN                                   = "Rotbühl Sindelfingen"
	ROTE_WEIL_IM_SCHONBUCH                                 = "Röte Weil im Schönbuch"
	ROTE_DOFFINGEN                                         = "Röte Döffingen"
	ROTEBUHLREINSBURGSTRASSE_STUTTGART                     = "Rotebühl-/Reinsburgstraße Stuttgart"
	ROTENBACH_NECKARTENZLINGEN                             = "Rotenbach Neckartenzlingen"
	ROTENBERG_STUTTGART                                    = "Rotenberg Stuttgart"
	ROTENBERGPLATZ_NURTINGEN                               = "Rotenbergplatz Nürtingen"
	ROTER_BAUM_OBERRIEXINGEN                               = "Roter Baum Oberriexingen"
	ROTER_BERG_SCHONAICH                                   = "Roter Berg Schönaich"
	ROTES_KREUZ_LIEBERSBRONN                               = "Rotes Kreuz Liebersbronn"
	ROTES_REUSCH_BARTENBACH_GOPPINGEN                      = "Rotes Reusch Bartenbach (Göppingen)"
	ROTESIEDLUNG_BURGSTALL_BURGSTETTEN                     = "Rötesiedlung Burgstall (Burgstetten)"
	ROTHWEG_DENKENDORF                                     = "Rothweg Denkendorf"
	ROTWEG_HOCHBERG                                        = "Rotweg Hochberg"
	ROTWIESENSTRASSE_STUTTGART                             = "Rotwiesenstraße Stuttgart"
	ROTZEIL_BONLANDEN                                      = "Rotzeil Bonlanden"
	RUBEZAHLWEG_STUTTGART                                  = "Rübezahlweg Stuttgart"
	RUDERSBERG_RUDERSBERG                                  = "Rudersberg Rudersberg"
	RUDERSBERGER_STRASSE_MIEDELSBACH                       = "Rudersberger Straße Miedelsbach"
	RUDOLFDIESELSTR_LAICHINGEN                             = "Rudolf-Diesel-Str. Laichingen"
	RUDOLFDIESELSTRASSE_ROMMELSHAUSEN                      = "Rudolf-Diesel-Straße Rommelshausen"
	RUDOLFSOPHIENSTIFT_STUTTGART                           = "Rudolf-Sophien-Stift Stuttgart"
	RUDOLFSHOHE_AICH                                       = "Rudolfshöhe Aich"
	RUHBANK_FERNSEHTURM_STUTTGART                          = "Ruhbank (Fernsehturm) Stuttgart"
	RUHRBERG_MUNCHINGEN                                    = "Rührberg Münchingen"
	RUHRSTRASSE_LUDWIGSBURG                                = "Ruhrstraße Ludwigsburg"
	RUIT_RUIT                                              = "Ruit Ruit"
	RUMOLDREALSCHULE_ROMMELSHAUSEN                         = "Rumold-Realschule Rommelshausen"
	RUNDSPORTHALLE_LUDWIGSBURG                             = "Rundsporthalle Ludwigsburg"
	RUNDSPORTHALLE_WAIBLINGEN                              = "Rundsporthalle Waiblingen"
	RUSSISCHE_KIRCHE_STUTTGART                             = "Russische Kirche Stuttgart"
	RUTESHEIM_LEONBERG                                     = "Rutesheim Leonberg"
	SAARSTRASSE_ASPERG                                     = "Saarstraße Asperg"
	SAARSTRASSE_KIRCHHEIM_T                                = "Saarstraße Kirchheim (T)"
	SACHSENHEIM_GROSSSACHSENHEIM                           = "Sachsenheim Großsachsenheim"
	SACHSENHEIMER_STRASSE_OBERRIEXINGEN                    = "Sachsenheimer Straße Oberriexingen"
	SACHSENHEIMER_WEG_BESIGHEIM                            = "Sachsenheimer Weg Besigheim"
	SAERSTRASSE_NURTINGEN                                  = "Säerstraße Nürtingen"
	SAGMUHLE_HEMMINGEN                                     = "Sägmühle Hemmingen"
	SALACH_SALACH                                          = "Salach Salach"
	SALAMANDERSTRASSE_FAURNDAU                             = "Salamanderstraße Faurndau"
	SALAMANDERWEG_STUTTGART                                = "Salamanderweg Stuttgart"
	SALIERSCHULE_WAIBLINGEN                                = "Salierschule Waiblingen"
	SALIERSTRASSE_WAIBLINGEN                               = "Salierstraße Waiblingen"
	SALZACKER_STUTTGART                                    = "Salzäcker Stuttgart"
	SALZACKER_VAIHINGEN_E                                  = "Salzäcker Vaihingen (E)"
	SALZBURGER_STRASSE_MAUBACH                             = "Salzburger Straße Maubach"
	SALZWIESENSTRASSE_STUTTGART                            = "Salzwiesenstraße Stuttgart"
	SAMARITERSTIFT_LEONBERG                                = "Samariterstift Leonberg"
	SANGERSTRASSE_EISLINGEN_F                              = "Sängerstraße Eislingen (F)"
	SANKT_MARTIN_HERRENBERG                                = "Sankt Martin Herrenberg"
	SANKT_PAULS_KIRCHE_GOPPINGEN                           = "Sankt Pauls Kirche Göppingen"
	SATTLEREI_EISELE_HESSIGHEIM                            = "Sattlerei Eisele Hessigheim"
	SAUERBRUNNEN_FAURNDAU                                  = "Sauerbrunnen Faurndau"
	SAUERHOFLE_MURRHARDT                                   = "Sauerhöfle Murrhardt"
	SAUSERHOF_GROSSBOTTWAR                                 = "Sauserhof Großbottwar"
	SAUSTEIGE_GOPPINGEN                                    = "Sausteige Göppingen"
	SCSTADION_GEISLINGEN_STEIGE                            = "SC-Stadion Geislingen (Steige)"
	SCHADBERG_KAISERSBACH                                  = "Schadberg Kaisersbach"
	SCHAFERSFELD_SCHULEN_LORCH                             = "Schäfersfeld Schulen Lorch"
	SCHAFGARTENSTRASSE_MUSBERG                             = "Schafgartenstraße Musberg"
	SCHAFGASSE_BOBLINGEN                                   = "Schafgasse Böblingen"
	SCHAFHAUS_KLEINASPACH                                  = "Schafhaus Kleinaspach"
	SCHAFHAUSER_STRASSE_AIDLINGEN                          = "Schafhauser Straße Aidlingen"
	SCHAFHOF_KIRCHHEIM_T                                   = "Schafhof Kirchheim (T)"
	SCHAFHOF_WELZHEIM                                      = "Schafhof Welzheim"
	SCHAFSTRASSE_ROMMELSHAUSEN                             = "Schafstraße Rommelshausen"
	SCHAFWASCHE_HEIMSHEIM                                  = "Schafwäsche Heimsheim"
	SCHAICHHOFSIEDLUNG_WEIL_IM_SCHONBUCH                   = "Schaichhofsiedlung Weil im Schönbuch"
	SCHAICHHOFSTRASSE_WEIL_IM_SCHONBUCH                    = "Schaichhofstraße Weil im Schönbuch"
	SCHANZLE_WAIBLINGEN                                    = "Schänzle Waiblingen"
	SCHARNHAUSER_BRUCKE_STUTTGART                          = "Scharnhäuser Brücke Stuttgart"
	SCHARNHAUSER_PARK_SCHARNHAUSER_PARK                    = "Scharnhauser Park Scharnhauser Park"
	SCHARNHAUSER_STRASSE_RUIT                              = "Scharnhäuser Straße Ruit"
	SCHATTENGRUND_STUTTGART                                = "Schattengrund Stuttgart"
	SCHAUBER_BESIGHEIM                                     = "Schäuber Besigheim"
	SCHAUCHERT_HEMMINGEN                                   = "Schauchert Hemmingen"
	SCHEERSTRDRENGEL_REALSCHULE_EISLINGEN_F                = "Scheerstr./Dr.Engel Realschule Eislingen (F)"
	SCHELLINGSTRASSE_NURTINGEN                             = "Schellingstraße Nürtingen"
	SCHELLINGSTRASSE_WAIBLINGEN                            = "Schellingstraße Waiblingen"
	SCHELMENACKER_HASLACH                                  = "Schelmenäcker Haslach"
	SCHELMENSTRASSEFRIEDHOF_BARTENBACH_GOPPINGEN           = "Schelmenstraße/Friedhof Bartenbach (Göppingen)"
	SCHELMENWASEN_STUTTGART                                = "Schelmenwasen Stuttgart"
	SCHELMENWASEN_NURTINGEN                                = "Schelmenwasen Nürtingen"
	SCHELMENWIESEN_GARTRINGEN                              = "Schelmenwiesen Gärtringen"
	SCHELZTOR_ESSLINGEN_N                                  = "Schelztor Esslingen (N)"
	SCHEMPPSTRASSE_STUTTGART                               = "Schemppstraße Stuttgart"
	SCHICKARDSTRASSE_BOBLINGEN                             = "Schickardstraße Böblingen"
	SCHICKHARDTSCHULE_STUTTGART                            = "Schickhardtschule Stuttgart"
	SCHIESSTAL_NECKARGRONINGEN                             = "Schießtal Neckargröningen"
	SCHIESSTALE_HERRENBERG                                 = "Schießtäle Herrenberg"
	SCHILDFARNWEG_STUTTGART                                = "Schildfarnweg Stuttgart"
	SCHILDWACHTWEG_GEISLINGEN_STEIGE                       = "Schildwachtweg Geislingen (Steige)"
	SCHILLERREALSCHULE_GOPPINGEN                           = "Schiller-Realschule Göppingen"
	SCHILLERHOHE_MARBACH_N                                 = "Schillerhöhe Marbach (N)"
	SCHILLERHOHE_BOSCH_GERLINGEN                           = "Schillerhöhe Bosch Gerlingen"
	SCHILLERPLATZ_BACKNANG                                 = "Schillerplatz Backnang"
	SCHILLERPLATZ_ESSLINGEN_N                              = "Schillerplatz Esslingen (N)"
	SCHILLERPLATZ_NURTINGEN                                = "Schillerplatz Nürtingen"
	SCHILLERPLATZ_SCHORNDORF                               = "Schillerplatz Schorndorf"
	SCHILLERPLATZ_KIRCHHEIM_N                              = "Schillerplatz Kirchheim (N)"
	SCHILLERPLATZ_BIETIGHEIM                               = "Schillerplatz Bietigheim"
	SCHILLERPLATZ_GOPPINGEN                                = "Schillerplatz Göppingen"
	SCHILLERSCHULE_BISSINGEN_LB                            = "Schillerschule Bissingen (LB)"
	SCHILLERSCHULE_BITTENFELD                              = "Schillerschule Bittenfeld"
	SCHILLERSTRASSE_PLATTENHARDT                           = "Schillerstraße Plattenhardt"
	SCHILLERSTRASSE_NEUHAUSEN_F                            = "Schillerstraße Neuhausen (F)"
	SCHILLERSTRASSE_SCHONAICH                              = "Schillerstraße Schönaich"
	SCHILLERSTRASSE_REICHENBACH_F                          = "Schillerstraße Reichenbach (F)"
	SCHILLERSTRASSE_NECKARTAILFINGEN                       = "Schillerstraße Neckartailfingen"
	SCHILLERSTRASSE_KLEINGLATTBACH                         = "Schillerstraße Kleinglattbach"
	SCHILLERSTRASSE_RUTESHEIM                              = "Schillerstraße Rutesheim"
	SCHILLERSTRASSE_KOHLBERG                               = "Schillerstraße Kohlberg"
	SCHILLERSTRASSE_KUCHEN_F                               = "Schillerstraße Kuchen (F)"
	SCHILLERSTRASSE_WASCHENBEUREN                          = "Schillerstraße Wäschenbeuren"
	SCHILLERSTRASSE_ZELL_A                                 = "Schillerstraße Zell (A)"
	SCHILLERSTRASSE_GOPPINGEN                              = "Schillerstraße Göppingen"
	SCHILLERSTRASSE_SCHWABISCH_GMUND                       = "Schillerstraße Schwäbisch Gmünd"
	SCHIPPERTSACKER_WAIBLINGEN                             = "Schippertsäcker Waiblingen"
	SCHLACHTHAUSBRUCKE_ESSLINGEN_N                         = "Schlachthausbrücke Esslingen (N)"
	SCHLACHTHOF_STUTTGART                                  = "Schlachthof Stuttgart"
	SCHLATER_STRASSE_SUSSEN                                = "Schlater Straße Süßen"
	SCHLATTERHOHE_B465_SCHOPFLOCH_LENNINGEN                = "Schlatterhöhe B465 Schopfloch (Lenningen)"
	SCHLATTERHOHE_L_1212_SCHOPFLOCH_LENNINGEN              = "Schlatterhöhe L 1212 Schopfloch (Lenningen)"
	SCHLECHTBACH_SCHLECHTBACH                              = "Schlechtbach Schlechtbach"
	SCHLEIERMACHERSTRASSE_LEONBERG                         = "Schleiermacherstraße Leonberg"
	SCHLICHTENER_STRASSE_SCHORNDORF                        = "Schlichtener Straße Schorndorf"
	SCHLIEFFENSTRASSE_LUDWIGSBURG                          = "Schlieffenstraße Ludwigsburg"
	SCHLIENZTOURS_ROMMELSHAUSEN                            = "Schlienz-Tours Rommelshausen"
	SCHLIERBACHER_DREIECK_KIRCHHEIM_T                      = "Schlierbacher Dreieck Kirchheim (T)"
	SCHLIPFWEG_SCHORNDORF                                  = "Schlipfweg Schorndorf"
	SCHLOSS_SCHOCKINGEN                                    = "Schloss Schöckingen"
	SCHLOSS_WALDENBUCH                                     = "Schloss Waldenbuch"
	SCHLOSS_URBACH                                         = "Schloss Urbach"
	SCHLOSS_FAVORITE_LUDWIGSBURG                           = "Schloss Favorite Ludwigsburg"
	SCHLOSS_FILSECK_FAURNDAU                               = "Schloss Filseck Faurndau"
	SCHLOSS_LAUTERECK_SULZBACH_M                           = "Schloss Lautereck Sulzbach (M)"
	SCHLOSSJOHANNESSTRASSE_STUTTGART                       = "Schloss-/Johannesstraße Stuttgart"
	SCHLOSSBERGALLEE_BONNIGHEIM                            = "Schlossbergallee Bönnigheim"
	SCHLOSSBERGHALLE_DETTINGEN_T                           = "Schlossberghalle Dettingen (T)"
	SCHLOSSERSTRASSE_NEUHAUSEN_F                           = "Schlosserstraße Neuhausen (F)"
	SCHLOSSFELD_BONNIGHEIM                                 = "Schlossfeld Bönnigheim"
	SCHLOSSGARTEN_NEIDLINGEN                               = "Schlossgärten Neidlingen"
	SCHLOSSGYMNASIUM_KIRCHHEIM_T                           = "Schlossgymnasium Kirchheim (T)"
	SCHLOSSHALDE_GEISLINGEN_STEIGE                         = "Schlosshalde Geislingen (Steige)"
	SCHLOSSHALDENSTRASSE_HEMMINGEN                         = "Schlosshaldenstraße Hemmingen"
	SCHLOSSHOF_ALDINGEN                                    = "Schlosshof Aldingen"
	SCHLOSSHOF_FORNSBACH                                   = "Schlosshof Fornsbach"
	SCHLOSSHOFSTRASSE_RECHBERGHAUSEN                       = "Schlosshofstraße Rechberghausen"
	SCHLOSSLESFELD_LUDWIGSBURG                             = "Schlösslesfeld Ludwigsburg"
	SCHLOSSMARKT_RECHBERGHAUSEN                            = "Schlossmarkt Rechberghausen"
	SCHLOSSMUHLE_KIRCHENKIRNBERG                           = "Schlossmühle Kirchenkirnberg"
	SCHLOSSPLATZ_NEUHAUSEN_F                               = "Schlossplatz Neuhausen (F)"
	SCHLOSSPLATZ_STUTTGART                                 = "Schlossplatz Stuttgart"
	SCHLOSSPLATZBAHNHOF_EISLINGEN_F                        = "Schlossplatz/Bahnhof Eislingen (F)"
	SCHLOSSSTRASSE_SCHORNDORF                              = "Schlossstraße Schorndorf"
	SCHLOSSSTRASSE_KLEININGERSHEIM                         = "Schlossstraße Kleiningersheim"
	SCHLOTTERBECKSTRASSE_STUTTGART                         = "Schlotterbeckstraße Stuttgart"
	SCHLOTWIESE_STUTTGART                                  = "Schlotwiese Stuttgart"
	SCHLUCHTWEG_WEILHEIM_T                                 = "Schluchtweg Weilheim (T)"
	SCHMELLBACHTAL_LEINFELDEN                              = "Schmellbachtal Leinfelden"
	SCHMELLENHOF_WUSTENROT                                 = "Schmellenhof Wüstenrot"
	SCHMIDENER_STRASSE_WAIBLINGEN                          = "Schmidener Straße Waiblingen"
	SCHMIDHAUSEN_ABZWEIG_BEILSTEIN                         = "Schmidhausen Abzweig Beilstein"
	SCHMIEDE_HERTMANNSWEILER                               = "Schmiede Hertmannsweiler"
	SCHNALLENACKER_MALMSHEIM                               = "Schnallenäcker Malmsheim"
	SCHNARRENBERG_STUTTGART                                = "Schnarrenberg Stuttgart"
	SCHNECKEN_PLATTENHARDT                                 = "Schnecken Plattenhardt"
	SCHNEIDER_NEUSTADT                                     = "Schneider Neustadt"
	SCHNEIDERACKERSTRASSE_STUTTGART                        = "Schneideräckerstraße Stuttgart"
	SCHOCKINGER_STRASSE_HEMMINGEN                          = "Schöckinger Straße Hemmingen"
	SCHONAICHER_FIRST_BOBLINGEN                            = "Schönaicher First Böblingen"
	SCHONAICHER_STRASSE_BOBLINGEN                          = "Schönaicher Straße Böblingen"
	SCHONBERG_STUTTGART                                    = "Schönberg Stuttgart"
	SCHONBERG_HOLZGERLINGEN                                = "Schönberg Holzgerlingen"
	SCHONBERGSTRASSE_KEMNAT                                = "Schönbergstraße Kemnat"
	SCHONBLICK_REICHENBACH_F                               = "Schönblick Reichenbach (F)"
	SCHONBLICK_GROTZINGEN                                  = "Schönblick Grötzingen"
	SCHONBLICKSTRASSE_WUSTENROT                            = "Schönblickstraße Wüstenrot"
	SCHONBRONN_GRAB                                        = "Schönbronn Grab"
	SCHONBUCHSTRASSE_PLATTENHARDT                          = "Schönbuchstraße Plattenhardt"
	SCHONBUCHSTRASSE_BOBLINGEN                             = "Schönbuchstraße Böblingen"
	SCHONBUCHSTRASSE_SPORTH_BOBLINGEN                      = "Schönbuchstraße (Sporth.) Böblingen"
	SCHONTALWEG_WIESENSTEIG                                = "Schöntalweg Wiesensteig"
	SCHOPFLENBERG_BEZGENRIET                               = "Schopflenberg Bezgenriet"
	SCHOPFLOCH_ELTINGEN                                    = "Schopfloch Eltingen"
	SCHORNBACHER_WEG_SCHORNDORF                            = "Schornbacher Weg Schorndorf"
	SCHORNDORF_SCHORNDORF                                  = "Schorndorf Schorndorf"
	SCHORNDORFER_STRASSE_GRUNBACH                          = "Schorndorfer Straße Grunbach"
	SCHORNDORFER_STRASSE_OBERESSLINGEN                     = "Schorndorfer Straße Oberesslingen"
	SCHORNDORFER_STRASSE_STEINENBERG                       = "Schorndorfer Straße Steinenberg"
	SCHORNDORFER_STRASSE_REICHENBACH_F                     = "Schorndorfer Straße Reichenbach (F)"
	SCHORNDORFER_STRASSE_URBACH                            = "Schorndorfer Straße Urbach"
	SCHORNDORFER_TOR_LUDWIGSBURG                           = "Schorndorfer Tor Ludwigsburg"
	SCHOSSHOFE_SINDELFINGEN                                = "Schoßhöfe Sindelfingen"
	SCHOTTSTRASSE_STUTTGART                                = "Schottstraße Stuttgart"
	SCHOZACHER_STRASSE_STUTTGART                           = "Schozacher Straße Stuttgart"
	SCHOZACHSTRASSE_WALDREMS                               = "Schozachstraße Waldrems"
	SCHRANNE_STUTTGART                                     = "Schranne Stuttgart"
	SCHUBARTSTRASSE_BISSINGEN_LB                           = "Schubartstraße Bissingen (LB)"
	SCHUBARTSTRASSE_FAURNDAU                               = "Schubartstraße Faurndau"
	SCHUBERTSTRASSE_WOLFSCHLUGEN                           = "Schubertstraße Wolfschlugen"
	SCHULSPORTZENTRUM_SIELMINGEN                           = "Schul-/Sportzentrum Sielmingen"
	SCHULBAD_GOPPINGEN                                     = "Schulbad Göppingen"
	SCHULE_ALTENRIET                                       = "Schule Altenriet"
	SCHULE_OBERENSINGEN                                    = "Schule Oberensingen"
	SCHULE_OBERSTENFELD                                    = "Schule Oberstenfeld"
	SCHULE_KIRCHBERG_M                                     = "Schule Kirchberg (M)"
	SCHULE_HESSIGHEIM                                      = "Schule Hessigheim"
	SCHULE_ALTHUTTE                                        = "Schule Althütte"
	SCHULE_DEIZISAU                                        = "Schule Deizisau"
	SCHULE_SCHANBACH                                       = "Schule Schanbach"
	SCHULE_DATZINGEN                                       = "Schule Dätzingen"
	SCHULE_SULZGRIES                                       = "Schule Sulzgries"
	SCHULE_FRICKENHAUSEN                                   = "Schule Frickenhausen"
	SCHULE_NECKARTAILFINGEN                                = "Schule Neckartailfingen"
	SCHULE_BEMPFLINGEN                                     = "Schule Bempflingen"
	SCHULE_SECHSELBERG                                     = "Schule Sechselberg"
	SCHULE_ALLMERSBACH_IM_TAL                              = "Schule Allmersbach im Tal"
	SCHULE_SPIEGELBERG                                     = "Schule Spiegelberg"
	SCHULE_LEUTENBACH                                      = "Schule Leutenbach"
	SCHULE_GROSSINGERSHEIM                                 = "Schule Großingersheim"
	SCHULE_GRONAU                                          = "Schule Gronau"
	SCHULE_HOHENHASLACH                                    = "Schule Hohenhaslach"
	SCHULE_ALFDORF                                         = "Schule Alfdorf"
	SCHULE_NEUFFEN                                         = "Schule Neuffen"
	SCHULE_SACHSENWEILER                                   = "Schule Sachsenweiler"
	SCHULE_MONCHBERG                                       = "Schule Mönchberg"
	SCHULE_GROSSERLACH                                     = "Schule Großerlach"
	SCHULE_AUFHAUSEN_GEISLINGEN                            = "Schule Aufhausen (Geislingen)"
	SCHULE_TREFFELHAUSEN                                   = "Schule Treffelhausen"
	SCHULE_DETTENHAUSEN                                    = "Schule Dettenhausen"
	SCHULE_WUSTENROT                                       = "Schule Wüstenrot"
	SCHULE_SCHALKSTETTEN                                   = "Schule Schalkstetten"
	SCHULEHALLENBAD_WEIL_IM_SCHONBUCH                      = "Schule/Hallenbad Weil im Schönbuch"
	SCHULERHOF_BURG                                        = "Schulerhof Bürg"
	SCHULMEISTERBUCHE_HERRENBERG                           = "Schulmeisterbuche Herrenberg"
	SCHULSTRASSE_TAILFINGEN                                = "Schulstraße Tailfingen"
	SCHULSTRASSE_PLUDERHAUSEN                              = "Schulstraße Plüderhausen"
	SCHULSTRASSE_KIRCHBERG_M                               = "Schulstraße Kirchberg (M)"
	SCHULSTRASSE_UHINGEN                                   = "Schulstraße Uhingen"
	SCHULSTRASSE_WANGEN_GP                                 = "Schulstraße Wangen (GP)"
	SCHULTHEISSSCHNEIDERSTRASSE_GEISLINGEN_STEIGE          = "Schultheiß-Schneider-Straße Geislingen (Steige)"
	SCHULZEDELITZSCHSTRASSE_STUTTGART                      = "Schulze-Delitzsch-Straße Stuttgart"
	SCHULZENTRUM_SCHWIEBERDINGEN                           = "Schulzentrum Schwieberdingen"
	SCHULZENTRUM_GROSSBOTTWAR                              = "Schulzentrum Großbottwar"
	SCHULZENTRUM_REICHENBACH_F                             = "Schulzentrum Reichenbach (F)"
	SCHULZENTRUM_WENDLINGEN_N                              = "Schulzentrum Wendlingen (N)"
	SCHULZENTRUM_NELLINGEN                                 = "Schulzentrum Nellingen"
	SCHULZENTRUM_RUDERSBERG                                = "Schulzentrum Rudersberg"
	SCHULZENTRUM_WELZHEIM                                  = "Schulzentrum Welzheim"
	SCHULZENTRUM_BONNIGHEIM                                = "Schulzentrum Bönnigheim"
	SCHULZENTRUM_RUTESHEIM                                 = "Schulzentrum Rutesheim"
	SCHULZENTRUM_STEINHEIM_M                               = "Schulzentrum Steinheim (M)"
	SCHULZENTRUM_DONZDORF                                  = "Schulzentrum Donzdorf"
	SCHULZENTRUM_BEILSTEIN                                 = "Schulzentrum Beilstein"
	SCHULZENTRUM_LAICHINGEN                                = "Schulzentrum Laichingen"
	SCHULZENTRUM_AM_BERG_WENDLINGEN_N                      = "Schulzentrum am Berg Wendlingen (N)"
	SCHULZENTRUM_MURKENBACH_BOBLINGEN                      = "Schulzentrum Murkenbach Böblingen"
	SCHULZENTRUM_NORD_SERACH                               = "Schulzentrum Nord Serach"
	SCHUMISBERG_LEONBERG                                   = "Schumisberg Leonberg"
	SCHURWALDHALLE_OBERBERKEN                              = "Schurwaldhalle Oberberken"
	SCHURWALDHOHE_BALTMANNSWEILER                          = "Schurwaldhöhe Baltmannsweiler"
	SCHURWALDSTRASSE_BEUTELSBACH                           = "Schurwaldstraße Beutelsbach"
	SCHURWALDWEG_BIRENBACH                                 = "Schurwaldweg Birenbach"
	SCHUTZENHOF_WASCHENBEUREN                              = "Schützenhof Wäschenbeuren"
	SCHUTZENSTRASSE_WENDLINGEN_N                           = "Schützenstraße Wendlingen (N)"
	SCHUTZENSTRASSE_GOPPINGEN                              = "Schützenstraße Göppingen"
	SCHUTZENWEG_DAGERSHEIM                                 = "Schützenweg Dagersheim"
	SCHWABBEBELSTRASSE_STUTTGART                           = "Schwab-/Bebelstraße Stuttgart"
	SCHWABREINSBURGSTRASSE_STUTTGART                       = "Schwab-/Reinsburgstraße Stuttgart"
	SCHWABENGALERIE_STUTTGART                              = "Schwabengalerie Stuttgart"
	SCHWABENLANDHALLE_FELLBACH                             = "Schwabenlandhalle Fellbach"
	SCHWABENSTRASSE_SCHONAICH                              = "Schwabenstraße Schönaich"
	SCHWABISCH_GMUND_SCHWABISCH_GMUND                      = "Schwäbisch Gmünd Schwäbisch Gmünd"
	SCHWABSTRASSE_BOBLINGEN                                = "Schwabstraße Böblingen"
	SCHWABSTRASSE_STUTTGART                                = "Schwabstraße Stuttgart"
	SCHWABSTRASSE_WAIBLINGEN                               = "Schwabstraße Waiblingen"
	SCHWAIKHEIM_SCHWAIKHEIM                                = "Schwaikheim Schwaikheim"
	SCHWALBENFLUG_GRAB                                     = "Schwalbenflug Grab"
	SCHWALBENHALDE_BESIGHEIM                               = "Schwalbenhälde Besigheim"
	SCHWALBENWEG_SIRNAU                                    = "Schwalbenweg Sirnau"
	SCHWALBENWEG_PLUDERHAUSEN                              = "Schwalbenweg Plüderhausen"
	SCHWAMMHOF_MURRHARDT                                   = "Schwammhof Murrhardt"
	SCHWANEN_WAIBLINGEN                                    = "Schwanen Waiblingen"
	SCHWANFELD_PLUDERHAUSEN                                = "Schwanfeld Plüderhausen"
	SCHWARENBERGSTRASSE_STUTTGART                          = "Schwarenbergstraße Stuttgart"
	SCHWARMERWEG_STUTTGART                                 = "Schwärmerweg Stuttgart"
	SCHWARZWALDSTRASSE_BIETIGHEIM                          = "Schwarzwaldstraße Bietigheim"
	SCHWARZWALDSTRASSE_GARTRINGEN                          = "Schwarzwaldstraße Gärtringen"
	SCHWEISSBRUCKE_ERDMANNHAUSEN                           = "Schweißbrücke Erdmannhausen"
	SCHWERTMUHLE_OBERESSLINGEN                             = "Schwertmühle Oberesslingen"
	SCHWERTSTRASSE_SINDELFINGEN                            = "Schwertstraße Sindelfingen"
	SCHWIEBERDINGEN_SCHWIEBERDINGEN                        = "Schwieberdingen Schwieberdingen"
	SCHWIEBERDINGER_STRASSE_MOGLINGEN                      = "Schwieberdinger Straße Möglingen"
	SCHWIEBERDINGER_STRASSE_HEMMINGEN                      = "Schwieberdinger Straße Hemmingen"
	SCHWILKENHOFSTRASSE_STUTTGART                          = "Schwilkenhofstraße Stuttgart"
	SEDANSTRASSE_SERSHEIM                                  = "Sedanstraße Sersheim"
	SEE_BERKHEIM                                           = "See Berkheim"
	SEE_BISSINGEN_AN_DER_TECK_T                            = "See Bissingen an der Teck (T)"
	SEE_WEIL_IM_SCHONBUCH                                  = "See Weil im Schönbuch"
	SEE_HEIMSHEIM                                          = "See Heimsheim"
	SEE_SEESTR_WEIL_IM_SCHONBUCH                           = "See (Seestr.) Weil im Schönbuch"
	SEEBRUCKENMUHLE_MUSBERG                                = "Seebruckenmühle Musberg"
	SEEBRUCKENMUHLE_KLARWERK_MUSBERG                       = "Seebruckenmühle (Klärwerk) Musberg"
	SEEGARTEN_LEONBERG                                     = "Seegarten Leonberg"
	SEEGER_DOFFINGEN                                       = "Seeger Döffingen"
	SEEHALDE_URBACH                                        = "Seehalde Urbach"
	SEEHALDENWEG_HOFEN                                     = "Seehaldenweg Höfen"
	SEEHOF_BACKNANG                                        = "Seehof Backnang"
	SEEHOFWEG_BACKNANG                                     = "Seehofweg Backnang"
	SEELACH_RUDERSBERG                                     = "Seelach Rudersberg"
	SEEMUHLE_STRUMPFELBACH_WEINSTADT                       = "Seemühle Strümpfelbach (Weinstadt)"
	SEEMUHLE_ROSSWAG                                       = "Seemühle Roßwag"
	SEEPLATZ_KORB                                          = "Seeplatz Korb"
	SEESTRASSE_HASLACH                                     = "Seestraße Häslach"
	SEESTRASSE_WINNENDEN                                   = "Seestraße Winnenden"
	SEESTRASSE_ROMMELSHAUSEN                               = "Seestraße Rommelshausen"
	SEESTRASSE_SCHONAICH                                   = "Seestraße Schönaich"
	SEESTRASSE_HOLZMADEN                                   = "Seestraße Holzmaden"
	SEESTRASSE_KLEINBOTTWAR                                = "Seestraße Kleinbottwar"
	SEESTRASSE_UHINGEN                                     = "Seestraße Uhingen"
	SEGELFALTERSTRASSE_STUTTGART                           = "Segelfalterstraße Stuttgart"
	SEHNINGEN_BAD_BOLL                                     = "Sehningen Bad Boll"
	SEIBOLDSWEILER_WELZHEIM                                = "Seiboldsweiler Welzheim"
	SEMINAR_BACKNANG                                       = "Seminar Backnang"
	SENEFELDERSTRASSE_STUTTGART                            = "Senefelderstraße Stuttgart"
	SENIORENZENTRUM_WAIBLINGEN                             = "Seniorenzentrum Waiblingen"
	SENIORENZENTRUM_HERRENBERG                             = "Seniorenzentrum Herrenberg"
	SERACHER_STRASSE_SERACH                                = "Seracher Straße Serach"
	SERSHEIM_SERSHEIM                                      = "Sersheim Sersheim"
	SEYFFERSTRASSE_STUTTGART                               = "Seyfferstraße Stuttgart"
	SIEBERSBACH_SULZBACH_M                                 = "Siebersbach Sulzbach (M)"
	SIEDLUNG_GERLINGEN                                     = "Siedlung Gerlingen"
	SIEDLUNG_AICH                                          = "Siedlung Aich"
	SIEDLUNG_WOLFSCHLUGEN                                  = "Siedlung Wolfschlugen"
	SIEDLUNG_SCHANBACH                                     = "Siedlung Schanbach"
	SIEDLUNG_OHMDEN                                        = "Siedlung Ohmden"
	SIEDLUNG_HAUSEN_BAD_UBERKINGEN                         = "Siedlung Hausen (Bad Überkingen)"
	SIEDLUNG_SCHNITTLINGEN                                 = "Siedlung Schnittlingen"
	SIEDLUNG_OTTENBACH                                     = "Siedlung Ottenbach"
	SIEDLUNG_BUNZWANGEN                                    = "Siedlung Bünzwangen"
	SIEDLUNG_GUSSENSTADT                                   = "Siedlung Gussenstadt"
	SIEGELSBERG_MURRHARDT                                  = "Siegelsberg Murrhardt"
	SIEGENBERGPLATZ_REICHENBACH_F                          = "Siegenbergplatz Reichenbach (F)"
	SIEGLESTRASSE_STUTTGART                                = "Sieglestraße Stuttgart"
	SIELMINGER_STRASSE_STETTEN_LEINFELDENECHT              = "Sielminger Straße Stetten (Leinfelden-Echt.)"
	SIEMENSSTRASSE_SCHWIEBERDINGEN                         = "Siemensstraße Schwieberdingen"
	SIEMENSSTRASSE_HOLZGERLINGEN                           = "Siemensstraße Holzgerlingen"
	SIEMENSSTRASSE_DECKENPFRONN                            = "Siemensstraße Deckenpfronn"
	SIEMENSSTRASSE_WANGEN_GP                               = "Siemensstraße Wangen (GP)"
	SIEMENSSTRASSE_SCHLIERBACH                             = "Siemensstraße Schlierbach"
	SIGMARINGER_STRASSE_STUTTGART                          = "Sigmaringer Straße Stuttgart"
	SILBERBURGREINSBURGSTRASSE_STUTTGART                   = "Silberburg-/Reinsburgstraße Stuttgart"
	SILBERSTOLLEN_WUSTENROT                                = "Silberstollen Wüstenrot"
	SILBERWALD_STUTTGART                                   = "Silberwald Stuttgart"
	SILBERWEG_BOBLINGEN                                    = "Silberweg Böblingen"
	SILCHERSCHULE_EISLINGEN_F                              = "Silcherschule Eislingen (F)"
	SILCHERSTRASSE_SCHOCKINGEN                             = "Silcherstraße Schöckingen"
	SILCHERSTRASSE_HERRENBERG                              = "Silcherstraße Herrenberg"
	SILCHERSTRASSE_SCHNAIT                                 = "Silcherstraße Schnait"
	SILCHERSTRASSE_BUNZWANGEN                              = "Silcherstraße Bünzwangen"
	SILCHERWEG_FELLBACH                                    = "Silcherweg Fellbach"
	SILLENBUCH_STUTTGART                                   = "Sillenbuch Stuttgart"
	SILVANERSTRASSE_ROSSWAG                                = "Silvanerstraße Roßwag"
	SINDELFINGEN_SINDELFINGEN                              = "Sindelfingen Sindelfingen"
	SINDELFINGER_STRASSE_BOBLINGEN                         = "Sindelfinger Straße Böblingen"
	SINDELFINGER_STRASSE_MAICHINGEN                        = "Sindelfinger Straße Maichingen"
	SINDLINGEN_AUSSIEDLERHOFE_UNTERJETTINGEN               = "Sindlingen Aussiedlerhöfe Unterjettingen"
	SINDLINGEN_ORTSMITTE_UNTERJETTINGEN                    = "Sindlingen Ortsmitte Unterjettingen"
	SINDLINGER_STRASSE_HASLACH                             = "Sindlinger Straße Haslach"
	SINDLINGER_STRASSE_OBERJETTINGEN                       = "Sindlinger Straße Oberjettingen"
	SINDLINGER_STRASSE_NEBRINGEN                           = "Sindlinger Straße Nebringen"
	SIRNAUER_BRUCKE_OBERESSLINGEN                          = "Sirnauer Brücke Oberesslingen"
	SIRNAUER_HOF_SIRNAU                                    = "Sirnauer Hof Sirnau"
	SKILIFT_TREFFELHAUSEN                                  = "Skilift Treffelhausen"
	SKILIFT_LORCH                                          = "Skilift Lorch"
	SOLITUDE_STUTTGART                                     = "Solitude Stuttgart"
	SOLITUDEALLEE_LUDWIGSBURG                              = "Solitudeallee Ludwigsburg"
	SOLITUDESTRASSE_EISLINGEN_F                            = "Solitudestraße Eislingen (F)"
	SOLO_MAICHINGEN                                        = "Solo Maichingen"
	SOMMERHALDE_WAIBLINGEN                                 = "Sommerhalde Waiblingen"
	SOMMERHOFENSTRASSE_SINDELFINGEN                        = "Sommerhofenstraße Sindelfingen"
	SOMMERRAIN_STUTTGART                                   = "Sommerrain Stuttgart"
	SOMMERRAIN_BACKNANG                                    = "Sommerrain Backnang"
	SONATENWEG_STUTTGART                                   = "Sonatenweg Stuttgart"
	SONNE_HOCHDORF_REMSECK                                 = "Sonne Hochdorf (Remseck)"
	SONNE_TREFFELHAUSEN                                    = "Sonne Treffelhausen"
	SONNE_LORCH                                            = "Sonne Lorch"
	SONNE_CENTER_GEISLINGEN_STEIGE                         = "Sonne Center Geislingen (Steige)"
	SONNENBERG_BONNIGHEIM                                  = "Sonnenberg Bönnigheim"
	SONNENBERG_STUTTGART                                   = "Sonnenberg Stuttgart"
	SONNENBERGSCHULE_AIDLINGEN                             = "Sonnenbergschule Aidlingen"
	SONNENBERGSTRASSE_SINDELFINGEN                         = "Sonnenbergstraße Sindelfingen"
	SONNENBRUCKE_RECHBERGHAUSEN                            = "Sonnenbrücke Rechberghausen"
	SONNENBRUNNEN_MOGLINGEN                                = "Sonnenbrunnen Möglingen"
	SONNENBUHL_STUTTGART                                   = "Sonnenbühl Stuttgart"
	SONNENHALDE_MUSBERG                                    = "Sonnenhalde Musberg"
	SONNENHALDE_SCHARNHAUSEN                               = "Sonnenhalde Scharnhausen"
	SONNENHALDE_HOFEN_BONNIGHEIM                           = "Sonnenhalde Hofen (Bönnigheim)"
	SONNENHOF_KLEINASPACH                                  = "Sonnenhof Kleinaspach"
	SONNENWEG_NOTZINGEN                                    = "Sonnenweg Notzingen"
	SONNENWEG_LORCH                                        = "Sonnenweg Lorch"
	SOS_KINDERDORF_OBERBERKEN                              = "SOS Kinderdorf Oberberken"
	SPAICHINGER_STRASSE_STUTTGART                          = "Spaichinger Straße Stuttgart"
	SPECHTSHOF_REICHENBACH_BERGLEN                         = "Spechtshof Reichenbach (Berglen)"
	SPIELKARTENMUSEUM_LEINFELDEN                           = "Spielkartenmuseum Leinfelden"
	SPIELPLATZ_KIRCHHEIM_T                                 = "Spielplatz Kirchheim (T)"
	SPITTLER_STIFT_SCHORNDORF                              = "Spittler Stift Schorndorf"
	SPITZACKER_NECKARTENZLINGEN                            = "Spitzacker Neckartenzlingen"
	SPITZACKER_PLATTENHARDT                                = "Spitzäcker Plattenhardt"
	SPITZHOLZSTRASSE_SINDELFINGEN                          = "Spitzholzstraße Sindelfingen"
	SPORTGELANDE_AISCHBACH_PEROUSE                         = "Sportgelände Aischbach Perouse"
	SPORTGELANDE_TENNWENGERT_OEFFINGEN                     = "Sportgelände Tennwengert Oeffingen"
	SPORTHALLE_WENDLINGEN_N                                = "Sporthalle Wendlingen (N)"
	SPORTHALLE_BUHL_RUTESHEIM                              = "Sporthalle Bühl Rutesheim"
	SPORTPARK_FEUERBACH_STUTTGART                          = "Sportpark Feuerbach Stuttgart"
	SPORTPARK_WEIL_WEIL                                    = "Sportpark Weil Weil"
	SPORTPLATZ_FLACHT                                      = "Sportplatz Flacht"
	SPORTPLATZ_RIELINGSHAUSEN                              = "Sportplatz Rielingshausen"
	SPORTPLATZ_PFLUGFELDEN                                 = "Sportplatz Pflugfelden"
	SPORTPLATZ_ALTENRIET                                   = "Sportplatz Altenriet"
	SPORTPLATZ_GEMMRIGHEIM                                 = "Sportplatz Gemmrigheim"
	SPORTPLATZ_LOCHGAU                                     = "Sportplatz Löchgau"
	SPORTPLATZ_BARTENBACH_GOPPINGEN                        = "Sportplatz Bartenbach (Göppingen)"
	SPORTPLATZ_WINZINGEN                                   = "Sportplatz Winzingen"
	SPORTPLATZE_GOPPINGEN                                  = "Sportplätze Göppingen"
	SPORTZENTRUM_MUNCHINGEN                                = "Sportzentrum Münchingen"
	SPORTZENTRUM_LEINFELDEN                                = "Sportzentrum Leinfelden"
	SSBZENTRUM_STUTTGART                                   = "SSB-Zentrum Stuttgart"
	ST_GOTTH_WEILERBACHWEG_GOPPINGEN                       = "St. Gotth. Weilerbachweg Göppingen"
	STAATSGALERIE_STUTTGART                                = "Staatsgalerie Stuttgart"
	STADION_SINDELFINGEN                                   = "Stadion Sindelfingen"
	STADION_WALDENBUCH                                     = "Stadion Waldenbuch"
	STADION_KIRCHHEIM_T                                    = "Stadion Kirchheim (T)"
	STADION_SCHONAICH                                      = "Stadion Schönaich"
	STADION_DENKENDORF                                     = "Stadion Denkendorf"
	STADION_KORNWESTHEIM                                   = "Stadion Kornwestheim"
	STADIONHALLE_MARBACH_N                                 = "Stadionhalle Marbach (N)"
	STADIONSTRASSE_ECHTERDINGEN                            = "Stadionstraße Echterdingen"
	STADTBADALTENWOHNANLAGE_DITZINGEN                      = "Stadtbad/Altenwohnanlage Ditzingen"
	STADTBAHNHOF_VAIHINGEN_E                               = "Stadtbahnhof Vaihingen (E)"
	STADTBIBLIOTHEK_STUTTGART                              = "Stadtbibliothek Stuttgart"
	STADTBIBLIOTHEK_EBERSBACH_F                            = "Stadtbibliothek Ebersbach (F)"
	STADTBUCHEREI_LEINFELDEN                               = "Stadtbücherei Leinfelden"
	STADTFRIEDHOF_HERRENBERG                               = "Stadtfriedhof Herrenberg"
	STADTGARTEN_WEIL_DER_STADT                             = "Stadtgarten Weil der Stadt"
	STADTGRENZE_STUTTGART                                  = "Stadtgrenze Stuttgart"
	STADTHALLE_SINDELFINGEN                                = "Stadthalle Sindelfingen"
	STADTHALLE_KORNTAL                                     = "Stadthalle Korntal"
	STADTHALLE_KIRCHHEIM_T                                 = "Stadthalle Kirchheim (T)"
	STADTHALLE_MURRHARDT                                   = "Stadthalle Murrhardt"
	STADTHALLE_VAIHINGEN_E                                 = "Stadthalle Vaihingen (E)"
	STADTHALLE_DONZDORF                                    = "Stadthalle Donzdorf"
	STADTHALLE_GOPPINGEN                                   = "Stadthalle Göppingen"
	STADTKIRCHE_GEISLINGEN_STEIGE                          = "Stadtkirche Geislingen (Steige)"
	STADTMITTE_WENDLINGEN_N                                = "Stadtmitte Wendlingen (N)"
	STADTMITTE_WAIBLINGEN                                  = "Stadtmitte Waiblingen"
	STADTMITTE_STUTTGART                                   = "Stadtmitte Stuttgart"
	STADTMITTE_ASPERG                                      = "Stadtmitte Asperg"
	STADTMITTE_EISLINGEN_F                                 = "Stadtmitte Eislingen (F)"
	STADTMITTE_EBERSBACH_F                                 = "Stadtmitte Ebersbach (F)"
	STADTMUHLE_BACKNANG                                    = "Stadtmühle Backnang"
	STADTPLATZ_WERNAU_N                                    = "Stadtplatz Wernau (N)"
	STADTWERKE_WAIBLINGEN                                  = "Stadtwerke Waiblingen"
	STADTWERKE_LUDWIGSBURG                                 = "Stadtwerke Ludwigsburg"
	STADTWERKE_HERRENBERG                                  = "Stadtwerke Herrenberg"
	STADTWERKE_GOPPINGEN                                   = "Stadtwerke Göppingen"
	STADTZENTRUM_FREIBERG_N                                = "Stadtzentrum Freiberg (N)"
	STAFFEL_HOHENHASLACH                                   = "Staffel Hohenhaslach"
	STAFFLENBERGSTRASSE_STUTTGART                          = "Stafflenbergstraße Stuttgart"
	STAHLACKERWEG_NECKARHALDE                              = "Stahlackerweg Neckarhalde"
	STAIGACKER_BACKNANG                                    = "Staigacker Backnang"
	STAMMHEIM_STUTTGART                                    = "Stammheim Stuttgart"
	STAMMHEIM_FRIEDHOF_STUTTGART                           = "Stammheim Friedhof Stuttgart"
	STAMMHEIM_SPORTHALLE_STUTTGART                         = "Stammheim Sporthalle Stuttgart"
	STANGEN_ECHTERDINGEN                                   = "Stangen Echterdingen"
	STANGENBACH_WUSTENROT                                  = "Stangenbach Wüstenrot"
	STANGENBACH_HAUS_WALDESRUH_WUSTENROT                   = "Stangenbach Haus Waldesruh Wüstenrot"
	STARENWEG_FREIBERG_N                                   = "Starenweg Freiberg (N)"
	STATTMANNSTRASSE_OBERBOIHINGEN                         = "Stattmannstraße Oberboihingen"
	STAUCHLE_HOLZGERLINGEN                                 = "Stäuchle Holzgerlingen"
	STAUFENECKSCHULE_SALACH                                = "Staufeneckschule Salach"
	STAUFENSTRASSE_ERKENBRECHTSWEILER                      = "Staufenstraße Erkenbrechtsweiler"
	STAUFENSTRASSE_ALBERSHAUSEN                            = "Staufenstraße Albershausen"
	STAUFERPARK_GOPPINGEN                                  = "Stauferpark Göppingen"
	STAUFERSCHULE_WAIBLINGEN                               = "Stauferschule Waiblingen"
	STAUFERSCHULE_WASCHENBEUREN                            = "Stauferschule Wäschenbeuren"
	STAUFERSCHULE_LORCH                                    = "Stauferschule Lorch"
	STAUFERSTRASSE_ALTBACH                                 = "Stauferstraße Altbach"
	STAUFERSTRASSE_WAIBLINGEN                              = "Stauferstraße Waiblingen"
	STAUFERSTRASSE_SCHMIDEN                                = "Stauferstraße Schmiden"
	STAUFERWEG_MAITIS                                      = "Stauferweg Maitis"
	STAUFFENBERGSTRASSE_OBERESSLINGEN                      = "Stauffenbergstraße Oberesslingen"
	STAUFFENBERGSTRASSE_KORNWESTHEIM                       = "Stauffenbergstraße Kornwestheim"
	STECKFELD_STUTTGART                                    = "Steckfeld Stuttgart"
	STEIGE_ERDMANNHAUSEN                                   = "Steige Erdmannhausen"
	STEIGE_HOHENHASLACH                                    = "Steige Hohenhaslach"
	STEIGE_NECKARREMS                                      = "Steige Neckarrems"
	STEIGHOF_ABZWEIG_TREFFELHAUSEN_LAUTERSTEIN             = "Steighof Abzweig Treffelhausen Lauterstein"
	STEIGSTRASSE_FELLBACH                                  = "Steigstraße Fellbach"
	STEINACHBRUCKE_NURTINGEN                               = "Steinachbrücke Nürtingen"
	STEINACKER_STUTTGART                                   = "Steinäcker Stuttgart"
	STEINACKERSTRASSE_REICHENBACH_F                        = "Steinäckerstraße Reichenbach (F)"
	STEINBACHBRUCKE_LOCHGAU                                = "Steinbachbrücke Löchgau"
	STEINBASS_SCHONAICH                                    = "Steinbass Schönaich"
	STEINBEISE_WELZHEIM                                    = "Steinbeise Welzheim"
	STEINBEISSTRASSE_FELLBACH                              = "Steinbeisstraße Fellbach"
	STEINBEISSTRASSE_ZELL_ESSLINGEN                        = "Steinbeisstraße Zell (Esslingen)"
	STEINBEISSTRASSE_BOBLINGEN                             = "Steinbeisstraße Böblingen"
	STEINBEISSTRASSE_LUDWIGSBURG                           = "Steinbeisstraße Ludwigsburg"
	STEINBEISSTRASSE_VAIHINGEN_E                           = "Steinbeisstraße Vaihingen (E)"
	STEINBEISSTRASSE_MARKGRONINGEN                         = "Steinbeisstraße Markgröningen"
	STEINBEISSTRASSE_SCHORNDORF                            = "Steinbeisstraße Schorndorf"
	STEINBEISSTRASSE_EISLINGEN_F                           = "Steinbeisstraße Eislingen (F)"
	STEINBEISWEG_NURTINGEN                                 = "Steinbeisweg Nürtingen"
	STEINBERG_MURRHARDT                                    = "Steinberg Murrhardt"
	STEINBERG_FRITZHOF_MURRHARDT                           = "Steinberg Fritzhof Murrhardt"
	STEINBERG_HALDE_MURRHARDT                              = "Steinberg Halde Murrhardt"
	STEINBRUCH_SCHOCKINGEN                                 = "Steinbruch Schöckingen"
	STEINBRUCK_WELZHEIM                                    = "Steinbruck Welzheim"
	STEINBUTTSTRASSE_STUTTGART                             = "Steinbuttstraße Stuttgart"
	STEINEN_NELLINGEN                                      = "Steinen Nellingen"
	STEINENKIRCH_STEINENKIRCH                              = "Steinenkirch Steinenkirch"
	STEINGARTENSCHULE_DONZDORF                             = "Steingartenschule Donzdorf"
	STEINGAUSTRASSE_KIRCHHEIM_T                            = "Steingaustraße Kirchheim (T)"
	STEINGRUBE_WAIBLINGEN                                  = "Steingrube Waiblingen"
	STEINHALDE_OBERESSLINGEN                               = "Steinhalde Oberesslingen"
	STEINHALDENFELD_STUTTGART                              = "Steinhaldenfeld Stuttgart"
	STEINHEIMER_STRASSE_POPPENWEILER                       = "Steinheimer Straße Poppenweiler"
	STEINHEIMER_STRASSE_BIETIGHEIM                         = "Steinheimer Straße Bietigheim"
	STEINIGE_HALDE_SPARWIESEN                              = "Steinige Halde Sparwiesen"
	STEINREINACH_KORB                                      = "Steinreinach Korb"
	STEINSTRASSE_HOLZHEIM_GOPPINGEN                        = "Steinstraße Holzheim (Göppingen)"
	STEINWERK_VAIHINGEN_E                                  = "Steinwerk Vaihingen (E)"
	STELLE_STUTTGART                                       = "Stelle Stuttgart"
	STERN_AICH                                             = "Stern Aich"
	STERN_UNTERENSINGEN                                    = "Stern Unterensingen"
	STERNBERGSTRASSE_FRICKENHAUSEN                         = "Sternbergstraße Frickenhausen"
	STERNHAULE_STUTTGART                                   = "Sternhäule Stuttgart"
	STERNKREUZUNG_GOPPINGEN                                = "Sternkreuzung Göppingen"
	STERNPLATZ_ZOH_GEISLINGEN_STEIGE                       = "Sternplatz ZOH Geislingen (Steige)"
	STETTENBEINSTEIN_ENDERSBACH                            = "Stetten-Beinstein Endersbach"
	STETTENER_STRASSE_ENDERSBACH                           = "Stettener Straße Endersbach"
	STETTINER_RING_BACKNANG                                = "Stettiner Ring Backnang"
	STETTINER_STRASSE_NURTINGEN                            = "Stettiner Straße Nürtingen"
	STETTINER_STRASSE_BOBLINGEN                            = "Stettiner Straße Böblingen"
	STETTINER_STRASSE_SCHWIEBERDINGEN                      = "Stettiner Straße Schwieberdingen"
	STIEGELPLATZ_MUNCHINGEN                                = "Stiegelplatz Münchingen"
	STIEGELWIESE_SUSSEN                                    = "Stiegelwiese Süßen"
	STIFTSGYMNASIUM_SINDELFINGEN                           = "Stiftsgymnasium Sindelfingen"
	STOCKACH_STUTTGART                                     = "Stöckach Stuttgart"
	STOCKACHSCHULE_WINNENDEN                               = "Stöckachschule Winnenden"
	STOCKENHOF_OSCHELBRONN_BERGLEN                         = "Stöckenhof Öschelbronn (Berglen)"
	STORCHEN_GOPPINGEN                                     = "Storchen Göppingen"
	STRASSENACKER_EGLOSHEIM                                = "Straßenäcker Eglosheim"
	STRASSENMEISTEREI_KIRCHHEIM_T                          = "Straßenmeisterei Kirchheim (T)"
	STRASSENMEISTEREI_HERRENBERG                           = "Straßenmeisterei Herrenberg"
	STRASSLE_MARKGRONINGEN                                 = "Sträßle Markgröningen"
	STRAUSSSTAFFEL_STUTTGART                               = "Straußstaffel Stuttgart"
	STREICH_VORDERWEISSBUCH                                = "Streich Vorderweißbuch"
	STRESEMANNSTRASSE_BACKNANG                             = "Stresemannstraße Backnang"
	STROHGAUSTRASSE_LEONBERG                               = "Strohgäustraße Leonberg"
	STROHHOF_ALTHUTTE                                      = "Strohhof Althütte"
	STROMBERGGYMNASIUM_VAIHINGEN_E                         = "Stromberggymnasium Vaihingen (E)"
	STRUBELMUHLE_ALFDORF                                   = "Strübelmühle Alfdorf"
	STRUDELBACHHALLE_WEISSACH_BB                           = "Strudelbachhalle Weissach (BB)"
	STRUMPFELBACH_STRUMPFELBACH_BACKNANG                   = "Strümpfelbach Strümpfelbach (Backnang)"
	STRUTSTRASSE_SCHLICHTEN                                = "Strutstraße Schlichten"
	STUIFENSTRASSE_PLOCHINGEN                              = "Stuifenstraße Plochingen"
	STUIFENSTRASSE_MAITIS                                  = "Stuifenstraße Maitis"
	STUMPENHOF_PLOCHINGEN                                  = "Stumpenhof Plochingen"
	STUTTGARTER_PLATZ_FELLBACH                             = "Stuttgarter Platz Fellbach"
	STUTTGARTER_STRASSE_BOBLINGEN                          = "Stuttgarter Straße Böblingen"
	STUTTGARTER_STRASSE_FREIBERG_N                         = "Stuttgarter Straße Freiberg (N)"
	STUTTGARTER_STRASSE_PLIENSAUVORSTADT                   = "Stuttgarter Straße Pliensauvorstadt"
	STUTTGARTER_STRASSE_ALTDORF_ES                         = "Stuttgarter Straße Altdorf (ES)"
	STUTTGARTER_STRASSE_OSCHELBRONN_GAUFELDEN              = "Stuttgarter Straße Öschelbronn (Gäufelden)"
	STUTTGARTER_STRASSE_WEILER_ZUM_STEIN                   = "Stuttgarter Straße Weiler zum Stein"
	STUTTGARTER_STRASSE_VAIHINGEN_E                        = "Stuttgarter Straße Vaihingen (E)"
	STUTTGARTER_STRASSE_SCHORNDORF                         = "Stuttgarter Straße Schorndorf"
	STUTTGARTER_STRASSE_REICHENBACH_F                      = "Stuttgarter Straße Reichenbach (F)"
	STUTTGARTER_STRASSE_ASPERG                             = "Stuttgarter Straße Asperg"
	STUTTGARTER_STRASSEEICHHOLZ_MAICHINGEN                 = "Stuttgarter Straße/Eichholz Maichingen"
	STUTTGARTER_WEG_MARKGRONINGEN                          = "Stuttgarter Weg Markgröningen"
	STUTTGARTERMARIENSTR_LORCH                             = "Stuttgarter-/Marienstr. Lorch"
	STUTTGARTERWESTSTRASSE_GRUNBACH                        = "Stuttgarter-/Weststraße Grunbach"
	SUD_HASLACH                                            = "Süd Häslach"
	SUD_RENNINGEN                                          = "Süd Renningen"
	SUD_KIRCHHEIM_T                                        = "Süd Kirchheim (T)"
	SUDBAHNHOF_BOBLINGEN                                   = "Südbahnhof Böblingen"
	SUDETENSTRASSE_JEBENHAUSEN                             = "Sudetenstraße Jebenhausen"
	SUDETENSTRASSE_DONZDORF                                = "Sudetenstraße Donzdorf"
	SUDHEIMER_PLATZ_STUTTGART                              = "Südheimer Platz Stuttgart"
	SUDSTRASSE_BISSINGEN_LB                                = "Südstraße Bissingen (LB)"
	SULZBACH_M_SULZBACH_M                                  = "Sulzbach (M) Sulzbach (M)"
	SULZBACHER_STRASSE_SPIEGELBERG                         = "Sulzbacher Straße Spiegelberg"
	SULZBACHTAL_SCHONAICH                                  = "Sulzbachtal Schönaich"
	SULZGRIESER_STEIGE_ESSLINGEN_N                         = "Sulzgrieser Steige Esslingen (N)"
	SUSSEN_SUSSEN                                          = "Süßen Süßen"
	SUTTNERSTRASSE_STUTTGART                               = "Suttnerstraße Stuttgart"
	SUWAG_PLEIDELSHEIM                                     = "SÜWAG Pleidelsheim"
	SYNERGIEPARK_STUTTGART                                 = "Synergiepark Stuttgart"
	TACHENBERGSTRASSE_KORNTAL                              = "Tachenbergstraße Korntal"
	TACHENHAUSERSTRASSE_OBERBOIHINGEN                      = "Tachenhäuserstraße Oberboihingen"
	TAFERHALDE_UNTERWEISSACH                               = "Täferhalde Unterweissach"
	TAGFALTERWEG_STUTTGART                                 = "Tagfalterweg Stuttgart"
	TAILFINGER_STRASSE_OSCHELBRONN_GAUFELDEN               = "Tailfinger Straße Öschelbronn (Gäufelden)"
	TALLANDHAUSSTRASSE_STUTTGART                           = "Tal-/Landhausstraße Stuttgart"
	TALOSTENDSTRASSE_STUTTGART                             = "Tal-/Ostendstraße Stuttgart"
	TALACKERSTRASSE_WARMBRONN                              = "Talackerstraße Warmbronn"
	TALACKERSTRASSE_WENDLINGEN_N                           = "Taläckerstraße Wendlingen (N)"
	TALAUE_BIRKMANNSWEILER                                 = "Talaue Birkmannsweiler"
	TALBACHBRUCKE_HOCHDORF_ES                              = "Talbachbrücke Hochdorf (ES)"
	TALDORFER_STRASSE_STUTTGART                            = "Taldorfer Straße Stuttgart"
	TALE_KAISERSBACH                                       = "Täle Kaisersbach"
	TALHOFE_ENZWEIHINGEN                                   = "Talhöfe Enzweihingen"
	TALSTRASSE_HOLZGERLINGEN                               = "Talstraße Holzgerlingen"
	TALSTRASSE_EHNINGEN                                    = "Talstraße Ehningen"
	TALSTRASSE_HOF_UND_LEMBACH                             = "Talstraße Hof und Lembach"
	TALSTRASSE_HOPFIGHEIM                                  = "Talstraße Höpfigheim"
	TALSTRASSE_WAIBLINGEN                                  = "Talstraße Waiblingen"
	TALSTRASSE_ST_BERNHARDT                                = "Talstraße St. Bernhardt"
	TALSTRASSE_RAIDWANGEN                                  = "Talstraße Raidwangen"
	TALSTRASSE_BOBLINGEN                                   = "Talstraße Böblingen"
	TALSTRASSE_LUDWIGSBURG                                 = "Talstraße Ludwigsburg"
	TALSTRASSE_SCHORNDORF                                  = "Talstraße Schorndorf"
	TALSTRASSE_KORNWESTHEIM                                = "Talstraße Kornwestheim"
	TALSTRASSE_EISLINGEN_F                                 = "Talstraße Eislingen (F)"
	TALSTRASSE_LORCH                                       = "Talstraße Lorch"
	TAMM_TAMM                                              = "Tamm Tamm"
	TAMMER_STRASSE_BISSINGEN_LB                            = "Tammer Straße Bissingen (LB)"
	TAMMER_STRASSE_EGLOSHEIM                               = "Tammer Straße Eglosheim"
	TANNBACHBRUCKE_MIEDELSBACH                             = "Tannbachbrücke Miedelsbach"
	TANNENBERG_BOBLINGEN                                   = "Tannenberg Böblingen"
	TANNENBERGSTRASSE_KIRCHHEIM_T                          = "Tannenbergstraße Kirchheim (T)"
	TANNENHOFE_SCHLIERBACH                                 = "Tannenhöfe Schlierbach"
	TANNHOF_PFAHLBRONN                                     = "Tannhof Pfahlbronn"
	TAPACHSTRASSE_STUTTGART                                = "Tapachstraße Stuttgart"
	TAUBENHOF_WELZHEIM                                     = "Taubenhof Welzheim"
	TAUNUSSTRASSE_BOBLINGEN                                = "Taunusstraße Böblingen"
	TAUSGYMNASIUM_BACKNANG                                 = "Tausgymnasium Backnang"
	TAUSSCHULE_BACKNANG                                    = "Tausschule Backnang"
	TECHNISCHE_AKADEMIE_NELLINGEN                          = "Technische Akademie Nellingen"
	TECHNOLOGIEPARK_BOBLINGEN                              = "Technologiepark Böblingen"
	TECHNOLOGIEPARK_MARBACH_N                              = "Technologiepark Marbach (N)"
	TECKHALLE_OWEN                                         = "Teckhalle Owen"
	TECKSTRASSE_KIRCHHEIM_T                                = "Teckstraße Kirchheim (T)"
	TEINACHER_STRASSE_EGLOSHEIM                            = "Teinacher Straße Eglosheim"
	TENNISPLATZ_STUTTGART                                  = "Tennisplatz Stuttgart"
	TENNISPLATZ_LUDWIGSBURG                                = "Tennisplatz Ludwigsburg"
	TENNISPLATZE_MOGLINGEN                                 = "Tennisplätze Möglingen"
	TERLANERSTRASSE_STUTTGART                              = "Terlanerstraße Stuttgart"
	TESSINER_STRASSE_DARMSHEIM                             = "Tessiner Straße Darmsheim"
	TEUCHELWEG_MAICHINGEN                                  = "Teuchelweg Maichingen"
	TEXTILZENTRUM_SINDELFINGEN                             = "Textilzentrum Sindelfingen"
	THHEUSSSTRASSE_HOCHDORF_EBERDINGEN                     = "Th.-Heuss-Straße Hochdorf (Eberdingen)"
	THSTORMSTRSIEDLUNG_STRASSDORF                          = "Th.-Storm-Str./Siedlung Straßdorf"
	THALESPLATZ_DITZINGEN                                  = "Thalesplatz Ditzingen"
	THEATER_IM_BAHNHOF_RECHBERGHAUSEN                      = "Theater im Bahnhof Rechberghausen"
	THEOLORCHWERKSTATTEN_BIETIGHEIM                        = "Theo-Lorch-Werkstätten Bietigheim"
	THEOLORCHWERKSTATTEN_LUDWIGSBURG                       = "Theo-Lorch-Werkstätten Ludwigsburg"
	THEODORHEUSSGYMNASIUM_OBERESSLINGEN                    = "Theodor-Heuss-Gymnasium Oberesslingen"
	THEODORHEUSSPLATZ_SCHELMENHOLZ                         = "Theodor-Heuss-Platz Schelmenholz"
	THEODORHEUSSSTRASSE_HILDRIZHAUSEN                      = "Theodor-Heuss-Straße Hildrizhausen"
	THEODORHEUSSSTRASSE_OBERRIEXINGEN                      = "Theodor-Heuss-Straße Oberriexingen"
	THERMALBAD_HEXENBUCKEL_BOBLINGEN                       = "Thermalbad (Hexenbuckel) Böblingen"
	THERMALBAD_STUTTGARTER_STR_BOBLINGEN                   = "Thermalbad (Stuttgarter Str.) Böblingen"
	THINGSTRASSE_STUTTGART                                 = "Thingstraße Stuttgart"
	THOMASHARDTER_STRASSE_HEGENLOHE                        = "Thomashardter Straße Hegenlohe"
	TIEFENBACHSCHULE_STUTTGART                             = "Tiefenbachschule Stuttgart"
	TIEFENMAD_KIRCHENKIRNBERG                              = "Tiefenmad Kirchenkirnberg"
	TIEFGARAGE_HOHENGEHREN                                 = "Tiefgarage Hohengehren"
	TIGERSTRASSE_STUTTGART                                 = "Tigerstraße Stuttgart"
	TIROLER_STRASSE_ELTINGEN                               = "Tiroler Straße Eltingen"
	TISCHARDTER_STRASSE_NURTINGEN                          = "Tischardter Straße Nürtingen"
	TISCHARDTER_STRASSE_KOHLBERG                           = "Tischardter Straße Kohlberg"
	TOBELKREUZUNG_SUSSEN                                   = "Tobelkreuzung Süßen"
	TOBELWASEN_WEILHEIM_T                                  = "Tobelwasen Weilheim (T)"
	TORFGRUBE_SCHOPFLOCH_LENNINGEN                         = "Torfgrube Schopfloch (Lenningen)"
	TORGAUER_WEG_HERRENBERG                                = "Torgauer Weg Herrenberg"
	TORLENSWEG_LEONBERG                                    = "Törlensweg Leonberg"
	TORSTRASSE_NEUSTADT                                    = "Torstraße Neustadt"
	TORSTRASSE_BAHNHOF_DETTENHAUSEN                        = "Torstraße (Bahnhof) Dettenhausen"
	TRANKE_STUTTGART                                       = "Tränke Stuttgart"
	TRAUBE_NEUHAUSEN_F                                     = "Traube Neuhausen (F)"
	TRAUBE_STRUMPFELBACH_WEINSTADT                         = "Traube Strümpfelbach (Weinstadt)"
	TRAUBE_LIPPOLDSWEILER                                  = "Traube Lippoldsweiler"
	TRAUBENSTRASSE_NECKARREMS                              = "Traubenstraße Neckarrems"
	TRAUZENBACH_GRAB                                       = "Trauzenbach Grab"
	TREIBERSTRASSE_STUTTGART                               = "Treiberstraße Stuttgart"
	TRIPSDRILL_ERLEBNISPARK_CLEEBRONN                      = "Tripsdrill Erlebnispark Cleebronn"
	TROPPEL_WEIL_IM_SCHONBUCH                              = "Troppel Weil im Schönbuch"
	TUBINGER_ALLEE_SINDELFINGEN                            = "Tübinger Allee Sindelfingen"
	TUBINGER_STRASSE_ECHTERDINGEN                          = "Tübinger Straße Echterdingen"
	TUBINGER_STRASSE_BOBLINGEN                             = "Tübinger Straße Böblingen"
	TUBINGER_STRASSE_DECKENPFRONN                          = "Tübinger Straße Deckenpfronn"
	TUBINGER_STRASSE_NECKARTAILFINGEN                      = "Tübinger Straße Neckartailfingen"
	TUBINGER_STRASSE_DETTENHAUSEN                          = "Tübinger Straße Dettenhausen"
	TULPENSTRASSE_SINDELFINGEN                             = "Tulpenstraße Sindelfingen"
	TUNNEL_OSTPORTAL_STUTTGART                             = "Tunnel Ostportal Stuttgart"
	TUNNELSTRASSE_STUTTGART                                = "Tunnelstraße Stuttgart"
	TURKHEIM_TURKHEIM                                      = "Türkheim Türkheim"
	TURMUHRENMUSEUM_MAINHARDT                              = "Turmuhrenmuseum Mainhardt"
	TURNACKERSTRASSE_BERNHAUSEN                            = "Turnackerstraße Bernhausen"
	TURNERSCHAFT_STADION_GOPPINGEN                         = "Turnerschaft Stadion Göppingen"
	TURNHALLE_HIRSCHLANDEN                                 = "Turnhalle Hirschlanden"
	TURNHALLE_ERDMANNHAUSEN                                = "Turnhalle Erdmannhausen"
	TURNHALLE_STEINENBERG                                  = "Turnhalle Steinenberg"
	TURNHALLE_REICHENBACH_IM_TALE_DEGGINGEN                = "Turnhalle Reichenbach im Täle (Deggingen)"
	TURNHALLE_UHINGEN                                      = "Turnhalle Uhingen"
	TUV_BERNHAUSEN                                         = "TÜV Bernhausen"
	TUXER_WEG_MAUBACH                                      = "Tuxer Weg Maubach"
	UDITORIUM_UHINGEN                                      = "Uditorium Uhingen"
	UFERSTRASSE_HOHENECK                                   = "Uferstraße Hoheneck"
	UFFKIRCHHOF_STUTTGART                                  = "Uff-Kirchhof Stuttgart"
	UHINGEN_UHINGEN                                        = "Uhingen Uhingen"
	UHLANDGYMNASIUM_KIRCHHEIM_T                            = "Uhlandgymnasium Kirchheim (T)"
	UHLANDSTRASSE_LUDWIGSBURG                              = "Uhlandstraße Ludwigsburg"
	UHLANDSTRASSE_HEININGEN_GP                             = "Uhlandstraße Heiningen (GP)"
	UHLANDSTRASSE_FAURNDAU                                 = "Uhlandstraße Faurndau"
	UHLANDSTRASSE_OBERWALDEN                               = "Uhlandstraße Oberwälden"
	UHLBACH_STUTTGART                                      = "Uhlbach Stuttgart"
	UHLBERG_PLATTENHARDT                                   = "Uhlberg Plattenhardt"
	ULMENSTRASSE_STUTTGART                                 = "Ulmenstraße Stuttgart"
	ULMENSTRASSE_MAICHINGEN                                = "Ulmenstraße Maichingen"
	ULMER_STRASSE_RAMTEL                                   = "Ulmer Straße Ramtel"
	ULMER_STRASSE_PLOCHINGEN                               = "Ulmer Straße Plochingen"
	ULMER_STRASSE_REICHENBACH_F                            = "Ulmer Straße Reichenbach (F)"
	ULMERHEININGER_STR_GOPPINGEN                           = "Ulmer-/Heininger Str. Göppingen"
	ULRICHSTRASSE_BESIGHEIM                                = "Ulrichstraße Besigheim"
	ULRICHSTRASSE_BEUTELSBACH                              = "Ulrichstraße Beutelsbach"
	ULRICHWEG_DOFFINGEN                                    = "Ulrichweg Döffingen"
	UMGELTERWEG_STUTTGART                                  = "Umgelterweg Stuttgart"
	UMSPANNWERK_SUSSEN                                     = "Umspannwerk Süßen"
	UNGEHEUERHOF_BACKNANG                                  = "Ungeheuerhof Backnang"
	UNGEHEUERHOF_ABZWEIG_BACKNANG                          = "Ungeheuerhof Abzweig Backnang"
	UNIVERSITAT_STUTTGART                                  = "Universität Stuttgart"
	UNIVERSITAT_SCHLEIFE_STUTTGART                         = "Universität (Schleife) Stuttgart"
	UNIVERSITAT_HOHENHEIM_STUTTGART                        = "Universität Hohenheim Stuttgart"
	UNTER_DEN_ARKADEN_ROMMELSHAUSEN                        = "Unter den Arkaden Rommelshausen"
	UNTERAICHEN_LEINFELDEN                                 = "Unteraichen Leinfelden"
	UNTERBERKEN_HERRENBACHSTRASSE_OBERBERKEN               = "Unterberken Herrenbachstraße Oberberken"
	UNTERBOHRINGEN_UNTERBOHRINGEN                          = "Unterböhringen Unterböhringen"
	UNTERDORF_WEILER_OB_HELFENSTN                          = "Unterdorf Weiler ob Helfenstn."
	UNTERDORFSTRASSE_KONGEN                                = "Unterdorfstraße Köngen"
	UNTERDRACKENSTEIN_KIRCHE_DRACKENSTEIN                  = "Unterdrackenstein Kirche Drackenstein"
	UNTERE_BRUCK_BORTLINGEN                                = "Untere Bruck Börtlingen"
	UNTERE_BUHLSTRASSE_ALFDORF                             = "Untere Bühlstraße Alfdorf"
	UNTERE_HALDE_STETTEN_LEINFELDENECHT                    = "Untere Halde Stetten (Leinfelden-Echt.)"
	UNTERE_HALDE_WEIL_IM_SCHONBUCH                         = "Untere Halde Weil im Schönbuch"
	UNTERE_MARBACHER_STRASSE_LUDWIGSBURG                   = "Untere Marbacher Straße Ludwigsburg"
	UNTERE_MAYENNER_STRASSE_WAIBLINGEN                     = "Untere Mayenner Straße Waiblingen"
	UNTERE_SCHLOSSSTRASSE_ALFDORF                          = "Untere Schlossstraße Alfdorf"
	UNTERE_STRASSE_ERKENBRECHTSWEILER                      = "Untere Straße Erkenbrechtsweiler"
	UNTERENSINGER_STRASSE_ZIZISHAUSEN                      = "Unterensinger Straße Zizishausen"
	UNTERLENNINGEN_UNTERLENNINGEN                          = "Unterlenningen Unterlenningen"
	UNTERM_GREUT_SCHOPFLOCH_LENNINGEN                      = "Unterm Greut Schopfloch (Lenningen)"
	UNTERNEUSTETTEN_KIRCHENKIRNBERG                        = "Unterneustetten Kirchenkirnberg"
	UNTERRIEDEN_GYMNASIUM_MAICHINGEN                       = "Unterrieden Gymnasium Maichingen"
	UNTERSCHONTAL_BACKNANG                                 = "Unterschöntal Backnang"
	UNTERTURKHEIM_STUTTGART                                = "Untertürkheim Stuttgart"
	UNTERTURKHEIM_ALTER_FRIEDHOF_STUTTGART                 = "Untertürkheim Alter Friedhof Stuttgart"
	UNTERTURKHEIM_FRIEDHOF_STUTTGART                       = "Untertürkheim Friedhof Stuttgart"
	UNTERTURKHEIM_KELTERPLATZ_STUTTGART                    = "Untertürkheim Kelterplatz Stuttgart"
	UNTERTURKHEIMER_STRASSE_FELLBACH                       = "Untertürkheimer Straße Fellbach"
	URACHER_WEG_NEUFFEN                                    = "Uracher Weg Neuffen"
	URACHSTRASSE_STUTTGART                                 = "Urachstraße Stuttgart"
	URBACH_URBACH                                          = "Urbach Urbach"
	URBAN_HARBOR_LUDWIGSBURG                               = "Urban Harbor Ludwigsburg"
	URBANSTRASSE_BEUTELSBACH                               = "Urbanstraße Beutelsbach"
	URSENWANG_ESCHENSTR_GOPPINGEN                          = "Ursenwang Eschenstr. Göppingen"
	URSENWANG_GRUNDSCHULE_GOPPINGEN                        = "Ursenwang Grundschule Göppingen"
	URSENWANG_KIEFERNSTEIGE_GOPPINGEN                      = "Ursenwang Kiefernsteige Göppingen"
	URSENWANG_LINDENPLATZ_GOPPINGEN                        = "Ursenwang Lindenplatz Göppingen"
	VAIHINGEN_STUTTGART                                    = "Vaihingen Stuttgart"
	VAIHINGEN_E_VAIHINGEN_E                                = "Vaihingen (E) Vaihingen (E)"
	VAIHINGEN_ALTER_FRIEDHOF_STUTTGART                     = "Vaihingen Alter Friedhof Stuttgart"
	VAIHINGEN_BF_OST_STUTTGART                             = "Vaihingen Bf (Ost) Stuttgart"
	VAIHINGEN_RATHAUS_STUTTGART                            = "Vaihingen Rathaus Stuttgart"
	VAIHINGEN_SCHILLERPLATZ_STUTTGART                      = "Vaihingen Schillerplatz Stuttgart"
	VAIHINGEN_VIADUKT_STUTTGART                            = "Vaihingen Viadukt Stuttgart"
	VAIHINGER_ECK_VAIHINGEN_E                              = "Vaihinger Eck Vaihingen (E)"
	VAIHINGER_STRASSE_STUTTGART                            = "Vaihinger Straße Stuttgart"
	VAIHINGER_STRASSE_SERSHEIM                             = "Vaihinger Straße Sersheim"
	VAIHINGER_STRASSE_NUSSDORF                             = "Vaihinger Straße Nussdorf"
	VARNBULERSTRASSE_HOFINGEN                              = "Varnbülerstraße Höfingen"
	VEILCHENSTRASSE_HOLZHAUSEN_UHINGEN                     = "Veilchenstraße Holzhausen (Uhingen)"
	VEILCHENWEG_LEINFELDEN                                 = "Veilchenweg Leinfelden"
	VENDELAUSTRASSE_NURTINGEN                              = "Vendelaustraße Nürtingen"
	VERDISTRASSE_NECKARHAUSEN                              = "Verdistraße Neckarhausen"
	VESOULER_STRASSE_GERLINGEN                             = "Vesouler Straße Gerlingen"
	VIADUKT_ENDERSBACH                                     = "Viadukt Endersbach"
	VIEHTRIEB_MALMSHEIM                                    = "Viehtrieb Malmsheim"
	VIEHWEIDE_SINDELFINGEN                                 = "Viehweide Sindelfingen"
	VIER_PAPPELN_SCHORNDORF                                = "Vier Pappeln Schorndorf"
	VIERGIEBELWEG_STUTTGART                                = "Viergiebelweg Stuttgart"
	VIKTORKOCHLWEG_STUTTGART                               = "Viktor-Köchl-Weg Stuttgart"
	VILLASTRASSE_STUTTGART                                 = "Villastraße Stuttgart"
	VILLENEUVESTRASSE_KORNWESTHEIM                         = "Villeneuvestraße Kornwestheim"
	VINZENZ_THERME_BAD_DITZENBACH                          = "Vinzenz Therme Bad Ditzenbach"
	VOGELSANG_STUTTGART                                    = "Vogelsang Stuttgart"
	VOGELSANG_NAGOLD                                       = "Vogelsang Nagold"
	VOGELSANGSTRASSE_DENKENDORF                            = "Vogelsangstraße Denkendorf"
	VOGGENHOF_ABZWEIG_ALTHUTTE                             = "Voggenhof Abzweig Althütte"
	VOGTHESSSCHULE_HERRENBERG                              = "Vogt-Heß-Schule Herrenberg"
	VOLKERBOHRINGERWEG_WIFLINGSHAUSEN                      = "Volker-Böhringer-Weg Wiflingshausen"
	VOLKSBANK_DETTINGEN_T                                  = "Volksbank Dettingen (T)"
	VOLKSBANK_ASPERG                                       = "Volksbank Asperg"
	VOLKSBANK_AUENDORF                                     = "Volksbank Auendorf"
	VOLKSBANK_ADELBERG                                     = "Volksbank Adelberg"
	VOLKSHOCHSCHULE_ESSLINGEN_N                            = "Volkshochschule Esslingen (N)"
	VOLLMWEGBIRKENWEG_NAGOLD                               = "Vollm.Weg/Birkenweg Nagold"
	VOLLZUGSANSTALT_HEIMSHEIM                              = "Vollzugsanstalt Heimsheim"
	VORDERBUCHELBERG_SPIEGELBERG                           = "Vorderbüchelberg Spiegelberg"
	VORDERBUCHELBERG_ABZWEIG_WUSTENROT                     = "Vorderbüchelberg Abzweig Wüstenrot"
	VORDERE_KARLSTRASSE_GOPPINGEN                          = "Vordere Karlstraße Göppingen"
	VORDERE_SCHMIEDGASSE_SCHWABISCH_GMUND                  = "Vordere Schmiedgasse Schwäbisch Gmünd"
	VORDERER_BERG_JEBENHAUSEN                              = "Vorderer Berg Jebenhausen"
	VORDERHUNDSBERG_WELZHEIM                               = "Vorderhundsberg Welzheim"
	VORDERSTEINENBERG_VORDERSTEINENBERG                    = "Vordersteinenberg Vordersteinenberg"
	VORDERWEISSBUCH_VORDERWEISSBUCH                        = "Vorderweißbuch Vorderweißbuch"
	VORDERWESTERMURR_MURRHARDT                             = "Vorderwestermurr Murrhardt"
	VORSTADT_NURTINGEN                                     = "Vorstadt Nürtingen"
	VOSSSTRASSE_GEISLINGEN_STEIGE                          = "Vossstraße Geislingen (Steige)"
	W__W_KORNWESTHEIM                                      = "W & W Kornwestheim"
	WAAGE_ERDMANNHAUSEN                                    = "Waage Erdmannhausen"
	WAAGE_AFFALTERBACH                                     = "Waage Affalterbach"
	WAGENBURGSTRASSE_STUTTGART                             = "Wagenburgstraße Stuttgart"
	WAGNERSTRASSE_ROMMELSHAUSEN                            = "Wagnerstraße Rommelshausen"
	WAGON_NAGOLD                                           = "Wagon Nagold"
	WAGRAINACKER_STUTTGART                                 = "Wagrainäcker Stuttgart"
	WAHLENHEIM_VORDERSTEINENBERG                           = "Wahlenheim Vordersteinenberg"
	WAIBLINGEN_WAIBLINGEN                                  = "Waiblingen Waiblingen"
	WAIBLINGER_STRASSE_BEINSTEIN                           = "Waiblinger Straße Beinstein"
	WAIBLINGERSCHORNDORFER_STR_FELLBACH                    = "Waiblinger-/Schorndorfer Str. Fellbach"
	WALA_BAD_BOLL                                          = "WALA Bad Boll"
	WALD_KIRCHHEIM_T                                       = "Wald Kirchheim (T)"
	WALDACKERWEG_WIFLINGSHAUSEN                            = "Waldackerweg Wiflingshausen"
	WALDALLEE_HOCHBERG                                     = "Waldallee Hochberg"
	WALDAU_STUTTGART                                       = "Waldau Stuttgart"
	WALDBURGSTRASSE_STUTTGART                              = "Waldburgstraße Stuttgart"
	WALDBURGSTRASSE_BOBLINGEN                              = "Waldburgstraße Böblingen"
	WALDCAFE_WANNENHOF_GOPPINGEN                           = "Waldcafe Wannenhof Göppingen"
	WALDECK_STUTTGART                                      = "Waldeck Stuttgart"
	WALDECKSCHULE_JEBENHAUSEN                              = "Waldeckschule Jebenhausen"
	WALDECKSTRASSE_ALBERSHAUSEN                            = "Waldeckstraße Albershausen"
	WALDECKSTRASSE_GOPPINGEN                               = "Waldeckstraße Göppingen"
	WALDENBRONN_KREUZUNG_HOHENKREUZ                        = "Wäldenbronn Kreuzung Hohenkreuz"
	WALDENBUCHER_PLATZ_STUTTGART                           = "Waldenbucher Platz Stuttgart"
	WALDENSERSTRASSE_PEROUSE                               = "Waldenserstraße Perouse"
	WALDENWEILER_SECHSELBERG                               = "Waldenweiler Sechselberg"
	WALDFRIEDHOF_LEONBERG                                  = "Waldfriedhof Leonberg"
	WALDFRIEDHOF_MAICHINGEN                                = "Waldfriedhof Maichingen"
	WALDFRIEDHOF_KIRCHHEIM_T                               = "Waldfriedhof Kirchheim (T)"
	WALDFRIEDHOF_NURTINGEN                                 = "Waldfriedhof Nürtingen"
	WALDFRIEDHOF_BOBLINGEN                                 = "Waldfriedhof Böblingen"
	WALDFRIEDHOF_HERRENBERG                                = "Waldfriedhof Herrenberg"
	WALDFRIEDHOF_BACKNANG                                  = "Waldfriedhof Backnang"
	WALDFRIEDHOF_STUTTGART                                 = "Waldfriedhof Stuttgart"
	WALDFRIEDHOF_GERLINGEN                                 = "Waldfriedhof Gerlingen"
	WALDFRIEDHOF_KOHLBERG                                  = "Waldfriedhof Kohlberg"
	WALDHAUSEN_WALDHAUSEN_B_SCHORND_LORCH                  = "Waldhausen Waldhausen b Schornd (Lorch)"
	WALDHEIM_SINDELFINGEN                                  = "Waldheim Sindelfingen"
	WALDHOF_BIETIGHEIM                                     = "Waldhof Bietigheim"
	WALDHORN_HOHENGEHREN                                   = "Waldhorn Hohengehren"
	WALDHORN_EISLINGEN_F                                   = "Waldhorn Eislingen (F)"
	WALDHORNSTRASSE_ECHTERDINGEN                           = "Waldhornstraße Echterdingen"
	WALDLUST_NAGOLD                                        = "Waldlust Nagold"
	WALDORFSCHULE_BONLANDEN                                = "Waldorfschule Bonlanden"
	WALDORFSCHULE_BOBLINGEN                                = "Waldorfschule Böblingen"
	WALDSCHULE_BISSINGEN_LB                                = "Waldschule Bissingen (LB)"
	WALDSEE_FORNSBACH                                      = "Waldsee Fornsbach"
	WALDSIEDLUNG_GERLINGEN                                 = "Waldsiedlung Gerlingen"
	WALDSTRASSE_PLATTENHARDT                               = "Waldstraße Plattenhardt"
	WALDSTRASSE_GROSSBOTTWAR                               = "Waldstraße Großbottwar"
	WALHEIM_WALHEIM                                        = "Walheim Walheim"
	WALKERSBACH_PLUDERHAUSEN                               = "Walkersbach Plüderhausen"
	WALKERSBACHER_TAL_LORCH                                = "Walkersbacher Tal Lorch"
	WALLGRABEN_STUTTGART                                   = "Wallgraben Stuttgart"
	WALLSTRASSE_WINNENDEN                                  = "Wallstraße Winnenden"
	WALTERFLEXSTRASSE_OSSWEIL                              = "Walter-Flex-Straße Oßweil"
	WALTERHELMESWEG_LEONBERG                               = "Walter-Helmes-Weg Leonberg"
	WALTERSBERG_MURRHARDT                                  = "Waltersberg Murrhardt"
	WANDERPARKPLATZ_ROTER_STICH_HOSSLINSWART               = "Wanderparkplatz Roter Stich Hößlinswart"
	WANDERWEG_MUSBERG                                      = "Wanderweg Musberg"
	WANGEN_KELTER_STUTTGART                                = "Wangen Kelter Stuttgart"
	WANGEN_MARKTPLATZ_STUTTGART                            = "Wangen Marktplatz Stuttgart"
	WANGENER_STRASSE_RECHBERGHAUSEN                        = "Wangener Straße Rechberghausen"
	WANGENER_STRASSE_GOPPINGEN                             = "Wangener Straße Göppingen"
	WANGENERLANDHAUSSTRASSE_STUTTGART                      = "Wangener-/Landhausstraße Stuttgart"
	WANKELSTRASSE_MALMSHEIM                                = "Wankelstraße Malmsheim"
	WANNENHOFE_AUFHAUSEN_GEISLINGEN                        = "Wannenhöfe Aufhausen (Geislingen)"
	WANNENRAIN_WEIL                                        = "Wannenrain Weil"
	WARNENBERG_OBERBOIHINGEN                               = "Warnenberg Oberboihingen"
	WARTBERGSTRASSE_MOGLINGEN                              = "Wartbergstraße Möglingen"
	WARTEHALLE_JESINGEN                                    = "Wartehalle Jesingen"
	WARTEHALLE_GERSTETTEN                                  = "Wartehalle Gerstetten"
	WARTH_OTLINGEN                                         = "Warth Ötlingen"
	WASCHERHOFSTRASSE_WASCHENBEUREN                        = "Wäscherhofstraße Wäschenbeuren"
	WASENFELD_ALLMERSBACH_IM_TAL                           = "Wasenfeld Allmersbach im Tal"
	WASENHOF_EISLINGEN_F                                   = "Wasenhof Eislingen (F)"
	WASENSTRASSE_STUTTGART                                 = "Wasenstraße Stuttgart"
	WASENSTRASSE_NECKARGRONINGEN                           = "Wasenstraße Neckargröningen"
	WASENSTRASSE_URBACH                                    = "Wasenstraße Urbach"
	WASENSTRASSE_JEBENHAUSEN                               = "Wasenstraße Jebenhausen"
	WASSERBERG_UHINGEN                                     = "Wasserberg Uhingen"
	WASSERBERGSTRASSE_GOPPINGEN                            = "Wasserbergstraße Göppingen"
	WASSERBERGWEG_PLUDERHAUSEN                             = "Wasserbergweg Plüderhausen"
	WASSERSTUBENWEG_WAIBLINGEN                             = "Wasserstubenweg Waiblingen"
	WASSERTURM_BESIGHEIM                                   = "Wasserturm Besigheim"
	WASSERWERK_NECKARTAILFINGEN                            = "Wasserwerk Neckartailfingen"
	WEBEREI_REICHENBACH_F                                  = "Weberei Reichenbach (F)"
	WEBERSTRASSE_NURTINGEN                                 = "Weberstraße Nürtingen"
	WEIDACH_STETTEN_LEINFELDENECHT                         = "Weidach Stetten (Leinfelden-Echt.)"
	WEIDACHTAL_STUTTGART                                   = "Weidachtal Stuttgart"
	WEIDENWEG_BACKNANG                                     = "Weidenweg Backnang"
	WEIHERSTRASSE_OBERESSLINGEN                            = "Weiherstraße Oberesslingen"
	WEIL_WEIL                                              = "Weil Weil"
	WEIL_DER_STADT_WEIL_DER_STADT                          = "Weil der Stadt Weil der Stadt"
	WEILDORF_BONDORF                                       = "Weildorf Bondorf"
	WEILEMER_PFAD_STUTTGART                                = "Weilemer Pfad Stuttgart"
	WEILENBERGER_HOF_UHINGEN                               = "Weilenberger Hof Uhingen"
	WEILER_BB_SCHAFHAUSEN                                  = "Weiler (BB) Schafhausen"
	WEILER_R_WEILER_SCHORNDORF                             = "Weiler (R) Weiler (Schorndorf)"
	WEILER_HELFENSTEIN_WEILER_OB_HELFENSTN                 = "Weiler Helfenstein Weiler ob Helfenstn."
	WEILER_LINDENHOF_GEISLINGEN_STEIGE                     = "Weiler Lindenhof Geislingen (Steige)"
	WEILERHAU_PLATTENHARDT                                 = "Weilerhau Plattenhardt"
	WEILHEIMER_STRASSE_ZELL_A                              = "Weilheimer Straße Zell (A)"
	WEILIMDORF_BF_STUTTGART                                = "Weilimdorf Bf Stuttgart"
	WEILIMDORF_BF_UNTERFUHRUNG_STUTTGART                   = "Weilimdorf Bf (Unterführung) Stuttgart"
	WEILIMDORF_LOWENMARKT_STUTTGART                        = "Weilimdorf Löwen-Markt Stuttgart"
	WEILSTRASSE_PLIENSAUVORSTADT                           = "Weilstraße Pliensauvorstadt"
	WEIMARSTRASSE_MARBACH_N                                = "Weimarstraße Marbach (N)"
	WEIMARSTRASSE_LUDWIGSBURG                              = "Weimarstraße Ludwigsburg"
	WEINBERGTECKSTRASSE_REICHENBACH_F                      = "Weinberg-/Teckstraße Reichenbach (F)"
	WEINBERGSTRASSE_HOCHDORF_ES                            = "Weinbergstraße Hochdorf (ES)"
	WEINBERGSTRASSE_REICHENBACH_F                          = "Weinbergstraße Reichenbach (F)"
	WEINGARTENTEGELBERG_GEISLINGEN_STEIGE                  = "Weingärten/Tegelberg Geislingen (Steige)"
	WEINGARTENSTRASSE_MONCHBERG                            = "Weingartenstraße Mönchberg"
	WEINGARTENSTRASSE_GAMMELSHAUSEN                        = "Weingartenstraße Gammelshausen"
	WEINGARTSTRASSE_DENKENDORF                             = "Weingartstraße Denkendorf"
	WEINSTEIGE_STUTTGART                                   = "Weinsteige Stuttgart"
	WEINSTEIGE_BEUTELSBACH                                 = "Weinsteige Beutelsbach"
	WEINSTEIGE_WEILHEIM_T                                  = "Weinsteige Weilheim (T)"
	WEINSTRASSE_ASPERG                                     = "Weinstraße Asperg"
	WEINSTRASSE_METTINGEN                                  = "Weinstraße Mettingen"
	WEISSACHER_STRASSE_HEIMERDINGEN                        = "Weissacher Straße Heimerdingen"
	WEISSACHSTRASSE_OBERWEISSACH                           = "Weissachstraße Oberweissach"
	WEISSENHOF_OST_LOCHGAU                                 = "Weißenhof Ost Löchgau"
	WEISSER_STEIN_PLOCHINGEN                               = "Weißer Stein Plochingen"
	WEIZENSTRASSE_KLEINGLATTBACH                           = "Weizenstraße Kleinglattbach"
	WELLARIUM_STEINHEIM_M                                  = "Wellarium Steinheim (M)"
	WELLINGEN_KIRCHLE_NOTZINGEN                            = "Wellingen Kirchle Notzingen"
	WELLINGEN_SCHLIERBACHER_STR_NOTZINGEN                  = "Wellingen Schlierbacher Str. Notzingen"
	WELLINGER_STRASSE_NOTZINGEN                            = "Wellinger Straße Notzingen"
	WELTESTRASSE_EISLINGEN_F                               = "Weltestraße Eislingen (F)"
	WELZHEIMER_STRASSE_SCHORNDORF                          = "Welzheimer Straße Schorndorf"
	WELZHEIMER_STRASSE_PFAHLBRONN                          = "Welzheimer Straße Pfahlbronn"
	WELZHEIMER_STRASSE_UNTERWEISSACH                       = "Welzheimer Straße Unterweissach"
	WELZHEIMER_STRASSE_GSCHWEND                            = "Welzheimer Straße Gschwend"
	WENDELKONIG_WAIBLINGEN                                 = "Wendelkönig Waiblingen"
	WENDEPLATTE_NECKARWEIHINGEN                            = "Wendeplatte Neckarweihingen"
	WENDEPLATTE_HOCHDORF_REMSECK                           = "Wendeplatte Hochdorf (Remseck)"
	WENDEPLATTE_MAINHARDT                                  = "Wendeplatte Mainhardt"
	WENDLINGEN_N_WENDLINGEN_N                              = "Wendlingen (N) Wendlingen (N)"
	WENGERTSTRASSE_SINDELFINGEN                            = "Wengertstraße Sindelfingen"
	WERASTRASSE_GEISLINGEN_STEIGE                          = "Werastraße Geislingen (Steige)"
	WERNAU_N_WERNAU_N                                      = "Wernau (N) Wernau (N)"
	WERNER_UND_PFLEIDERER_STUTTGART                        = "Werner und Pfleiderer Stuttgart"
	WERNERSTRASSE_FELLBACH                                 = "Wernerstraße Fellbach"
	WERNERSTRASSE_LUDWIGSBURG                              = "Wernerstraße Ludwigsburg"
	WERT_DEIZISAU                                          = "Wert Deizisau"
	WERTSTOFFHOF_GOPPINGEN                                 = "Wertstoffhof Göppingen"
	WESERSTRASSE_STUTTGART                                 = "Weserstraße Stuttgart"
	WESPENWEG_STUTTGART                                    = "Wespenweg Stuttgart"
	WEST_REUDERN                                           = "West Reudern"
	WESTBAHNHOF_STUTTGART                                  = "Westbahnhof Stuttgart"
	WESTBAHNHOF_SCHLEIFE_STUTTGART                         = "Westbahnhof (Schleife) Stuttgart"
	WESTENDSTRASSE_BISSINGEN_LB                            = "Westendstraße Bissingen (LB)"
	WESTFALENSTRASSE_OSSWEIL                               = "Westfalenstraße Oßweil"
	WESTRANDALLEE_SCHONAICH                                = "Westrandallee Schönaich"
	WESTSTADTKIRCHE_LUDWIGSBURG                            = "Weststadt-Kirche Ludwigsburg"
	WETTBACHPLATZ_SINDELFINGEN                             = "Wettbachplatz Sindelfingen"
	WETTE_LOCHGAU                                          = "Wette Löchgau"
	WETTEMARKT_OSSWEIL                                     = "Wettemarkt Oßweil"
	WETTERWARTE_SCHNITTLINGEN                              = "Wetterwarte Schnittlingen"
	WETZSTEINSTOLLEN_JUX                                   = "Wetzsteinstollen Jux"
	WHG_GOPPINGEN                                          = "WHG Göppingen"
	WIDDERSTALL_MERKLINGEN                                 = "Widderstall Merklingen"
	WIDDUMHOF_RUTESHEIM                                    = "Widdumhof Rutesheim"
	WIEGEHALLE_PLEIDELSHEIM                                = "Wiegehalle Pleidelsheim"
	WIEGEHALLERIEDBACHSTRASSE_PLEIDELSHEIM                 = "Wiegehalle/Riedbachstraße Pleidelsheim"
	WIELANDSHOHE_STUTTGART                                 = "Wielandshöhe Stuttgart"
	WIELANDSTRASSE_STUTTGART                               = "Wielandstraße Stuttgart"
	WIELANDSTRASSEBF_HOLZHEIM_GOPPINGEN                    = "Wielandstraße/Bf Holzheim (Göppingen)"
	WIESENACKERSTRASSE_HEIMERDINGEN                        = "Wiesenäckerstraße Heimerdingen"
	WIESENSTEIGER_STRASSE_GOSBACH                          = "Wiesensteiger Straße Gosbach"
	WIESENSTR_SCHORNDORF                                   = "Wiesenstr. Schorndorf"
	WIESENSTRASSE_ROMMELSHAUSEN                            = "Wiesenstraße Rommelshausen"
	WIESENTAL_WASCHENBEUREN                                = "Wiesental Wäschenbeuren"
	WIESENTALSTRASSE_SCHNAIT                               = "Wiesentalstraße Schnait"
	WIESENTALSTRASSE_HERTMANNSWEILER                       = "Wiesentalstraße Hertmannsweiler"
	WIESLAUFWEG_SCHORNDORF                                 = "Wieslaufweg Schorndorf"
	WILDENSTEINSTRASSE_STUTTGART                           = "Wildensteinstraße Stuttgart"
	WILDGANSWEG_STUTTGART                                  = "Wildgansweg Stuttgart"
	WILDPARADIES_TRIPSDRILL_CLEEBRONN                      = "Wildparadies Tripsdrill Cleebronn"
	WILHELMOLGASTRASSE_STUTTGART                           = "Wilhelm-/Olgastraße Stuttgart"
	WILHELMFEINSTRASSE_LUDWIGSBURG                         = "Wilhelm-Fein-Straße Ludwigsburg"
	WILHELMFOLLSTRASSE_STRUMPFELBACH_BACKNANG              = "Wilhelm-Föll-Straße Strümpfelbach (Backnang)"
	WILHELMGEIGERPLATZ_STUTTGART                           = "Wilhelm-Geiger-Platz Stuttgart"
	WILHELMHASPELSTRASSE_SINDELFINGEN                      = "Wilhelm-Haspel-Straße Sindelfingen"
	WILHELMKOPPSTRASSE_PEROUSE                             = "Wilhelm-Kopp-Straße Perouse"
	WILHELMMAIERSTRASSE_KONGEN                             = "Wilhelm-Maier-Straße Köngen"
	WILHELMPFITZERSTRASSE_FELLBACH                         = "Wilhelm-Pfitzer-Straße Fellbach"
	WILHELMRAABESTRASSE_WENDLINGEN_N                       = "Wilhelm-Raabe-Straße Wendlingen (N)"
	WILHELMA_STUTTGART                                     = "Wilhelma Stuttgart"
	WILHELMSBAU_STUTTGART                                  = "Wilhelmsbau Stuttgart"
	WILHELMSHILFE_GOPPINGEN                                = "Wilhelmshilfe Göppingen"
	WILHELMSHILFE_BARTENBACH_GOPPINGEN                     = "Wilhelmshilfe Bartenbach (Göppingen)"
	WILHELMSHOHE_GEISLINGEN_STEIGE                         = "Wilhelmshöhe Geislingen (Steige)"
	WILHELMSPLATZ_ASPERG                                   = "Wilhelmsplatz Asperg"
	WILHELMSPLATZ_HOCHDORF_REMSECK                         = "Wilhelmsplatz Hochdorf (Remseck)"
	WILHELMSTRASSE_NEUHAUSEN_F                             = "Wilhelmstraße Neuhausen (F)"
	WILHELMSTRASSE_WEIL_IM_SCHONBUCH                       = "Wilhelmstraße Weil im Schönbuch"
	WILHELMSTRASSE_KLEINSACHSENHEIM                        = "Wilhelmstraße Kleinsachsenheim"
	WILHELMSTRASSE_BISSINGEN_AN_DER_TECK_T                 = "Wilhelmstraße Bissingen an der Teck (T)"
	WILHELMSTRASSE_GEISLINGEN_STEIGE                       = "Wilhelmstraße Geislingen (Steige)"
	WILHELMSTRASSE_LORCH                                   = "Wilhelmstraße Lorch"
	WILLIBLEICHERSTRASSE_KIRCHHEIM_T                       = "Willi-Bleicher-Straße Kirchheim (T)"
	WIMPFENER_STRASSE_STUTTGART                            = "Wimpfener Straße Stuttgart"
	WINDHALMWEG_STUTTGART                                  = "Windhalmweg Stuttgart"
	WINKELS_GROSSSACHSENHEIM                               = "Winkels Großsachsenheim"
	WINNENDEN_WINNENDEN                                    = "Winnenden Winnenden"
	WINNENDER_STRASSE_WAIBLINGEN                           = "Winnender Straße Waiblingen"
	WINNENDER_STRASSE_SCHWAIKHEIM                          = "Winnender Straße Schwaikheim"
	WINNENDER_STRASSE_LEUTENBACH                           = "Winnender Straße Leutenbach"
	WINTERBACH_WINTERBACH                                  = "Winterbach Winterbach"
	WIRTSCHAFTSSCHULE_LAICHINGEN                           = "Wirtschaftsschule Laichingen"
	WITHAU_STUTTGART                                       = "Withau Stuttgart"
	WITHAUWEG_STUTTGART                                    = "Withauweg Stuttgart"
	WITTENBERGER_STRASSE_HERRENBERG                        = "Wittenberger Straße Herrenberg"
	WITTUMHALLE_URBACH                                     = "Wittumhalle Urbach"
	WOHRTOURSDEPOT_WEISSACH_BB                             = "Wöhr-Tours-Depot Weissach (BB)"
	WOLFHIRTHSTRASSE_SIRNAU                                = "Wolf-Hirth-Straße Sirnau"
	WOLFBUSCH_STUTTGART                                    = "Wolfbusch Stuttgart"
	WOLFENMUHLE_SCHONAICH                                  = "Wolfenmühle Schönaich"
	WOLFSKLINGENWEG_WINNENDEN                              = "Wolfsklingenweg Winnenden"
	WOLFSOLDEN_AFFALTERBACH                                = "Wolfsölden Affalterbach"
	WOLFSTRASSE_ALTENHEIM_SINDELFINGEN                     = "Wolfstraße Altenheim Sindelfingen"
	WOLKBAD_GEISLINGEN_STEIGE                              = "Wölkbad Geislingen (Steige)"
	WOLLINSTRASSE_STUTTGART                                = "Wollinstraße Stuttgart"
	WUNNEBAD_WINNENDEN                                     = "Wunnebad Winnenden"
	WUNNENSTEINSTRASSE_STUTTGART                           = "Wunnensteinstraße Stuttgart"
	WURMBRUCKE_SCHAFHAUSEN                                 = "Würmbrücke Schafhausen"
	WURMBRUCKE_EHNINGEN                                    = "Würmbrücke Ehningen"
	WURMLINGER_STRASSE_STUTTGART                           = "Wurmlinger Straße Stuttgart"
	WURMSTRASSE_ALTDORF_BB                                 = "Würmstraße Altdorf (BB)"
	WURMTALSTRASSE_HAUSEN_WEIL_DER_STADT                   = "Würmtalstraße Hausen (Weil der Stadt)"
	WURTTEMBERGSTRASSE_DITZINGEN                           = "Württembergstraße Ditzingen"
	WURTTEMBERGSTRASSE_WEIL                                = "Württembergstraße Weil"
	WUSTENROT_LUDWIGSBURG                                  = "Wüstenrot Ludwigsburg"
	YHAUSER_GEISLINGEN_STEIGE                              = "Y-Häuser Geislingen (Steige)"
	YITZHAKRABINSTRASSE_STUTTGART                          = "Yitzhak-Rabin-Straße Stuttgart"
	ZABERGAUSTRASSE_NEUHAUSEN_F                            = "Zabergäustraße Neuhausen (F)"
	ZABERGAUSTRASSE_STUTTGART                              = "Zabergäustraße Stuttgart"
	ZABERSTRASSE_OBERJESINGEN                              = "Zaberstraße Oberjesingen"
	ZACHERSMUHLE_ADELBERG                                  = "Zachersmühle Adelberg"
	ZAHNNOPPERSTRASSE_STUTTGART                            = "Zahn-Nopper-Straße Stuttgart"
	ZAHNKLINIK_BISSINGEN_LB                                = "Zahnklinik Bissingen (LB)"
	ZAHNRADBAHNHOF_STUTTGART                               = "Zahnradbahnhof Stuttgart"
	ZAHRINGER_STRASSE_LINDORF                              = "Zähringer Straße Lindorf"
	ZAUNACKER_ECHTERDINGEN                                 = "Zaunacker Echterdingen"
	ZAZENHAUSEN_STUTTGART                                  = "Zazenhausen Stuttgart"
	ZAZENHAUSEN_KRONE_STUTTGART                            = "Zazenhausen Krone Stuttgart"
	ZAZENHAUSEN_STEIGLE_STUTTGART                          = "Zazenhausen Steigle Stuttgart"
	ZAZENHAUSEN_VIADUKT_STUTTGART                          = "Zazenhausen Viadukt Stuttgart"
	ZEHNTSCHEUER_SCHONAICH                                 = "Zehntscheuer Schönaich"
	ZEHNTSCHEUER_POPPENWEILER                              = "Zehntscheuer Poppenweiler"
	ZEIL_KLEINSACHSENHEIM                                  = "Zeil Kleinsachsenheim"
	ZEISIGWEG_HERRENBERG                                   = "Zeisigweg Herrenberg"
	ZEISIGWEG_EISLINGEN_F                                  = "Zeisigweg Eislingen (F)"
	ZEISSSTRASSE_DITZINGEN                                 = "Zeissstraße Ditzingen"
	ZELL_ZELL_ESSLINGEN                                    = "Zell Zell (Esslingen)"
	ZELL_OPPENWEILER                                       = "Zell Oppenweiler"
	ZELL_RECHBERGHAUSER_STR_BORTLINGEN                     = "Zell Rechberghäuser Str. Börtlingen"
	ZELLER_STRASSE_OHMDEN                                  = "Zeller Straße Ohmden"
	ZELLER_STRASSE_HATTENHOFEN                             = "Zeller Straße Hattenhofen"
	ZELLERSTRASSE_STUTTGART                                = "Zellerstraße Stuttgart"
	ZENTRUM_SULZGRIES                                      = "Zentrum Sulzgries"
	ZENTRUM_ZIEGELFELD_HERRENBERG                          = "Zentrum Ziegelfeld Herrenberg"
	ZEPPELINKIRCHSTRASSE_DEIZISAU                          = "Zeppelin-/Kirchstraße Deizisau"
	ZEPPELINSCHULE_FELLBACH                                = "Zeppelinschule Fellbach"
	ZEPPELINSTRASSE_OBERESSLINGEN                          = "Zeppelinstraße Oberesslingen"
	ZEPPELINSTRASSE_KEMNAT                                 = "Zeppelinstraße Kemnat"
	ZEPPELINSTRASSE_KORNWESTHEIM                           = "Zeppelinstraße Kornwestheim"
	ZEPPELINSTRASSE_DEIZISAU                               = "Zeppelinstraße Deizisau"
	ZIEGELBACHKLINGE_EISLINGEN_F                           = "Ziegelbachklinge Eislingen (F)"
	ZIEGELBACHSTRASSE_EISLINGEN_F                          = "Ziegelbachstraße Eislingen (F)"
	ZIEGELBERGWEG_UHINGEN                                  = "Ziegelbergweg Uhingen"
	ZIEGELEISTRASSE_FRICKENHAUSEN                          = "Ziegeleistraße Frickenhausen"
	ZIEGELEISTRASSE_MUNCHINGEN                             = "Ziegeleistraße Münchingen"
	ZIEGELHOF_HOCHDORF_ES                                  = "Ziegelhof Hochdorf (ES)"
	ZIEGELHUTTE_KAISERSBACH                                = "Ziegelhütte Kaisersbach"
	ZIEGELSTRASSE_RECHBERGHAUSEN                           = "Ziegelstraße Rechberghausen"
	ZIEGELSTRASSE_SALACH                                   = "Ziegelstraße Salach"
	ZIEGELWASEN_KIRCHHEIM_T                                = "Ziegelwasen Kirchheim (T)"
	ZIEGELWASENSTRASSE_GRAFENBERG                          = "Ziegelwasenstraße Grafenberg"
	ZIEGERHOF_LENGLINGEN                                   = "Ziegerhof Lenglingen"
	ZIMMERBACHSTRASSE_OBERESSLINGEN                        = "Zimmerbachstraße Oberesslingen"
	ZIMMERER_PFAD_GROSSSACHSENHEIM                         = "Zimmerer Pfad Großsachsenheim"
	ZIMMERSCHLAG_BOBLINGEN                                 = "Zimmerschlag Böblingen"
	ZIMMERSTRASSE_SINDELFINGEN                             = "Zimmerstraße Sindelfingen"
	ZINSHOLZ_RUIT                                          = "Zinsholz Ruit"
	ZOB_SINDELFINGEN                                       = "ZOB Sindelfingen"
	ZOB_ESSLINGEN_N                                        = "ZOB Esslingen (N)"
	ZOB_NAGOLD                                             = "ZOB Nagold"
	ZOB_LAICHINGEN                                         = "ZOB Laichingen"
	ZOLLBERG_ZOLLBERG                                      = "Zollberg Zollberg"
	ZOLLERNLICHTENSTEINSTRASSE_DITZINGEN                   = "Zollern-/Lichtensteinstraße Ditzingen"
	ZOLLERNPLATZ_ZOLLBERG                                  = "Zollernplatz Zollberg"
	ZOLLHAUSWEG_ZOLLBERG                                   = "Zollhausweg Zollberg"
	ZUCKERBERG_STUTTGART                                   = "Zuckerberg Stuttgart"
	ZUFFENHAUSEN_STUTTGART                                 = "Zuffenhausen Stuttgart"
	ZUFFENHAUSEN_FRIEDHOF_STUTTGART                        = "Zuffenhausen Friedhof Stuttgart"
	ZUFFENHAUSEN_KELTERPLATZ_STUTTGART                     = "Zuffenhausen Kelterplatz Stuttgart"
	ZUFFENHAUSEN_PORSCHE_STUTTGART                         = "Zuffenhausen Porsche Stuttgart"
	ZUFFENHAUSEN_RATHAUS_STUTTGART                         = "Zuffenhausen Rathaus Stuttgart"
	ZUFFENHAUSEN_REIBEDANZ_STUTTGART                       = "Zuffenhausen Reibedanz Stuttgart"
	ZUFFENHAUSER_HEIDE_STUTTGART                           = "Zuffenhäuser Heide Stuttgart"
	ZUM_KLEINEN_FELDLE_STETTEN_I_R                         = "Zum Kleinen Feldle Stetten (i. R.)"
	ZUR_ANHOHE_STUTTGART                                   = "Zur Anhöhe Stuttgart"
	ZURICHER_STRASSE_STUTTGART                             = "Züricher Straße Stuttgart"
	ZURICHER_STRASSE_BACKNANG                              = "Züricher Straße Backnang"
	ZWERCHWEG_HERRENBERG                                   = "Zwerchweg Herrenberg"
	ZWERENBERG_SULZBACH_M                                  = "Zwerenberg Sulzbach (M)"
	ZWINGELH_GROSSASPACHER_STRASSE_KIRCHBERG_M             = "Zwingelh. Großaspacher Straße Kirchberg (M)"
	ZWINGELHAUSEN_ABZWEIG_GROSSASPACH                      = "Zwingelhausen Abzweig Großaspach"
	ZWINGELHAUSEN_KIRCHBERGER_STR_KIRCHBERG_M              = "Zwingelhausen Kirchberger Str. Kirchberg (M)"
	ZWINGELHAUSER_STRASSE_KIRCHBERG_M                      = "Zwingelhäuser Straße Kirchberg (M)"
)

var stationNameIdMap = map[string]string{
	AALSTRASSE_STUTTGART:                                   "de:08111:2480",
	ABELSTRASSE_LUDWIGSBURG:                                "de:08118:5503",
	ABZW_HOHENEGARTEN_MAINHARDT:                            "de:08127:27509",
	ABZW_HOHENEGARTEN_B14_MAINHARDT:                        "de:08127:28273",
	ABZW_REHNENHAUWETZGAU_SCHWABISCH_GMUND:                 "de:08136:3058",
	ABZWEIG_ROHRBRONN:                                      "de:08119:3701",
	ABZWEIG_HOLZGERLINGEN:                                  "de:08115:4206",
	ABZWEIG_SCHLATTSTALL:                                   "de:08116:4419",
	ABZWEIG_NEUWEILER:                                      "de:08115:4789",
	ABZWEIG_MONCHBERG:                                      "de:08115:4853",
	ABZWEIG_HEBSACK:                                        "de:08119:5197",
	ABZWEIG_AURICH:                                         "de:08118:5752",
	ABZWEIG_DATZINGEN:                                      "de:08115:5941",
	ABZWEIG_ALTERSBERG_KAISERSBACH:                         "de:08119:5093",
	ABZWEIG_B466_REICHENBACH_IM_TALE_DEGGINGEN:             "de:08117:49",
	ABZWEIG_BAHNHOF_MARKGRONINGEN:                          "de:08118:3306",
	ABZWEIG_BEIM_KLOSTER_ADELBERG:                          "de:08117:7654",
	ABZWEIG_DAFERN_LIPPOLDSWEILER:                          "de:08119:5993",
	ABZWEIG_DEGENFELD_LAUTERSTEIN:                          "de:08117:330",
	ABZWEIG_DRACKENSTEIN_GOSBACH:                           "de:08117:57",
	ABZWEIG_EISENBACHSEE_PFAHLBRONN:                        "de:08119:6942",
	ABZWEIG_GREUTHOF_VORDERSTEINENBERG:                     "de:08119:5026",
	ABZWEIG_HOHREIN_GOPPINGEN:                              "de:08117:1303",
	ABZWEIG_KORNBERGHALLE_DURNAU:                           "de:08117:3328",
	ABZWEIG_NEUFURSTENHUTTE_GROSSERLACH:                    "de:08119:4908",
	ABZWEIG_OBERWALDEN_WANGEN_GP:                           "de:08117:5016",
	ABZWEIG_RAUHER_KAPF_BOBLINGEN:                          "de:08115:4617",
	ABZWEIG_REICHENBACH_DONZDORF:                           "de:08117:7620",
	ABZWEIG_SCHMALENBERG_RUDERSBERG:                        "de:08119:5125",
	ABZWEIG_STEINBACH_RUDERSBERG:                           "de:08119:5124",
	ABZWEIG_WALTERTAL_HOHENSTADT_GP:                        "de:08117:3506",
	ABZWEIG_WINZINGEN_DONZDORF:                             "de:08117:7625",
	ABZWEIGUNG_BAIERECK_THOMASHARDT:                        "de:08116:4502",
	ABZWEIGUNG_HEPSISAU_WEILHEIM_T:                         "de:08116:4399",
	ACHALMSTRASSE_HOLZGERLINGEN:                            "de:08115:3244",
	ACHALMSTRASSE_NURTINGEN:                                "de:08116:4448",
	ADELSBACH_WINNENDEN:                                    "de:08119:3758",
	ADELSTETTEN_PFAHLBRONN:                                 "de:08119:5105",
	ADENAUERBRUCKE_OBEN_OBERESSLINGEN:                      "de:08116:4143",
	ADLER_HEININGEN_GP:                                     "de:08117:3307",
	ADLER_ALTENSTADT_GEISLINGEN_STEIGE:                     "de:08117:10",
	ADLERAPOTHEKE_STRASSDORF:                               "de:08136:3508",
	ADLERBRUCKE_PLUDERHAUSEN:                               "de:08119:5242",
	ADLERPLATZ_HOCHBERG:                                    "de:08118:5550",
	ADLERSTRASSE_FELLBACH:                                  "de:08119:2646",
	ADLERSTRASSE_WERNAU_N:                                  "de:08116:6723",
	ADLERSTRASSE_HERRENBERG:                                "de:08115:7057",
	ADLERSTRASSE_OTTENBACH:                                 "de:08117:6016",
	ADLERWEG_GEMMRIGHEIM:                                   "de:08118:6813",
	ADMIRALWEG_STUTTGART:                                   "de:08111:201",
	ADOLFKOLPINGSTRASSE_DONZDORF:                           "de:08117:7632",
	AFFSTATTER_TAL_HERRENBERG:                              "de:08115:7065",
	AGENTUR_FUR_ARBEIT_BIETIGHEIM:                          "de:08118:3343",
	AGENTUR_FUR_ARBEIT_LUDWIGSBURG:                         "de:08118:5529",
	AGNESWEG_HOLZGERLINGEN:                                 "de:08115:6530",
	AHORNSTRASSE_PLUDERHAUSEN:                              "de:08119:5234",
	AHRENWEG_STUTTGART:                                     "de:08111:2535",
	AICHELBACH_OPPENWEILER:                                 "de:08119:7485",
	AICHENBACHHOF_PLUDERHAUSEN:                             "de:08119:5241",
	AICHHOLZHOF_UNTERWEISSACH:                              "de:08119:5994",
	AICHSTRUT_WELZHEIM:                                     "de:08119:5065",
	AIDLINGER_STRASSE_DAGERSHEIM:                           "de:08115:4764",
	AIDLINGER_STRASSE_DEUFRINGEN:                           "de:08115:4779",
	AKADEMIEGARTEN_NEUHAUSEN_F:                             "de:08116:2893",
	AKAZIENWEG_HERRENBERG:                                  "de:08115:7067",
	ALBBLICK_ZELL_ESSLINGEN:                                "de:08116:4021",
	ALBBLICK_STUTTGART:                                     "de:08111:6016",
	ALBERTEINSTEINGYMNASIUM_BOBLINGEN:                      "de:08115:4680",
	ALBERTLUTHULIPLATZ_STUTTGART:                           "de:08111:2476",
	ALBERTSCHAFFLESTRASSE_STUTTGART:                        "de:08111:2189",
	ALBERTSCHWEITZERGYMNASIUM_LEONBERG:                     "de:08115:4547",
	ALBERTSCHWEITZERSCHULE_GOPPINGEN:                       "de:08117:4904",
	ALBERTSCHWEITZERSTRASSE_BESIGHEIM:                      "de:08118:3468",
	ALBERTSCHWEITZERSTRASSE_HOFINGEN:                       "de:08115:6816",
	ALBRECHTHAUTGASSE_KUPPINGEN:                            "de:08115:4844",
	ALBRECHTSTRASSE_LUDWIGSBURG:                            "de:08118:5537",
	ALBSCHULE_STUTTGART:                                    "de:08111:6570",
	ALBSPORTHALLE_BOHMENKIRCH:                              "de:08117:316",
	ALBSTRASSE_REICHENBACH_F:                               "de:08116:4237",
	ALBSTRASSE_KORNWESTHEIM:                                "de:08118:5471",
	ALBSTRASSE_GROSSBETTLINGEN:                             "de:08116:5901",
	ALBSTRASSE_NEUFFEN:                                     "de:08116:7081",
	ALBSTRASSE_SPARWIESEN:                                  "de:08117:7013",
	ALBSTRASSE_SIEDLERSTUBE_NURTINGEN:                      "de:08116:2929",
	ALDINGER_STRASSE_KORNWESTHEIM:                          "de:08118:5446",
	ALDINGER_STRASSE_LUDWIGSBURG:                           "de:08118:5522",
	ALDINGER_STRASSE_HEGNACH:                               "de:08119:5546",
	ALEMANNENSTRASSE_SIELMINGEN:                            "de:08116:2030",
	ALEMANNENSTRASSE_OEFFINGEN:                             "de:08119:2374",
	ALEXANDRINENPLATZ_HOCHBERG:                             "de:08118:5551",
	ALFREDDIEBOLDWEG_WAIBLINGEN:                            "de:08119:3684",
	ALFREDLEIKAMGARTEN_KORB:                                "de:08119:5414",
	ALLEEALTER_WEG_TAMM:                                    "de:08118:6886",
	ALLEEBRACHTERSTRASSE_TAMM:                              "de:08118:6885",
	ALLEENFELD_FREUDENTAL:                                  "de:08118:7658",
	ALLEENSTRASSE_ZELL_ESSLINGEN:                           "de:08116:4047",
	ALLMANDKLINGE_HOHENHASLACH:                             "de:08118:3567",
	ALLMENDSTRASSE_MAICHINGEN:                              "de:08115:7029",
	ALLMERSBACH_AM_WEINBERG_ALLMERSBACH_AM_WEINBERG_ASPACH: "de:08119:5360",
	ALM_MURRHARDT:                                          "de:08119:4922",
	ALPSEEWEG_STUTTGART:                                    "de:08111:6274",
	ALTBACH_ALTBACH:                                        "de:08116:1800",
	ALTBACHER_HOF_ALTBACH:                                  "de:08116:3800",
	ALTBACHER_WEG_LIEBERSBRONN:                             "de:08116:4072",
	ALTDORFER_STRASSE_HILDRIZHAUSEN:                        "de:08115:4797",
	ALTE_BERGSTRASSE_DEIZISAU:                              "de:08116:6879",
	ALTE_BUNDESSTRASSE_WAIBLINGEN:                          "de:08119:5955",
	ALTE_HEUSTEIGE_ZELL_ESSLINGEN:                          "de:08116:3807",
	ALTE_KELTER_FELLBACH:                                   "de:08119:2504",
	ALTE_KIRCHE_JEBENHAUSEN:                                "de:08117:2102",
	ALTE_KRONE_STUTTGART:                                   "de:08111:2508",
	ALTE_MUHLE_HAFNERHASLACH:                               "de:08118:5735",
	ALTE_PLOCHINGER_STEIGE_KIRCHHEIM_T:                     "de:08116:3950",
	ALTE_POST_DETTENHAUSEN:                                 "de:08416:6353",
	ALTE_RENNINGER_STRASSE_111_WEIL_DER_STADT:              "de:08115:6472",
	ALTE_RENNINGER_STRASSE_19_WEIL_DER_STADT:               "de:08115:6474",
	ALTE_RENNINGER_STRASSE_49_WEIL_DER_STADT:               "de:08115:6473",
	ALTE_SCHIFFFAHRT_METTINGEN:                             "de:08116:3992",
	ALTE_SCHULE_NEUSTADT:                                   "de:08119:7545",
	ALTE_STEIGE_GERLINGEN:                                  "de:08118:2301",
	ALTE_STEIGE_LIEBERSBRONN:                               "de:08116:4104",
	ALTE_STEIGE_SCHORNDORF:                                 "de:08119:6991",
	ALTE_STRASSE_FORNSBACH:                                 "de:08119:4927",
	ALTE_TALSTRASSE_ST_BERNHARDT:                           "de:08116:4136",
	ALTE_WAIBLINGER_STRASSE_NEUSTADT:                       "de:08119:5147",
	ALTE_WINNENDER_STEIGE_WAIBLINGEN:                       "de:08119:3720",
	ALTENBERGSTAFFEL_STUTTGART:                             "de:08111:162",
	ALTENBURG_STUTTGART:                                    "de:08111:2251",
	ALTENHEIM_PLATTENHARDT:                                 "de:08116:2056",
	ALTENHEIM_FELLBACH:                                     "de:08119:5432",
	ALTENHEIM_BURGHALDE_SINDELFINGEN:                       "de:08115:4628",
	ALTENSTEIGER_STRASSE_NAGOLD:                            "de:08235:10112",
	ALTENZENTRUM_HOLZGERLINGEN:                             "de:08115:3228",
	ALTER_BAHNHOF_HEININGEN_GP:                             "de:08117:3320",
	ALTER_BAHNHOFWEG_HOFINGEN:                              "de:08115:6051",
	ALTER_FRIEDHOF_DARMSHEIM:                               "de:08115:4637",
	ALTER_FRIEDHOF_BOBLINGEN:                               "de:08115:4686",
	ALTER_GUTSHOF_STUTTGART:                                "de:08111:2428",
	ALTER_OSSWEILER_WEG_LUDWIGSBURG:                        "de:08118:7435",
	ALTER_SPORTPLATZ_WENDLINGEN_N:                          "de:08116:4293",
	ALTER_WEG_TAMM:                                         "de:08118:6887",
	ALTES_RATHAUS_MAICHINGEN:                               "de:08115:3180",
	ALTES_RATHAUS_HILDRIZHAUSEN:                            "de:08115:4798",
	ALTES_RATHAUS_UNTERJETTINGEN:                           "de:08115:4863",
	ALTES_RATHAUS_KEMNAT:                                   "de:08116:5017",
	ALTES_RATHAUS_BITTENFELD:                               "de:08119:5156",
	ALTES_RATHAUS_WEISSENSTEIN:                             "de:08117:7666",
	ALTES_RATHAUS_DETTENHAUSEN:                             "de:08416:6362",
	ALTES_SCHULHAUS_ENSINGEN:                               "de:08118:6964",
	ALTFURSTENHUTTE_GROSSERLACH:                            "de:08119:4969",
	ALTSTADT_LEONBERG:                                      "de:08115:2352",
	ALTSTADT_WALDENBUCH:                                    "de:08115:3129",
	ALTSTADTGARAGE_HERRENBERG:                              "de:08115:3225",
	ALTVATERWEG_KIRCHHEIM_T:                                "de:08116:4328",
	ALUP_KONGEN:                                            "de:08116:4091",
	AM_AUSSEREN_GRABEN_STUTTGART:                           "de:08111:192",
	AM_BERGWALD_STUTTGART:                                  "de:08111:2544",
	AM_BISMARCKTURM_STUTTGART:                              "de:08111:2453",
	AM_BRENNER_LIPPOLDSWEILER:                              "de:08119:5384",
	AM_BRONNBACH_WEILER_SCHORNDORF:                         "de:08119:3754",
	AM_BRUNNELE_ALBERSHAUSEN:                               "de:08117:2018",
	AM_DRESSELBACH_SACHSENWEILER:                           "de:08119:5348",
	AM_ESCHELBACH_HOLZGERLINGEN:                            "de:08115:6532",
	AM_FELDRAND_STUTTGART:                                  "de:08111:2592",
	AM_FRIEDHOF_PLEIDELSHEIM:                               "de:08118:3582",
	AM_FRIEDHOF_GROSSSACHSENHEIM:                           "de:08118:7088",
	AM_GALGENBERG_GOPPINGEN:                                "de:08117:9302",
	AM_KATZENBACH_WAIBLINGEN:                               "de:08119:5160",
	AM_KINDELBERG_RENNINGEN:                                "de:08115:3188",
	AM_KIRCHPLATZ_STUTTGART:                                "de:08111:199",
	AM_KREUZWIESLE_LORCH:                                   "de:08136:7043",
	AM_KRIEGSBERGTURM_STUTTGART:                            "de:08111:2245",
	AM_KUGELRAIN_BAIERECK:                                  "de:08117:7669",
	AM_LINDLE_BONLANDEN:                                    "de:08116:2082",
	AM_MAURENER_BERG_DITZINGEN:                             "de:08118:3019",
	AM_MEERBACH_BARTENBACH_GOPPINGEN:                       "de:08117:4917",
	AM_MITTELKAI_STUTTGART:                                 "de:08111:4067",
	AM_MUHLBACH_SCHORNDORF:                                 "de:08119:7569",
	AM_OCHSENWALD_STUTTGART:                                "de:08111:2615",
	AM_PFAFFENWALD_HOCHDORF_EBERDINGEN:                     "de:08118:6952",
	AM_ROMERKASTELL_STUTTGART:                              "de:08111:2250",
	AM_SCHATTWALD_STUTTGART:                                "de:08111:2596",
	AM_SCHLOSS_BONNIGHEIM:                                  "de:08118:5698",
	AM_SCHLOSSBERG_HOFINGEN:                                "de:08115:4572",
	AM_SCHONEN_RAIN_ESSLINGEN_N:                            "de:08116:4000",
	AM_SONNENBERG_LUDWIGSBURG:                              "de:08118:3464",
	AM_TURMLE_UNTERMBERG:                                   "de:08118:5691",
	AM_UNTEREN_SCHLOSSBERG_NECKARREMS:                      "de:08118:5767",
	AM_WARMEN_MUHLHAUSEN_IM_TALE:                           "de:08117:3230",
	AM_WASEN_OTLINGEN:                                      "de:08116:4386",
	AM_WASSERTURM_MARKGRONINGEN:                            "de:08118:3312",
	AM_WIESENBACH_HOLZHAUSEN_UHINGEN:                       "de:08117:2034",
	AM_WIESENGRUND_BERKHEIM:                                "de:08116:4149",
	AM_WIESENRAIN_KIRCHHEIM_T:                              "de:08116:4307",
	AM_ZIPFELBACH_BITTENFELD:                               "de:08119:5155",
	AMSELWEG_WELZHEIM:                                      "de:08119:5128",
	AMSELWEG_GEMMRIGHEIM:                                   "de:08118:6811",
	AMSELWEG_HERRENBERG:                                    "de:08115:7040",
	AMSTERDAMER_STRASSE_BOBLINGEN:                          "de:08115:4697",
	AMSTETTER_STRASSE_STUTTGART:                            "de:08111:4068",
	AMTSGERICHT_NURTINGEN:                                  "de:08116:4373",
	AMUNDSENSTRASSE_SINDELFINGEN:                           "de:08115:3229",
	AN_DER_WETTE_KORNWESTHEIM:                              "de:08118:5460",
	ANHAUSERSTRASSE_HEGENSBERG:                             "de:08116:4100",
	ANNASCHIEBERWEG_ESSLINGEN_N:                            "de:08116:4028",
	ANNEFRANKSCHULE_WENDLINGEN_N:                           "de:08116:7577",
	ANNEFRANKSTRASSE_HOLZGERLINGEN:                         "de:08115:6531",
	ANTONSCHMIDTSTRASSE_WAIBLINGEN:                         "de:08119:3686",
	ANTONIAVISCONTISTRASSE_BIETIGHEIM:                      "de:08118:5690",
	ANTWERPENER_STRASSE_STUTTGART:                          "de:08111:35",
	AOKWALDHORN_BACKNANG:                                   "de:08119:3605",
	APOTHEKE_OTLINGEN:                                      "de:08116:4822",
	APOTHEKE_BEILSTEIN:                                     "de:08125:4673",
	ARBEITERZENTRUM_BOBLINGEN:                              "de:08115:3101",
	ARBEITSAGENTURPOST_WAIBLINGEN:                          "de:08119:7093",
	ARNDTSPITTASTRASSE_STUTTGART:                           "de:08111:218",
	ARSENALPLATZ_LUDWIGSBURG:                               "de:08118:3408",
	ARZTEHAUS_MURRHARDT:                                    "de:08119:4913",
	ASANG_STUTTGART:                                        "de:08111:2532",
	ASEMWALD_STUTTGART:                                     "de:08111:2573",
	ASPACHER_BRUCKE_BACKNANG:                               "de:08119:5329",
	ASPEN_STUTTGART:                                        "de:08111:2511",
	ASPEN_HOCHDORF_ES:                                      "de:08116:6772",
	ASPENWALDSTRASSE_STUTTGART:                             "de:08111:2418",
	ASPERG_ASPERG:                                          "de:08118:7400",
	ASPERGER_STRASSE_STUTTGART:                             "de:08111:2459",
	ASPERGLEN_ASPERGLEN:                                    "de:08119:5115",
	AUBERLENWEG_STUTTGART:                                  "de:08111:5854",
	AUCHTERTSTRASSE_SCHLIERBACH:                            "de:08117:7801",
	AUENWALDHALLE_UNTERBRUDEN:                              "de:08119:5383",
	AUERHAHNSTRASSE_SCHORNDORF:                             "de:08119:5976",
	AUF_DEM_BUHL_LORCH:                                     "de:08136:7004",
	AUF_DEM_KIES_BESIGHEIM:                                 "de:08118:5715",
	AUF_DEM_LERCHENBERG_NURTINGEN:                          "de:08116:4404",
	AUF_DEM_WASEN_LUDWIGSBURG:                              "de:08118:3475",
	AUF_DEN_RUDERN_BURGSTALL_BURGSTETTEN:                   "de:08119:3545",
	AUF_DER_GANS_STUTTGART:                                 "de:08111:2380",
	AUFHAUSEN_AUFHAUSEN_GEISLINGEN:                         "de:08117:257",
	AUFHAUSER_STR_NELLINGEN_UL:                             "de:08425:2620",
	AUGSBURGER_PLATZ_STUTTGART:                             "de:08111:33",
	AUGUSTBEBELSTRASSE_EGLOSHEIM:                           "de:08118:3409",
	AUGUSTBRANDLESTRASSE_FELLBACH:                          "de:08119:2516",
	AUGUSTINUM_STUTTGART:                                   "de:08111:2546",
	AUSSERE_HAUPTSTRASSE_GEISLINGEN_STEIGE:                 "de:08117:95",
	AUSSIEDLERHOFE_MOGLINGEN:                               "de:08118:3430",
	AUSSIEDLERHOFE_GROSSBOTTWAR:                            "de:08118:5619",
	AUSSIEDLERHOFE_HOFEN_BONNIGHEIM:                        "de:08118:6721",
	AUSSIEDLERHOFE_HASLACH:                                 "de:08115:6889",
	AUSSIEDLERHOFE_STOTTEN:                                 "de:08117:320",
	AUSSIEDLERHOFE_HEININGEN_GP:                            "de:08117:3306",
	AUSTRASSE_GRUNBACH:                                     "de:08119:3745",
	AUSTRASSE_KIRCHHEIM_T:                                  "de:08116:3947",
	AUSTRASSE_VAIHINGEN_E:                                  "de:08118:5746",
	AUSTRASSE_WENDLINGEN_N:                                 "de:08116:6864",
	AUTALHALLE_BAD_UBERKINGEN:                              "de:08117:47",
	AUTENSTRASSE_DITZINGEN:                                 "de:08118:3000",
	AUTMUT_TISCHARDT:                                       "de:08116:2943",
	AUWIESEN_STUTTGART:                                     "de:08111:271",
	AUWIESENBRUCKE_BIETIGHEIM:                              "de:08118:5648",
	AUWIESENSCHULE_NECKARTENZLINGEN:                        "de:08116:4310",
	AVE_MARIA_DEGGINGEN:                                    "de:08117:52",
	B_297HOHE_LINDE_KIRNECK:                                "de:08136:7002",
	B_297KORNSTR_OBERKIRNECK:                               "de:08136:7023",
	B14_MAINHARDT:                                          "de:08127:28102",
	B28_BERNECK:                                            "de:08235:10156",
	B312_NECKARTAILFINGEN:                                  "de:08116:5819",
	B466_HAUSEN_BAD_UBERKINGEN:                             "de:08117:48",
	BAACH_SCHNAIT:                                          "de:08119:5952",
	BAACHER_STRASSE_HOFEN:                                  "de:08119:5217",
	BAACHER_WEG_HERTMANNSWEILER:                            "de:08119:6664",
	BACHBRUCKE_SCHLIERBACH:                                 "de:08117:7804",
	BACHHALDE_STUTTGART:                                    "de:08111:6272",
	BACHLENWEG_STUTTGART:                                   "de:08111:2520",
	BACHSTRASSE_SINDELFINGEN:                               "de:08115:4610",
	BACHSTRASSE_ALTDORF_BB:                                 "de:08115:4795",
	BACHSTRASSE_GROSSBOTTWAR:                               "de:08118:5606",
	BACHSTRASSE_WEISSACH_BB:                                "de:08115:5912",
	BACHSTRASSE_SCHLAT:                                     "de:08117:3020",
	BACKHAUS_DACHTEL:                                       "de:08115:4181",
	BACKHAUS_OBERBERKEN:                                    "de:08119:4947",
	BACKHAUS_GROSSHEPPACH:                                  "de:08119:5170",
	BACKHAUS_ESCHENBACH:                                    "de:08117:3316",
	BACKHAUS_OBERWALDEN:                                    "de:08117:5015",
	BACKHAUSLE_HANWEILER:                                   "de:08119:5311",
	BACKNANG_BACKNANG:                                      "de:08119:7600",
	BACKNANGER_STRASSE_WINNENDEN:                           "de:08119:3634",
	BACKNANGER_STRASSE_RUDERSBERG:                          "de:08119:5035",
	BAD_LEINFELDEN:                                         "de:08116:7102",
	BAD_BAD_UBERKINGEN:                                     "de:08117:46",
	BAD_CANNSTATT_STUTTGART:                                "de:08111:6333",
	BAD_CANNSTATT_WILHELMSPLATZ_STUTTGART:                  "de:08111:6332",
	BADBERKAWEG_MARBACH_N:                                  "de:08118:3594",
	BADEPARK_BIETIGHEIM:                                    "de:08118:5701",
	BADEZENTRUM_SINDELFINGEN:                               "de:08115:4626",
	BADSTRASSE_ALTBACH:                                     "de:08116:3898",
	BADSTRASSE_LIPPOLDSWEILER:                              "de:08119:4982",
	BADWIESEN_KIRCHHEIM_T:                                  "de:08116:4361",
	BAHNBRUCKE_BONDORF:                                     "de:08115:4761",
	BAHNHOF_BERNHAUSEN:                                     "de:08116:2006",
	BAHNHOF_HOLZMADEN:                                      "de:08116:4324",
	BAHNHOF_STEINHEIM_M:                                    "de:08118:5602",
	BAHNHOF_DONZDORF:                                       "de:08117:7619",
	BAHNHOF_ROHRDORF_CW:                                    "de:08235:10165",
	BAHNHOF_ARENA_LUDWIGSBURG:                              "de:08118:7401",
	BAHNHOF_DAMMSTRASSE_WAIBLINGEN:                         "de:08119:7602",
	BAHNHOF_FLUGFELD_BOBLINGEN:                             "de:08115:7921",
	BAHNHOF_INDUSTRIESTR_RENNINGEN:                         "de:08115:3190",
	BAHNHOF_PLOCHINGER_STRASSE_NURTINGEN:                   "de:08116:4375",
	BAHNHOF_POST_NURTINGEN:                                 "de:08116:2988",
	BAHNHOF_SCHLUSSELACKER_MAICHINGEN:                      "de:08115:3179",
	BAHNHOF_SUD_OBERESSLINGEN:                              "de:08116:3326",
	BAHNHOF_TROLLINGERWEG_METTINGEN:                        "de:08116:3993",
	BAHNHOF_UHLANDSTRASSE_FRICKENHAUSEN:                    "de:08116:2822",
	BAHNHOF_WARTHSTRASSE_KORNTAL:                           "de:08118:2631",
	BAHNHOFSALLEE_BAD_BOLL:                                 "de:08117:2106",
	BAHNHOFSSTEG_GOPPINGEN:                                 "de:08117:3385",
	BAHNHOFSTRASSE_NEUHAUSEN_F:                             "de:08116:2901",
	BAHNHOFSTRASSE_NURTINGEN:                               "de:08116:2966",
	BAHNHOFSTRASSE_MOGLINGEN:                               "de:08118:3432",
	BAHNHOFSTRASSE_ERDMANNHAUSEN:                           "de:08118:3551",
	BAHNHOFSTRASSE_WENDLINGEN_N:                            "de:08116:3834",
	BAHNHOFSTRASSE_UNTERLENNINGEN:                          "de:08116:4283",
	BAHNHOFSTRASSE_WEILHEIM_T:                              "de:08116:4326",
	BAHNHOFSTRASSE_WEISSACH_BB:                             "de:08115:4533",
	BAHNHOFSTRASSE_HEMMINGEN:                               "de:08118:4558",
	BAHNHOFSTRASSE_SCHONAICH:                               "de:08115:4806",
	BAHNHOFSTRASSE_PLUDERHAUSEN:                            "de:08119:5245",
	BAHNHOFSTRASSE_NELLMERSBACH:                            "de:08119:5321",
	BAHNHOFSTRASSE_ENDERSBACH:                              "de:08119:6837",
	BAHNHOFSTRASSE_SERSHEIM:                                "de:08118:7417",
	BAHNHOFSTRASSE_GINGEN_F:                                "de:08117:17",
	BAHNHOFSTRASSE_MUHLHAUSEN_IM_TALE:                      "de:08117:59",
	BAHNHOFSTRASSE_BIRENBACH:                               "de:08117:1103",
	BAHNHOFSTRASSE_WASCHENBEUREN:                           "de:08117:1107",
	BAHNUBERFUHRUNG_GERADSTETTEN:                           "de:08119:3748",
	BAHNUNTERFUHRUNG_WINTERBACH:                            "de:08119:4083",
	BAJASTRASSE_WAIBLINGEN:                                 "de:08119:3685",
	BALZHOLZ_RATHAUS_BEUREN:                                "de:08116:4438",
	BALZHOLZER_STRASSE_BEUREN:                              "de:08116:5864",
	BANK_NECKARGRONINGEN:                                   "de:08118:4891",
	BANNHALDE_GROSSSACHSENHEIM:                             "de:08118:5888",
	BANNMUHLE_REICHENBACH_F:                                "de:08116:4506",
	BANRAIN_URBACH:                                         "de:08119:6691",
	BARBAROSSATHERMEN_GOPPINGEN:                            "de:08117:9401",
	BARBAROSSASEE_GOPPINGEN:                                "de:08117:1050",
	BARENWIESEN_SERACH:                                     "de:08116:4133",
	BARTENBACH_SULZBACH_M:                                  "de:08119:4912",
	BARTENWEG_SINDELFINGEN:                                 "de:08115:5929",
	BAUERNBRESLAUER_STRASSE_DITZINGEN:                      "de:08118:3020",
	BAUHOF_NURTINGEN:                                       "de:08116:2941",
	BAUHOF_GOPPINGEN:                                       "de:08117:1035",
	BAUMGARTEN_SALACH:                                      "de:08117:6019",
	BAUMGARTENWEG_STUTTGART:                                "de:08111:2434",
	BAUMWASENSTRASSE_SCHORNDORF:                            "de:08119:7570",
	BAUSCHE_WELZHEIM:                                       "de:08119:5055",
	BBW_SCHELMENHOLZ:                                       "de:08119:6660",
	BEBELSTRASSE_GEISLINGEN_STEIGE:                         "de:08117:43",
	BEERHALDENSTRASSE_ENZWEIHINGEN:                         "de:08118:7416",
	BEETHOVENSTRASSE_STUTTGART:                             "de:08111:212",
	BEETHOVENSTRASSE_PLOCHINGEN:                            "de:08116:4231",
	BEETHOVENSTRASSE_MERKLINGEN:                            "de:08115:4589",
	BEETHOVENSTRASSE_NEUWEILER:                             "de:08115:4821",
	BEETHOVENSTRASSE_UNTERWEISSACH:                         "de:08119:5287",
	BEETHOVENSTRASSE_HERRENBERG:                            "de:08115:7004",
	BEETHOVENSTRASSE_BITTENFELD:                            "de:08119:7550",
	BEETHOVENSTRASSE_LORCH:                                 "de:08136:7006",
	BEETHOVENSTRASSE_35_LORCH:                              "de:08136:7007",
	BEETHOVENSTRASSE_ABZW_LORCH:                            "de:08136:7008",
	BEIHINGER_PLATZ_FREIBERG_N:                             "de:08118:3519",
	BEIHINGER_STRASSE_HOHENECK:                             "de:08118:3472",
	BEIHINGER_STRASSE_PLEIDELSHEIM:                         "de:08118:5569",
	BEILSTEINER_STRASSE_OBERSTENFELD:                       "de:08118:5593",
	BEILSTEINER_STRASSE_STUTTGART:                          "de:08111:6267",
	BEIM_KREISEL_ENDERSBACH:                                "de:08119:6838",
	BEIM_WASSERTURM_WAIBLINGEN:                             "de:08119:3632",
	BEIM_WASSERTURM_BACKNANG:                               "de:08119:6938",
	BELFORTER_PLATZ_LEONBERG:                               "de:08115:2345",
	BEMPFLINGEN_BEMPFLINGEN:                                "de:08116:5772",
	BENNINGEN_N_BENNINGEN_N:                                "de:08118:1500",
	BENZACH_BEUTELSBACH:                                    "de:08119:5396",
	BENZACH_BILDUNGSZENTRUM_BEUTELSBACH:                    "de:08119:5397",
	BENZENHOFWEG_KIRCHHEIM_T:                               "de:08116:4353",
	BENZSTRASSE_URBACH:                                     "de:08119:6627",
	BENZSTRASSE_HERRENBERG:                                 "de:08115:7054",
	BERGFRIEDHOF_STUTTGART:                                 "de:08111:77",
	BERGHAUSTRASSE_STUTTGART:                               "de:08111:6007",
	BERGHEIMER_HOF_STUTTGART:                               "de:08111:146",
	BERGSIEDLUNG_DACHTEL:                                   "de:08115:4189",
	BERGSIEDLUNG_KEPLERSTRASSE_DACHTEL:                     "de:08115:4190",
	BERGSTAFFELSTRASSE_STUTTGART:                           "de:08111:90",
	BERGSTRASSE_KORNTAL:                                    "de:08118:2626",
	BERGSTRASSE_SULZGRIES:                                  "de:08116:4118",
	BERGSTRASSE_WERNAU_N:                                   "de:08116:4246",
	BERGSTRASSE_RECHBERGHAUSEN:                             "de:08117:4040",
	BERGSTRASSE_LORCH:                                      "de:08136:7009",
	BERGWALD_GECHINGEN:                                     "de:08235:1500",
	BERLINER_PLATZ_LUDWIGSBURG:                             "de:08118:7433",
	BERLINER_PLATZ_HOHE_STRASSE_STUTTGART:                  "de:08111:55",
	BERLINER_PLATZ_LIEDERHALLE_STUTTGART:                   "de:08111:6073",
	BERLINER_RING_BACKNANG:                                 "de:08119:5330",
	BERLINER_STRASSE_MAICHINGEN:                            "de:08115:3184",
	BERLINER_STRASSE_HERRENBERG:                            "de:08115:4513",
	BERLINER_STRASSE_BOBLINGEN:                             "de:08115:4700",
	BERLINER_STRASSE_ALDINGEN:                              "de:08118:5657",
	BERLINER_STRASSE_ASPERG:                                "de:08118:6947",
	BERNHALDEN_OPPENWEILER:                                 "de:08119:4956",
	BERNHAUSER_STRASSE_STUTTGART:                           "de:08111:2037",
	BERNHAUSER_WEG_NEUHAUSEN_F:                             "de:08116:2895",
	BERNSTEINSTRASSE_STUTTGART:                             "de:08111:6137",
	BERTHABENZSTRASSE_ENSINGEN:                             "de:08118:7579",
	BERUFLICHES_SCHULZENTRUM_ZELL_ESSLINGEN:                "de:08116:4170",
	BERUFSBILDUNGSZENTRUM_STUTTGART:                        "de:08111:2582",
	BERUFSSCHULE_PATTONVILLE:                               "de:08118:3574",
	BERUFSSCHULZENTRUM_WAIBLINGEN:                          "de:08119:3650",
	BERUFSSCHULZENTRUM_GEISLINGEN_STEIGE:                   "de:08117:118",
	BERUFSSCHULZENTRUM_ODE_GOPPINGEN:                       "de:08117:9896",
	BERWINKEL_SULZBACH_M:                                   "de:08119:4906",
	BESIGHEIM_BESIGHEIM:                                    "de:08118:5583",
	BESIGHEIMER_STRASSE_EGLOSHEIM:                          "de:08118:5535",
	BESIGHEIMER_STRASSE_LOCHGAU:                            "de:08118:5703",
	BESIGHEIMER_STRASSE_FREUDENTAL:                         "de:08118:5706",
	BESIGHEIMER_STRASSE_OTTMARSHEIM:                        "de:08118:6524",
	BESKIDENSTRASSE_STUTTGART:                              "de:08111:36",
	BETHEL_WELZHEIM:                                        "de:08119:5058",
	BETRIEBSHOF_WARMBRONN:                                  "de:08115:6734",
	BETTACKER_NEBRINGEN:                                    "de:08115:7920",
	BETTLINGER_FORUM_GROSSBETTLINGEN:                       "de:08116:5900",
	BETZ_KONGEN:                                            "de:08116:4090",
	BETZSTRASSE_GOPPINGEN:                                  "de:08117:1036",
	BEUNDSTRASSE_EISLINGEN_F:                               "de:08117:1045",
	BEURENER_STEIGE_BEUREN:                                 "de:08116:3933",
	BEUTELSBACH_BEUTELSBACH:                                "de:08119:3711",
	BEZIRKSAMT_BEZGENRIET:                                  "de:08117:2103",
	BEZIRKSAMT_WISSGOLDINGEN:                               "de:08136:7821",
	BIEGEL_BACKNANG:                                        "de:08119:3661",
	BIELER_WEG_STUTTGART:                                   "de:08111:375",
	BIETIGHEIM_BIETIGHEIM:                                  "de:08118:1400",
	BIETIGHEIMER_STRASSE_LUDWIGSBURG:                       "de:08118:5505",
	BIETIGHEIMER_STRASSE_GROSSINGERSHEIM:                   "de:08118:5566",
	BIETIGHEIMER_STRASSE_KLEINSACHSENHEIM:                  "de:08118:5669",
	BIETIGHEIMER_STRASSE_METTERZIMMERN:                     "de:08118:6719",
	BIETIGHEIMER_STRASSE_MURR:                              "de:08118:6999",
	BIHLPLATZ_STUTTGART:                                    "de:08111:15",
	BILDACKERSTRASSE_HOHENACKER:                            "de:08119:5150",
	BILDSTOCKLE_LEONBERG:                                   "de:08115:2320",
	BILDUNGSCAMPUS_GROSSSACHSENHEIM:                        "de:08118:5638",
	BILDUNGSZENTRUM_WEILHEIM_T:                             "de:08116:3893",
	BILDUNGSZENTRUM_MARKGRONINGEN:                          "de:08118:5860",
	BILDUNGSZENTRUM_COTTENWEILER:                           "de:08119:5995",
	BILDUNGSZENTRUM_II_WINNENDEN:                           "de:08119:5307",
	BILDUNGSZENTRUM_WEST_LUDWIGSBURG:                       "de:08118:7419",
	BILFINGER_STRASSE_FREIBERG_N:                           "de:08118:3523",
	BINSENSTRASSE_SCHALKSTETTEN:                            "de:08425:2578",
	BIRKACH_FRIEDHOF_STUTTGART:                             "de:08111:2557",
	BIRKACH_WEST_STUTTGART:                                 "de:08111:2559",
	BIRKACHER_STRASSE_STUTTGART:                            "de:08111:2545",
	BIRKENALLEE_BAHNUNTERFUHRUNG_PLUDERHAUSEN:              "de:08119:6683",
	BIRKENKOPF_STUTTGART:                                   "de:08111:2451",
	BIRKENSTRASSE_HOLZGERLINGEN:                            "de:08115:4857",
	BIRKENWEG_WAIBLINGEN:                                   "de:08119:3721",
	BIRKENWEG_ALTBACH:                                      "de:08116:3835",
	BIRKENWEG_BISSINGEN_LB:                                 "de:08118:7429",
	BIRKENWEG_GEISLINGEN_STEIGE:                            "de:08117:104",
	BIRKENWEISSBUCH_VORDERWEISSBUCH:                        "de:08119:5265",
	BIRKHAU_AFFALTERBACH:                                   "de:08118:3600",
	BIRKHECKENSTRASSE_STUTTGART:                            "de:08111:2558",
	BIRKHOF_KAISERSBACH:                                    "de:08119:6924",
	BIRKHOF_DEGGINGEN:                                      "de:08117:10075",
	BIRKHOFE_ROSSWALDEN:                                    "de:08117:7414",
	BISMARCKPLATZ_STUTTGART_WEST:                           "de:08111:6069",
	BISMARCKSTAFFEL_STUTTGART:                              "de:08111:2261",
	BISMARCKSTRASSE_ESSLINGEN_N:                            "de:08116:4005",
	BISMARCKSTRASSE_KIRCHHEIM_T:                            "de:08116:4332",
	BISSINGER_STRASSE_EGLOSHEIM:                            "de:08118:5584",
	BISSINGER_STRASSE_UNTERMBERG:                           "de:08118:5645",
	BISSINGER_STRASSE_OCHSENWANG:                           "de:08116:7821",
	BITTENFELDER_STRASSE_HOHENACKER:                        "de:08119:5152",
	BLAMMERBERGSTRASSE_WEIL_DER_STADT:                      "de:08115:6471",
	BLATTERT_MURR:                                          "de:08118:5597",
	BLEICHEREI_UHINGEN:                                     "de:08117:2037",
	BLEICHSTRASSE_GOPPINGEN:                                "de:08117:9025",
	BLICK_STUTTGART:                                        "de:08111:265",
	BLOSENBERGKIRCHE_LEONBERG:                              "de:08115:2356",
	BLUHENDES_BAROCK_LUDWIGSBURG:                           "de:08118:5497",
	BLUMENWILHELMSTRASSE_REICHENBACH_F:                     "de:08116:6496",
	BLUMENSTRASSE_BACKNANG:                                 "de:08119:3606",
	BLUMENSTRASSE_UNTERENSINGEN:                            "de:08116:4289",
	BLUMENSTRASSE_REICHENBACH_F:                            "de:08116:4508",
	BLUMENSTRASSE_WAIBLINGEN:                               "de:08119:5144",
	BLUMENSTRASSE_KLEININGERSHEIM:                          "de:08118:5568",
	BOBLINGEN_BOBLINGEN:                                    "de:08115:7100",
	BODELSCHWINGHSCHULE_MURRHARDT:                          "de:08119:4914",
	BOEHRINGER_AREAL_GOPPINGEN:                             "de:08117:2003",
	BOHNAU_KIRCHHEIM_T:                                     "de:08116:4823",
	BOHNAUHAUS_KIRCHHEIM_T:                                 "de:08116:3952",
	BOHNREUTE_KIRCHENKIRNBERG:                              "de:08119:5080",
	BOHRINGSWEILER_GROSSERLACH:                             "de:08119:4970",
	BOLLER_STRASSE_DURNAU:                                  "de:08117:3331",
	BOLZPLATZ_HOHENGEHREN:                                  "de:08116:7934",
	BOLZSTRASSE_KORNWESTHEIM:                               "de:08118:5452",
	BOLZSTRASSE_BIETIGHEIM:                                 "de:08118:5658",
	BOMBACH_AICH:                                           "de:08116:4475",
	BONDORF_BONDORF:                                        "de:08115:4510",
	BONHOLZ_WALDENBUCH:                                     "de:08115:2815",
	BONHOLZ_ALFDORF:                                        "de:08119:3257",
	BONHOLZ_ABZWEIG_ALFDORF:                                "de:08119:5104",
	BONLANDENER_STRASSE_BONLANDEN:                          "de:08116:2015",
	BONLANDER_STRASSE_ECHTERDINGEN:                         "de:08116:2101",
	BOPSER_STUTTGART:                                       "de:08111:6160",
	BOPSERWALDSTRASSE_GERLINGEN:                            "de:08118:2302",
	BORKUMSTRASSE_NEUWSIEDL_STUTTGART:                      "de:08111:2460",
	BORSENPLATZ_STUTTGART:                                  "de:08111:229",
	BORSIGSTRASSE_STUTTGART:                                "de:08111:111",
	BORSIGSTRASSE_BISSINGEN_LB:                             "de:08118:7474",
	BOSCH_RUTESHEIM:                                        "de:08115:4557",
	BOSCH_TOR_III_SCHWIEBERDINGEN:                          "de:08118:3321",
	BOSCH_TOR_IV_SCHWIEBERDINGEN:                           "de:08118:6926",
	BOSCHHOF_TURKHEIM:                                      "de:08117:512",
	BOSKOPWEG_WAIBLINGEN:                                   "de:08119:3692",
	BOTNANG_STUTTGART:                                      "de:08111:6503",
	BOTNANG_FREIBAD_STUTTGART:                              "de:08111:2421",
	BOTNANGER_SATTEL_STUTTGART:                             "de:08111:6214",
	BOTTROPER_STRASSE_STUTTGART:                            "de:08111:2260",
	BOTTWARTALSTRASSE_HOHENECK:                             "de:08118:3425",
	BRAHMSSTRASSE_KIRCHHEIM_T:                              "de:08116:4215",
	BRANDENBURGER_STRASSE_OBERESSLINGEN:                    "de:08116:4001",
	BRANDHOF_L_1080_GSCHWEND:                               "de:08136:6538",
	BRAUEREI_SCHWIEBERDINGEN:                               "de:08118:3332",
	BRAUEREIPLATZ_MAGSTADT:                                 "de:08115:3298",
	BRAUNGARTWEG_ZOLLBERG:                                  "de:08116:4176",
	BRECH_PFAHLBRONN:                                       "de:08119:5101",
	BRECH_KAISERLINDE_PFAHLBRONN:                           "de:08119:5102",
	BREECH_KRZG_RATTENHARZ_BORTLINGEN:                      "de:08117:4035",
	BREECH_WENDEPLATTE_BORTLINGEN:                          "de:08117:4050",
	BREITENF_REINHOLDMAIERPLATZ_WELZHEIM:                   "de:08119:5132",
	BREITENFURST_LORCHER_STRASSE_WELZHEIM:                  "de:08119:5087",
	BREITENFURST_SCHULHAUS_WELZHEIM:                        "de:08119:5056",
	BREITENSTEINER_STRASSE_BOBLINGEN:                       "de:08115:3137",
	BREITER_WEG_BOHMENKIRCH:                                "de:08117:217",
	BREITER_WEG_EBHAUSEN:                                   "de:08235:10208",
	BREITESTRASSE_HEININGEN_GP:                             "de:08117:3311",
	BREITFELDERHOF_OTTENBACH:                               "de:08117:6013",
	BREITWIESEN_GERLINGEN:                                  "de:08118:143",
	BREITWIESENHAUS_GERLINGEN:                              "de:08118:6909",
	BREND_PFAHLBRONN:                                       "de:08119:6642",
	BRENDLE_GROSSMARKT_STUTTGART:                           "de:08111:81",
	BRENNERSTRASSE_ELTINGEN:                                "de:08115:4543",
	BRESLAUER_STRASSE_BOBLINGEN:                            "de:08115:3106",
	BRESLAUER_STRASSE_SCHWIEBERDINGEN:                      "de:08118:6957",
	BRESLAUER_STRASSE_BEILSTEIN:                            "de:08125:4676",
	BRESLAUER_STRASSEKINDERGARTEN_DITZINGEN:                "de:08118:3021",
	BREUNINGERLAND_LUDWIGSBURG:                             "de:08118:3448",
	BREUNINGERLAND_SINDELFINGEN:                            "de:08115:6043",
	BRIEFZENTRUM_SALACH:                                    "de:08117:7608",
	BROMBERGERHOFE_HOHENHASLACH:                            "de:08118:3503",
	BRONNTOR_HERRENBERG:                                    "de:08115:4739",
	BRONNWIESENWEG_RUDERSBERG:                              "de:08119:7529",
	BRUCH_BRUCH_WEISSACH:                                   "de:08119:4989",
	BRUCH_KAISERSBACH:                                      "de:08119:5077",
	BRUCK_LORCH:                                            "de:08136:7000",
	BRUCKEN_UNTERLENNINGEN:                                 "de:08116:5780",
	BRUCKEN_UNTERE_STRASSE_UNTERLENNINGEN:                  "de:08116:4282",
	BRUCKENSTRASSE_WERNAU_N:                                "de:08116:6722",
	BRUCKENSTRASSE_HESSIGHEIM:                              "de:08118:3499",
	BRUCKENSTRASSE_ALDINGEN:                                "de:08118:3572",
	BRUCKENSTRASSE_WINNENDEN:                               "de:08119:3696",
	BRUCKENSTRASSE_BIRENBACH:                               "de:08117:82",
	BRUCKENSTRASSE_GINGEN_F:                                "de:08117:86",
	BRUCKENSTRASSE_GOPPINGEN:                               "de:08117:3003",
	BRUCKENWEG_GRUIBINGEN:                                  "de:08117:3234",
	BRUCKLE_WALDENBRONN:                                    "de:08116:4128",
	BRUCKWIESEN_GEISLINGEN_STEIGE:                          "de:08117:11",
	BRUCKWIESEN_SPIELPLATZ_HATTENHOFEN:                     "de:08117:2062",
	BRUCKWIESENSTRASSE_HATTENHOFEN:                         "de:08117:2051",
	BRUDERHAUS_STUTTGART:                                   "de:08111:2206",
	BRUHL_WERNAU_N:                                         "de:08116:4249",
	BRUHL_LINDORF:                                          "de:08116:4347",
	BRUHL_FREILICHTMUSEUM_BEUREN:                           "de:08116:3930",
	BRUHLSIEDLUNG_NEUHAUSEN_F:                              "de:08116:2902",
	BRUHLSTRASSE_MOGLINGEN:                                 "de:08118:3465",
	BRUHLSTRASSE_GROSSINGERSHEIM:                           "de:08118:5567",
	BRUHLSTRASSE_DECKENPFRONN:                              "de:08115:7011",
	BRUHLSTRASSE_MOTZINGEN:                                 "de:08115:7020",
	BRUHLSTRASSE_22_MAGSTADT:                               "de:08115:4600",
	BRUHLWEG_SALACH:                                        "de:08117:6022",
	BRUNNEN_WALDENBRONN:                                    "de:08116:4127",
	BRUNNENGARTEN_WIESENSTEIG:                              "de:08117:65",
	BRUNNENSTRASSE_WELZHEIM:                                "de:08119:7582",
	BRUNNENSTRASSE_GINGEN_F:                                "de:08117:122",
	BRUNNENSTRASSE_BARTENBACH_GOPPINGEN:                    "de:08117:4918",
	BRUNNENWIESEN_SPIELBERG:                                "de:08118:3346",
	BRUNNER_STRASSE_LUDWIGSBURG:                            "de:08118:5534",
	BUBENBAD_STUTTGART:                                     "de:08111:124",
	BUCH_HOLZGERLINGEN:                                     "de:08115:4786",
	BUCH_BERLINER_STRASSE_BIETIGHEIM:                       "de:08118:5477",
	BUCH_BRESLAUER_STRASSE_BIETIGHEIM:                      "de:08118:5679",
	BUCH_BUCHSCHULE_BIETIGHEIM:                             "de:08118:5678",
	BUCH_BUCHZENTRUM_BIETIGHEIM:                            "de:08118:3385",
	BUCH_DRESDNER_STRASSE_BIETIGHEIM:                       "de:08118:5676",
	BUCH_GRONINGER_WEG_BIETIGHEIM:                          "de:08118:5675",
	BUCH_POSENER_STRASSE_BIETIGHEIM:                        "de:08118:5680",
	BUCH_SUCYSTRASSE_BIETIGHEIM:                            "de:08118:5681",
	BUCHENBRONN_FORSTHAUS_EBERSBACH_F:                      "de:08117:7673",
	BUCHENBRONN_GEMEINDEHAUS_EBERSBACH_F:                   "de:08117:7638",
	BUCHENBRONN_KONIGSEICHE_EBERSBACH_F:                    "de:08117:7674",
	BUCHENBRONN_SIEDLUNG_EBERSBACH_F:                       "de:08117:7672",
	BUCHENBRONNER_STRASSE_EBERSBACH_F:                      "de:08117:8012",
	BUCHENGEHREN_PFAHLBRONN:                                "de:08119:6781",
	BUCHENGEHREN_SAGMUHLE_PFAHLBRONN:                       "de:08119:6782",
	BUCHENRAIN_OCHSENBACH:                                  "de:08118:5883",
	BUCHENSTRASSE_STEINACH_BERGLEN:                         "de:08119:5293",
	BUCHENSTRASSE_FAURNDAU:                                 "de:08117:5004",
	BUCHENWEG_WAIBLINGEN:                                   "de:08119:3725",
	BUCHENWEG_SCHORNDORF:                                   "de:08119:5222",
	BUCHHALDENSCHULE_AIDLINGEN:                             "de:08115:7027",
	BUCHHALDENSTRASSE_SCHNAIT:                              "de:08119:5403",
	BUCHRAIN_HOLZHEIM_GOPPINGEN:                            "de:08117:3009",
	BUCHRAIN_FRIEDHOF_STUTTGART:                            "de:08111:6200",
	BUCHSENSTRASSE_STUTTGART:                               "de:08111:69",
	BUCHSTEINER_GINGEN_F:                                   "de:08117:18",
	BUCHWALD_STUTTGART:                                     "de:08111:2187",
	BUDAPESTER_PLATZ_STUTTGART:                             "de:08111:299",
	BUHL_KORB:                                              "de:08119:3689",
	BUHLACKER_SCHONAICH:                                    "de:08115:4817",
	BUHLACKERSTRASSE_STETTEN_I_R:                           "de:08119:3771",
	BUHLBRONN_BUHLBRONN:                                    "de:08119:5262",
	BUHLENSTRASSE_HOLZGERLINGEN:                            "de:08115:4787",
	BUHLER_STRASSE_BOBLINGEN:                               "de:08115:4710",
	BUHLFELD_OPPENWEILER:                                   "de:08119:4950",
	BUHNERSTRASSE_SCHMIDEN:                                 "de:08119:2376",
	BULKESWEG_KIRCHHEIM_T:                                  "de:08116:4330",
	BUNSENSTR_OBERENSINGEN:                                 "de:08116:2994",
	BUNZWANGER_STRASSE_ALBERSHAUSEN:                        "de:08117:2045",
	BUNZWANGER_STRASSE_UHINGEN:                             "de:08117:9825",
	BUOCHER_STRASSE_BREUNINGSWEILER:                        "de:08119:5423",
	BUOLWEG_ENZWEIHINGEN:                                   "de:08118:7448",
	BURG_ESSLINGEN_N:                                       "de:08116:4123",
	BURG_BURG:                                              "de:08119:5272",
	BURGERHAUS_BEUREN:                                      "de:08116:2980",
	BURGERHAUS_MAICHINGEN:                                  "de:08115:3215",
	BURGERHAUS_AICHELBERG_GP:                               "de:08117:2061",
	BURGERHAUS_RSKN_RUDERN:                                 "de:08116:4114",
	BURGERMUHLE_BONNIGHEIM:                                 "de:08118:3593",
	BURGERSEE_KIRCHHEIM_T:                                  "de:08116:4382",
	BURGERZENTRUM_HEGENLOHE:                                "de:08116:6970",
	BURGERZENTRUM_HALLENBAD_WAIBLINGEN:                     "de:08119:3627",
	BURGERZENTRUM_REMSBRUCKE_WAIBLINGEN:                    "de:08119:5145",
	BURGGYMNASIUM_SCHORNDORF:                               "de:08119:5960",
	BURGHALDENFRIEDHOF_SINDELFINGEN:                        "de:08115:4629",
	BURGHALDENSTRASSE_POPPENWEILER:                         "de:08118:5561",
	BURGHOLZ_PFAHLBRONN:                                    "de:08119:6643",
	BURGHOLZHOF_STUTTGART:                                  "de:08111:2470",
	BURGKLINGE_GERLINGEN:                                   "de:08118:2303",
	BURGPLATZ_BACKNANG:                                     "de:08119:5345",
	BURGPLATZ_BONNIGHEIM:                                   "de:08118:5700",
	BURGSTALL_UHINGEN:                                      "de:08117:2157",
	BURGSTALL_M_BURGSTALL_BURGSTETTEN:                      "de:08119:7500",
	BURGSTALLER_STRASSE_ERBSTETTEN:                         "de:08119:3471",
	BURGSTALLER_STRASSE_KIRCHBERG_M:                        "de:08119:3528",
	BURGSTRASSE_ECHTERDINGEN:                               "de:08116:2145",
	BURGSTRASSE_URBACH:                                     "de:08119:6698",
	BURGUNDERSTRASSE_BEUTELSBACH:                           "de:08119:3761",
	BURGUNDERSTRASSE_METTINGEN:                             "de:08116:4011",
	BURGUNDERWEG_HANWEILER:                                 "de:08119:5310",
	BURKHARDSHOF_BIRKMANNSWEILER:                           "de:08119:5992",
	BURKHARDTSMUHLE_WALDENBUCH:                             "de:08115:3151",
	BUSBAHNHOF_WELZHEIM:                                    "de:08119:5060",
	BUSHOF_ABZWEIG_SULZBACH_M:                              "de:08119:7509",
	BUSNAUER_PLATZ_STUTTGART:                               "de:08111:2590",
	BUSNAUER_STRASSE_41_WARMBRONN:                          "de:08115:4602",
	BUSNAUER_STRASSE_6_WARMBRONN:                           "de:08115:4601",
	BUSSBACHSTRASSE_STUTTGART:                              "de:08111:2432",
	BUTZBERG_GRAB:                                          "de:08119:7538",
	CAFE_ROMMEL_SCHWAIKHEIM:                                "de:08119:5418",
	CAFE_STOLL_OBERENSINGEN:                                "de:08116:2926",
	CALWER_BRUCKE_SINDELFINGEN:                             "de:08115:4624",
	CALWER_STRASSE_SINDELFINGEN:                            "de:08115:4611",
	CALWER_STRASSE_BOBLINGEN:                               "de:08115:4620",
	CALWER_STRASSE_OBERJESINGEN:                            "de:08115:4840",
	CALWER_STRASSE_HULB_BOBLINGEN:                          "de:08115:4621",
	CANNSTATTER_PLATZ_FELLBACH:                             "de:08119:6041",
	CANNSTATTER_STRASSE_METTINGEN:                          "de:08116:4010",
	CANNSTATTER_STRASSE_ALDINGEN:                           "de:08118:5588",
	CANNSTATTER_WASEN_STUTTGART:                            "de:08111:2343",
	CARLBENZSTRASSE_OTTMARSHEIM:                            "de:08118:6515",
	CARLBENZSTRASSE_LAICHINGEN:                             "de:08425:2585",
	CARLKAELBLESTRASSE_BACKNANG:                            "de:08119:3662",
	CARLSCHMINCKESTRASSE_ELTINGEN:                          "de:08115:6733",
	CARLZEISSSTRASSE_HARTHAUSEN:                            "de:08116:2081",
	CARLZEISSSTRASSE_GEBERSHEIM:                            "de:08115:5824",
	CARLZEISSSTRASSE_OTTMARSHEIM:                           "de:08118:6516",
	CHARLOTTENPLATZ_ESSLINGEN_N:                            "de:08116:3805",
	CHARLOTTENPLATZ_STUTTGART:                              "de:08111:6075",
	CHAUSSEEBERGSTRASSE_BESIGHEIM:                          "de:08118:7087",
	CHRISTKONIGSKIRCHE_GOPPINGEN:                           "de:08117:4910",
	CHRISTOFSHOF_WALDHAUSEN_GEISLINGEN:                     "de:08117:10214",
	CHRISTOFSTRASSE_WAIBLINGEN:                             "de:08119:3698",
	CHRISTOPHSBAD_GOPPINGEN:                                "de:08117:7002",
	CHRISTOPHSTRASSE_BACKNANG:                              "de:08119:7512",
	CHRISTOPHSTRASSE_GOPPINGEN:                             "de:08117:4003",
	COMBURGSTRASSE_LUDWIGSBURG:                             "de:08118:4893",
	CRONHUTTE_KAISERSBACH:                                  "de:08119:5280",
	DACHSWALD_STUTTGART:                                    "de:08111:2605",
	DACHSWALDWEG_STUTTGART:                                 "de:08111:2609",
	DAFERN_LIPPOLDSWEILER:                                  "de:08119:4980",
	DAIMLERPLATZ_STUTTGART:                                 "de:08111:323",
	DAIMLERSTEG_SINDELFINGEN:                               "de:08115:2210",
	DAIMLERSTRASSE_OEFFINGEN:                               "de:08119:2371",
	DAIMLERSTRASSE_SCHONAICH:                               "de:08115:3214",
	DAIMLERSTRASSE_MOGLINGEN:                               "de:08118:3433",
	DAIMLERSTRASSE_LUDWIGSBURG:                             "de:08118:5490",
	DAIMLERSTRASSE_HOLZGERLINGEN:                           "de:08115:6533",
	DAIMLERSTRASSE_HERRENBERG:                              "de:08115:7002",
	DAMMSTRASSE_NURTINGEN:                                  "de:08116:4459",
	DANZIGER_PLATZ_WAIBLINGEN:                              "de:08119:5389",
	DANZIGER_STRASSE_MUNCHINGEN:                            "de:08118:4521",
	DANZIGER_STRASSE_BOBLINGEN:                             "de:08115:4711",
	DANZIGER_STRASSE_LUDWIGSBURG:                           "de:08118:5533",
	DANZIGER_STRASSE_KIRCHBERG_M:                           "de:08119:6859",
	DANZIGER_STRASSE_EISLINGEN_F:                           "de:08117:1021",
	DANZIGER_WEG_SERSHEIM:                                  "de:08118:6701",
	DATZINGER_STRASSE_SCHAFHAUSEN:                          "de:08115:4053",
	DAUERNBERG_SPIEGELBERG:                                 "de:08119:4957",
	DAUERNBERG_ABZWEIG_SPIEGELBERG:                         "de:08119:4954",
	DEGERLOCH_STUTTGART:                                    "de:08111:6165",
	DEGERLOCH_ALBSTRASSE_STUTTGART:                         "de:08111:2594",
	DEGERLOCH_EPPLESTRASSE_STUTTGART:                       "de:08111:6184",
	DEGERLOCH_ZOB_STUTTGART:                                "de:08111:6163",
	DENKENDORFER_STRASSE_NELLINGEN:                         "de:08116:5005",
	DENTELTAL_SPIEGELBERG:                                  "de:08119:7580",
	DESSAUER_STRASSE_STUTTGART:                             "de:08111:2259",
	DETTENHAUSEN_DETTENHAUSEN:                              "de:08416:6349",
	DETTINGEN_T_DETTINGEN_T:                                "de:08116:5778",
	DEUTSCHES_ROTES_KREUZ_SECHSELBERG:                      "de:08119:4985",
	DEVIZESSTRASSE_WAIBLINGEN:                              "de:08119:3628",
	DIAKONIE_STETTEN_I_R:                                   "de:08119:4933",
	DIAKONISSENMUTTERHAUS_AIDLINGEN:                        "de:08115:4765",
	DIEB_HEIMSHEIM:                                         "de:08236:3418",
	DIEGELSBERG_UHINGEN:                                    "de:08117:9823",
	DIEPOLDSTRASSE_BERNHAUSEN:                              "de:08116:2077",
	DIESELSTRASSE_OEFFINGEN:                                "de:08119:2370",
	DIESELSTRASSE_RUTESHEIM:                                "de:08115:4556",
	DIESELSTRASSE_OPPENWEILER:                              "de:08119:7487",
	DIESELSTRASSE_GOPPINGEN:                                "de:08117:3302",
	DIETRICHBONHOEFFERSTRASSE_GERADSTETTEN:                 "de:08119:3753",
	DIEZENHALDE_ZENTRUM_BOBLINGEN:                          "de:08115:6751",
	DIEZENHALDE_ZENTRUM_SUD_BOBLINGEN:                      "de:08115:3139",
	DIEZENHALDENWEG_BOBLINGEN:                              "de:08115:4695",
	DILLMANNSTRASSE_STUTTGART:                              "de:08111:2194",
	DINGLESMAD_GSCHWEND:                                    "de:08136:6537",
	DINKELSTRASSE_KARLSHOF_STUTTGART:                       "de:08111:2572",
	DITZENBACHER_STR_AUENDORF:                              "de:08117:3227",
	DITZENBRUNNER_STRASSE_DITZINGEN:                        "de:08118:6493",
	DITZENBRUNNERKNIELSTRASSE_DITZINGEN:                    "de:08118:3015",
	DITZINGEN_DITZINGEN:                                    "de:08118:7000",
	DITZINGER_STRASSE_STUTTGART:                            "de:08111:2621",
	DLW_BIETIGHEIM:                                         "de:08118:5617",
	DOBELSTRASSE_STUTTGART:                                 "de:08111:159",
	DOFFINGER_STRASSE_DARMSHEIM:                            "de:08115:4195",
	DOGGENBURG_STUTTGART:                                   "de:08111:2190",
	DOLLENSTRASSE_ALFDORF:                                  "de:08119:6641",
	DOMO_SINDELFINGEN:                                      "de:08115:4609",
	DORFBRUNNEN_HAFNERHASLACH:                              "de:08118:5714",
	DORFGEMEINSCHAFT_TENNENTAL_DECKENPFRONN:                "de:08115:7072",
	DORFHALLE_STEINBACH:                                    "de:08119:5343",
	DORFPLATZ_RIET:                                         "de:08118:5748",
	DORFPLATZ_GUSSENSTADT:                                  "de:08135:1147",
	DORFSTRASSE_NECKARREMS:                                 "de:08118:4883",
	DORFSTRASSE_PFLUGFELDEN:                                "de:08118:5496",
	DORFSTRASSE_ROSSWALDEN:                                 "de:08117:7403",
	DORFWIESEN_SCHMIDEN:                                    "de:08119:2375",
	DORNHALDENSTRASSE_STUTTGART:                            "de:08111:6099",
	DORNIERSTRASSE_DITZINGEN:                               "de:08118:3036",
	DORNIERSTRASSE_BOBLINGEN:                               "de:08115:4721",
	DORNIERSTRASSE_SIRNAU:                                  "de:08116:6759",
	DORNROSCHENWEG_STUTTGART:                               "de:08111:6172",
	DOROTHEENSTRASSE_STUTTGART:                             "de:08111:6023",
	DOROTHEENWEG_WERNAU_N:                                  "de:08116:4244",
	DRHERBERTKONIGPLATZ_GOPPINGEN:                          "de:08117:1054",
	DRPFEIFFERSTRASSE_GOPPINGEN:                            "de:08117:9201",
	DREIECK_WEILER_EBERSBACH:                               "de:08117:7407",
	DRESCHERSTRASSE_RUTESHEIM:                              "de:08115:4559",
	DRESCHHALLE_GUNDELBACH:                                 "de:08118:5881",
	DRESDENER_RING_BACKNANG:                                "de:08119:3657",
	DRIEFBRUNNEN_PLATTENHARDT:                              "de:08116:2025",
	DRK_SINDELFINGEN:                                       "de:08115:4663",
	DROSSELWEG_SINDELFINGEN:                                "de:08115:2211",
	DROSSELWEG_FELLBACH:                                    "de:08119:2500",
	DROSSELWEG_DOFFINGEN:                                   "de:08115:6748",
	DUGENDORF_SALACH:                                       "de:08117:7607",
	DULKHAUSLE_WIFLINGSHAUSEN:                              "de:08116:4107",
	DURERSTRASSE_GOPPINGEN:                                 "de:08117:9010",
	DURERWEG_LORCH:                                         "de:08136:7011",
	DURNAUER_STRASSE_BAD_BOLL:                              "de:08117:3330",
	DURNAUER_WEG_STUTTGART:                                 "de:08111:2576",
	DURRMARBACHER_WEG_BISSINGEN_LB:                         "de:08118:3394",
	DURRROSENSTRASSE_BISSINGEN_LB:                          "de:08118:3388",
	DURRBACH_WEILER_SCHORNDORF:                             "de:08119:6796",
	DURRBACHSTRASSE_STUTTGART:                              "de:08111:2521",
	DURRLEWANG_STUTTGART:                                   "de:08111:2613",
	EBELSTRASSE_HOHENECK:                                   "de:08118:5526",
	EBENE_ESSLINGEN_N:                                      "de:08116:4110",
	EBENESCHULE_RECHBERGHAUSEN:                             "de:08117:4019",
	EBERDINGER_STRASSE_HOCHDORF_EBERDINGEN:                 "de:08118:5876",
	EBERHARDBAUERSTRASSE_PLIENSAUVORSTADT:                  "de:08116:4037",
	EBERSBACH_F_EBERSBACH_F:                                "de:08117:151",
	EBERSBACHER_STRASSE_SCHLIERBACH:                        "de:08117:7404",
	EBERSHALDENFRIEDHOF_ESSLINGEN_N:                        "de:08116:4032",
	EBITZWEG_STUTTGART:                                     "de:08111:1604",
	EBNI_KAISERSBACH:                                       "de:08119:5028",
	EBNISEE_KAISERSBACH:                                    "de:08119:5029",
	EBNISEESTRASSE_STUTTGART:                               "de:08111:89",
	EBNISEESTRASSE_ALTHUTTE:                                "de:08119:3729",
	EBNISEESTRASSE_OBERWEISSACH:                            "de:08119:4988",
	ECHTERDINGEN_ECHTERDINGEN:                              "de:08116:7003",
	ECKARTSWEILER_WELZHEIM:                                 "de:08119:5074",
	ECKENERSTRASSE_SIRNAU:                                  "de:08116:4286",
	ECKWALDEN_ORTSEINGANG_BAD_BOLL:                         "de:08117:2111",
	EDEKA_UNTERWEISSACH:                                    "de:08119:5288",
	EDELMANNSHOF_RUDERSBERG:                                "de:08119:7568",
	EDELWEISSWEG_STUTTGART:                                 "de:08111:2487",
	EDENHOF_LORCH:                                          "de:08136:7001",
	EGELSBERG_EGELSBERGSTRASSE_WEILHEIM_T:                  "de:08116:3899",
	EGELSBERG_HOCHHAUS_WEILHEIM_T:                          "de:08116:2817",
	EGELSBERG_TECKSTRASSE_WEILHEIM_T:                       "de:08116:3900",
	EGELSEE_RIELINGSHAUSEN:                                 "de:08118:5368",
	EGELSEE_HEIMSHEIM:                                      "de:08236:3507",
	EGERT_ZELL_ESSLINGEN:                                   "de:08116:3851",
	EHNINGEN_EHNINGEN:                                      "de:08115:5773",
	EHRENHALDE_STUTTGART:                                   "de:08111:2193",
	EIBENWEG_NURTINGEN:                                     "de:08116:2984",
	EICHBUHL_WANGEN_GP:                                     "de:08117:5011",
	EICHE_SCHONAICH:                                        "de:08115:4808",
	EICHENDORFFSCHULE_BOBLINGEN:                            "de:08115:4669",
	EICHENDORFFSCHULE_ZOLLBERG:                             "de:08116:4175",
	EICHENDORFFSTRASSE_NECKARTAILFINGEN:                    "de:08116:3938",
	EICHENDORFFSTRASSE_ZOLLBERG:                            "de:08116:4174",
	EICHENDORFFSTRASSE_KIRCHHEIM_T:                         "de:08116:4269",
	EICHENDORFFSTRASSE_NURTINGEN:                           "de:08116:4446",
	EICHENHOF_EISLINGEN_F:                                  "de:08117:6003",
	EICHENPFADLE_DAGERSHEIM:                                "de:08115:4655",
	EICHENWEG_KORNWESTHEIM:                                 "de:08118:5470",
	EICHENWEG_HERRENBERG:                                   "de:08115:7066",
	EICHENWEG_BISSINGEN_LB:                                 "de:08118:7426",
	EICHGRABEN_MARBACH_N:                                   "de:08118:3544",
	EICHHALDE_BAD_BOLL:                                     "de:08117:2109",
	EICHHOLZ_SINDELFINGEN:                                  "de:08115:4603",
	EICHWALD_BACKNANG:                                      "de:08119:6688",
	EICHWASEN_NECKARTENZLINGEN:                             "de:08116:4308",
	EIERHAUSLE_SCHWIEBERDINGEN:                             "de:08118:3359",
	EIERHOF_WELZHEIM:                                       "de:08119:5072",
	EIERNEST_STUTTGART:                                     "de:08111:6098",
	EINFELDSTRASSE_HOLZHAUSEN_UHINGEN:                      "de:08117:2040",
	EINKAUFSZENTRUM_GERLINGEN:                              "de:08118:4552",
	EINKAUFSZENTRUM_OBERJETTINGEN:                          "de:08115:4741",
	EINKAUFSZENTRUM_GOSBACH:                                "de:08117:3510",
	EINKAUFSZENTRUM_NAGOLD:                                 "de:08235:10122",
	EINSTEINSTRASSE_GULTSTEIN:                              "de:08115:3223",
	EINSTEINSTRASSE_KIRCHHEIM_T:                            "de:08116:3955",
	EINSTEINSTRASSE_SERSHEIM:                               "de:08118:6959",
	EINTRACHT_BACKNANG:                                     "de:08119:5337",
	EISBERG_L362_NAGOLD:                                    "de:08235:9350",
	EISENBAHNBRUCKE_GEISLINGEN_STEIGE:                      "de:08117:199",
	EISENBAHNSTRASSE_PLOCHINGEN:                            "de:08116:4227",
	EISENSCHMIEDMUHLE_MURRHARDT:                            "de:08119:4924",
	EISENTALSTRASSE_WAIBLINGEN:                             "de:08119:5956",
	EISGASSE_HEMMINGEN:                                     "de:08118:4570",
	EISLINGEN_F_EISLINGEN_F:                                "de:08117:155",
	EISLINGER_STRASSE_SALACH:                               "de:08117:6007",
	ELBENPLATZ_BOBLINGEN:                                   "de:08115:3105",
	ELBESTRASSE_STUTTGART:                                  "de:08111:275",
	ELISABETHKRANZSTRASSE_LUDWIGSBURG:                      "de:08118:3476",
	ELISABETHSELBERTGYMNASIUM_BERNHAUSEN:                   "de:08116:2049",
	ELLENTAL_BIETIGHEIM:                                    "de:08118:1910",
	ELLENTALGYMNASIEN_BIETIGHEIM:                           "de:08118:5647",
	ELLENWEILER_OPPENWEILER:                                "de:08119:4903",
	ELSABRANDSTROMSTRASSE_BOBLINGEN:                        "de:08115:4676",
	ELSABRANDSTROMSTRASSE_HOFINGEN:                         "de:08115:6817",
	ELSENHALDE_SCHONAICH:                                   "de:08115:4805",
	ELSERRING_BESIGHEIM:                                    "de:08118:3469",
	ELSTERSTAFFEL_STUTTGART:                                "de:08111:318",
	ELSTERWEG_DOFFINGEN:                                    "de:08115:6749",
	ELTINGER_STRASSE_SINDELFINGEN:                          "de:08115:2213",
	ELTINGER_STRASSE_STUTTGART:                             "de:08111:6210",
	ELWERTSTRASSE_STUTTGART:                                "de:08111:2358",
	EMERHOLZ_STUTTGART:                                     "de:08111:193",
	EMILMUNZSTRASSE_WAIBLINGEN:                             "de:08119:7094",
	ENBW_CITY_STUTTGART:                                    "de:08111:2586",
	ENDERSBACH_ENDERSBACH:                                  "de:08119:7704",
	ENDERSBACHER_STRASSE_STETTEN_I_R:                       "de:08119:4934",
	ENDERSBACHER_STRASSE_BEINSTEIN:                         "de:08119:5165",
	ENDWIESENSTRASSE_HOPFIGHEIM:                            "de:08118:3597",
	ENGELBERG_LEONBERG:                                     "de:08115:2321",
	ENGELBERG_WINTERBACH:                                   "de:08119:4082",
	ENGELBERG_WALDORFSCHULE_WINTERBACH:                     "de:08119:3774",
	ENGELBOLDSTRASSE_STUTTGART:                             "de:08111:8",
	ENSINGER_STRASSE_KLEINGLATTBACH:                        "de:08118:3353",
	ENTENACKER_BISSINGEN_LB:                                "de:08118:7427",
	ENZBACHWEG_SCHLAT:                                      "de:08117:3026",
	ENZBRUCKE_BIETIGHEIM:                                   "de:08118:5694",
	ENZBRUCKE_VAIHINGEN_E:                                  "de:08118:5751",
	ENZENHARDTPLATZ_NURTINGEN:                              "de:08116:4460",
	ENZPLATZ_BESIGHEIM:                                     "de:08118:5563",
	ENZSTRASSE_KORNWESTHEIM:                                "de:08118:5445",
	ENZSTRASSE_UNTERMBERG:                                  "de:08118:7477",
	ENZWEIHINGEN_B10_ENZWEIHINGEN:                          "de:08118:5744",
	EPPLERINWEG_SCHORNDORF:                                 "de:08119:5977",
	ERDBEERWEG_STUTTGART:                                   "de:08111:2439",
	ERDMANNHAUSEN_ERDMANNHAUSEN:                            "de:08118:3514",
	ERDMANNHAUSER_STRASSE_MARBACH_N:                        "de:08118:3532",
	ERGENZINGEN_ERGENZINGEN:                                "de:08416:2052",
	ERHOLUNGSHEIM_GULTSTEIN:                                "de:08115:4836",
	ERICHSCHUMMSTRASSE_MURRHARDT:                           "de:08119:4917",
	ERLACH_STETTEN_LEINFELDENECHT:                          "de:08116:2124",
	ERLACH_GROSSERLACH:                                     "de:08119:7514",
	ERLACH_ABZWEIG_GROSSERLACH:                             "de:08119:4962",
	ERLENHOF_ODERNHARDT:                                    "de:08119:5201",
	ERLENSTRASSE_SCHORNDORF:                                "de:08119:6693",
	ERLIGHEIM_ERLIGHEIM:                                    "de:08118:5695",
	ERLIGHEIMER_STRASSE_LOCHGAU:                            "de:08118:7423",
	ERNSTBAUERPLATZ_RENNINGEN:                              "de:08115:3318",
	ERNSTREUTERPLATZ_STUTTGART:                             "de:08111:6794",
	ERNSTSIGLEGYMNASIUM_KORNWESTHEIM:                       "de:08118:3424",
	ERWINSCHOETTLEPLATZ_STUTTGART:                          "de:08111:6017",
	ESCHENBACHER_STR_GAMMELSHAUSEN:                         "de:08117:3325",
	ESCHENBRUNNLESTRASSE_SINDELFINGEN:                      "de:08115:4652",
	ESCHENRIED_SINDELFINGEN:                                "de:08115:4605",
	ESCHENSTRUET_SULZBACH_M:                                "de:08119:7508",
	ESCHENWEG_SCHELMENHOLZ:                                 "de:08119:6661",
	ESELSHALDEN_WELZHEIM:                                   "de:08119:5054",
	ESELSMUHLE_MUSBERG:                                     "de:08116:2098",
	ESELWEG_STUTTGART:                                      "de:08111:2541",
	ESSEGGER_STRASSE_SINDELFINGEN:                          "de:08115:4645",
	ESSLINGEN_N_ESSLINGEN_N:                                "de:08116:7800",
	ESSLINGER_STRASSE_FELLBACH:                             "de:08119:38",
	ESSLINGER_STRASSE_WOLFSCHLUGEN:                         "de:08116:2918",
	ESSLINGER_STRASSE_DEIZISAU:                             "de:08116:4054",
	ESSLINGER_STRASSE_DENKENDORF:                           "de:08116:5007",
	ESSLINGER_STRASSE_PLOCHINGEN:                           "de:08116:5822",
	ESSLINGER_WEG_MAGSTADT:                                 "de:08115:3301",
	ESZET_STUTTGART:                                        "de:08111:267",
	ETTERSTRASSE_SCHLATTSTALL:                              "de:08116:5896",
	ETZELWEG_KOHLBERG:                                      "de:08116:2821",
	ETZWIESENBRUCKE_BACKNANG:                               "de:08119:5323",
	EUGENBOLZSTRASSE_ST_BERNHARDT:                          "de:08116:4131",
	EUGENBOLZSTRASSE_OST_BOBLINGEN:                         "de:08115:4703",
	EUGENSPLATZ_STUTTGART:                                  "de:08111:6121",
	EUGENSTRASSE_FELLBACH:                                  "de:08119:2645",
	EUGENSTRASSE_GOPPINGEN:                                 "de:08117:4907",
	EULENBERG_MUSBERG:                                      "de:08116:2107",
	EULENHOF_KAISERSBACH:                                   "de:08119:7530",
	EUROPAPLATZ_STUTTGART:                                  "de:08111:2584",
	EV_AKADEMIEREHAKLINIK_BAD_BOLL:                         "de:08117:3336",
	EV_KINDERGARTEN_GINGEN_F:                               "de:08117:87",
	EV_KINDERGARTEN_LORCH:                                  "de:08136:7012",
	EV_KIRCHE_SACHSENWEILER:                                "de:08119:5349",
	EVANGELISCHE_KIRCHE_SCHONAICH:                          "de:08115:3216",
	EVANGELISCHE_KIRCHE_UNTERJETTINGEN:                     "de:08115:4859",
	EVANGELISCHE_KIRCHE_STETTEN_I_R:                        "de:08119:4929",
	EWSARENALORCHER_STR_GOPPINGEN:                          "de:08117:4901",
	EWSARENANORDL_RINGSTR_GOPPINGEN:                        "de:08117:4006",
	EYBTALHALLE_EYBACH:                                     "de:08117:216",
	FA_ALLGAIER_UHINGEN:                                    "de:08117:2008",
	FABRIKSTRASSE_WENDLINGEN_N:                             "de:08116:3831",
	FABRIKSTRASSE_OTLINGEN:                                 "de:08116:4339",
	FABRIKSTRASSE_NEIDLINGEN:                               "de:08116:4406",
	FABRIKSTRASSE_BISSINGEN_AN_DER_TECK_T:                  "de:08116:4412",
	FABRIKSTRASSE_UHINGEN:                                  "de:08117:2036",
	FAHRGERGASSE_GEMMRIGHEIM:                               "de:08118:5720",
	FAHRZEUGWERKE_EISLINGEN_F:                              "de:08117:7602",
	FALBENHENNENSTRASSE_STUTTGART:                          "de:08111:2199",
	FALKENSTRASSE_JEBENHAUSEN:                              "de:08117:2098",
	FALKENWEG_EGLOSHEIM:                                    "de:08118:5513",
	FARBERSTRASSE_KUCHEN_F:                                 "de:08117:23",
	FASANENHOF_STUTTGART:                                   "de:08111:364",
	FASANENHOF_ABZWEIG_STUTTGART:                           "de:08111:6580",
	FASANENSTRASSE_JEBENHAUSEN:                             "de:08117:2153",
	FAULHABER_SCHONAICH:                                    "de:08115:3134",
	FAURNDAU_FAURNDAU:                                      "de:08117:153",
	FAUSLERSTRASSE_JESINGEN:                                "de:08116:4387",
	FAUSTSTRASSE_STUTTGART:                                 "de:08111:5",
	FAUTSPACH_SECHSELBERG:                                  "de:08119:5301",
	FAVORITEPARK_EGLOSHEIM:                                 "de:08118:7403",
	FEHRBELLINER_STR_STUTTGART:                             "de:08111:2629",
	FELDBERGSTRASSE_SINDELFINGEN:                           "de:08115:2212",
	FELDBERGSTRASSE_BOBLINGEN:                              "de:08115:4668",
	FELDWIESEN_BUNZWANGEN:                                  "de:08117:8002",
	FELLBACH_FELLBACH:                                      "de:08119:6500",
	FELLBACHER_STRASSE_ROMMELSHAUSEN:                       "de:08119:5431",
	FELLBACHER_STRASSE_OSSWEIL:                             "de:08118:5536",
	FERDINANDPORSCHESTR_NAGOLD:                             "de:08235:10092",
	FERNSEHTURM_STUTTGART:                                  "de:08111:2564",
	FESSLER_MUHLE_SERSHEIM:                                 "de:08118:6961",
	FESTHALLE_RUTESHEIM:                                    "de:08115:5828",
	FESTHALLE_HOLZHAUSEN_UHINGEN:                           "de:08117:2033",
	FESTPLATZ_MARKGRONINGEN:                                "de:08118:3581",
	FESTPLATZ_BALTMANNSWEILER:                              "de:08116:7822",
	FEUERBACH_STUTTGART:                                    "de:08111:6157",
	FEUERBACH_BOSCH_STUTTGART:                              "de:08111:6423",
	FEUERBACH_FRIEDHOF_STUTTGART:                           "de:08111:2422",
	FEUERBACH_PFOSTENWALDLE_STUTTGART:                      "de:08111:152",
	FEUERBACHER_STRASSE_LEONBERG:                           "de:08115:6055",
	FEUERBACHER_WEG_STUTTGART:                              "de:08111:2454",
	FEUERSEE_KAPPISHAUSERN:                                 "de:08116:4492",
	FEUERSEE_STUTTGART:                                     "de:08111:6221",
	FEUERWACHE_KUCHEN_F:                                    "de:08117:24",
	FEUERWACHE_GOPPINGEN:                                   "de:08117:4001",
	FEUERWEHR_GERLINGEN:                                    "de:08118:2299",
	FEUERWEHR_ALTBACH:                                      "de:08116:3817",
	FEUERWEHR_BACKNANG:                                     "de:08119:5336",
	FEUERWEHR_URBACH:                                       "de:08119:6696",
	FEUERWEHR_KIRCHBERG_M:                                  "de:08119:7096",
	FEUERWEHR_RECHBERGHAUSEN:                               "de:08117:4016",
	FEUERWEHRGERATEHAUS_GARTRINGEN:                         "de:08115:7052",
	FEUERWEHRHAUS_HORRHEIM:                                 "de:08118:5737",
	FEUERWEHRHAUS_ESCHENBACH:                               "de:08117:3318",
	FEUERWEHRMAGAZIN_ENSINGEN:                              "de:08118:5920",
	FICHTESTRASSE_STUTTGART:                                "de:08111:2403",
	FILDERAIRPORTAREAL_BERNHAUSEN:                          "de:08116:2157",
	FILDERBAHNSTRASSE_STUTTGART:                            "de:08111:316",
	FILDERBLICKWEG_STUTTGART:                               "de:08111:2542",
	FILDERKLINIK_BONLANDEN:                                 "de:08116:2016",
	FILDERSTADT_BERNHAUSEN:                                 "de:08116:1905",
	FILDORADO_BONLANDEN:                                    "de:08116:2017",
	FILHARMONIE_BERNHAUSEN:                                 "de:08116:2076",
	FILSBRUCKE_EBERSBACH_F:                                 "de:08117:7409",
	FILSECKSTRASSE_FAURNDAU:                                "de:08117:7010",
	FINANZAMT_WAIBLINGEN:                                   "de:08119:3687",
	FINANZAMT_ESSLINGEN_N:                                  "de:08116:4095",
	FINKENBERG_WAIBLINGEN:                                  "de:08119:3688",
	FINKENSTRASSE_PLATTENHARDT:                             "de:08116:2074",
	FINKENWEG_SIRNAU:                                       "de:08116:4052",
	FINKENWEG_ALBERSHAUSEN:                                 "de:08117:2019",
	FIRMA_BINZ_LORCH:                                       "de:08136:7051",
	FIRMA_FESTO_BERKHEIM:                                   "de:08116:4146",
	FIRMA_ZUGEL_WUSTENROT:                                  "de:08125:983",
	FISCHERPFAD_BIETIGHEIM:                                 "de:08118:7472",
	FISCHERWORTH_GROSSINGERSHEIM:                           "de:08118:3485",
	FISCHSTRASSE_GOPPINGEN:                                 "de:08117:2038",
	FLACHTER_STRASSE_STUTTGART:                             "de:08111:2623",
	FLACHTER_STRASSE_WEISSACH_BB:                           "de:08115:4568",
	FLANDERNSTRASSE_ST_BERNHARDT:                           "de:08116:4130",
	FLATTICHSCHULE_MUNCHINGEN:                              "de:08118:3568",
	FLEINSBACH_BERNHAUSEN:                                  "de:08116:2071",
	FLICKENWIESEN_GROSSBOTTWAR:                             "de:08118:6971",
	FLUGFELDALLEE_BOBLINGEN:                                "de:08115:3145",
	FLUGHAFENMESSE_ECHTERDINGEN:                            "de:08116:2103",
	FOHRENBUHL_BACKNANG:                                    "de:08119:5342",
	FOHRICH_STUTTGART:                                      "de:08111:154",
	FORCHENRAINSTRASSE_GERLINGEN:                           "de:08118:5917",
	FORCHENWALDSTRASSE_SCHELMENHOLZ:                        "de:08119:5309",
	FORELLENWEG_STUTTGART:                                  "de:08111:2481",
	FORNSBACH_FORNSBACH:                                    "de:08119:5787",
	FORNSBACHER_WEG_BACKNANG:                               "de:08119:6937",
	FORSCHUNG_DENKENDORF:                                   "de:08116:5014",
	FORSTBODEN_GROSSASPACH:                                 "de:08119:3674",
	FORSTHAUS_I_STUTTGART:                                  "de:08111:2416",
	FORSTHAUS_II_STUTTGART:                                 "de:08111:2342",
	FORSTHAUS_PARKPLATZ_STUTTGART:                          "de:08111:2425",
	FORSTSTRASSE_STETTEN_LEINFELDENECHT:                    "de:08116:2136",
	FORSTSTRASSE_KAISERSBACH:                               "de:08119:7542",
	FORUM_AM_SCHLOSSPARK_LUDWIGSBURG:                       "de:08118:7516",
	FRANK_LEINFELDEN:                                       "de:08116:7173",
	FRANKENSTRASSE_STUTTGART:                               "de:08111:2467",
	FRANKENSTRASSE_SINDELFINGEN:                            "de:08115:3206",
	FRANKENWEILER_GRAB:                                     "de:08119:5141",
	FRANKFURTER_STRASSE_EGLOSHEIM:                          "de:08118:3423",
	FRAUBRONNSTRASSE_STUTTGART:                             "de:08111:2553",
	FRAUENKOPF_STUTTGART:                                   "de:08111:2540",
	FRAUENKREUZ_LEONBERG:                                   "de:08115:5998",
	FRAUENLANDERSTRASSE_STETTEN_I_R:                        "de:08119:7974",
	FRAUENSTRASSE_GEISLINGEN_STEIGE:                        "de:08117:42",
	FREIBAD_KIRCHHEIM_T:                                    "de:08116:2814",
	FREIBAD_BOBLINGEN:                                      "de:08115:3107",
	FREIBAD_STETTEN_I_R:                                    "de:08119:3773",
	FREIBAD_WAIBLINGEN:                                     "de:08119:5404",
	FREIBAD_BONNIGHEIM:                                     "de:08118:5729",
	FREIBAD_GOPPINGEN:                                      "de:08117:1006",
	FREIBAD_UHINGEN:                                        "de:08117:2030",
	FREIBADFITKOM_BESIGHEIM:                                "de:08118:6963",
	FREIBERG_STUTTGART:                                     "de:08111:6287",
	FREIBERG_N_FREIBERG_N:                                  "de:08118:1503",
	FREIBERG_SCHULZENTRUM_STUTTGART:                        "de:08111:6286",
	FREIBERGER_STRASSE_BIETIGHEIM:                          "de:08118:3342",
	FREIBERGSTRASSE_STUTTGART:                              "de:08111:276",
	FREIBURGER_ALLEE_BOBLINGEN:                             "de:08115:4705",
	FREIHOFPLATZ_STUTTGART:                                 "de:08111:198",
	FREILICHTMUSEUM_BEUREN:                                 "de:08116:3939",
	FREITAGSHOF_WERNAU_N:                                   "de:08116:4369",
	FREIWALDAUSTRASSE_KIRCHHEIM_T:                          "de:08116:4329",
	FREIZEITPARK_KORNWESTHEIM:                              "de:08118:6810",
	FREIZEITZENTRUM_BUOCH:                                  "de:08119:5187",
	FREUDENSTADTER_STRASSE_BOBLINGEN:                       "de:08115:4704",
	FREUDENTALER_STRASSE_LOCHGAU:                           "de:08118:5705",
	FREYTAGWEG_STUTTGART:                                   "de:08111:361",
	FRICKENHAUSEN_FRICKENHAUSEN:                            "de:08116:5793",
	FRICKENHAUSER_STRASSE_TISCHARDT:                        "de:08116:7805",
	FRIDINGER_STRASSE_STUTTGART:                            "de:08111:2436",
	FRIEDENSKIRCHE_URBACH:                                  "de:08119:6682",
	FRIEDENSSCHULE_NEUSTADT:                                "de:08119:5149",
	FRIEDENSTRASSE_STUTTGART:                               "de:08111:2057",
	FRIEDENSTRASSE_5_LUDWIGSBURG:                           "de:08118:5488",
	FRIEDENSTRASSE_52_LUDWIGSBURG:                          "de:08118:5489",
	FRIEDHOF_MUSBERG:                                       "de:08116:310",
	FRIEDHOF_SIELMINGEN:                                    "de:08116:2045",
	FRIEDHOF_ECHTERDINGEN:                                  "de:08116:2102",
	FRIEDHOF_RENNINGEN:                                     "de:08115:3319",
	FRIEDHOF_BACKNANG:                                      "de:08119:3607",
	FRIEDHOF_ROMMELSHAUSEN:                                 "de:08119:3737",
	FRIEDHOF_ZELL_ESSLINGEN:                                "de:08116:4018",
	FRIEDHOF_OHMDEN:                                        "de:08116:4390",
	FRIEDHOF_ZIZISHAUSEN:                                   "de:08116:4454",
	FRIEDHOF_MUNCHINGEN:                                    "de:08118:4560",
	FRIEDHOF_WEIL_DER_STADT:                                "de:08115:4586",
	FRIEDHOF_ALTDORF_BB:                                    "de:08115:4794",
	FRIEDHOF_WELZHEIM:                                      "de:08119:5129",
	FRIEDHOF_WINNENDEN:                                     "de:08119:5207",
	FRIEDHOF_UNTERBRUDEN:                                   "de:08119:5381",
	FRIEDHOF_KORNWESTHEIM:                                  "de:08118:5447",
	FRIEDHOF_POPPENWEILER:                                  "de:08118:5558",
	FRIEDHOF_BESIGHEIM:                                     "de:08118:5718",
	FRIEDHOF_GROSSINGERSHEIM:                               "de:08118:5933",
	FRIEDHOF_WAIBLINGEN:                                    "de:08119:5958",
	FRIEDHOF_NUFRINGEN:                                     "de:08115:6480",
	FRIEDHOF_THOMASHARDT:                                   "de:08116:6494",
	FRIEDHOF_HOLZGERLINGEN:                                 "de:08115:6529",
	FRIEDHOF_NECKARREMS:                                    "de:08118:6825",
	FRIEDHOF_WENDLINGEN_N:                                  "de:08116:6865",
	FRIEDHOF_GRAB:                                          "de:08119:6996",
	FRIEDHOF_BITTENFELD:                                    "de:08119:7551",
	FRIEDHOF_ROSSWAG:                                       "de:08118:7841",
	FRIEDHOF_DEGGINGEN:                                     "de:08117:53",
	FRIEDHOF_GEISLINGEN_STEIGE:                             "de:08117:96",
	FRIEDHOF_BOHMENKIRCH:                                   "de:08117:201",
	FRIEDHOF_WASCHENBEUREN:                                 "de:08117:1109",
	FRIEDHOF_HATTENHOFEN:                                   "de:08117:2054",
	FRIEDHOF_HEININGEN_GP:                                  "de:08117:3312",
	FRIEDHOF_BORTLINGEN:                                    "de:08117:4034",
	FRIEDHOF_NORD_EISLINGEN_F:                              "de:08117:1028",
	FRIEDHOF_NORD_GOPPINGEN:                                "de:08117:9012",
	FRIEDHOF_SANKT_PETER_BIETIGHEIM:                        "de:08118:3333",
	FRIEDHOF_SUD_EISLINGEN_F:                               "de:08117:1044",
	FRIEDHOF_SUD_GOPPINGEN:                                 "de:08117:9301",
	FRIEDHOFSTRASSE_STUTTGART:                              "de:08111:303",
	FRIEDRICHBREININGSTRASSE_BESIGHEIM:                     "de:08118:5716",
	FRIEDRICHEBERTSTRASSE_SINDELFINGEN:                     "de:08115:4604",
	FRIEDRICHEBERTSTRASSE_BIETIGHEIM:                       "de:08118:7440",
	FRIEDRICHLISTPLATZ_BOBLINGEN:                           "de:08115:3112",
	FRIEDRICHLISTSTRASSE_FELLBACH:                          "de:08119:2496",
	FRIEDRICHLISTSTRASSE_GROSSASPACH:                       "de:08119:3673",
	FRIEDRICHLISTSTRASSE_WERNAU_N:                          "de:08116:4247",
	FRIEDRICHSCHILLERGYMNASIUM_FELLBACH:                    "de:08119:5436",
	FRIEDRICHSCHILLERSCHULE_NEUHAUSEN_F:                    "de:08116:4191",
	FRIEDRICHSTRASSE_SIELMINGEN:                            "de:08116:2031",
	FRIEDRICHSTRASSE_SCHMIDEN:                              "de:08119:2379",
	FRIEDRICHSTRASSE_LUDWIGSBURG:                           "de:08118:7436",
	FRIEDRICHSTRASSE_GOPPINGEN:                             "de:08117:4002",
	FRIEDRICHSWAHL_STUTTGART:                               "de:08111:110",
	FRIESENSTRASSE_OSSWEIL:                                 "de:08118:5515",
	FRITZWALTERWEG_STUTTGART:                               "de:08111:6328",
	FROBELSCHULE_HERRENBERG:                                "de:08115:7055",
	FROBELSTRASSE_BERNHAUSEN:                               "de:08116:2012",
	FROBELSTRASSE_WAIBLINGEN:                               "de:08119:3652",
	FRONACKERSTRASSE_WAIBLINGEN:                            "de:08119:4875",
	FRONACKERSTRASSE_SINDELFINGEN:                          "de:08115:3203",
	FRONTALSTRASSE_DECKENPFRONN:                            "de:08115:7049",
	FROSCHEGERT_GROTZINGEN:                                 "de:08116:2135",
	FRUHLINGSTRASSE_ALBERSHAUSEN:                           "de:08117:2024",
	FRUHLINGSTRASSE_ESCHENBACH:                             "de:08117:3314",
	FRUHMESSHOF_KIRCHBERG_M:                                "de:08119:5366",
	FRUWIRTHSTRASSE_STUTTGART:                              "de:08111:2556",
	FUCHSBAUEINKAUFSMARKT_DITZINGEN:                        "de:08118:3034",
	FUCHSECKSATTEL_AUENDORF:                                "de:08117:3225",
	FUCHSECKSTRASSE_GOPPINGEN:                              "de:08117:3004",
	FUCHSECKSTRASSE_SUSSEN:                                 "de:08117:7613",
	FUCHSGRUBE_WAIBLINGEN:                                  "de:08119:5407",
	FUCHSLOCHER_DONNSTETTEN:                                "de:08415:28678",
	FUCHSRAIN_STUTTGART:                                    "de:08111:2188",
	FUCHSRAIN_HARDT:                                        "de:08116:2949",
	FULLERSTRASSE_GERLINGEN:                                "de:08118:2304",
	FURFELDER_STRASSE_STUTTGART:                            "de:08111:6292",
	FURSTENHOF_GROSSASPACH:                                 "de:08119:5364",
	FURTBACHSCHULE_MOGLINGEN:                               "de:08118:5861",
	FURTHMUHLE_AIDLINGEN:                                   "de:08115:4766",
	GABLENBERG_STUTTGART:                                   "de:08111:2200",
	GAILDORFER_STRASSE_BACKNANG:                            "de:08119:4949",
	GAISBURG_STUTTGART:                                     "de:08111:79",
	GAISERPLATZ_KIRCHHEIM_T:                                "de:08116:3958",
	GALERIE_WAIBLINGEN:                                     "de:08119:3654",
	GALGENBERG_WAIBLINGEN:                                  "de:08119:3645",
	GALGENBERGBRUCKE_WAIBLINGEN:                            "de:08119:5163",
	GAMMELSHAUSER_STR_DURNAU:                               "de:08117:3327",
	GAMP_ROHRDORF_CW:                                       "de:08235:10166",
	GANSBUHL_SCHORNDORF:                                    "de:08119:7632",
	GANSFUSSALLEE_LUDWIGSBURG:                              "de:08118:3411",
	GANSSEE_BOBLINGEN:                                      "de:08115:3142",
	GANSWASEN_PLUDERHAUSEN:                                 "de:08119:5235",
	GARTENFREUNDE_GOPPINGEN:                                "de:08117:1064",
	GARTENSTADT_LEONBERG:                                   "de:08115:2322",
	GARTENSTADT_OBERESSLINGEN:                              "de:08116:4024",
	GARTENSTRASSE_FELLBACH:                                 "de:08119:2515",
	GARTENSTRASSE_BACKNANG:                                 "de:08119:3660",
	GARTENSTRASSE_NABERN:                                   "de:08116:4411",
	GARTENSTRASSE_BEUREN:                                   "de:08116:4437",
	GARTENSTRASSE_HESSIGHEIM:                               "de:08118:5576",
	GARTENSTRASSE_HOCHDORF_EBERDINGEN:                      "de:08118:6949",
	GARTENSTRASSE_LUDWIGSBURG:                              "de:08118:7563",
	GARTENSTRASSE_GECHINGEN:                                "de:08235:2540",
	GARTRINGEN_GARTRINGEN:                                  "de:08115:5774",
	GARTRINGER_STRASSE_NUFRINGEN:                           "de:08115:6479",
	GASTHAUS_KRONE_OTTENBACH:                               "de:08117:6017",
	GAUFELDEN_GAUFELDEN:                                    "de:08115:5776",
	GAUSMANNSWEILER_WELZHEIM:                               "de:08119:5075",
	GAUSSSTRASSE_STUTTGART:                                 "de:08111:2402",
	GAUSSWEG_GOPPINGEN:                                     "de:08117:9004",
	GEBENWEILER_KAISERSBACH:                                "de:08119:5067",
	GEBERSHEIMER_STRASSE_RUTESHEIM:                         "de:08115:5831",
	GEBERSHEIMERROSEGGERSTRASSE_ELTINGEN:                   "de:08115:4548",
	GEGEN_EICH_OSSWEIL:                                     "de:08118:7454",
	GEHRENWALD_STUTTGART:                                   "de:08111:5868",
	GEIBELSTRASSE_STUTTGART:                                "de:08111:2408",
	GEISINGER_DORFPLATZ_FREIBERG_N:                         "de:08118:3522",
	GEISLINGEN_STEIGE_GEISLINGEN_STEIGE:                    "de:08117:161",
	GEISLINGEN_WEST_GEISLINGEN_STEIGE:                      "de:08117:160",
	GEISLINGER_STRASSE_LAICHINGEN:                          "de:08425:2660",
	GEISLINGER_STRASSE_21_ELTINGEN:                         "de:08115:4562",
	GEISLINGER_STRASSE_3_ELTINGEN:                          "de:08115:4563",
	GEISLINGER_STRASSE_51_ELTINGEN:                         "de:08115:4561",
	GEISSBERG_BIRENBACH:                                    "de:08117:1101",
	GEISSHALDE_DEUFRINGEN:                                  "de:08115:4478",
	GELEENER_STRASSE_BOBLINGEN:                             "de:08115:4699",
	GELEENER_STRASSEKINDERGARTEN_BOBLINGEN:                 "de:08115:4698",
	GEMEINDEHALLE_KIRCHHEIM_N:                              "de:08118:3682",
	GEMEINDEHALLE_ALTBACH:                                  "de:08116:3804",
	GEMEINDEHALLE_KIRCHENKIRNBERG:                          "de:08119:5081",
	GEMEINDEHALLE_HOFEN:                                    "de:08119:5273",
	GEMEINDEHAUS_BOBLINGEN:                                 "de:08115:4687",
	GEMEINDEHAUS_BUOCH:                                     "de:08119:5186",
	GEMEINDEHAUS_ESCHENBACH:                                "de:08117:3315",
	GEMEINDELANDERWEG_ALBERSHAUSEN:                         "de:08117:2021",
	GEMEINDEZENTRUM_OEFFINGEN:                              "de:08119:2381",
	GEMEINDEZENTRUM_BERKHEIM:                               "de:08116:4148",
	GEMEINDEZENTRUM_OBERJETTINGEN:                          "de:08115:4860",
	GEMEINDEZENTRUM_KUCHEN_F:                               "de:08117:15",
	GEMEINSCHAFTSSCHULE_BAD_BOLL:                           "de:08117:3332",
	GENERATIONENHAUS_STUTTGART:                             "de:08111:207",
	GENERATORSTRASSE_STUTTGART:                             "de:08111:2346",
	GEORGELSERWEG_SUSSEN:                                   "de:08117:7600",
	GEORGIIWALDSTADION_LIEBERSBRONN:                        "de:08116:4106",
	GERADSTETTEN_GERADSTETTEN:                              "de:08119:1702",
	GERBERSTRASSE_BACKNANG:                                 "de:08119:5335",
	GERIATRISCHES_ZENTRUM_KENNENBURG:                       "de:08116:4097",
	GERLINGEN_GERLINGEN:                                    "de:08118:7140",
	GERLINGER_STRASSE_RAMTEL:                               "de:08115:2338",
	GERLINGER_STRASSE_STUTTGART:                            "de:08111:2620",
	GERLINGER_TOR_GERLINGEN:                                "de:08118:2314",
	GERMANENSTRASSE_HOLZGERLINGEN:                          "de:08115:4880",
	GEROKSRUHE_STUTTGART:                                   "de:08111:126",
	GEROKSTRASSE_BISSINGEN_LB:                              "de:08118:7406",
	GEROKWEG_KUCHEN_F:                                      "de:08117:25",
	GERTRUDBAUMERWEG_BACKNANG:                              "de:08119:3675",
	GESCHICHTSHAUS_OWEN:                                    "de:08116:4280",
	GESCHWISTERSCHOLLGYMNASIUM_STUTTGART:                   "de:08111:6135",
	GESTUTSHOF_SCHARNHAUSEN:                                "de:08116:2915",
	GESTUTSWEG_WEIL:                                        "de:08116:4016",
	GESUNDHEITSZENTRUM_BACKNANG:                            "de:08119:4971",
	GEWERBE_TURKHEIM:                                       "de:08117:275",
	GEWERBEGEBIET_STETTEN_LEINFELDENECHT:                   "de:08116:2156",
	GEWERBEGEBIET_ALFDORF:                                  "de:08119:3256",
	GEWERBEGEBIET_FREUDENTAL:                               "de:08118:5708",
	GEWERBEGEBIET_NUFRINGEN:                                "de:08115:6483",
	GEWERBEGEBIET_THOMASHARDT:                              "de:08116:6523",
	GEWERBEGEBIET_NELLMERSBACH:                             "de:08119:7445",
	GEWERBEGEBIET_SPARWIESEN:                               "de:08117:5020",
	GEWERBEGEBIET_MERKLINGEN:                               "de:08425:2635",
	GEWERBEGEBIET_STEIGE_RUTESHEIM:                         "de:08115:5833",
	GEWERBEGEBIET_SUD_HEIMERDINGEN:                         "de:08118:3334",
	GEWERBEGEBIET_SUD_DITZINGEN:                            "de:08118:7001",
	GEWERBEPARK_ESCHENBACH:                                 "de:08117:3319",
	GEWERBESTRASSE_STUTTGART:                               "de:08111:2699",
	GEWERBESTRASSE_AKADEMIE_WAIBLINGEN:                     "de:08119:5924",
	GIEBEL_STUTTGART:                                       "de:08111:144",
	GIESHOF_SPIEGELBERG:                                    "de:08119:4960",
	GINGEN_F_GINGEN_F:                                      "de:08117:158",
	GISELASTRASSE_WAIBLINGEN:                               "de:08119:3655",
	GLASHUTTE_WALDENBUCH:                                   "de:08115:3120",
	GLASPALAST_SINDELFINGEN:                                "de:08115:2240",
	GLAUNERWEG_STUTTGART:                                   "de:08111:2537",
	GLEMSAUE_DITZINGEN:                                     "de:08118:3001",
	GLEMSBRUCKE_UNTERRIEXINGEN:                             "de:08118:5637",
	GLEMSECK_LEONBERG:                                      "de:08115:2324",
	GLEMSECKSTRASSE_ELTINGEN:                               "de:08115:2340",
	GLEMSTAL_GERLINGEN:                                     "de:08118:2305",
	GLEMSTAL_SCHWIEBERDINGEN:                               "de:08118:3322",
	GLOCKENSTRASSE_STUTTGART:                               "de:08111:263",
	GMEINWEILER_KAISERSBACH:                                "de:08119:5066",
	GNESENER_STRASSE_STUTTGART:                             "de:08111:321",
	GOCKELHOF_KIRCHENKIRNBERG:                              "de:08119:5985",
	GOCKELHOF_KREUZUNG_FORNSBACH:                           "de:08119:5086",
	GOERDELERSTRASSE_KORNWESTHEIM:                          "de:08118:5443",
	GOETHEHOLDERLINSTR_MOTZINGEN:                           "de:08115:7018",
	GOETHESTR_FAURNDAU:                                     "de:08117:2005",
	GOETHESTRASSE_GERLINGEN:                                "de:08118:149",
	GOETHESTRASSE_WENDLINGEN_N:                             "de:08116:3826",
	GOETHESTRASSE_DAGERSHEIM:                               "de:08115:4635",
	GOETHESTRASSE_KOHLBERG:                                 "de:08116:6703",
	GOETHESTRASSE_ALDINGEN:                                 "de:08118:6829",
	GOETHESTRASSE_MOTZINGEN:                                "de:08115:7019",
	GOLDACKER_STEINENBRONN:                                 "de:08115:7111",
	GOLDBERG_BOBLINGEN:                                     "de:08115:3212",
	GOLDBERG_REALSCHULE_SINDELFINGEN:                       "de:08115:4636",
	GOLDBERG_WASSERTURM_SINDELFINGEN:                       "de:08115:4614",
	GOLDBERGGYMNASIUM_SINDELFINGEN:                         "de:08115:3204",
	GOLDBODEN_WINTERBACH:                                   "de:08119:4080",
	GOLDMORGEN_DETTINGEN_T:                                 "de:08116:4278",
	GOLDMUHLESTRASSE_SINDELFINGEN:                          "de:08115:4657",
	GOLFPARK_GOPPINGEN:                                     "de:08117:9028",
	GOLFPLATZ_LEONBERG:                                     "de:08115:2325",
	GOLFPLATZ_EGLOSHEIM:                                    "de:08118:5508",
	GOLLENHOFER_STRASSE_WEILER_ZUM_STEIN:                   "de:08119:6768",
	GOPPINGEN_GOPPINGEN:                                    "de:08117:154",
	GOPPINGER_STRASSE_RAMTEL:                               "de:08115:2339",
	GOPPINGER_STRASSE_ZELL_A:                               "de:08117:2117",
	GOTTLIEBDAIMLERSCHULE_SINDELFINGEN:                     "de:08115:3205",
	GOTTLIEBDAIMLERSTRASSE_NAGOLD:                          "de:08235:10093",
	GOTTLOBBAUKNECHTPLATZ_WELZHEIM:                         "de:08119:5063",
	GOTTLOBKAMMPLATZ_SCHORNDORF:                            "de:08119:5970",
	GOTTSCHEER_STRASSE_SINDELFINGEN:                        "de:08115:4659",
	GRABENSTR_STADTM_VAIHINGEN_E:                           "de:08118:5631",
	GRABENSTRASSE_MARBACH_N:                                "de:08118:3543",
	GRABENSTRASSE_WEIL_DER_STADT:                           "de:08115:4581",
	GRABENSTRASSE_PLUDERHAUSEN:                             "de:08119:5988",
	GRABENSTRASSE_GARTRINGEN:                               "de:08115:7073",
	GRABENSTRASSE_BAHNHOF_SCHORNDORF:                       "de:08119:5224",
	GRAFDEGENFELDSTRASSE_RECHBERGHAUSEN:                    "de:08117:4024",
	GRAFENBERGER_STRASSE_KLEINBETTLINGEN:                   "de:08116:4464",
	GRAFENBERGWEG_SCHORNDORF:                               "de:08119:6994",
	GRAFENWEG_HERRENBERG:                                   "de:08115:7006",
	GRAIRICH_KAISERSBACH:                                   "de:08119:5068",
	GRAUBACHTAL_HATTENHOFEN:                                "de:08117:2059",
	GRAUHALDE_SCHORNDORF:                                   "de:08119:5959",
	GRAUHALDENHOF_RUDERSBERG:                               "de:08119:4999",
	GREUT_STUTTGART:                                        "de:08111:195",
	GREUTACKERSTRASSE_MONCHBERG:                            "de:08115:7023",
	GREUTHOFLE_VORDERSTEINENBERG:                           "de:08119:6669",
	GREUTTERSTRASSE_STUTTGART:                              "de:08111:2630",
	GRIEBENACKER_PLATTENHARDT:                              "de:08116:6727",
	GROENERSTRASSE_LUDWIGSBURG:                             "de:08118:5474",
	GRONINGER_STRASSEBAUMHALDE_DITZINGEN:                   "de:08118:3011",
	GRONINGERHALDENSTRASSE_DITZINGEN:                       "de:08118:3012",
	GRORACH_NEUENHAUS:                                      "de:08116:4476",
	GROSSE_GASSE_GINGEN_F:                                  "de:08117:22",
	GROSSER_FORST_NURTINGEN:                                "de:08116:2991",
	GROSSGLOCKNERSTRASSE_STUTTGART:                         "de:08111:86",
	GROSSHOCHBERG_SPIEGELBERG:                              "de:08119:5045",
	GROSSHOCHBERGER_STRASSE_SPIEGELBERG:                    "de:08119:4961",
	GROSSSACHSENHEIMER_STRASSE_KLEINSACHSENHEIM:            "de:08118:5671",
	GROTZINGER_STRASSE_OBERENSINGEN:                        "de:08116:2924",
	GROTZSTRASSE_BISSINGEN_LB:                              "de:08118:5686",
	GROTZSTRASSE_FRIEDHOF_BISSINGEN_LB:                     "de:08118:6717",
	GRUIBINGER_STRASSE_BAD_BOLL:                            "de:08117:2108",
	GRUNACKERSTRASSE_SINDELFINGEN:                          "de:08115:5921",
	GRUNBACH_GRUNBACH:                                      "de:08119:1703",
	GRUNBACH_DONZDORF:                                      "de:08117:7644",
	GRUND_UND_HAUPTSCH_GOLDBERG_SINDELFINGEN:               "de:08115:3236",
	GRUNDGENSSTRASSE_STUTTGART:                             "de:08111:6203",
	GRUNDSCHULE_GRUNBACH:                                   "de:08119:5183",
	GRUNDSCHULE_KLEINGLATTBACH:                             "de:08118:5653",
	GRUNDSCHULE_SPARWIESEN:                                 "de:08117:121",
	GRUNDSCHULE_GRUIBINGEN:                                 "de:08117:3233",
	GRUNDSCHULE_ESCHENBACH:                                 "de:08117:3317",
	GRUNDSCHULE_BAD_DITZENBACH:                             "de:08117:3505",
	GRUNDSCHULE_RECHBERGHAUSEN:                             "de:08117:4018",
	GRUNENBERGSTRASSE_GUTENBERG:                            "de:08116:4420",
	GRUNLINGWEG_STUTTGART:                                  "de:08111:2562",
	GRUNWIESENSTRASSE_BIETIGHEIM:                           "de:08118:7441",
	GSCHWENDER_STRASSE_WELZHEIM:                            "de:08119:6686",
	GULTSTEIN_GULTSTEIN:                                    "de:08115:3231",
	GUNTTERSTRASSE_MARBACH_N:                               "de:08118:5372",
	GUSTAVRAUSTRASSE_BIETIGHEIM:                            "de:08118:3344",
	GUSTAVWERNERSCHULE_WALDDORF:                            "de:08415:29044",
	GUTENBERGER_STRASSE_DONNSTETTEN:                        "de:08415:28646",
	GUTENBERGER_STRASSE_17_OBERLENNINGEN:                   "de:08116:5897",
	GUTENBERGER_STRASSE_47_OBERLENNINGEN:                   "de:08116:7589",
	GUTENBERGER_STRASSE_66_OBERLENNINGEN:                   "de:08116:7583",
	GUTENBERGSTRASSE_SCHMIDEN:                              "de:08119:5427",
	GUTENHALDE_BONLANDEN:                                   "de:08116:2018",
	GUTTENBRUNNSTRASSE_SINDELFINGEN:                        "de:08115:4644",
	GYMNASIUM_ASPERG:                                       "de:08118:3400",
	GYMNASIUM_MURRHARDT:                                    "de:08119:4918",
	GYMNASIUM_WEIL_DER_STADT:                               "de:08115:6477",
	GYMNASIUM_HOLZGERLINGEN:                                "de:08115:6506",
	GYMNASIUM_ECHTERDINGEN:                                 "de:08116:7185",
	GYMNASIUM_KORNTAL:                                      "de:08118:7606",
	HABICHTWEG_KIRCHHEIM_T:                                 "de:08116:4351",
	HAFENBAHNSTRASSE_STUTTGART:                             "de:08111:3997",
	HAFNERSTRASSE_NEUENHAUS:                                "de:08116:4477",
	HAGENBUCH_DONZDORF:                                     "de:08117:7626",
	HAGER_SULZBACH_M:                                       "de:08119:7510",
	HAGHOF_PFAHLBRONN:                                      "de:08119:5089",
	HAGWEG_UNTERKIRNECK:                                    "de:08136:7028",
	HAHNWEIDSTRASSE_KIRCHHEIM_T:                            "de:08116:3960",
	HAIER_HAIERSCHULE_FAURNDAU:                             "de:08117:5006",
	HAIER_WALDORFSCHULE_FAURNDAU:                           "de:08117:5005",
	HAIGST_STUTTGART:                                       "de:08111:165",
	HAILFINGER_STRASSE_BONDORF:                             "de:08115:7034",
	HAILINGSTRASSE_GOPPINGEN:                               "de:08117:1051",
	HAINBACH_WALDENBRONN:                                   "de:08116:2788",
	HALDE_OTLINGEN:                                         "de:08116:4343",
	HALDE_NECKARHAUSEN:                                     "de:08116:4481",
	HALDE_HEMMINGEN:                                        "de:08118:7460",
	HALDENRAINWEG_ALTBACH:                                  "de:08116:3818",
	HALDENSTRASSE_OTLINGEN:                                 "de:08116:3959",
	HALLENBAD_WALDENBUCH:                                   "de:08115:3122",
	HALLENBAD_MAICHINGEN:                                   "de:08115:3208",
	HALLENBAD_HERRENBERG:                                   "de:08115:3233",
	HALLENBAD_BOBLINGEN:                                    "de:08115:4692",
	HALLENBAD_KORNWESTHEIM:                                 "de:08118:5440",
	HALLENBAD_BISSINGEN_LB:                                 "de:08118:7430",
	HALLENBAD_EISLINGEN_F:                                  "de:08117:1013",
	HALLSCHLAG_STUTTGART:                                   "de:08111:2473",
	HAMMERSCHLAG_SCHORNDORF:                                "de:08119:5804",
	HAMMERSCHMIEDE_MURRHARDT:                               "de:08119:5295",
	HANDELSTRASSE_STUTTGART:                                "de:08111:2419",
	HANDWERKSTRASSE_STUTTGART:                              "de:08111:2700",
	HANFSTRASSE_DETTINGEN_T:                                "de:08116:4273",
	HANGELSTEIN_KIRCHE_ZELL_ESSLINGEN:                      "de:08116:4020",
	HANGLICH_WANGEN_GP:                                     "de:08117:5007",
	HANGWEIDE_ROMMELSHAUSEN:                                "de:08119:4935",
	HANNSKLEMMSTRASSE_OST_BOBLINGEN:                        "de:08115:4724",
	HANNSKLEMMSTRASSE_WEST_BOBLINGEN:                       "de:08115:4719",
	HANSBRUMMERPLATZ_LEINFELDEN:                            "de:08116:2100",
	HANSREHNSTIFT_STUTTGART:                                "de:08111:2610",
	HANSSEYFFSTRASSE_GOPPINGEN:                             "de:08117:74",
	HANSTHOMAWEG_KIRCHHEIM_T:                               "de:08116:4371",
	HANSVONHUTTENPLATZ_HARDT:                               "de:08116:2950",
	HARBACH_MURRHARDT:                                      "de:08119:5984",
	HARDT_HARDT:                                            "de:08116:2923",
	HARDTHOF_SCHWIEBERDINGEN:                               "de:08118:5626",
	HARDTLINDE_MURR:                                        "de:08118:5600",
	HARDTSCHULE_EBERSBACH_F:                                "de:08117:7411",
	HARDTWALDHALLE_KLEINASPACH:                             "de:08119:6651",
	HARFENBRUCKE_SINDELFINGEN:                              "de:08115:3242",
	HARTENECKSTRASSE_FREIBERG_N:                            "de:08118:3525",
	HARTENECKSTRASSE_LUDWIGSBURG:                           "de:08118:5500",
	HARTHAUSER_WEG_SIELMINGEN:                              "de:08116:2032",
	HARTWEG_BESIGHEIM:                                      "de:08118:5730",
	HASELBACHMUHLE_ABZWEIG_SULZBACH_M:                      "de:08119:7488",
	HASELSTEINSTRASSE_BREUNINGSWEILER:                      "de:08119:5422",
	HASENHALDE_BACKNANG:                                    "de:08119:3668",
	HASENHEIM_HARTHAUSEN:                                   "de:08116:2086",
	HASENHOF_WALDENBUCH:                                    "de:08115:3131",
	HASLACH_DARMSHEIM:                                      "de:08115:7836",
	HATTENHOFER_STRASSE_SCHLIERBACH:                        "de:08117:7802",
	HAUBENWASEN_PFAHLBRONN:                                 "de:08119:6968",
	HAUBERSBRONN_HAUBERSBRONN:                              "de:08119:5051",
	HAUFFSTRASSE_KORNTAL:                                   "de:08118:2632",
	HAUFFSTRASSE_OBERENSINGEN:                              "de:08116:2952",
	HAUFFSTRASSE_KORNWESTHEIM:                              "de:08118:3414",
	HAUFFSTRASSE_RECHBERGHAUSEN:                            "de:08117:1201",
	HAUGENNEST_OBERBOIHINGEN:                               "de:08116:7958",
	HAUPTBAHNHOF_OBEN_STUTTGART:                            "de:08111:6115",
	HAUPTBAHNHOF_TIEF_STUTTGART:                            "de:08111:6118",
	HAUPTBF_ARNULFKLETTPLATZ_STUTTGART:                     "de:08111:6112",
	HAUPTFRIEDHOF_STUTTGART:                                "de:08111:2492",
	HAUPTSTRBAHNHOF_EISLINGEN_F:                            "de:08117:1014",
	HAUPTSTRASSE_BERNHAUSEN:                                "de:08116:2008",
	HAUPTSTRASSE_BONLANDEN:                                 "de:08116:2019",
	HAUPTSTRASSE_GROSSBOTTWAR:                              "de:08118:3577",
	HAUPTSTRASSE_MERKLINGEN:                                "de:08115:4587",
	HAUPTSTRASSE_DAGERSHEIM:                                "de:08115:4631",
	HAUPTSTRASSE_HEGNACH:                                   "de:08119:4882",
	HAUPTSTRASSE_KAISERSBACH:                               "de:08119:5076",
	HAUPTSTRASSE_URBACH:                                    "de:08119:5134",
	HAUPTSTRASSE_LEUTENBACH:                                "de:08119:5314",
	HAUPTSTRASSE_RIELINGSHAUSEN:                            "de:08118:5369",
	HAUPTSTRASSE_PLEIDELSHEIM:                              "de:08118:5570",
	HAUPTSTRASSE_HOPFIGHEIM:                                "de:08118:5573",
	HAUPTSTRASSE_UNTERRIEXINGEN:                            "de:08118:5636",
	HAUPTSTRASSE_GROSSSACHSENHEIM:                          "de:08118:5677",
	HAUPTSTRASSE_PEROUSE:                                   "de:08115:5837",
	HAUPTSTRASSE_WEIL_IM_SCHONBUCH:                         "de:08115:5947",
	HAUPTSTRASSE_HATTENHOFEN:                               "de:08117:2053",
	HAUPTSTRASSE_MACHTOLSHEIM:                              "de:08425:2640",
	HAUPTSTRASSE_77_WARMBRONN:                              "de:08115:4538",
	HAUPTSTRASSE_99_NECKARWEIHINGEN:                        "de:08118:5486",
	HAUS_AM_REMSUFER_NECKARREMS:                            "de:08118:6852",
	HAUS_AM_WEINBERG_STUTTGART:                             "de:08111:2519",
	HAUS_AN_DER_METTER_BIETIGHEIM:                          "de:08118:3390",
	HAUS_DER_BURGER_ALDINGEN:                               "de:08118:5589",
	HAUS_DER_FEUERWEHR_ALDINGEN:                            "de:08118:6973",
	HAUS_DER_JUGEND_NECKARGRONINGEN:                        "de:08118:6827",
	HAUS_DER_VEREINE_DAGERSHEIM:                            "de:08115:4630",
	HAUS_FRANK_MARKGRONINGEN:                               "de:08118:3310",
	HAUSACKER_SIELMINGEN:                                   "de:08116:2085",
	HAUSEN_MURRHARDT:                                       "de:08119:4923",
	HAUSENER_STRASSE_MERKLINGEN:                            "de:08115:4590",
	HAUSENRING_STUTTGART:                                   "de:08111:2619",
	HAUSGARTEN_WAIBLINGEN:                                  "de:08119:3648",
	HAUSWEINBERG_BEINSTEIN:                                 "de:08119:6638",
	HAYDNSTRASSE_HOFINGEN:                                  "de:08115:6814",
	HAYDNSTRASSEBADER_SCHORNDORF:                           "de:08119:5211",
	HEBSACKER_STRASSE_STUTTGART:                            "de:08111:2477",
	HECKBACHSTRASSE_KLEINHEPPACH:                           "de:08119:6806",
	HECKENGAUSTRASSE_NAGOLD:                                "de:08235:10118",
	HECKENROSENWEG_NAGOLD:                                  "de:08235:10119",
	HECKENWEG_BESIGHEIM:                                    "de:08118:3473",
	HECKENWEG_HERTMANNSWEILER:                              "de:08119:7456",
	HEDELFINGEN_STUTTGART:                                  "de:08111:227",
	HEDELFINGER_STRASSE_STUTTGART:                          "de:08111:226",
	HEERSTRASSE_STUTTGART:                                  "de:08111:2593",
	HEERSTRASSE_GULTSTEIN:                                  "de:08115:3219",
	HEERWEG_GROSSBETTLINGEN:                                "de:08116:4312",
	HEGELSEIDENSTRASSE_STUTTGART:                           "de:08111:6068",
	HEGELGYMNASIUM_STUTTGART:                               "de:08111:6013",
	HEGELSTRASSE_KIRCHHEIM_T:                               "de:08116:4338",
	HEGENLOHER_STRASSE_THOMASHARDT:                         "de:08116:4503",
	HEGNACHER_HOHE_WAIBLINGEN:                              "de:08119:4877",
	HEGNACHER_STRASSE_HOHENACKER:                           "de:08119:7596",
	HEIDEHOFSTRASSE_STUTTGART:                              "de:08111:6122",
	HEIDESTRASSE_NECKARHALDE:                               "de:08116:4119",
	HEIDHOFE_ALLEENSTRASSE_BOHMENKIRCH:                     "de:08117:10218",
	HEIDHOFE_SCHAFHAUS_BOHMENKIRCH:                         "de:08117:215",
	HEIGELINSTRASSE_STUTTGART:                              "de:08111:2587",
	HEILBAD_HOHENECK:                                       "de:08118:5528",
	HEILBRONNER_STRASSE_LUDWIGSBURG:                        "de:08118:3412",
	HEILIGENACKER_GEISLINGEN_STEIGE:                        "de:08117:120",
	HEIMATMUSEUM_UNTERWEISSACH:                             "de:08119:5388",
	HEIMBACH_JEBENHAUSEN:                                   "de:08117:2189",
	HEIMBERG_STUTTGART:                                     "de:08111:2423",
	HEIMERDINGEN_HEIMERDINGEN:                              "de:08118:4536",
	HEININGERJAHNSTRASSE_GOPPINGEN:                         "de:08117:3303",
	HEINKELSTRASSE_KIRCHHEIM_T:                             "de:08116:3956",
	HEINKELSTRASSE_BOBLINGEN:                               "de:08115:4723",
	HEINKELSTRASSE_LOCHGAU:                                 "de:08118:7565",
	HEINLESMUHLE_VORDERSTEINENBERG:                         "de:08119:6648",
	HEINRICHESSIGSTRASSE_LEONBERG:                          "de:08115:2363",
	HELENELANGESCHULE_MARKGRONINGEN:                        "de:08118:3311",
	HELENENBURGWEG_BIETIGHEIM:                              "de:08118:5666",
	HELFFERICHSTRASSE_STUTTGART:                            "de:08111:2244",
	HELLERSHOF_VORDERSTEINENBERG:                           "de:08119:5281",
	HELLERSHOF_KREUZUNG_VORDERSTEINENBERG:                  "de:08119:6915",
	HEMMINGEN_HEMMINGEN:                                    "de:08118:4527",
	HEMMINGER_STRASSE_STUTTGART:                            "de:08111:2357",
	HEMMINGER_STRASSE_HEIMERDINGEN:                         "de:08118:4529",
	HENNERSDORFER_STRASSE_NEUWEILER:                        "de:08115:4819",
	HENRIETTENSTRASSE_KIRCHHEIM_T:                          "de:08116:4385",
	HERBSTHALDE_STUTTGART:                                  "de:08111:2265",
	HERDERPLATZ_STUTTGART:                                  "de:08111:216",
	HERDERSTRASSE_OBERESSLINGEN:                            "de:08116:4200",
	HERDWEG_DITZINGEN:                                      "de:08118:3033",
	HERDWEG_JEBENHAUSEN:                                    "de:08117:2100",
	HERDWEG_POSTAMT_BOBLINGEN:                              "de:08115:4690",
	HERMANNESSIGSTRASSE_SCHWIEBERDINGEN:                    "de:08118:4525",
	HERMANNHESSEREALSCHULE_GOPPINGEN:                       "de:08117:3383",
	HERMANNHESSESTRASSE_AIDLINGEN:                          "de:08115:4778",
	HERMANNHESSESTRASSE_NECKARWEIHINGEN:                    "de:08118:6848",
	HERMANNLONSSTRASSE_NURTINGEN:                           "de:08116:4427",
	HERMANNLONSSTRASSE_GERLINGEN:                           "de:08118:6895",
	HERMANNLONSWEG_WINNENDEN:                               "de:08119:5218",
	HERMANNSCHWABHALLE_WINNENDEN:                           "de:08119:4896",
	HERRENBACHSTRASSE_BAIERECK:                             "de:08117:7677",
	HERRENBERG_HERRENBERG:                                  "de:08115:4512",
	HERRENBERGER_STRASSE_MAICHINGEN:                        "de:08115:4643",
	HERRENBERGER_STRASSE_BOBLINGEN:                         "de:08115:4716",
	HERRENBERGER_STRASSE_KAYH:                              "de:08115:4832",
	HERRENBERGER_STRASSE_OBERJETTINGEN:                     "de:08115:7016",
	HERRENWIESENWEG_NECKARHALDE:                            "de:08116:4121",
	HERRSCHAFTSWEG_PFLUGFELDEN:                             "de:08118:5495",
	HERTICHSTRASSE_33_ELTINGEN:                             "de:08115:4540",
	HERTICHSTRASSE_73_ELTINGEN:                             "de:08115:6505",
	HERWEGHSTRASSE_STUTTGART:                               "de:08111:6511",
	HERZOGPHILIPPPLATZ_PARKSIEDLUNG:                        "de:08116:2935",
	HERZOGWEG_HERRENBERG:                                   "de:08115:7005",
	HESLACH_VOGELRAIN_STUTTGART:                            "de:08111:43",
	HESSIGHEIMER_STRASSE_MUNDELSHEIM:                       "de:08118:5575",
	HEUBACHSTRASSE_GOPPINGEN:                               "de:08117:1065",
	HEUBERGSTRASSE_KORNWESTHEIM:                            "de:08118:3413",
	HEUMADEN_STUTTGART:                                     "de:08111:136",
	HEUMADEN_KAISERSBACH:                                   "de:08119:6671",
	HEUMADEN_BOCKELSTRASSE_STUTTGART:                       "de:08111:6133",
	HEUMADEN_ROSE_STUTTGART:                                "de:08111:2538",
	HEUMADEN_SCHULE_STUTTGART:                              "de:08111:2539",
	HEUSEE_PLUDERHAUSEN:                                    "de:08119:5236",
	HEUSTEIGSTRASSE_BOBLINGEN:                              "de:08115:4712",
	HEUTINGSHEIMER_STRASSE_STUTTGART:                       "de:08111:103",
	HEUWEG_RUTESHEIM:                                       "de:08115:5834",
	HEYDSCHULE_MARKGRONINGEN:                               "de:08118:5862",
	HIEBERSCHULE_UHINGEN:                                   "de:08117:2009",
	HILDENSTRASSE_GUNDELBACH:                               "de:08118:5882",
	HILDRIZHAUSER_STRASSE_HERRENBERG:                       "de:08115:7012",
	HILLERLOCHGAUER_STRASSE_BIETIGHEIM:                     "de:08118:5661",
	HILLERPLATZ_BIETIGHEIM:                                 "de:08118:5664",
	HIMMELSLEITER_STUTTGART:                                "de:08111:289",
	HINDENBURGSTRASSE_NELLINGEN:                            "de:08116:2910",
	HINDENBURGSTRASSE_MOGLINGEN:                            "de:08118:3442",
	HINDENBURGSTRASSE_WEISSACH_BB:                          "de:08115:4571",
	HINDENBURGSTRASSE_HERRENBERG:                           "de:08115:4802",
	HINDENBURGSTRASSE_DENKENDORF:                           "de:08116:5023",
	HINTERBUCHELBERG_MURRHARDT:                             "de:08119:5057",
	HINTERE_SCHMIEDGASSE_SCHWABISCH_GMUND:                  "de:08136:3028",
	HINTERE_TOS_BACKNANG:                                   "de:08119:6666",
	HINTERER_HOLZWEG_RUDERN:                                "de:08116:4116",
	HINTERSTEINENBERG_SUD_VORDERSTEINENBERG:                "de:08119:6668",
	HINTERWESTERMURR_FORNSBACH:                             "de:08119:5303",
	HIRSCH_GOSBACH:                                         "de:08117:70",
	HIRSCH_HAUSEN_BAD_UBERKINGEN:                           "de:08117:90",
	HIRSCH_SCHNITTLINGEN:                                   "de:08117:7680",
	HIRSCHBERGER_STRASSE_RAMTEL:                            "de:08115:2326",
	HIRSCHBERGSTRASSE_GERLINGEN:                            "de:08118:2306",
	HIRSCHBERGSTRASSE_EGLOSHEIM:                            "de:08118:3456",
	HIRSCHGARTENWEG_MOGLINGEN:                              "de:08118:3435",
	HIRSCHHOF_LENGLINGEN:                                   "de:08117:11209",
	HIRSCHLANDER_STRASSE_DITZINGEN:                         "de:08118:3028",
	HIRSCHLANDER_STRASSE_HOFINGEN:                          "de:08115:6050",
	HIRSCHLANDHOF_OBERESSLINGEN:                            "de:08116:4073",
	HIRSCHLANDSTRASSE_OBERESSLINGEN:                        "de:08116:4062",
	HIRSCHPLATZ_FAURNDAU:                                   "de:08117:7008",
	HIRSCHSTRASSE_OSSWEIL:                                  "de:08118:6833",
	HIRSCHSTRASSE_ECHTERDINGEN:                             "de:08116:7178",
	HIRSCHSTRASSE_FAURNDAU:                                 "de:08117:7009",
	HOCHBERGER_STRASSE_HOCHDORF_REMSECK:                    "de:08118:5552",
	HOCHBERGER_STRASSE_BITTENFELD:                          "de:08119:5553",
	HOCHBERGER_STRASSE_POPPENWEILER:                        "de:08118:6820",
	HOCHBERGER_WALD_HOCHBERG:                               "de:08118:6823",
	HOCHDORFER_STRASSE_NOTZINGEN:                           "de:08116:4219",
	HOCHDORFER_STRASSE_BITTENFELD:                          "de:08119:5554",
	HOCHDORFER_STRASSE_EBERDINGEN:                          "de:08118:5623",
	HOCHHAUS_REICHENBACH_F:                                 "de:08116:4238",
	HOCHHAUS_LORCH:                                         "de:08136:7015",
	HOCHSCHULE_ESSLINGEN_N:                                 "de:08116:4022",
	HOCHSCHULE_GOPPINGEN:                                   "de:08117:3301",
	HOCHSCHULZENTRUM_ESSLINGEN_N:                           "de:08116:3998",
	HOCHSTETTER_WEG_ZOLLBERG:                               "de:08116:3920",
	HOCHWACHTTURM_WAIBLINGEN:                               "de:08119:5387",
	HOCHWANG_RATHAUS_OBERLENNINGEN:                         "de:08116:5769",
	HOCHWIESENSTRASSE_BONDORF:                              "de:08115:4760",
	HOF_STETTEN_LEINFELDENECHT:                             "de:08116:2115",
	HOF_SACHSENWEILER:                                      "de:08119:5347",
	HOFEN_STUTTGART:                                        "de:08111:272",
	HOFEN_HOFEN_BONNIGHEIM:                                 "de:08118:5696",
	HOFENER_STRASSE_BONNIGHEIM:                             "de:08118:5697",
	HOFERSTRASSE_LUDWIGSBURG:                               "de:08118:5538",
	HOFFELD_STUTTGART:                                      "de:08111:2569",
	HOFGARTENSTR_ROMMELSHAUSEN:                             "de:08119:3755",
	HOFHALDE_BARTENBACH_GOPPINGEN:                          "de:08117:4912",
	HOFINGEN_HOFINGEN:                                      "de:08115:1003",
	HOFMEISTERSCHAUWERK_SINDELFINGEN:                       "de:08115:4653",
	HOFSTRASSE_HEGENSBERG:                                  "de:08116:4204",
	HOFSTRASSE_DARMSHEIM:                                   "de:08115:4638",
	HOFWIESENSTRASSE_GERLINGEN:                             "de:08118:2307",
	HOGY_GOPPINGEN:                                         "de:08117:9011",
	HOHBAUM_PLUDERHAUSEN:                                   "de:08119:5250",
	HOHBERGSCHULE_PLUDERHAUSEN:                             "de:08119:5243",
	HOHBERGSTRASSE_WEITMARS:                                "de:08136:7036",
	HOHE_EICHE_STUTTGART:                                   "de:08111:6571",
	HOHEBILDSTRASSE_WALDDORF:                               "de:08415:29043",
	HOHENACKERSTRASSE_SCHMIDEN:                             "de:08119:6521",
	HOHENBRACH_GRAB:                                        "de:08119:5106",
	HOHENFREIBAD_STUTTGART:                                 "de:08111:2455",
	HOHENKREUZ_HOHENKREUZ:                                  "de:08116:4124",
	HOHENRAINSTRASSE_NECKARWEIHINGEN:                       "de:08118:5487",
	HOHENSTADTER_WEG_GOPPINGEN:                             "de:08117:4920",
	HOHENSTANGE_FRANKFURTER_STRASSE_TAMM:                   "de:08118:3460",
	HOHENSTANGE_STUTTGARTER_STRASSE_TAMM:                   "de:08118:3452",
	HOHENSTANGE_TUBINGER_STRASSE_TAMM:                      "de:08118:3446",
	HOHENSTANGE_ULMER_STRASSE_TAMM:                         "de:08118:3453",
	HOHENSTAUFENSTRASSE_ASPERG:                             "de:08118:3407",
	HOHENSTAUFENSTRASSE_SCHORNDORF:                         "de:08119:5979",
	HOHENSTEINSTRASSE_STUTTGART:                            "de:08111:6109",
	HOHENSTEINSTRASSE_GOPPINGEN:                            "de:08117:9013",
	HOHENSTRASSE_FELLBACH:                                  "de:08119:37",
	HOHENWEG_OBERBERKEN:                                    "de:08119:4946",
	HOHENZOLLERNPLATZ_LUDWIGSBURG:                          "de:08118:5517",
	HOHENZOLLERNSTRASSE_PLEIDELSHEIM:                       "de:08118:3565",
	HOHENZOLLERNSTRASSE_HOLZGERLINGEN:                      "de:08115:4208",
	HOHENZOLLERNSTRASSE_SCHAFHAUSEN:                        "de:08115:4579",
	HOHERBAUMWEG_NAGOLD:                                    "de:08235:4401",
	HOHES_GESTADE_NURTINGEN:                                "de:08116:2996",
	HOHEWARTSTRASSE_STUTTGART:                              "de:08111:6155",
	HOHLWEG_STETTEN_LEINFELDENECHT:                         "de:08116:2111",
	HOHNWEILER_FORSTSTRASSE_LIPPOLDSWEILER:                 "de:08119:5996",
	HOHNWEILER_RATHAUS_LIPPOLDSWEILER:                      "de:08119:4981",
	HOHNWEILER_RATHAUSSTRASSE_LIPPOLDSWEILER:               "de:08119:4979",
	HOHREIN_ORTSMITTE_GOPPINGEN:                            "de:08117:4916",
	HOHROT_GROSSASPACH:                                     "de:08119:5358",
	HOHWEG_FLACHT:                                          "de:08115:4567",
	HOLBEINWEG_KIRCHHEIM_T:                                 "de:08116:4357",
	HOLDERACKER_STUTTGART:                                  "de:08111:2617",
	HOLDERBUSCHLE_GROSSSACHSENHEIM:                         "de:08118:5891",
	HOLDERLINPLATZ_STUTTGART:                               "de:08111:6070",
	HOLDERLINSTRASSE_STUTTGART:                             "de:08111:2195",
	HOLDERLINSTRASSE_MARBACH_N:                             "de:08118:3596",
	HOLDERLINWEG_URBACH:                                    "de:08119:6695",
	HOLDERWEG_STETTEN_LEINFELDENECHT:                       "de:08116:2112",
	HOLDIS_PFAHLBRONN:                                      "de:08119:6784",
	HOLL_AICHELBERG_AICHWALD:                               "de:08116:4167",
	HOLZBERGWEG_SCHORNDORF:                                 "de:08119:5256",
	HOLZGERLINGEN_HOLZGERLINGEN:                            "de:08115:4193",
	HOLZGERLINGER_STRASSE_ALTDORF_BB:                       "de:08115:4793",
	HOLZHAUSER_STRASSE_WANGEN_GP:                           "de:08117:5008",
	HOLZHEIMER_STRASSE_EISLINGEN_F:                         "de:08117:1047",
	HOLZLACHSTRASSE_FRICKENHAUSEN:                          "de:08116:2824",
	HOLZSTEIG_GULTSTEIN:                                    "de:08115:3234",
	HOLZWEILERHOF_WINZERHAUSEN:                             "de:08118:5620",
	HONOLDWEG_STUTTGART:                                    "de:08111:2404",
	HOPFIGHEIMER_STRASSE_STEINHEIM_M:                       "de:08118:5599",
	HOPFIGHEIMER_STRASSE_BIETIGHEIM:                        "de:08118:7414",
	HORBSTRASSE_RUIT:                                       "de:08116:5022",
	HORDTHOF_MURRHARDT:                                     "de:08119:5140",
	HORNBACH_ALDINGEN:                                      "de:08118:3570",
	HORNBERGSTRLACHEBRUCKLE_DITZINGEN:                      "de:08118:3016",
	HORNBERGSTRASSE_KORNWESTHEIM:                           "de:08118:5465",
	HORNLE_DREIBRONNENSTRASSE_MARBACH_N:                    "de:08118:3539",
	HORNLE_HOCHHAUS_MARBACH_N:                              "de:08118:3540",
	HORNLE_WIESBADENER_PLATZ_MARBACH_N:                     "de:08118:3541",
	HORNLESWEG_KOHLBERG:                                    "de:08116:4321",
	HORRENWINKEL_STEINHEIM_M:                               "de:08118:6987",
	HORSCHBACHSCHULE_MURRHARDT:                             "de:08119:4915",
	HOSSLINSWART_HOSSLINSWART:                              "de:08119:5199",
	HUBERTUSWEG_WIFLINGSHAUSEN:                             "de:08116:4140",
	HUGELSTRASSE_UNTERBRUDEN:                               "de:08119:5377",
	HUGOWOLFWEG_KIRCHHEIM_T:                                "de:08116:4364",
	HUGSTRASSE_GOPPINGEN:                                   "de:08117:9204",
	HULB_BOBLINGEN:                                         "de:08115:7115",
	HULBE_SCHWIEBERDINGEN:                                  "de:08118:3348",
	HULBEN_HOLZGERLINGEN:                                   "de:08115:6525",
	HULBEN_NORD_HOLZGERLINGEN:                              "de:08115:6526",
	HUMBOLDTSTRASSE_LINSENHOFEN:                            "de:08116:4431",
	HUMMELBAUM_RENNINGEN:                                   "de:08115:3320",
	HUMMELBUHL_SULZBACH_M:                                  "de:08119:4911",
	HUMPFENTALSTRASSE_NURTINGEN:                            "de:08116:2963",
	HUNDSACKER_STRUMPFELBACH_WEINSTADT:                     "de:08119:5180",
	HUNDSBERG_VORDERSTEINENBERG:                            "de:08119:6925",
	HUNGERBUHLSTRASSE_SCHORNDORF:                           "de:08119:5223",
	HUOBER_ERDMANNHAUSEN:                                   "de:08118:5371",
	HUSARENHOF_BESIGHEIM:                                   "de:08118:5719",
	HUTTENBUHL_VORDERSTEINENBERG:                           "de:08119:6644",
	HUTTENBUHLSEE_ABZW_VORDERSTEINENBERG:                   "de:08119:6914",
	HUTTLEN_SPIEGELBERG:                                    "de:08119:5047",
	HUTWIESENSTRASSE_MAGSTADT:                              "de:08115:3297",
	IBM_EHNINGEN:                                           "de:08115:3218",
	IHINGER_HOF_RENNINGEN:                                  "de:08115:3178",
	IKEA_LUDWIGSBURG:                                       "de:08118:3467",
	IKEA_SINDELFINGEN:                                      "de:08115:4803",
	ILSFELDER_STRASSE_OTTMARSHEIM:                          "de:08118:3502",
	ILTISWEG_FAURNDAU:                                      "de:08117:7001",
	IM_ALTEN_SEE_STEINENBRONN:                              "de:08115:7108",
	IM_BACKENSCHLAG_BONDORF:                                "de:08115:7079",
	IM_BAUMSTUCKLE_WAIBLINGEN:                              "de:08119:3693",
	IM_BORNRAIN_MOGLINGEN:                                  "de:08118:3436",
	IM_BRAUNKIEL_ALTBACH:                                   "de:08116:3801",
	IM_CHAUSSEEFELD_STUTTGART:                              "de:08111:2440",
	IM_DEGEN_STUTTGART:                                     "de:08111:82",
	IM_ELSENTAL_STUTTGART:                                  "de:08111:2578",
	IM_ENGENLAUCH_BARTENBACH_GOPPINGEN:                     "de:08117:4013",
	IM_FELDLE_BIETIGHEIM:                                   "de:08118:3384",
	IM_GAILER_GECHINGEN:                                    "de:08235:1530",
	IM_GANSLESGRUND_NURTINGEN:                              "de:08116:2995",
	IM_GRUND_BAD_UBERKINGEN:                                "de:08117:45",
	IM_HAG_STUTTGART:                                       "de:08111:2507",
	IM_HEGNACH_EBERSBACH_F:                                 "de:08117:8013",
	IM_HILLER_MIEDELSBACH:                                  "de:08119:5109",
	IM_HIMMERREICH_STUTTGART:                               "de:08111:2450",
	IM_HOLDER_PLIENSAUVORSTADT:                             "de:08116:6770",
	IM_HUMMERHOLZ_WEILER_ZUM_STEIN:                         "de:08119:6769",
	IM_KAISEMER_STUTTGART:                                  "de:08111:2246",
	IM_KAPF_RUIT:                                           "de:08116:2933",
	IM_KOLLER_SIELMINGEN:                                   "de:08116:2159",
	IM_KOPFERT_STUTTGART:                                   "de:08111:2094",
	IM_KUSTERFELD_BACKNANG:                                 "de:08119:3728",
	IM_LANGEN_FELD_MARKGRONINGEN:                           "de:08118:6977",
	IM_LAUCH_STUTTGART:                                     "de:08111:2431",
	IM_LETTEN_EHNINGEN:                                     "de:08115:3238",
	IM_MADER_STUTTGART:                                     "de:08111:2533",
	IM_MOLDENGRABEN_KORNWESTHEIM:                           "de:08118:3363",
	IM_MORGEN_ALBERSHAUSEN:                                 "de:08117:2020",
	IM_MUNZEN_KIRCHHEIM_T:                                  "de:08116:3951",
	IM_RAISER_STUTTGART:                                    "de:08111:2472",
	IM_SAMANN_WAIBLINGEN:                                   "de:08119:5412",
	IM_SEE_KLEINGLATTBACH:                                  "de:08118:5731",
	IM_SEELE_HERRENBERG:                                    "de:08115:7056",
	IM_STEINGRABEN_HERRENBERG:                              "de:08115:7039",
	IM_STOCKACH_WEILHEIM_T:                                 "de:08116:2831",
	IM_TAL_MUNKLINGEN:                                      "de:08115:4591",
	IM_UNTEREN_ZEHEN_WIFLINGSHAUSEN:                        "de:08116:3912",
	IM_VOGELSANG_NEUENHAUS:                                 "de:08116:2962",
	IM_VOGELSANG_HERRENBERG:                                "de:08115:7042",
	IM_WALDECK_ASPERG:                                      "de:08118:5634",
	IM_WEIZEN_KORNWESTHEIM:                                 "de:08118:5810",
	IM_WIESENGRUND_BACKNANG:                                "de:08119:3613",
	IM_WIESENGRUND_KORNWESTHEIM:                            "de:08118:5450",
	IMENTAL_UNTERJETTINGEN:                                 "de:08115:4864",
	IMMANUELDORNFELDSTRASSE_NECKARWEIHINGEN:                "de:08118:7451",
	IMMELMANNZENTRUM_GOPPINGEN:                             "de:08117:1008",
	IMMO_NAGOLD:                                            "de:08235:10145",
	IN_DEN_AULESWIESEN_ALLMERSBACH_IM_TAL:                  "de:08119:7480",
	IN_DEN_BEETEN_GROSSINGERSHEIM:                          "de:08118:5940",
	IN_DEN_ENTENACKERN_STUTTGART:                           "de:08111:2093",
	IN_DEN_FRAUENGARTEN_GROSSBOTTWAR:                       "de:08118:3578",
	IN_DEN_HOFACKERN_WEILER_ZUM_STEIN:                      "de:08119:6765",
	IN_DEN_MESSENWIESEN_ROSSWALDEN:                         "de:08117:7408",
	IN_DEN_STEGWIESEN_STUTTGART:                            "de:08111:3996",
	IN_DEN_WEINGARTEN_EISLINGEN_F:                          "de:08117:1032",
	IN_DER_MISSE_EBHAUSEN:                                  "de:08235:10206",
	IN_DER_WERRE_STUTTGART:                                 "de:08111:2433",
	INDEXSTRASSEBAHNHOF_OBERESSLINGEN:                      "de:08116:3810",
	INDUSTRIEGEBIET_PLATTENHARDT:                           "de:08116:2026",
	INDUSTRIEGEBIET_AICHSCHIESS:                            "de:08116:4154",
	INDUSTRIEGEBIET_WENDLINGEN_N:                           "de:08116:4287",
	INDUSTRIEGEBIET_NABERN:                                 "de:08116:4410",
	INDUSTRIEGEBIET_UNTERRIEXINGEN:                         "de:08118:5635",
	INDUSTRIEGEBIET_BIRKMANNSWEILER:                        "de:08119:5968",
	INDUSTRIEGEBIET_ECHTERDINGEN:                           "de:08116:7101",
	INDUSTRIEGEBIET_ALBERSHAUSEN:                           "de:08117:2023",
	INDUSTRIEGEBIET_BORTLINGEN:                             "de:08117:4036",
	INDUSTRIEGEBIET_OST_NEUHAUSEN_F:                        "de:08116:2904",
	INDUSTRIEGEBIET_OST_DITZINGEN:                          "de:08118:5844",
	INDUSTRIEGEBIET_OST_GOPPINGEN:                          "de:08117:3007",
	INDUSTRIEGEBIET_SUD_JEBENHAUSEN:                        "de:08117:2118",
	INDUSTRIEGEBIET_WEST_NEUHAUSEN_F:                       "de:08116:2905",
	INDUSTRIEGEBIETFA_KNECHT_LORCH:                         "de:08136:7013",
	INDUSTRIESTRASSE_STUTTGART:                             "de:08111:2601",
	INDUSTRIESTRASSE_BACKNANG:                              "de:08119:4972",
	INDUSTRIESTRASSE_MERKLINGEN:                            "de:08115:5839",
	INDUSTRIESTRASSE_SERSHEIM:                              "de:08118:7415",
	INGPARK_NORD_NAGOLD:                                    "de:08235:10401",
	INGPARK_SUD_NAGOLD:                                     "de:08235:10399",
	INSELBAD_ZIZISHAUSEN:                                   "de:08116:4301",
	INSELSTRASSE_STUTTGART:                                 "de:08111:83",
	ISARSTRASSE_WALDREMS:                                   "de:08119:3665",
	ISEGRIMWEG_STUTTGART:                                   "de:08111:137",
	ISELSHAUSER_STRASSE_MOTZINGEN:                          "de:08115:7051",
	JAGERHAUS_LIEBERSBRONN:                                 "de:08116:4105",
	JAHNHALLE_ENDERSBACH:                                   "de:08119:5188",
	JAHNSTRASSE_KORNTAL:                                    "de:08118:2636",
	JAHNSTRASSE_WEILHEIM_T:                                 "de:08116:4394",
	JAHNSTRASSE_BESIGHEIM:                                  "de:08118:5581",
	JAHNSTRASSE_POPPENWEILER:                               "de:08118:6985",
	JAHNSTRASSE_WALDDORF:                                   "de:08415:29042",
	JAKOBBUTTERSTRASSE_SCHMIDEN:                            "de:08119:2378",
	JAKOBRAIBLEANLAGE_SCHORNDORF:                           "de:08119:5219",
	JAKOBSTRASSE_BERKHEIM:                                  "de:08116:4151",
	JELINSTRASSE_MOHRINGEN:                                 "de:08111:2585",
	JENISCHSTRASSE_LUDWIGSBURG:                             "de:08118:5539",
	JENNERSTRASSE_KUPPINGEN:                                "de:08115:4845",
	JESISTRASSE_WAIBLINGEN:                                 "de:08119:5950",
	JETTINGER_STRASSE_KUPPINGEN:                            "de:08115:4868",
	JOHANNESKEPLERGYMNASIUM_LEONBERG:                       "de:08115:2329",
	JOHANNESKEPLERSCHULE_MAGSTADT:                          "de:08115:3302",
	JOHANNESTAUFERKIRCHE_MAGSTADT:                          "de:08115:3295",
	JOHANNESGRABEN_STUTTGART:                               "de:08111:2580",
	JOHANNESKIRCHE_BERNHAUSEN:                              "de:08116:1998",
	JOHANNESKIRCHE_BACKNANG:                                "de:08119:3666",
	JOHANNESKIRCHE_STUTTGART:                               "de:08111:6268",
	JOHANNESSTRASSE_ZELL_ESSLINGEN:                         "de:08116:4173",
	JOHANNESSTRASSE_SCHORNDORF:                             "de:08119:5191",
	JOHANNESSTRASSE_KORNWESTHEIM:                           "de:08118:5461",
	JOHANNESSTRASSE_HERRENBERG:                             "de:08115:7074",
	JOHNFKENNEDYSTR_GOPPINGEN:                              "de:08117:9015",
	JOSEPHHAYDNSTRASSE_WEIL_DER_STADT:                      "de:08115:6478",
	JUGENDDORF_VAIHINGEN_E:                                 "de:08118:6980",
	JUGENDHERBERGE_MURRHARDT:                               "de:08119:5981",
	JUGENDHERBERGE_HOHENSTAUFEN:                            "de:08117:1305",
	JUNGBORN_NURTINGEN:                                     "de:08116:2809",
	JUNGE_WEINBERGE_WAIBLINGEN:                             "de:08119:3723",
	JUNKERSSTRASSE_BOBLINGEN:                               "de:08115:3224",
	JURASTRASSE_STUTTGART:                                  "de:08111:356",
	JUSIWEG_ZOLLBERG:                                       "de:08116:4180",
	JUX_JUX:                                                "de:08119:5046",
	K_1206_BAHNHOF_REICHENBACH_F:                           "de:08116:3941",
	KFERDBRAUNSTRASSE_BACKNANG:                             "de:08119:6912",
	KHORNSCHUCHSTRASSE_URBACH:                              "de:08119:6625",
	KAISERLUDWIGSTRASSE_MURRHARDT:                          "de:08119:5312",
	KAISERBAU_AGENTUR_FUR_ARBEIT_GOPPINGEN:                 "de:08117:1004",
	KAISEREICHE_SCHLICHTEN:                                 "de:08119:6501",
	KAISERSTRASSE_NELLINGEN:                                "de:08116:2981",
	KAISERSTRASSE_BEUTELSBACH:                              "de:08119:3764",
	KALBERSTELLE_WEIL_IM_SCHONBUCH:                         "de:08115:7036",
	KALKOFEN_ENDERSBACH:                                    "de:08119:6839",
	KALKOFENSTRASSE_HERRENBERG:                             "de:08115:3235",
	KALLENBERG_ALTHUTTE:                                    "de:08119:4992",
	KALLENBERG_LACKFABRIK_MUNCHINGEN:                       "de:08118:5627",
	KALLENBERG_RASTHAUS_MUNCHINGEN:                         "de:08118:4516",
	KALTENTAL_STUTTGART:                                    "de:08111:6009",
	KALTENTHALSTRASSE_ALDINGEN:                             "de:08118:6828",
	KAMERALAMTSSTRASSE_STUTTGART:                           "de:08111:200",
	KAMMGARNSPINNEREI_BIETIGHEIM:                           "de:08118:5646",
	KAMMGARNSPINNEREI_B27_BIETIGHEIM:                       "de:08118:7450",
	KANTSTRASSE_SIELMINGEN:                                 "de:08116:2033",
	KANTSTRASSE_GEISLINGEN_STEIGE:                          "de:08117:116",
	KANTSTRASSE_GOPPINGEN:                                  "de:08117:1000",
	KAPELLE_PLOCHINGEN:                                     "de:08116:4229",
	KAPELLENBERG_DOFFINGEN:                                 "de:08115:4769",
	KAPELLENSTRASSE_WENDLINGEN_N:                           "de:08116:4295",
	KAPELLENSTRASSE_SUSSEN:                                 "de:08117:7610",
	KAPF_KAPFHOFWEG_VORDERSTEINENBERG:                      "de:08119:5097",
	KAPF_MOHNWIESENWEG_VORDERSTEINENBERG:                   "de:08119:5096",
	KAPFFSTRASSE_URBACH:                                    "de:08119:6629",
	KAPPELBERGSTRASSE_FELLBACH:                             "de:08119:2502",
	KAPPELGASSE_SCHWABISCH_GMUND:                           "de:08136:3078",
	KARLGEORGPFLEIDERERSTRASSE_HERTMANNSWEILER:             "de:08119:5869",
	KARLJOOSSTRASSE_KORNWESTHEIM:                           "de:08118:5485",
	KARLMAIALLEE_BIETIGHEIM:                                "de:08118:7446",
	KARLOLGAKRANKENHAUS_STUTTGART:                          "de:08111:76",
	KARLPFAFFSTRASSE_STUTTGART:                             "de:08111:164",
	KARLPFAFFSTRASSE_PLIENSAUVORSTADT:                      "de:08116:4036",
	KARLSHOF_KLEINASPACH:                                   "de:08119:6656",
	KARLSHOHE_LUDWIGSBURG:                                  "de:08118:5519",
	KARLSPLATZ_LUDWIGSBURG:                                 "de:08118:3462",
	KARLSRUHER_ALLEE_PFLUGFELDEN:                           "de:08118:5457",
	KARLSTRASSE_MUSBERG:                                    "de:08116:309",
	KARLSTRASSE_BERNHAUSEN:                                 "de:08116:2009",
	KARLSTRASSE_BOBLINGEN:                                  "de:08115:3103",
	KARLSTRASSE_BISSINGEN_LB:                               "de:08118:5683",
	KARLSTRASSE_ROMMELSHAUSEN:                              "de:08119:5949",
	KARNSBERG_MURRHARDT:                                    "de:08119:5059",
	KAROLINGERSTRASSE_SCHMIDEN:                             "de:08119:5426",
	KARPATENWEG_OTLINGEN:                                   "de:08116:4263",
	KASBACH_MURRHARDT:                                      "de:08119:5298",
	KASERNE_MOHRINGEN_STUTTGART:                            "de:08111:2038",
	KASPARSWALD_STETTEN_LEINFELDENECHT:                     "de:08116:2114",
	KATH_KINDERGARTEN_LORCH:                                "de:08136:7016",
	KATH_KIRCHE_GINGEN_F:                                   "de:08117:19",
	KATHARINENHOF_STRUMPFELBACH_BACKNANG:                   "de:08119:4898",
	KATHARINENHOSPITAL_STUTTGART:                           "de:08111:2203",
	KATHARINENSTAFFEL_ESSLINGEN_N:                          "de:08116:4094",
	KATHARINENSTRASSE_EGLOSHEIM:                            "de:08118:5511",
	KATHARINENSTRASSE_GUNDELBACH:                           "de:08118:6976",
	KATHEKOLLWITZSTRASSE_HERRENBERG:                        "de:08115:3237",
	KATHOLISCHE_KIRCHE_DETTINGEN_T:                         "de:08116:4277",
	KATHOLISCHE_KIRCHE_BEMPFLINGEN:                         "de:08116:4314",
	KATHOLISCHE_KIRCHE_AIDLINGEN:                           "de:08115:4725",
	KATHOLISCHE_KIRCHE_GROSSHEPPACH:                        "de:08119:5171",
	KATHOLISCHE_KIRCHE_HEGNACH:                             "de:08119:5429",
	KATHOLISCHE_KIRCHE_OBERSTENFELD:                        "de:08118:5613",
	KATHOLISCHE_KIRCHE_UNTERJETTINGEN:                      "de:08115:7015",
	KATHOLISCHE_KIRCHE_BESIGHEIM:                           "de:08118:7085",
	KATHOLISCHE_KIRCHE_ZELL_A:                              "de:08117:2114",
	KATHOLISCHE_KIRCHE_GRUIBINGEN:                          "de:08117:3239",
	KATZENBACHSTRASSE_STUTTGART:                            "de:08111:6011",
	KATZENKOPF_WALDENBRONN:                                 "de:08116:4171",
	KATZENLOCH_GEISLINGEN_STEIGE:                           "de:08117:105",
	KEGELPLATZ_STETTEN_I_R:                                 "de:08119:4930",
	KEHLSTRASSE_VAIHINGEN_E:                                "de:08118:6941",
	KEILBERGSTRASSE_BOBLINGEN:                              "de:08115:3149",
	KELTENBURGSTRASSE_BOBLINGEN:                            "de:08115:3136",
	KELTENMUSEUM_HOCHDORF_EBERDINGEN:                       "de:08118:6830",
	KELTENSTRASSE_RENNINGEN:                                "de:08115:3194",
	KELTER_HOF_UND_LEMBACH:                                 "de:08118:3336",
	KELTER_FRICKENHAUSEN:                                   "de:08116:4430",
	KELTER_STETTEN_I_R:                                     "de:08119:4928",
	KELTER_BRUCH_WEISSACH:                                  "de:08119:4990",
	KELTER_STRUMPFELBACH_WEINSTADT:                         "de:08119:5177",
	KELTER_WINNENDEN:                                       "de:08119:5306",
	KELTER_SCHNAIT:                                         "de:08119:5400",
	KELTER_STEINHEIM_M:                                     "de:08118:5603",
	KELTER_GROSSBOTTWAR:                                    "de:08118:5609",
	KELTERSTADION_ROMMELSHAUSEN:                            "de:08119:3742",
	KELTERPLATZ_BESIGHEIM:                                  "de:08118:5582",
	KELTERPLATZ_OBERBRUDEN:                                 "de:08119:7458",
	KELTERRAIN_ECHTERDINGEN:                                "de:08116:2142",
	KELTERSCHULE_NECKARREMS:                                "de:08118:5591",
	KELTERSIEDLUNG_SCHLECHTBACH:                            "de:08119:5120",
	KELTERSTRASSE_STUTTGART:                                "de:08111:370",
	KELTERSTRASSE_ROMMELSHAUSEN:                            "de:08119:3741",
	KELTERSTRASSE_GRUNBACH:                                 "de:08119:3749",
	KELTERSTRASSE_FRICKENHAUSEN:                            "de:08116:5859",
	KELTERSTRASSE_BISSINGEN_LB:                             "de:08118:6720",
	KELTERWEG_BURGSTALL_BURGSTETTEN:                        "de:08119:3509",
	KELTERWEG_SCHONAICH:                                    "de:08115:4804",
	KELTERWIESENBACH_GERADSTETTEN:                          "de:08119:3746",
	KEMNATER_STRASSE_STUTTGART:                             "de:08111:6136",
	KENNENBURG_KENNENBURG:                                  "de:08116:4098",
	KEPLERSTRASSE_OBERESSLINGEN:                            "de:08116:4168",
	KEPLERSTRASSE_RUTESHEIM:                                "de:08115:5935",
	KEPLERSTRASSE_DEIZISAU:                                 "de:08116:6998",
	KEPLERSTRASSE_GOPPINGEN:                                "de:08117:9003",
	KERNERSTEG_ALDINGEN:                                    "de:08118:5587",
	KERSCHENSTEINERSCHULE_STUTTGART:                        "de:08111:6179",
	KESSELTOBELSTRASSE_FAURNDAU:                            "de:08117:1203",
	KIEBITZWEG_KIRCHHEIM_T:                                 "de:08116:4153",
	KIENBACHSTRASSE_STUTTGART:                              "de:08111:264",
	KIENBACHSTRASSE_FELLBACH:                               "de:08119:2644",
	KIESACKERFA_WOLLIN_WALDHAUSEN_B_SCHORND_LORCH:          "de:08136:7032",
	KIESELHOF_MURRHARDT:                                    "de:08119:5139",
	KILLESBERG_STUTTGART:                                   "de:08111:2235",
	KIMMICHSWEILER_WEG_LIEBERSBRONN:                        "de:08116:4003",
	KINDERGARTEN_HOHENGEHREN:                               "de:08116:4078",
	KINDERGARTEN_FORNSBACH:                                 "de:08119:4926",
	KINDERGARTEN_METTERZIMMERN:                             "de:08118:5667",
	KINDERGARTEN_JUX:                                       "de:08119:6672",
	KINDERGARTEN_OCHSENBACH:                                "de:08118:7661",
	KINDERGARTEN_ROSSWALDEN:                                "de:08117:7405",
	KINDERGARTEN_BAIERECK:                                  "de:08117:7676",
	KINDERGARTEN_NAGOLD:                                    "de:08235:10124",
	KINDERGARTEN_ALBSTRASSE_ALDINGEN:                       "de:08118:3586",
	KINDERGARTEN_LEONBERGER_STRASSE_ALDINGEN:               "de:08118:3587",
	KIRBACHHOF_OCHSENBACH:                                  "de:08118:5713",
	KIRCHACKERSTRASSE_ST_BERNHARDT:                         "de:08116:4132",
	KIRCHBERG_GECHINGEN:                                    "de:08235:1535",
	KIRCHBERG_M_KIRCHBERG_M:                                "de:08119:7502",
	KIRCHE_HARTHAUSEN:                                      "de:08116:2024",
	KIRCHE_MUSBERG:                                         "de:08116:2108",
	KIRCHE_MUNCHINGEN:                                      "de:08118:2638",
	KIRCHE_RUIT:                                            "de:08116:2932",
	KIRCHE_WEILER_SCHORNDORF:                               "de:08119:3710",
	KIRCHE_LIEBERSBRONN:                                    "de:08116:4102",
	KIRCHE_OTLINGEN:                                        "de:08116:4340",
	KIRCHE_SCHLICHTEN:                                      "de:08119:4501",
	KIRCHE_HEGENLOHE:                                       "de:08116:4505",
	KIRCHE_DAGERSHEIM:                                      "de:08115:4632",
	KIRCHE_HASLACH:                                         "de:08115:4743",
	KIRCHE_GULTSTEIN:                                       "de:08115:4835",
	KIRCHE_AFFSTATT:                                        "de:08115:4846",
	KIRCHE_NEUENHAUS:                                       "de:08116:5000",
	KIRCHE_OCHSENBACH:                                      "de:08118:5712",
	KIRCHE_STEINENBRONN:                                    "de:08115:7106",
	KIRCHE_MUHLHAUSEN_IM_TALE:                              "de:08117:66",
	KIRCHE_SCHLAT:                                          "de:08117:3018",
	KIRCHE_GAMMELSHAUSEN:                                   "de:08117:3323",
	KIRCHE_HOHENSTADT_GP:                                   "de:08117:3501",
	KIRCHE_ROSSWALDEN:                                      "de:08117:7406",
	KIRCHE_NENNINGEN:                                       "de:08117:7682",
	KIRCHE_SCHLIERBACH:                                     "de:08117:7803",
	KIRCHE_WEITMARS:                                        "de:08136:7037",
	KIRCHE_MACHTOLSHEIM:                                    "de:08425:2641",
	KIRCHE_SANKT_MAGNUS_WERNAU_N:                           "de:08116:6726",
	KIRCHHEIM_N_KIRCHHEIM_N:                                "de:08118:5723",
	KIRCHHEIM_T_KIRCHHEIM_T:                                "de:08116:4211",
	KIRCHHEIMER_STRHEGELSTRASSE_WEILHEIM_T:                 "de:08116:4395",
	KIRCHHEIMER_STRASSE_STUTTGART:                          "de:08111:2430",
	KIRCHHEIMER_STRASSE_WENDLINGEN_N:                       "de:08116:3827",
	KIRCHHEIMER_STRASSE_NOTZINGEN:                          "de:08116:4216",
	KIRCHHEIMER_STRASSE_NURTINGEN:                          "de:08116:4376",
	KIRCHHEIMER_STRASSE_BONNIGHEIM:                         "de:08118:5726",
	KIRCHHEIMER_STRASSE_KONGEN:                             "de:08116:7819",
	KIRCHHEIMER_STRASSE_ZELL_A:                             "de:08117:2115",
	KIRCHPLATZ_WELZHEIM:                                    "de:08119:5061",
	KIRCHPLATZ_OBERBRUDEN:                                  "de:08119:5379",
	KIRCHPLATZ_GRUNBACH:                                    "de:08119:5805",
	KIRCHPLATZ_BAD_BOLL:                                    "de:08117:3334",
	KIRCHPLATZRATHAUS_RECHBERGHAUSEN:                       "de:08117:4020",
	KIRCHSTRASSE_BOHMENKIRCH:                               "de:08117:204",
	KIRCHSTRASSE_HEININGEN_GP:                              "de:08117:3308",
	KIRCHSTRASSENELKENWEG_DEIZISAU:                         "de:08116:6877",
	KIRCHTALSTRASSE_STUTTGART:                              "de:08111:107",
	KIRCHTALSTRASSE_KORNWESTHEIM:                           "de:08118:5448",
	KIRSCHENHARDTHOF_BURGSTALL_BURGSTETTEN:                 "de:08119:3617",
	KIRSCHHALDENWEG_BESIGHEIM:                              "de:08118:3501",
	KITTENESHALDE_KIRCHHEIM_T:                              "de:08116:3949",
	KITZENER_STRASSE_OTTENBACH:                             "de:08117:6015",
	KLAFFENBACH_RUDERSBERG:                                 "de:08119:4998",
	KLAFFENSTEINSTRASSE_BOBLINGEN:                          "de:08115:3125",
	KLAIBERSTRASSE_ENZWEIHINGEN:                            "de:08118:6955",
	KLEIDERWERK_ADLER_NECKARTENZLINGEN:                     "de:08116:4311",
	KLEINBOTTWAR_KLEINBOTTWAR:                              "de:08118:5604",
	KLEINBOTTWARER_STRASSE_STEINHEIM_M:                     "de:08118:5922",
	KLEINER_OSTRING_STUTTGART:                              "de:08111:2382",
	KLEINER_SCHLOSSPLATZ_STUTTGART:                         "de:08111:70",
	KLEINERLACH_GROSSERLACH:                                "de:08119:4967",
	KLEINFELDFRIEDHOF_FELLBACH:                             "de:08119:2499",
	KLEINGARTEN_SCHORNDORF:                                 "de:08119:5225",
	KLEINGARTENANLAGE_JEBENHAUSEN:                          "de:08117:2099",
	KLEINGLATTBACHER_STRASSE_ENSINGEN:                      "de:08118:5654",
	KLEINHEPPACHER_STRASSE_GROSSHEPPACH:                    "de:08119:5172",
	KLEINHEPPACHER_STRASSE_BEINSTEIN:                       "de:08119:6636",
	KLEINHOHENHEIM_STUTTGART:                               "de:08111:2186",
	KLEINSACHSENHEIMER_STRASSE_METTERZIMMERN:               "de:08118:5668",
	KLEINSACHSENHEIMER_STRASSE_GROSSSACHSENHEIM:            "de:08118:5674",
	KLEISTSTRASSE_STUTTGART:                                "de:08111:2248",
	KLINGEN_MURRHARDT:                                      "de:08119:5296",
	KLINGENMUHLE_WELZHEIM:                                  "de:08119:5284",
	KLINGENSTRASSE_MUSBERG:                                 "de:08116:2110",
	KLINGENSTRASSE_AFFALTERBACH:                            "de:08118:3601",
	KLINGENSTRASSE_BITTENFELD:                              "de:08119:5157",
	KLINIK_AM_EICHERT_GOPPINGEN:                            "de:08117:9006",
	KLINIK_SCHILLERHOHE_GERLINGEN:                          "de:08118:2312",
	KLINIKSCHULE_MARKGRONINGEN:                             "de:08118:3329",
	KLINIKUM_LUDWIGSBURG:                                   "de:08118:5498",
	KLINIKUM_ESSLINGEN_OBERESSLINGEN:                       "de:08116:6763",
	KLINKERNFELD_BERNHAUSEN:                                "de:08116:2087",
	KLOPFERBACH_GROSSASPACH:                                "de:08119:5357",
	KLOSTER_LORCH:                                          "de:08136:7017",
	KLOSTERGARTEN_SINDELFINGEN:                             "de:08115:4607",
	KLOSTERPLATZ_OEFFINGEN:                                 "de:08119:2373",
	KLOSTERSTRASSE_STETTEN_I_R:                             "de:08119:4931",
	KLOSTERWEG_KONGEN:                                      "de:08116:6777",
	KNAPPENWEG_STUTTGART:                                   "de:08111:2604",
	KNIEBISSTRASSE_MAGSTADT:                                "de:08115:3300",
	KOHLENWEG_ALDINGEN:                                     "de:08118:5586",
	KOHLERWEG_HEGENSBERG:                                   "de:08116:4099",
	KOHLPLATTE_UNTERJETTINGEN:                              "de:08115:7022",
	KOHLSTRASSE_LUDWIGSBURG:                                "de:08118:5455",
	KOLPINGSIEDLUNG_STUTTGART:                              "de:08111:2563",
	KOLPINGSTRASSE_WENDLINGEN_N:                            "de:08116:3832",
	KONIGWILHELMPLATZ_MARBACH_N:                            "de:08118:3537",
	KONIGSALLEE_WEIL:                                       "de:08116:4015",
	KONIGSBAU_PLUDERHAUSEN:                                 "de:08119:5251",
	KONIGSBERGER_STRASSE_EHNINGEN:                          "de:08115:4757",
	KONIGSBERGER_STRASSE_WERNAU_N:                          "de:08116:6724",
	KONIGSBERGER_STRASSE_HERRENBERG:                        "de:08115:7038",
	KONIGSBRONNHOF_RUDERSBERG:                              "de:08119:5033",
	KONIGSKNOLL_SINDELFINGEN:                               "de:08115:4660",
	KONIGSTRASSLE_STUTTGART:                                "de:08111:2565",
	KONRADHAUSSMANNWEG_SCHORNDORF:                          "de:08119:7575",
	KONRADWIDERHOLTHALLE_KIRCHHEIM_T:                       "de:08116:4333",
	KORBER_STEIGE_WAIBLINGEN:                               "de:08119:5408",
	KORBER_STRASSE_KLEINHEPPACH:                            "de:08119:6633",
	KORNBECKSTRASSE_LUDWIGSBURG:                            "de:08118:5501",
	KORNBERGSATTEL_GRUIBINGEN:                              "de:08117:3235",
	KORNBERGWEG_PLOCHINGEN:                                 "de:08116:7823",
	KORNHAUSPLATZ_GOPPINGEN:                                "de:08117:1058",
	KORNTAL_KORNTAL:                                        "de:08118:7603",
	KORNTALER_STRASSE_STUTTGART:                            "de:08111:102",
	KORNTALER_STRASSE_LEONBERG:                             "de:08115:2366",
	KORNTALERSCHUBARTSTRASSE_DITZINGEN:                     "de:08118:3014",
	KORNWESTHEIM_KORNWESTHEIM:                              "de:08118:1402",
	KORNWESTHEIMER_STRASSE_STUTTGART:                       "de:08111:105",
	KORNWESTHEIMER_STRASSE_MUNCHINGEN:                      "de:08118:4519",
	KORSCHSTRASSE_ZELL_ESSLINGEN:                           "de:08116:4169",
	KORSCHTAL_DEIZISAU:                                     "de:08116:5016",
	KOTTWEIL_STEINACH_BERGLEN:                              "de:08119:5279",
	KRAFTWERK_MUNSTER_STUTTGART:                            "de:08111:280",
	KRAHERWALD_STUTTGART:                                   "de:08111:2401",
	KRANKENHAUS_LEONBERG:                                   "de:08115:2327",
	KRANKENHAUS_MARBACH_N:                                  "de:08118:3533",
	KRANKENHAUS_SINDELFINGEN:                               "de:08115:4666",
	KRANKENHAUS_BIETIGHEIM:                                 "de:08118:5656",
	KRANKENHAUS_VAIHINGEN_E:                                "de:08118:5873",
	KRANKENHAUS_GEISLINGEN_STEIGE:                          "de:08117:212",
	KRANKENHAUS_SEEBACH_GEISLINGEN_STEIGE:                  "de:08117:4",
	KRANSTRASSE_STUTTGART:                                  "de:08111:2347",
	KRAPFENREUT_EBERSBACH_F:                                "de:08117:7639",
	KREHLSTRASSE_STUTTGART:                                 "de:08111:6201",
	KREHWINKEL_ASPERGLEN:                                   "de:08119:5114",
	KREISBERUFSSCHULE_LUDWIGSBURG:                          "de:08118:3454",
	KREISBERUFSSCHULE_BIETIGHEIM:                           "de:08118:5616",
	KREISBERUFSSCHULZENTRUM_BACKNANG:                       "de:08119:3611",
	KREISKRANKENHAUS_BOBLINGEN:                             "de:08115:4677",
	KREISKRANKENHAUS_B28_HERRENBERG:                        "de:08115:4837",
	KREISLERSTRASSE_SCHOPFLOCH_LENNINGEN:                   "de:08116:4423",
	KREISSPARKASSE_BALTMANNSWEILER:                         "de:08116:4075",
	KREISSPARKASSE_PLOCHINGEN:                              "de:08116:4226",
	KREISSPARKASSE_OBERBOIHINGEN:                           "de:08116:4299",
	KREISSPARKASSE_BORTLINGEN:                              "de:08117:4033",
	KREISSTRASSE_UNTERKIRNECK:                              "de:08136:7029",
	KREISTIERHEIM_BOBLINGEN:                                "de:08115:4717",
	KREISVERKEHR_SCHANBACH:                                 "de:08116:4156",
	KREISVERKEHR_HIRSCHLANDEN:                              "de:08118:6490",
	KREISVERKEHR_GAMMELSHAUSEN:                             "de:08117:3322",
	KREMSER_STRASSE_BOBLINGEN:                              "de:08115:4707",
	KRETTENHOF_GOPPINGEN:                                   "de:08117:1104",
	KREUZACKER_BONLANDEN:                                   "de:08116:2091",
	KREUZBRUNNEN_SCHARNHAUSER_PARK:                         "de:08116:2972",
	KREUZSTRASSE_STUTTGART:                                 "de:08111:2701",
	KREUZSTRASSE_OBERESSLINGEN:                             "de:08116:4029",
	KREUZSTRASSE_DATZINGEN:                                 "de:08115:4773",
	KREUZSTRASSE_KAISERSBACH:                               "de:08119:5071",
	KREUZSTRASSE_OBERSTENFELD:                              "de:08118:5614",
	KREUZUNG_AICHSCHIESS:                                   "de:08116:4155",
	KREUZUNG_HEININGEN_BACKNANG:                            "de:08119:4974",
	KREUZUNG_AICHELBERG_GP:                                 "de:08117:2060",
	KREUZUNG_ADELBERG_RECHBERGHAUSEN:                       "de:08117:4028",
	KREUZUNG_ESCHENBACHE_EISLINGEN_F:                       "de:08117:6012",
	KREUZUNG_GRUIBINGEN_AUENDORF_GAMMELSHAUSEN:             "de:08117:3224",
	KREUZWEGACKER_STEINHEIM_M:                              "de:08118:6995",
	KRIEGERDENKMAL_HEGENSBERG:                              "de:08116:4101",
	KRIEGERDENKMAL_GROTZINGEN:                              "de:08116:4472",
	KRIEGSBERG_HOHENACKER:                                  "de:08119:7548",
	KRING_STEINENBRONN:                                     "de:08115:7118",
	KROATENHOF_NURTINGEN:                                   "de:08116:2953",
	KRONE_GRONAU:                                           "de:08118:3347",
	KRONE_BAIERECK:                                         "de:08117:7663",
	KRONEALTENSTADT_GEISLINGEN_STEIGE:                      "de:08117:9",
	KRONENBRUNNEN_PLATTENHARDT:                             "de:08116:6735",
	KRONENPLATZ_WINNENDEN:                                  "de:08119:3637",
	KRONENSTRASSE_BERKHEIM:                                 "de:08116:4152",
	KRONENSTRASSE_AICHELBERG_AICHWALD:                      "de:08116:4165",
	KRONENZENTRUM_BIETIGHEIM:                               "de:08118:5659",
	KRONERHOF_AUFHAUSEN_GEISLINGEN:                         "de:08117:511",
	KRUICHLING_KIRCHHEIM_T:                                 "de:08116:3944",
	KRUMMBACHTAL_GERLINGEN:                                 "de:08118:2317",
	KRUMMENACKER_KRUMMENACKER:                              "de:08116:4111",
	KRUMMHARDT_ABZWEIG_AICHSCHIESS:                         "de:08116:4163",
	KRUMMHARDT_ORT_KRUMMHARDT:                              "de:08116:4164",
	KRUMMW_SCHWARZENGASSE_EISLINGEN_F:                      "de:08117:6011",
	KRUMMWALDEN_BRUCKENSTR_EISLINGEN_F:                     "de:08117:6010",
	KUCHEN_KUCHEN_F:                                        "de:08117:159",
	KUCHENGRUNDOVR_BACKNANG:                                "de:08119:3639",
	KUHWASEN_STUTTGART:                                     "de:08111:2437",
	KULLENWIESEN_HARDT:                                     "de:08116:2951",
	KUNKELINHALLE_SCHORNDORF:                               "de:08119:5961",
	KUNKELINSCHULE_SCHORNDORF:                              "de:08119:3716",
	KUNSTAKADEMIE_STUTTGART:                                "de:08111:2241",
	KUNSTHALLE_GOPPINGEN:                                   "de:08117:4004",
	KUPPINGER_STRASSE_NUFRINGEN:                            "de:08115:4736",
	KURKLINIK_BAD_DITZENBACH:                               "de:08117:54",
	KURMARKER_KASERNE_STUTTGART:                            "de:08111:6202",
	KURSAAL_STUTTGART:                                      "de:08111:322",
	KURTSCHUMACHERSTRASSE_STUTTGART:                        "de:08111:2581",
	KURZACH_NASSACH:                                        "de:08119:5049",
	KURZE_STRASSE_GEMMRIGHEIM:                              "de:08118:6812",
	LACHENSTRASSE_SCHONAICH:                                "de:08115:3258",
	LACHENTORSTRASSE_HOFINGEN:                              "de:08115:6049",
	LAIBLINGER_WEG_SCHWIEBERDINGEN:                         "de:08118:3349",
	LAICHLESTRASSE_GERLINGEN:                               "de:08118:320",
	LAIHLE_STUTTGART:                                       "de:08111:2420",
	LAILENSACKER_PLATTENHARDT:                              "de:08116:6730",
	LAMBERTWEG_STUTTGART:                                   "de:08111:2614",
	LAMM_NECKARGRONINGEN:                                   "de:08118:4890",
	LAMM_GOSBACH:                                           "de:08117:3504",
	LAMM_TREFFELHAUSEN:                                     "de:08117:7675",
	LAMMTAL_GARTRINGEN:                                     "de:08115:7153",
	LANDAUER_STRASSE_STUTTGART:                             "de:08111:151",
	LANDHAUS_STUTTGART:                                     "de:08111:6352",
	LANDHAUSKREUZUNG_STUTTGART:                             "de:08111:6353",
	LANDHAUSSIEDLUNG_MAICHINGEN:                            "de:08115:3207",
	LANDRATSAMT_LUDWIGSBURG:                                "de:08118:5530",
	LANDRATSAMT_GOPPINGEN:                                  "de:08117:4902",
	LANGE_ANWANDEN_SINDELFINGEN:                            "de:08115:4656",
	LANGE_MORGEN_WEILHEIM_T:                                "de:08116:4396",
	LANGE_STRASSE_SIELMINGEN:                               "de:08116:2035",
	LANGE_STRASSE_MUNDELSHEIM:                              "de:08118:5574",
	LANGE_WEIDEN_WINNENDEN:                                 "de:08119:5308",
	LANGENBERG_WELZHEIM:                                    "de:08119:7525",
	LANGENBUHL_ELTINGEN:                                    "de:08115:7305",
	LANGER_WEG_WALDENBRONN:                                 "de:08116:4126",
	LANGHANS_BEILSTEIN:                                     "de:08125:4672",
	LANGMANTEL_HOHENHASLACH:                                "de:08118:5547",
	LANNERSTRASSE_OPPELSBOHM:                               "de:08119:5276",
	LAPP_KABEL_STUTTGART:                                   "de:08111:2599",
	LAUBERT_ERBSTETTEN:                                     "de:08119:3518",
	LAUBSANGERWEG_FELLBACH:                                 "de:08119:2648",
	LAUCHHAU_STUTTGART:                                     "de:08111:2591",
	LAUERHALDE_WARMBRONN:                                   "de:08115:4539",
	LAUFENMUHLE_WELZHEIM:                                   "de:08119:5127",
	LAUSTRASSE_STUTTGART:                                   "de:08111:6164",
	LAUTERGARTEN_DONZDORF:                                  "de:08117:7624",
	LAUTERN_SULZBACH_M:                                     "de:08119:5042",
	LAUTERTAL_SULZBACH_M:                                   "de:08119:4953",
	LAUTERTAL_WUSTENROT:                                    "de:08125:1991",
	LECHTSTRASSE_NECKARWEIHINGEN:                           "de:08118:3426",
	LEDERBERG_STUTTGART:                                    "de:08111:2536",
	LEDERGASSE_SCHWABISCH_GMUND:                            "de:08136:3037",
	LEDERSTRASSE_SCHORNDORF:                                "de:08119:6992",
	LEERWASEN_NEIDLINGEN:                                   "de:08116:4408",
	LEHENBRUCKE_ASPERG:                                     "de:08118:6946",
	LEHENSTRASSE_STUTTGART:                                 "de:08111:2414",
	LEHENWEILER_AIDLINGEN:                                  "de:08115:4775",
	LEHENWEILER_ABZWEIG_AIDLINGEN:                          "de:08115:4774",
	LEHLESTRASSE_FAURNDAU:                                  "de:08117:2004",
	LEHNENBERG_REICHENBACH_BERGLEN:                         "de:08119:5290",
	LEHNENBERG_KREUZUNG_REICHENBACH_BERGLEN:                "de:08119:5202",
	LEIBNIZSTRASSE_STUTTGART:                               "de:08111:2407",
	LEIBNIZSTRASSE_KORNWESTHEIM:                            "de:08118:3364",
	LEINFELDEN_LEINFELDEN:                                  "de:08116:175",
	LEINFELDENER_STRASSE_STUTTGART:                         "de:08111:2566",
	LEINFELDER_HAUS_LEINFELDEN:                             "de:08116:7176",
	LEINTEL_REICHENBACH_F:                                  "de:08116:3903",
	LEINTELSTRASSE_EBERSBACH_F:                             "de:08117:8015",
	LEIPZIGER_PLATZ_STUTTGART:                              "de:08111:2050",
	LEIPZIGER_STRASSE_SINDELFINGEN:                         "de:08115:4615",
	LEMBERGSCHULE_NAGOLD:                                   "de:08235:10135",
	LENAUDENKMAL_ESSLINGEN_N:                               "de:08116:4063",
	LENAUSTRASSE_ALBERSHAUSEN:                              "de:08117:2026",
	LENBACHSTRASSE_GOPPINGEN:                               "de:08117:1062",
	LENZHALDE_ESSLINGEN_N:                                  "de:08116:3999",
	LENZHALDE_SCHARNHAUSEN:                                 "de:08116:5012",
	LENZHALDE_KORNWESTHEIM:                                 "de:08118:6835",
	LEOCENTER_ELTINGEN:                                     "de:08115:2328",
	LEOBAD_ELTINGEN:                                        "de:08115:7304",
	LEOBENER_STRASSE_STUTTGART:                             "de:08111:156",
	LEONARDODAVINCIPLATZ_BOBLINGEN:                         "de:08115:3141",
	LEONBERG_LEONBERG:                                      "de:08115:7301",
	LEONBERGER_DREIECK_ELTINGEN:                            "de:08115:2359",
	LEONBERGER_STRASSE_ELTINGEN:                            "de:08115:2360",
	LEONBERGER_STRASSE_GERLINGEN:                           "de:08118:4554",
	LEONBERGER_STRASSE_HEIMSHEIM:                           "de:08236:3514",
	LERCHENACKER_OBERESSLINGEN:                             "de:08116:3459",
	LERCHENACKER_BACKNANG:                                  "de:08119:6933",
	LERCHENBERG_ORTSMITTE_GOPPINGEN:                        "de:08117:4915",
	LERCHENRAINSCHULE_STUTTGART:                            "de:08111:6100",
	LERCHENSTRASSE_OBERBOIHINGEN:                           "de:08116:4298",
	LERCHENSTRASSE_WEILHEIM_T:                              "de:08116:4362",
	LERCHENSTRASSE_GROSSASPACH:                             "de:08119:5356",
	LERCHENWEG_WALDENBUCH:                                  "de:08115:3123",
	LERCHENWIESEN_STUTTGART:                                "de:08111:2560",
	LESSINGSTRASSE_SCHONAICH:                               "de:08115:4814",
	LESSINGSTRASSE_RUTESHEIM:                               "de:08115:5835",
	LESSINGSTRASSE_GOPPINGEN:                               "de:08117:9002",
	LETTENACKER_PLOCHINGEN:                                 "de:08116:4232",
	LETTENSTICH_WELZHEIM:                                   "de:08119:7527",
	LETTENSTRASSE_NEUHAUSEN_F:                              "de:08116:2897",
	LEUTENBACHER_STRASSE_NELLMERSBACH:                      "de:08119:5320",
	LIBANONSTRASSE_STUTTGART:                               "de:08111:2201",
	LICHTENBERGER_STRASSE_OBERSTENFELD:                     "de:08118:5594",
	LICHTENSTEINSTRASSE_BACKNANG:                           "de:08119:5351",
	LICHTENSTERN_GYMNASIUM_GROSSSACHSENHEIM:                "de:08118:6929",
	LIEBENAU_WALDENBUCH:                                    "de:08115:3118",
	LIEBFRAUENKIRCHE_EISLINGEN_F:                           "de:08117:1010",
	LIEMANNSKLINGE_ABZWEIG_SULZBACH_M:                      "de:08119:7507",
	LIEMERSBACH_GROSSERLACH:                                "de:08119:7505",
	LILIENSTRASSE_HOLZGERLINGEN:                            "de:08115:4872",
	LILIENSTRASSE_MALMSHEIM:                                "de:08115:5850",
	LILIENSTRASSE_ROMMELSHAUSEN:                            "de:08119:6954",
	LILIENTHALSTRASSE_STUTTGART:                            "de:08111:2348",
	LIMBURGSCHULEBISSINGER_STRASSE_WEILHEIM_T:              "de:08116:3896",
	LIMBURGSCHULEKELTERSTRASSE_WEILHEIM_T:                  "de:08116:5908",
	LIMESWEG_GRAB:                                          "de:08119:5143",
	LIMESWEG_MAITIS:                                        "de:08117:1216",
	LINDACH_RECHBERGHAUSEN:                                 "de:08117:4043",
	LINDACHSCHULE_STETTEN_LEINFELDENECHT:                   "de:08116:2122",
	LINDE_HARTHAUSEN:                                       "de:08116:2023",
	LINDE_KONGEN:                                           "de:08116:4089",
	LINDE_UNTERENSINGEN:                                    "de:08116:4093",
	LINDE_KAPPISHAUSERN:                                    "de:08116:4493",
	LINDE_JEBENHAUSEN:                                      "de:08117:2101",
	LINDENMUSEUM_STUTTGART:                                 "de:08111:2196",
	LINDENHOF_MOTZINGEN:                                    "de:08115:6739",
	LINDENHOF_GEISLINGEN_STEIGE:                            "de:08117:119",
	LINDENHOF_BOHMENKIRCH:                                  "de:08117:205",
	LINDENPLATZ_OBERENSINGEN:                               "de:08116:2925",
	LINDENPLATZ_HEBSACK:                                    "de:08119:3706",
	LINDENPLATZ_NEUFFEN:                                    "de:08116:4442",
	LINDENPLATZ_UNTERWEISSACH:                              "de:08119:5030",
	LINDENPLATZ_STEINBACH:                                  "de:08119:6046",
	LINDENSCHULE_GEISLINGEN_STEIGE:                         "de:08117:143",
	LINDENSTRASSE_BEMPFLINGEN:                              "de:08116:4315",
	LINDENSTRASSE_KORNWESTHEIM:                             "de:08118:5467",
	LINDENSTRASSE_HOCHBERG:                                 "de:08118:6822",
	LINDENSTRASSE_STEINENBRONN:                             "de:08115:7120",
	LINDENSTRASSE_BISSINGEN_LB:                             "de:08118:7425",
	LINDENSTRASSE_HEININGEN_GP:                             "de:08117:3378",
	LINDENSTRASSE_WISSGOLDINGEN:                            "de:08136:7822",
	LINDENSTRASSE_MERKLINGEN:                               "de:08425:2636",
	LINDENTAL_SCHLECHTBACH:                                 "de:08119:5117",
	LINDENTALER_STRASSE_SCHLECHTBACH:                       "de:08119:6675",
	LINDENWEG_SCHORNDORF:                                   "de:08119:5973",
	LINDORFER_WEG_KIRCHHEIM_T:                              "de:08116:3945",
	LINDPAINTNERSTRASSE_STUTTGART:                          "de:08111:6213",
	LINGWIESEN_MUNCHINGEN:                                  "de:08118:7940",
	LINSENHOFEN_LINSENHOFEN:                                "de:08116:4433",
	LINSENHOFER_STRASSE_BEUREN:                             "de:08116:4434",
	LISEMEITNERGYMNASIUM_ALDINGEN:                          "de:08118:5590",
	LISEMEITNERSTR_OBERENSINGEN:                            "de:08116:2993",
	LISEMEITNERSTRASSE_GOPPINGEN:                           "de:08117:9021",
	LISEMEITNERSTRASSE_NAGOLD:                              "de:08235:4403",
	LISTSTRASSE_STUTTGART:                                  "de:08111:344",
	LISZTSTRASSE_MURRHARDT:                                 "de:08119:4945",
	LOBENROT_ORT_LOBENROT:                                  "de:08116:4161",
	LOBENROTER_HOF_SCHANBACH:                               "de:08116:4160",
	LOCHGAUER_STRASSE_BIETIGHEIM:                           "de:08118:5660",
	LOCHLESACKER_FREIBERG_N:                                "de:08118:6940",
	LOCHLESBERG_DOFFINGEN:                                  "de:08115:4768",
	LOHACKERSTRASSE_STUTTGART:                              "de:08111:6173",
	LOHMUHLE_WUSTENROT:                                     "de:08125:978",
	LORCH_LORCH:                                            "de:08136:7005",
	LORCHER_STRASSE_PFAHLBRONN:                             "de:08119:5092",
	LORTZINGSTRASSE_ROMMELSHAUSEN:                          "de:08119:3735",
	LOSACKER_STUTTGART:                                     "de:08111:319",
	LOSBURGSTRASSE_ALTBACH:                                 "de:08116:3802",
	LOSCHER_OST_MOGLINGEN:                                  "de:08118:3437",
	LOSCHER_WEST_MOGLINGEN:                                 "de:08118:3438",
	LOTTER_LUDWIGSBURG:                                     "de:08118:3410",
	LOWEN_WOLFSCHLUGEN:                                     "de:08116:2919",
	LOWEN_KUCHEN_F:                                         "de:08117:13",
	LOWEN_BORTLINGEN:                                       "de:08117:4032",
	LOWENBRUNNEN_FELLBACH:                                  "de:08119:3726",
	LOWENKELLER_SCHORNDORF:                                 "de:08119:4943",
	LOWENSTEINER_STRASSE_SPIEGELBERG:                       "de:08119:4959",
	LOWENTOR_STUTTGART:                                     "de:08111:6261",
	LOWENTORBRUCKE_STUTTGART:                               "de:08111:6114",
	LUDWIGFINCKHSTRASSE_WENDLINGEN_N:                       "de:08116:3830",
	LUDWIGHERRSTRASSE_KORNWESTHEIM:                         "de:08118:5464",
	LUDWIGUHLANDSCHULE_HEIMSHEIM:                           "de:08236:3515",
	LUDWIGSBURG_LUDWIGSBURG:                                "de:08118:7402",
	LUDWIGSBURGER_STRASSE_OEFFINGEN:                        "de:08119:2384",
	LUDWIGSBURGER_STRASSE_SCHWIEBERDINGEN:                  "de:08118:3328",
	LUDWIGSBURGER_STRASSE_MOGLINGEN:                        "de:08118:3439",
	LUDWIGSBURGER_STRASSE_TAMM:                             "de:08118:3451",
	LUDWIGSBURGER_STRASSE_WAIBLINGEN:                       "de:08119:4876",
	LUDWIGSBURGER_STRASSE_KORNWESTHEIM:                     "de:08118:5459",
	LUDWIGSBURGER_STRASSE_HOHENECK:                         "de:08118:5525",
	LUDWIGSBURGER_STRASSE_BISSINGEN_LB:                     "de:08118:5643",
	LUDWIGSBURGER_STRASSE_GROSSSACHSENHEIM:                 "de:08118:5649",
	LUDWIGSTRASSE_MUSBERG:                                  "de:08116:2109",
	LUDWIGSTRASSE_FAURNDAU:                                 "de:08117:7016",
	LUFTFRACHTZENTRUM_WEST_BERNHAUSEN:                      "de:08116:2090",
	LUGINSLAND_STUTTGART:                                   "de:08111:2506",
	LUGLENWEG_HERRENBERG:                                   "de:08115:3248",
	LUIKENWEG_STUTTGART:                                    "de:08111:2525",
	LUTHERKIRCHE_FELLBACH:                                  "de:08119:6040",
	LUTZENBERG_ALTHUTTE:                                    "de:08119:4991",
	LUTZESTRASSE_SCHNAIT:                                   "de:08119:5401",
	MAGELLANSTRASSE_ECHTERDINGEN:                           "de:08116:2141",
	MAGSTADT_MAGSTADT:                                      "de:08115:3196",
	MAGSTADTER_STRASSE_SCHAFHAUSEN:                         "de:08115:4584",
	MAHDENTALSTRASSE_SINDELFINGEN:                          "de:08115:2225",
	MAHDERKLINGE_STUTTGART:                                 "de:08111:2427",
	MAICHINGEN_MAICHINGEN:                                  "de:08115:3198",
	MAICKLERSTRASSE_FELLBACH:                               "de:08119:2639",
	MAIENPLATZ_BOBLINGEN:                                   "de:08115:3138",
	MAIENWALTERSTRASSE_SULZGRIES:                           "de:08116:4112",
	MAIERHOFE_WEILHEIM_T:                                   "de:08116:3897",
	MAIERHOFSTRASSE_LORCH:                                  "de:08136:7045",
	MAILLE_ESSLINGEN_N:                                     "de:08116:2800",
	MAINZER_ALLEE_LUDWIGSBURG:                              "de:08118:7432",
	MALMSHEIM_MALMSHEIM:                                    "de:08115:1301",
	MANFREDVONARDENNEALLEE_BACKNANG:                        "de:08119:7738",
	MANFREDWORNERSTRASSE_GOPPINGEN:                         "de:08117:9027",
	MANN__HUMMEL_LUDWIGSBURG:                               "de:08118:5475",
	MANNENBERG_RUDERSBERG:                                  "de:08119:4994",
	MANNENBERG_IM_HAU_RUDERSBERG:                           "de:08119:4993",
	MANNENWEILER_GRAB:                                      "de:08119:7519",
	MANNHOLZ_PFAHLBRONN:                                    "de:08119:6645",
	MANOLZWEILER_WINTERBACH:                                "de:08119:4081",
	MANOSQUER_STRASSE_LEINFELDEN:                           "de:08116:2129",
	MANSLER_STRASSE_KORB:                                   "de:08119:5417",
	MANZELLER_WEG_STUTTGART:                                "de:08111:3995",
	MANZEN_MANZENSTR_GOPPINGEN:                             "de:08117:3017",
	MANZEN_OBERE_SIEDLUNG_GOPPINGEN:                        "de:08117:3012",
	MANZEN_SPITZENBERGSTR_GOPPINGEN:                        "de:08117:3023",
	MANZEN_SPORTPLATZ_GOPPINGEN:                            "de:08117:3024",
	MARABUSTRASSE_STUTTGART:                                "de:08111:2483",
	MARBACH_N_MARBACH_N:                                    "de:08118:7503",
	MARBACHER_STRASSE_KIRCHBERG_M:                          "de:08119:3526",
	MARBACHER_STRASSE_AFFALTERBACH:                         "de:08118:3599",
	MARBACHER_STRASSE_HEININGEN_BACKNANG:                   "de:08119:3670",
	MARBACHER_STRASSE_GROSSASPACH:                          "de:08119:6518",
	MARCOPOLOWEG_STUTTGART:                                 "de:08111:97",
	MARCONIBRUCKE_STUTTGART:                                "de:08111:2464",
	MARCONISTRASSE_STUTTGART:                               "de:08111:2463",
	MARGARETENSTRASSE_PLUDERHAUSEN:                         "de:08119:5240",
	MARGARETENWEG_GERLINGEN:                                "de:08118:6897",
	MARIECURIESCHULE_RAMTEL:                                "de:08115:4546",
	MARIENSILBERBURGSTRASSE_STUTTGART:                      "de:08111:2410",
	MARIENHEIM_BACKNANG:                                    "de:08119:5346",
	MARIENHOSPITAL_STUTTGART:                               "de:08111:6101",
	MARIENPLATZ_STUTTGART:                                  "de:08111:6019",
	MARIENSTRASSE_STUTTGART:                                "de:08111:2054",
	MARIENSTRASSE_WERNAU_N:                                 "de:08116:4245",
	MARIENSTRASSE_UHINGEN:                                  "de:08117:2042",
	MARIENSTRASSE_FAURNDAU:                                 "de:08117:7006",
	MARKGRABEN_HEBSACK:                                     "de:08119:3707",
	MARKGRONINGER_STRASSE_EGLOSHEIM:                        "de:08118:3422",
	MARKGRONINGER_STRASSE_MOGLINGEN:                        "de:08118:3440",
	MARKGRONINGER_STRASSE_ASPERG:                           "de:08118:5481",
	MARKLINEUM_GOPPINGEN:                                   "de:08117:2002",
	MARKTBAUERNSTRASSE_DITZINGEN:                           "de:08118:3009",
	MARKTGASSE_WAIBLINGEN:                                  "de:08119:5386",
	MARKTPLATZ_WEILHEIM_T:                                  "de:08116:3894",
	MARKTPLATZ_ESSLINGEN_N:                                 "de:08116:4034",
	MARKTPLATZ_KIRCHHEIM_T:                                 "de:08116:4213",
	MARKTPLATZ_NECKARTENZLINGEN:                            "de:08116:4486",
	MARKTPLATZ_WEISSACH_BB:                                 "de:08115:4534",
	MARKTPLATZ_SULZBACH_M:                                  "de:08119:4905",
	MARKTPLATZ_FORNSBACH:                                   "de:08119:4925",
	MARKTPLATZ_ALFDORF:                                     "de:08119:5099",
	MARKTPLATZ_RUDERSBERG:                                  "de:08119:5119",
	MARKTPLATZ_HORRHEIM:                                    "de:08118:5738",
	MARKTPLATZ_WASCHENBEUREN:                               "de:08117:1106",
	MARKTPLATZ_GOPPINGEN:                                   "de:08117:3002",
	MARKTPLATZ_SALACH:                                      "de:08117:6004",
	MARKTPLATZ_SPARWIESEN:                                  "de:08117:7014",
	MARKTPLATZ_HEIMSHEIM:                                   "de:08236:3513",
	MARKTPLATZ_GSCHWEND:                                    "de:08136:6512",
	MARKTPLATZ_BOBLINGER_STRASSE_SINDELFINGEN:              "de:08115:5926",
	MARKTPLATZ_RATHAUS_SINDELFINGEN:                        "de:08115:5927",
	MARKTSTRASSE_BONLANDEN:                                 "de:08116:2078",
	MARKTSTRASSE_GOPPINGEN:                                 "de:08117:1037",
	MARKUSKIRCHE_STUTTGART:                                 "de:08111:2411",
	MARQUARDTSCHULE_PLOCHINGEN:                             "de:08116:2816",
	MARRENSTRASSE_DONZDORF:                                 "de:08117:7621",
	MARTINLUTHERSTRASSE_ECHTERDINGEN:                       "de:08116:2134",
	MARTINLUTHERSTRASSE_LUDWIGSBURG:                        "de:08118:5510",
	MARTINSHAUS_BESIGHEIM:                                  "de:08118:3561",
	MARTINSKIRCHE_STUTTGART:                                "de:08111:2252",
	MARTINSKIRCHE_KIRCHHEIM_T:                              "de:08116:4317",
	MARTINSTRASSE_NUSSDORF:                                 "de:08118:5625",
	MARXENHOF_KIRCHENKIRNBERG:                              "de:08119:5987",
	MASSSTABFABRIK_MAINHARDT:                               "de:08127:28268",
	MATHILDENSTRASSE_STUTTGART:                             "de:08111:150",
	MAUBACH_MAUBACH:                                        "de:08119:7601",
	MAUBACHER_STRASSE_BACKNANG:                             "de:08119:7513",
	MAUERACKER_TAILFINGEN:                                  "de:08115:7930",
	MAULBRONNER_STRASSE_HORRHEIM:                           "de:08118:5739",
	MAURENER_STRASSE_EHNINGEN:                              "de:08115:4753",
	MAURENER_WEG_BOBLINGEN:                                 "de:08115:4696",
	MAURER_STEINENBRONN:                                    "de:08115:7119",
	MAUSERSTRASSE_OSSWEIL:                                  "de:08118:5999",
	MAXBORNGYMNASIUM_BACKNANG:                              "de:08119:3609",
	MAXEYTHDAIMLERSTRASSE_WINNENDEN:                        "de:08119:3641",
	MAXEYTHSEE_STUTTGART:                                   "de:08111:6273",
	MAXEYTHSTRASSE_WINNENDEN:                               "de:08119:3635",
	MAXEYTHSTRASSE_WERNAU_N:                                "de:08116:4248",
	MAXEYTHSTRASSE_KORNWESTHEIM:                            "de:08118:5456",
	MAXPLANCKINSTITUTE_STUTTGART:                           "de:08111:2589",
	MAXPLANCKSTRASSE_FELLBACH:                              "de:08119:2643",
	MAXPLANCKSTRASSE_OBERESSLINGEN:                         "de:08116:4043",
	MAXPLANCKSTRASSE_NEUFFEN:                               "de:08116:4443",
	MAXPLANCKSTRASSE_KORNWESTHEIM:                          "de:08118:6809",
	MAYBACHSTRASSE_STUTTGART:                               "de:08111:158",
	MAYBACHSTRASSE_SCHONAICH:                               "de:08115:3133",
	MAYBACHSTRASSE_KORNWESTHEIM:                            "de:08118:5468",
	MAYBACHSTRASSE_FRICKENHAUSEN:                           "de:08116:5902",
	MAYBACHSTRASSE_RECHBERGHAUSEN:                          "de:08117:1204",
	MEDIUS_KLINIK_RUIT:                                     "de:08116:2913",
	MEDIUS_KLINIK_KIRCHHEIM_T:                              "de:08116:4265",
	MEDIUS_KLINIK_NURTINGEN:                                "de:08116:4445",
	MEHRZWECKHALLE_OSSWEIL:                                 "de:08118:6800",
	MEISENWEG_WAIBLINGEN:                                   "de:08119:3656",
	MELLIBEESESTRASSE_SINDELFINGEN:                         "de:08115:3241",
	MENDELSSOHNSTRASSE_WAIBLINGEN:                          "de:08119:7091",
	MENZELSTRASSE_FREIBERG_N:                               "de:08118:3560",
	MERCEDESBENZ_STUTTGART:                                 "de:08111:4066",
	MERCEDESBENZ_HST_A_SINDELFINGEN:                        "de:08115:2222",
	MERCEDESBENZ_HST_B_SINDELFINGEN:                        "de:08115:2226",
	MERCEDESBENZ_HST_C_SINDELFINGEN:                        "de:08115:2228",
	MERCEDESBENZ_HST_D_SINDELFINGEN:                        "de:08115:2229",
	MERCEDESBENZ_HST_E_SINDELFINGEN:                        "de:08115:2230",
	MERCEDESBENZ_HST_F_SINDELFINGEN:                        "de:08115:2231",
	MERCEDESBENZ_HST_I_SINDELFINGEN:                        "de:08115:2234",
	MERCEDESBENZ_HST_K_SINDELFINGEN:                        "de:08115:2237",
	MERCEDESBENZ_HST_N_SINDELFINGEN:                        "de:08115:2281",
	MERCEDESBENZ_HST_O_SINDELFINGEN:                        "de:08115:2282",
	MERCEDESBENZ_HST_P_SINDELFINGEN:                        "de:08115:2283",
	MERCEDESBENZ_P305_SINDELFINGEN:                         "de:08115:3246",
	MERCEDESBENZ_P307_SINDELFINGEN:                         "de:08115:6746",
	MERCEDESBENZ_TOR_I_SINDELFINGEN:                        "de:08115:4612",
	MERCEDESBENZ_TOR_III_SINDELFINGEN:                      "de:08115:4551",
	MERCEDESBENZ_TOR_VII_SINDELFINGEN:                      "de:08115:3202",
	MERCEDESBENZ_WELT_STUTTGART:                            "de:08111:2331",
	MERCEDESSTR_UHINGEN:                                    "de:08117:2007",
	MERCEDESSTRASSE_STUTTGART:                              "de:08111:30",
	MERCEDESSTRASSE_HIRSCHLANDEN:                           "de:08118:6491",
	MERKELSCHES_BAD_ESSLINGEN_N:                            "de:08116:4006",
	MERKLINGEN_MERKLINGEN:                                  "de:08425:2626",
	MERKLINGER_STRASSE_WEIL_DER_STADT:                      "de:08115:4575",
	MERKLINGER_STRASSE_MACHTOLSHEIM:                        "de:08425:2634",
	MESSE_WEST_ECHTERDINGEN:                                "de:08116:7182",
	MESSEHALLE_SINDELFINGEN:                                "de:08115:2214",
	METLANGEN_ABZWEIG_SCHWABISCH_GMUND:                     "de:08136:3503",
	METLANGEN_ORT_SCHWABISCH_GMUND:                         "de:08136:3502",
	METTELBACH_KIRCHENKIRNBERG:                             "de:08119:6674",
	METTELBERG_FORNSBACH:                                   "de:08119:5305",
	METTERBRUCKE_GROSSSACHSENHEIM:                          "de:08118:5650",
	METTINGEN_METTINGEN:                                    "de:08116:1801",
	METZELHOF_LORCH:                                        "de:08136:7022",
	METZELWIESEN_MERKLINGEN:                                "de:08115:5841",
	METZGERHAU_STUTTGART:                                   "de:08111:2415",
	METZINGER_STRASSE_KOHLBERG:                             "de:08116:3821",
	METZINGER_STRASSE_NURTINGEN:                            "de:08116:4305",
	METZINGER_STRASSE_NECKARTENZLINGEN:                     "de:08116:4309",
	METZINGER_STRASSE_BONLANDEN:                            "de:08116:5798",
	METZINGER_STRASSE_GROSSBETTLINGEN:                      "de:08116:6956",
	METZSTRASSE_STUTTGART:                                  "de:08111:27",
	MICHELAU_SCHLECHTBACH:                                  "de:08119:5791",
	MICHELAU_ORTSMITTE_SCHLECHTBACH:                        "de:08119:5113",
	MICHELAUER_STRASSE_STEINENBERG:                         "de:08119:5112",
	MICHELBERGGYMNASIUM_GEISLINGEN_STEIGE:                  "de:08117:84",
	MIEDELSBACHSTEINENBERG_MIEDELSBACH:                     "de:08119:5790",
	MIEDELSBACHER_STRASSE_HAUBERSBRONN:                     "de:08119:5108",
	MIKROZENTRUM_WAIBLINGEN:                                "de:08119:5406",
	MILANWEG_KIRCHHEIM_T:                                   "de:08116:4352",
	MILCHHOF_STUTTGART:                                     "de:08111:297",
	MILCHZENTRALE_GEISLINGEN_STEIGE:                        "de:08117:115",
	MILLOCKERSTRASSE_STUTTGART:                             "de:08111:211",
	MINERALBADER_STUTTGART:                                 "de:08111:28",
	MINIGOLFPLATZ_RUTESHEIM:                                "de:08115:5825",
	MITTE_SCHLAITDORF:                                      "de:08116:2062",
	MITTE_SCHWIEBERDINGEN:                                  "de:08118:3327",
	MITTE_GERADSTETTEN:                                     "de:08119:3703",
	MITTE_REUDERN:                                          "de:08116:4380",
	MITTE_NEIDLINGEN:                                       "de:08116:4407",
	MITTE_NUFRINGEN:                                        "de:08115:4735",
	MITTE_HAUBERSBRONN:                                     "de:08119:5803",
	MITTE_PATTONVILLE:                                      "de:08118:6982",
	MITTEHELENENSTRASSE_MARKGRONINGEN:                      "de:08118:6888",
	MITTEVOLKSBANK_MARKGRONINGEN:                           "de:08118:3317",
	MITTEWERNERSTRASSE_MARKGRONINGEN:                       "de:08118:3313",
	MITTELBRUDEN_OBERBRUDEN:                                "de:08119:5378",
	MITTELFISCHBACH_ABZWEIGUNG_GROSSERLACH:                 "de:08119:4964",
	MITTELSCHONTAL_BACKNANG:                                "de:08119:6632",
	MITTELWEILER_PFAHLBRONN:                                "de:08119:6780",
	MITTENBUHL_DOFFINGEN:                                   "de:08115:6752",
	MITTLERE_BRUCKE_SCHORNDORF:                             "de:08119:5050",
	MITTLERE_MUHLE_HOLZGERLINGEN:                           "de:08115:4785",
	MITTLERE_UFERSTRASSE_SCHORNDORF:                        "de:08119:5254",
	MITTNACHTSTRASSE_STUTTGART:                             "de:08111:296",
	MOGLINGER_STRASSE_ASPERG:                               "de:08118:3402",
	MOGLINGER_STRASSE_PFLUGFELDEN:                          "de:08118:3429",
	MOHRINGEN_BF_MOHRINGEN:                                 "de:08111:6169",
	MOHRINGEN_FREIBAD_STUTTGART:                            "de:08111:6171",
	MOHRINGEN_RATHAUS_STUTTGART:                            "de:08111:2583",
	MOHRINGER_STRASSE_SIELMINGEN:                           "de:08116:2143",
	MOLTKESTRASSE_KUCHEN_F:                                 "de:08117:12",
	MONCHFELD_STUTTGART:                                    "de:08111:6288",
	MONCHHOF_KAISERSBACH:                                   "de:08119:5078",
	MONCHSBERGSTRASSE_FREIBERG_N:                           "de:08118:5565",
	MONCHSBRUCKE_SCHORNDORF:                                "de:08119:5220",
	MONCHSBRUNNEN_SINDELFINGEN:                             "de:08115:7201",
	MONCHWEG_STUTTGART:                                     "de:08111:2527",
	MONREPOS_EGLOSHEIM:                                     "de:08118:3417",
	MONSHEIMER_WEG_STUTTGART:                               "de:08111:2635",
	MORBACH_GRAB:                                           "de:08119:7515",
	MORIKESCHULE_ELTINGEN:                                  "de:08115:2341",
	MORIKESCHULE_OTLINGEN:                                  "de:08116:4346",
	MORIKESTRASSE_STUTTGART:                                "de:08111:2409",
	MORIKESTRASSE_LUDWIGSBURG:                              "de:08118:5476",
	MORIKESTRASSE_OBERRIEXINGEN:                            "de:08118:6514",
	MORIKESTRASSE_GROSSBETTLINGEN:                          "de:08116:6846",
	MORIKESTRASSE_GOPPINGEN:                                "de:08117:1053",
	MORIKESTRASSE_HEININGEN_GP:                             "de:08117:3310",
	MORSESTRASSE_STUTTGART:                                 "de:08111:6464",
	MOSELSTRASSE_STUTTGART:                                 "de:08111:373",
	MOTORSTRASSE_24_STUTTGART:                              "de:08111:2350",
	MOTZINGER_STRASSE_UNTERJETTINGEN:                       "de:08115:4865",
	MOTZINGER_STRASSE_OSCHELBRONN_GAUFELDEN:                "de:08115:7928",
	MOTZINGER_STRASSE_NAGOLD:                               "de:08235:10072",
	MOZARTSCHULE_NEUHAUSEN_F:                               "de:08116:2898",
	MOZARTSTRASSE_HOCHDORF_ES:                              "de:08116:3910",
	MOZARTSTRASSE_FELLBACH:                                 "de:08119:6039",
	MOZARTSTRASSE_EISLINGEN_F:                              "de:08117:1040",
	MOZARTSTRASSE_GOPPINGEN:                                "de:08117:4007",
	MUCKENSTURM_STUTTGART:                                  "de:08111:2488",
	MUHLE_ALDINGEN:                                         "de:08118:3571",
	MUHLENBUCKEL_BERNHAUSEN:                                "de:08116:2088",
	MUHLHAUSEN_STUTTGART:                                   "de:08111:6270",
	MUHLHAUSEN_SCHLOSS_STUTTGART:                           "de:08111:6271",
	MUHLHAUSENER_STRASSE_LEHNINGEN:                         "de:08236:3607",
	MUHLHAUSER_STR_GRUIBINGEN:                              "de:08117:3231",
	MUHLHAUSER_STRASSE_KORNWESTHEIM:                        "de:08118:5451",
	MUHLHAUSER_STRASSE_ROSSWAG:                             "de:08118:6850",
	MUHLHAUSER_STRASSE_OSSWEIL:                             "de:08118:7556",
	MUHLSTEG_STUTTGART:                                     "de:08111:281",
	MUHLSTRASSE_NURTINGEN:                                  "de:08116:2813",
	MUHLSTRASSE_OBERRIEXINGEN:                              "de:08118:5641",
	MUHLSTRASSE_SUSSEN:                                     "de:08117:303",
	MUHLSTRASSE_WALDHAUSEN_B_SCHORND_LORCH:                 "de:08136:7048",
	MUHLWEG_HERRENBERG:                                     "de:08115:4848",
	MUHLWEG_BEINSTEIN:                                      "de:08119:5951",
	MUHLWEG_ALDINGEN:                                       "de:08118:6832",
	MUHLWIESEN_BIETIGHEIM:                                  "de:08118:7471",
	MULLERHEIM_MUNCHINGEN:                                  "de:08118:4518",
	MUNCHINGEN_MUNCHINGEN:                                  "de:08118:5632",
	MUNCHINGER_STRASSE_DITZINGEN:                           "de:08118:3029",
	MUNCHINGER_STRASSE_MARKGRONINGEN:                       "de:08118:3314",
	MUNCHINGER_STRASSE_HEMMINGEN:                           "de:08118:6930",
	MUNSTER_STUTTGART:                                      "de:08111:1504",
	MUNSTER_RATHAUS_STUTTGART:                              "de:08111:277",
	MUNSTER_VIADUKT_STUTTGART:                              "de:08111:278",
	MUNZENHALDE_PLUDERHAUSEN:                               "de:08119:5249",
	MURKENBACHWEG_BOBLINGEN:                                "de:08115:4688",
	MURRBADER_BACKNANG:                                     "de:08119:7599",
	MURRBRUCKE_BURGSTALL_BURGSTETTEN:                       "de:08119:3508",
	MURRER_STRASSE_STEINHEIM_M:                             "de:08118:5601",
	MURRHARDT_MURRHARDT:                                    "de:08119:4916",
	MURRHARDTER_STRASSE_SULZBACH_M:                         "de:08119:4910",
	MURRHARLE_MURRHARDT:                                    "de:08119:5138",
	MURRTALSCHULE_OPPENWEILER:                              "de:08119:4951",
	MUTZENREIS_ZOLLBERG:                                    "de:08116:2802",
	MUTZENREISSTRASSE_ZOLLBERG:                             "de:08116:2805",
	NACHBARSCHAFTSSCHULE_THOMASHARDT:                       "de:08116:4504",
	NACHBARSCHAFTSSCHULE_BRETZENACKER:                      "de:08119:5277",
	NACHTIGALLENWEG_WAIBLINGEN:                             "de:08119:5413",
	NAGELESTAL_KIRCHHEIM_T:                                 "de:08116:4383",
	NAGELESTRASSE_STUTTGART:                                "de:08111:341",
	NAGELSTRASSE_SCHARNHAUSEN:                              "de:08116:2900",
	NAGOLDER_STEIG_MOTZINGEN:                               "de:08115:7050",
	NASSACH_NASSACH:                                        "de:08119:5048",
	NASSACH_BOLZPLATZ_UHINGEN:                              "de:08117:7679",
	NASSACH_FRIEDHOF_UHINGEN:                               "de:08117:7678",
	NASSACH_UNTERHUTT_UHINGEN:                              "de:08117:7664",
	NASSACHMUHLE_UHINGEN:                                   "de:08117:9822",
	NASSACHMUHLE_SCHULE_UHINGEN:                            "de:08117:9826",
	NASTPLATZ_STUTTGART:                                    "de:08111:2258",
	NASTSTRASSE_LUDWIGSBURG:                                "de:08118:7555",
	NATURFREUNDEHAUS_SECHSELBERG:                           "de:08119:4984",
	NATURFREUNDEHAUS_WELZHEIM:                              "de:08119:7818",
	NATURSCHUTZZENTRUM_SCHOPFLOCH_LENNINGEN:                "de:08116:3822",
	NEBRINGER_STRASSE_BONDORF:                              "de:08115:5946",
	NECKAR_FORUM_ESSLINGEN_N:                               "de:08116:4033",
	NECKARBRUCKE_NURTINGEN:                                 "de:08116:2927",
	NECKARBRUCKE_LUDWIGSBURG:                               "de:08118:3418",
	NECKARBRUCKE_NECKARTAILFINGEN:                          "de:08116:6860",
	NECKARFREIBAD_ESSLINGEN_N:                              "de:08116:4041",
	NECKARHALDE_STUTTGART:                                  "de:08111:87",
	NECKARHALDE_NECKARREMS:                                 "de:08118:4885",
	NECKARPARK_STUTTGART:                                   "de:08111:6080",
	NECKARPARK_STADION_STUTTGART:                           "de:08111:6330",
	NECKARSCHULE_ALDINGEN:                                  "de:08118:4886",
	NECKARSTRASSE_HEGNACH:                                  "de:08119:4878",
	NECKARSTRASSE_KORNWESTHEIM:                             "de:08118:5444",
	NECKARSTRASSE_ASPERG:                                   "de:08118:5478",
	NECKARSTRASSE_LUDWIGSBURG:                              "de:08118:5514",
	NECKARTAILFINGER_STRASSE_SCHLAITDORF:                   "de:08116:2070",
	NECKARTENZLINGER_STRASSE_BEMPFLINGEN:                   "de:08116:4313",
	NECKARTOR_STUTTGART:                                    "de:08111:25",
	NECKARWEIHINGER_STRASSE_OSSWEIL:                        "de:08118:6883",
	NECKARZENTRUM_HOCHBERG:                                 "de:08118:5629",
	NECKLINSBERG_ASPERGLEN:                                 "de:08119:5267",
	NECKLINSBERG_KREUZUNG_ASPERGLEN:                        "de:08119:5266",
	NEIDLINGER_STRASSE_WEILHEIM_T:                          "de:08116:4398",
	NEL_MEZZO_GEISLINGEN_STEIGE:                            "de:08117:6",
	NELKENSTRASSE_AFFSTATT:                                 "de:08115:4854",
	NELLINGER_LINDE_NELLINGEN:                              "de:08116:5003",
	NELLINGER_STRASSE_BERKHEIM:                             "de:08116:4147",
	NELLMERSBACH_NELLMERSBACH:                              "de:08119:1601",
	NETZ__IHK_NAGOLD:                                       "de:08235:9333",
	NETZESTRASSE_GRUNBUHL:                                  "de:08118:5523",
	NEUE_BAHNHOFSTRASSE_VAIHINGEN_E:                        "de:08118:5880",
	NEUE_RAMTELSTRASSE_RAMTEL:                              "de:08115:2351",
	NEUE_ROMMELSHAUSER_STRASSE_WAIBLINGEN:                  "de:08119:3640",
	NEUE_STRASSE_LIEBERSBRONN:                              "de:08116:4103",
	NEUE_STRASSE_STEINBACH:                                 "de:08119:6844",
	NEUE_WEINSTEIGE_61_STUTTGART:                           "de:08111:161",
	NEUENBUHL_FLACHT:                                       "de:08115:4578",
	NEUER_FRIEDHOF_LUDWIGSBURG:                             "de:08118:5499",
	NEUER_FRIEDHOF_SCHORNDORF:                              "de:08119:5975",
	NEUER_FRIEDHOF_BISSINGEN_LB:                            "de:08118:7422",
	NEUER_MARKT_LEINFELDEN:                                 "de:08116:2069",
	NEUES_RATHAUS_SCHONAICH:                                "de:08115:4828",
	NEUFFEN_NEUFFEN:                                        "de:08116:5794",
	NEUFFENER_STRASSE_LINSENHOFEN:                          "de:08116:2819",
	NEUFFENSTRASSE_WENDLINGEN_N:                            "de:08116:4474",
	NEUFFENSTRASSE_REICHENBACH_F:                           "de:08116:6504",
	NEUFURSTENHUTTE_GROSSERLACH:                            "de:08119:4968",
	NEUGEREUT_STUTTGART:                                    "de:08111:2484",
	NEUHAUSER_STRASSE_MUNKLINGEN:                           "de:08115:4577",
	NEUHAUSER_STRASSE_BERNHAUSEN:                           "de:08116:2043",
	NEUHAUSER_STRASSE_DENKENDORF:                           "de:08116:5009",
	NEUHOF_KIRCHBERG_M:                                     "de:08119:3529",
	NEUHOF_WELZHEIM:                                        "de:08119:5088",
	NEULAUTERN_WUSTENROT:                                   "de:08125:1022",
	NEUSATZ_BESIGHEIM:                                      "de:08118:5580",
	NEUSCHONTAL_BACKNANG:                                   "de:08119:3511",
	NEUSTADTHOHENACKER_NEUSTADT:                            "de:08119:1602",
	NEUSTADTER_STRASSE_WAIBLINGEN:                          "de:08119:5146",
	NEUWEILER_STRASSE_KLEINSACHSENHEIM:                     "de:08118:5672",
	NEUWIESEN_B466_GEISLINGEN_STEIGE:                       "de:08117:44",
	NEUWIESEN_GEWERBEGEBIET_GEISLINGEN_STEIGE:              "de:08117:114",
	NEUWIRTSH_PORSCHEP_STUTTGART:                           "de:08111:1403",
	NEUWIRTSHAUSKREUZUNG_STUTTGART:                         "de:08111:95",
	NIEBUHRWEG_STUTTGART:                                   "de:08111:96",
	NIEDERE_KLINGE_GEMMRIGHEIM:                             "de:08118:5721",
	NIEDERHOFEN_ELTINGEN:                                   "de:08115:4541",
	NIEDERREUTIN_BONDORF:                                   "de:08115:7046",
	NIKOLAUSOTTOSTRASSE_ECHTERDINGEN:                       "de:08116:2132",
	NIKOLAUSPFLEGE_STUTTGART:                               "de:08111:2452",
	NIKOMODIESTRASSE_34_OBERLENNINGEN:                      "de:08116:7584",
	NIXENWEG_STUTTGART:                                     "de:08111:6275",
	NOBELSTRASSE_STUTTGART:                                 "de:08111:6018",
	NOLDEWEG_LORCH:                                         "de:08136:7003",
	NORD_BONLANDEN:                                         "de:08116:2021",
	NORD_WEIL_IM_SCHONBUCH:                                 "de:08115:3108",
	NORD_MAICHINGEN:                                        "de:08115:3197",
	NORD_RUDERSBERG:                                        "de:08119:5801",
	NORD_HIRSCHLANDEN:                                      "de:08118:6492",
	NORDWESTRING_BERNHAUSEN:                                "de:08116:2079",
	NORDWESTUMFAHRUNG_BERNHAUSEN:                           "de:08116:2084",
	NORDBAHNHOF_KLEINGLATTBACH:                             "de:08118:5766",
	NORDBAHNHOF_STUTTGART:                                  "de:08111:6295",
	NORDHALDENSTRASSE_BEUTELSBACH:                          "de:08119:3763",
	NORDSEESCHWIEBERD_STRASSE_STUTTGART:                    "de:08111:2461",
	NORMANNENSTRASSE_WEIL_DER_STADT:                        "de:08115:6470",
	NOTZINGER_STRASSE_KIRCHHEIM_T:                          "de:08116:2827",
	NOVIZENWEGTHW_NEUHAUSEN_F:                              "de:08116:2896",
	NUFRINGEN_NUFRINGEN:                                    "de:08115:5775",
	NURNBERGER_STRASSE_STUTTGART:                           "de:08111:34",
	NURTINGEN_NURTINGEN:                                    "de:08116:2931",
	NURTINGER_STRASSE_GROSSBETTLINGEN:                      "de:08116:4463",
	NURTINGER_STRASSE_GROTZINGEN:                           "de:08116:4471",
	NURTINGER_STRASSE_NECKARTAILFINGEN:                     "de:08116:4482",
	NURTINGER_STRASSE_BOBLINGEN:                            "de:08115:4671",
	NUSSBAUMWEG_OBERSTENFELD:                               "de:08118:5611",
	NUSSDORFER_STRASSE_AURICH:                              "de:08118:5885",
	NUSSDORFER_STRASSE_EBERDINGEN:                          "de:08118:6468",
	NUSSSTRASSE_SINDELFINGEN:                               "de:08115:4658",
	OB_DEM_STEINBACH_STUTTGART:                             "de:08111:2579",
	OBERAICHEN_LEINFELDEN:                                  "de:08116:2105",
	OBERAICHEN_WALDHEIM_LEINFELDEN:                         "de:08116:2106",
	OBERAMTEIGASSE_BESIGHEIM:                               "de:08118:3588",
	OBERBOIHINGEN_OBERBOIHINGEN:                            "de:08116:5770",
	OBERBOIHINGER_STRASSE_ZIZISHAUSEN:                      "de:08116:7936",
	OBERBRUDENER_STRASSE_STEINBACH:                         "de:08119:6918",
	OBERDORF_WEILER_OB_HELFENSTN:                           "de:08117:109",
	OBERE_ACKER_KAYH:                                       "de:08115:4830",
	OBERE_BAHNHOFSTRASSE_WAIBLINGEN:                        "de:08119:5925",
	OBERE_GASSE_NELLINGEN_UL:                               "de:08425:2621",
	OBERE_LIESACKER_ALTDORF_ES:                             "de:08116:7089",
	OBERE_MARBACHER_STRASSE_LUDWIGSBURG:                    "de:08118:5483",
	OBERE_MONCHHALDE_STUTTGART:                             "de:08111:2243",
	OBERE_SEEWIESEN_HOPFIGHEIM:                             "de:08118:3592",
	OBERE_STRASSE_DETTINGEN_T:                              "de:08116:4275",
	OBERE_ZIEGELEI_STUTTGART:                               "de:08111:6320",
	OBERER_EISBERGWEG_ZOLLBERG:                             "de:08116:5024",
	OBERER_GRABEN_NEUFFEN:                                  "de:08116:7070",
	OBERER_ROSBERG_WAIBLINGEN:                              "de:08119:3724",
	OBERER_SCHADBERG_KAISERSBACH:                           "de:08119:5282",
	OBERER_WASEN_WELZHEIM:                                  "de:08119:5123",
	OBERESSLINGEN_OBERESSLINGEN:                            "de:08116:1802",
	OBERFISCHBACH_GROSSERLACH:                              "de:08119:4966",
	OBERHOF_OBERHOF:                                        "de:08116:4070",
	OBERLENNINGEN_OBERLENNINGEN:                            "de:08116:4418",
	OBERNDORF_RUDERSBERG:                                   "de:08119:7567",
	OBERNDORF_BRUNNENPLATZ_RUDERSBERG:                      "de:08119:5025",
	OBERNDORF_MANNENBERGER_STRASSE_RUDERSBERG:              "de:08119:7663",
	OBERNEUSTETTEN_KIRCHENKIRNBERG:                         "de:08119:5084",
	OBERRIEXINGER_STRASSE_GROSSSACHSENHEIM:                 "de:08118:6944",
	OBERSCHONTAL_BACKNANG:                                  "de:08119:5354",
	OBERSCHONTAL_ABZWEIG_BACKNANG:                          "de:08119:5352",
	OBERSTEINENBERG_WELZHEIM:                               "de:08119:7526",
	OBERTALWEG_WALDENBRONN:                                 "de:08116:4129",
	OBERTURKH_BF_GOPPINGER_STR_STUTTGART:                   "de:08111:2543",
	OBERTURKHEIM_STUTTGART:                                 "de:08111:6091",
	OBERWIESENSTRASSE_55_STUTTGART:                         "de:08111:2435",
	OBSTHALDE_BEINSTEIN:                                    "de:08119:6634",
	OCHSEN_SPIELBERG:                                       "de:08118:5711",
	OCHSENTROG_BOBLINGEN:                                   "de:08115:4714",
	ODERNHARDT_ODERNHARDT:                                  "de:08119:5278",
	OESCHSCHULEN_EISLINGEN_F:                               "de:08117:1025",
	OFELE_OBERJETTINGEN:                                    "de:08115:4759",
	OHMSTRASSE_ZUFFENHAUSEN_BF_STUTTGART:                   "de:08111:6467",
	OLGAKEPLERSTRASSE_DEIZISAU:                             "de:08116:6880",
	OLGAZEPPELINSTRASSE_DEIZISAU:                           "de:08116:6875",
	OLGAECK_STUTTGART:                                      "de:08111:6119",
	OLMUHLE_WEISSACH_BB:                                    "de:08115:4532",
	OPPENWEILER_OPPENWEILER:                                "de:08119:5786",
	ORT_AICH:                                               "de:08116:2001",
	ORTSEINGANG_HEGENSBERG:                                 "de:08116:4203",
	ORTSMITTE_ERBSTETTEN:                                   "de:08119:3513",
	ORTSMITTE_KLEINASPACH:                                  "de:08119:3672",
	ORTSMITTE_REICHENBACH_F:                                "de:08116:4234",
	ORTSMITTE_OTLINGEN:                                     "de:08116:4341",
	ORTSMITTE_LINDORF:                                      "de:08116:4349",
	ORTSMITTE_WARMBRONN:                                    "de:08115:4537",
	ORTSMITTE_FLACHT:                                       "de:08115:4566",
	ORTSMITTE_MERKLINGEN:                                   "de:08115:4588",
	ORTSMITTE_OPPENWEILER:                                  "de:08119:4901",
	ORTSMITTE_ALTHUTTE:                                     "de:08119:4996",
	ORTSMITTE_GROTZINGEN:                                   "de:08116:5001",
	ORTSMITTE_ALLMERSBACH_IM_TAL:                           "de:08119:5032",
	ORTSMITTE_SCHLECHTBACH:                                 "de:08119:5118",
	ORTSMITTE_HOHENACKER:                                   "de:08119:5151",
	ORTSMITTE_BIRKMANNSWEILER:                              "de:08119:5204",
	ORTSMITTE_OPPELSBOHM:                                   "de:08119:5268",
	ORTSMITTE_BAACH_WINNENDEN:                              "de:08119:5274",
	ORTSMITTE_HOFEN:                                        "de:08119:5275",
	ORTSMITTE_ENDERSBACH:                                   "de:08119:5393",
	ORTSMITTE_MURR:                                         "de:08118:5598",
	ORTSMITTE_AURICH:                                       "de:08118:5753",
	ORTSMITTE_GROSSINGERSHEIM:                              "de:08118:5934",
	ORTSMITTE_HOHENGEHREN:                                  "de:08116:7935",
	ORTSMITTE_EYBACH:                                       "de:08117:208",
	ORTSMITTE_WALDHAUSEN_GEISLINGEN:                        "de:08117:213",
	ORTSMITTE_MAITIS:                                       "de:08117:1206",
	ORTSMITTE_LENGLINGEN:                                   "de:08117:1208",
	ORTSMITTE_ALBERSHAUSEN:                                 "de:08117:2044",
	ORTSMITTE_BAD_BOLL:                                     "de:08117:2107",
	ORTSMITTE_DURNAU:                                       "de:08117:3339",
	ORTSMITTE_MERKLINGEN_7:                                 "de:08425:2632",
	ORTSMITTEKIRCHE_WALDHAUSEN_B_SCHORND_LORCH:             "de:08136:7033",
	ORTSMITTEKORNSTR_OBERKIRNECK:                           "de:08136:7024",
	OSCHELBRONN_OSCHELBRONN_BERGLEN:                        "de:08119:5270",
	OSKAR_FRECH_SEEBAD_SCHORNDORF:                          "de:08119:5212",
	OSKARSCHLEMMERSTRASSE_STUTTGART:                        "de:08111:2236",
	OSSMANNSWEILER_EYBACH:                                  "de:08117:187",
	OSSWEILER_WEG_ALDINGEN:                                 "de:08118:3585",
	OST_REUDERN:                                            "de:08116:4381",
	OSTENDPLATZ_STUTTGART:                                  "de:08111:6060",
	OSTERBRONNSTRASSE_STUTTGART:                            "de:08111:2612",
	OSTERFELD_STUTTGART:                                    "de:08111:6027",
	OSTERHOLZ_LUDWIGSBURG:                                  "de:08118:5532",
	OSTERHOLZALLEE_LUDWIGSBURG:                             "de:08118:5540",
	OSTERREICHISCHER_PLATZ_STUTTGART:                       "de:08111:6020",
	OSTERWIESENWEG_KLEINGLATTBACH:                          "de:08118:5884",
	OSTFILDERN_NELLINGEN:                                   "de:08116:2974",
	OSTFRIEDHOF_OSSWEIL:                                    "de:08118:6000",
	OSTHEIM_LEOVETTERBAD_STUTTGART:                         "de:08111:78",
	OSTPREUSSENWEG_NURTINGEN:                               "de:08116:4461",
	OSTRING_HOCHDORF_ES:                                    "de:08116:3911",
	OSTSIEDLUNG_NURTINGEN:                                  "de:08116:4378",
	OSTSTRASSE_LUDWIGSBURG:                                 "de:08118:5531",
	OSWALDSTRASSE_MAGSTADT:                                 "de:08115:3296",
	OTLINGEN_OTLINGEN:                                      "de:08116:4345",
	OTLINGER_STRASSE_NOTZINGEN:                             "de:08116:4368",
	OTTOHAHNSTRASSE_NURTINGEN:                              "de:08116:2986",
	OTTOHAHNSTRASSE_KIRCHHEIM_T:                            "de:08116:3953",
	OTTOHAHNSTRASSE_NECKARWEIHINGEN:                        "de:08118:6847",
	OTTOHIRSCHBRUCKEN_STUTTGART:                            "de:08111:2534",
	OTTOLILIENTHALSTRASSE_BOBLINGEN:                        "de:08115:4722",
	OTTOMUHLSCHLEGELHAUS_ENDERSBACH:                        "de:08119:5174",
	OWEN_OWEN:                                              "de:08116:5779",
	PADAGOGISCHE_HOCHSCHULE_LUDWIGSBURG:                    "de:08118:5543",
	PALAZZO_FREIBERG_N:                                     "de:08118:7560",
	PALMENWALDSTRASSE_BRUHL:                                "de:08116:6842",
	PANORAMA_THERME_BEUREN:                                 "de:08116:4436",
	PANORAMAKLINIK_ESSLINGEN_N:                             "de:08116:4122",
	PANORAMASTR_RECHBERGHAUSEN:                             "de:08117:4025",
	PANORAMASTRASSE_PLOCHINGEN:                             "de:08116:4230",
	PANORAMASTRASSE_BISSINGEN_LB:                           "de:08118:5684",
	PANORAMASTRASSE_LORCH:                                  "de:08136:7018",
	PANORAMASTRASSEBAHNHOF_FAURNDAU:                        "de:08117:5003",
	PANZERKASERNE_BOBLINGEN:                                "de:08115:5807",
	PAPIERFABRIK_GEMMRIGHEIM:                               "de:08118:5722",
	PAPPELWEG_MARKGRONINGEN:                                "de:08118:3307",
	PARACELSUSSTRASSE_OBERESSLINGEN:                        "de:08116:4004",
	PARADIESWEG_RUDERN:                                     "de:08116:4115",
	PARKACKER_BIETIGHEIM:                                   "de:08118:3383",
	PARKPLATZ_SCHLOSSGYMNASIUM_KIRCHHEIM_T:                 "de:08116:4335",
	PARKPLATZ_ZOLLSTOCK_SPIEGELBERG:                        "de:08119:7449",
	PARKSIEDLUNG_PARKSIEDLUNG:                              "de:08116:2970",
	PARKSTRASSE_PARKSIEDLUNG:                               "de:08116:2911",
	PARKSTRASSE_PLIENSAUVORSTADT:                           "de:08116:4038",
	PARKSTRASSE_LANDRATSAMT_BOBLINGEN:                      "de:08115:3126",
	PAULGERHARDTHAUS_BESIGHEIM:                             "de:08118:3536",
	PAULKOEPFFWEG_GOPPINGEN:                                "de:08117:9018",
	PAULLINCKESTRASSE_STUTTGART:                            "de:08111:2429",
	PAULINENSTRASSE_NELLINGEN:                              "de:08116:2967",
	PAULUSKIRCHE_FELLBACH:                                  "de:08119:2501",
	PAULUSKIRCHE_GEISLINGEN_STEIGE:                         "de:08117:5",
	PAYERSTRASSE_STUTTGART:                                 "de:08111:125",
	PEREGRINASTRASSE_STUTTGART:                             "de:08111:6166",
	PERONNASPLATZ_NEUHAUSEN_F:                              "de:08116:2903",
	PEROUSER_STRASSE_MALMSHEIM:                             "de:08115:5845",
	PESTALOZZISCHULE_STUTTGART:                             "de:08111:6015",
	PESTALOZZISTRASSE_ENDERSBACH:                           "de:08119:3769",
	PETERHEBELSTRASSE_NECKARWEIHINGEN:                      "de:08118:6818",
	PETERSHALDE_PFAHLBRONN:                                 "de:08119:5090",
	PFAFFENWEG_STUTTGART:                                   "de:08111:343",
	PFAFFLESTRASSESCHILLERSCHULE_LORCH:                     "de:08136:7046",
	PFARRACKER_NEUSTADT:                                    "de:08119:7547",
	PFARRGARTENSTRASSE_KIRCHBERG_M:                         "de:08119:6858",
	PFARRHAUS_GUNDELBACH:                                   "de:08118:5740",
	PFARRSTRASSE_BIETIGHEIM:                                "de:08118:5662",
	PFARRWIESENGYMNASIUM_SINDELFINGEN:                      "de:08115:2216",
	PFEIFFERKLINGE_PLIENSAUVORSTADT:                        "de:08116:6969",
	PFINGSTHALDE_EYBACH:                                    "de:08117:209",
	PFINGSTWASEN_GOPPINGEN:                                 "de:08117:7003",
	PFITZERSTRASSE_SCHWABISCH_GMUND:                        "de:08136:3054",
	PFLEGEHEIM_WINTERBACH:                                  "de:08119:3713",
	PFLEGEHEIM_SINDELFINGEN:                                "de:08115:4627",
	PFLEGEHEIM_SCHLIERBACH:                                 "de:08117:7798",
	PFLUGFELDER_STRASSE_STUTTGART:                          "de:08111:6466",
	PFLUGFELDER_STRASSE_KORNWESTHEIM:                       "de:08118:6807",
	PFORZHEIMER_STRASSE_HORRHEIM:                           "de:08118:7811",
	PFOSTENBERG_PLOCHINGEN:                                 "de:08116:4255",
	PFULLINGER_STRASSE_STUTTGART:                           "de:08111:3308",
	PIEMONTESER_STRASSE_ERDMANNHAUSEN:                      "de:08118:3563",
	PLACKERTSTRASSE_BUNZWANGEN:                             "de:08117:8003",
	PLAISIRSCHULE_BACKNANG:                                 "de:08119:5332",
	PLANCKSTRASSE_STUTTGART:                                "de:08111:2310",
	PLANCKWEG_GOPPINGEN:                                    "de:08117:79",
	PLATTENWALD_BACKNANG:                                   "de:08119:5325",
	PLATZ_HOLZMADEN:                                        "de:08116:4393",
	PLEIDELSHEIMER_STRASSE_BIETIGHEIM:                      "de:08118:7408",
	PLIENINGEN_STUTTGART:                                   "de:08111:6555",
	PLIENINGEN_GARBE_STUTTGART:                             "de:08111:6554",
	PLIENINGEN_POST_STUTTGART:                              "de:08111:2039",
	PLIENINGEN_SEEMUHLENWEG_STUTTGART:                      "de:08111:2976",
	PLIENINGER_STRASSE_STUTTGART:                           "de:08111:353",
	PLIENINGER_STRASSE_SCHARNHAUSEN:                        "de:08116:2916",
	PLIENINGER_STRASSE_NORD_SCHARNHAUSEN:                   "de:08116:2899",
	PLIENSAUFRIEDHOF_ZOLLBERG:                              "de:08116:2804",
	PLIENSAUTURM_ESSLINGEN_N:                               "de:08116:2806",
	PLOCHINGEN_PLOCHINGEN:                                  "de:08116:7802",
	PLOCHINGER_STR_DEIZISAU:                                "de:08116:4060",
	PLOCHINGER_STRASSE_KONGEN:                              "de:08116:4092",
	PLUDERHAUSEN_PLUDERHAUSEN:                              "de:08119:5761",
	POLIZEIHOCHSCHULE_HERRENBERG:                           "de:08115:4737",
	POLIZEIPRASIDIUM_EINSATZ_GOPPINGEN:                     "de:08117:3305",
	POLIZEIWACHE_KIRCHHEIM_T:                               "de:08116:4268",
	POPPENWEILER_STRASSE_HOCHDORF_REMSECK:                  "de:08118:5556",
	PORSCHE_WEISSACH_BB:                                    "de:08115:4564",
	PORSCHE_GROSSSACHSENHEIM:                               "de:08118:7084",
	PORSCHESTRASSE_HERRENBERG:                              "de:08115:7007",
	POST_RENNINGEN:                                         "de:08115:3193",
	POST_DAGERSHEIM:                                        "de:08115:3232",
	POST_GUTENBERG:                                         "de:08116:4421",
	POST_OSCHELBRONN_GAUFELDEN:                             "de:08115:4748",
	POST_OPPENWEILER:                                       "de:08119:4899",
	POST_SCHORNBACH:                                        "de:08119:5260",
	POST_FREUDENTAL:                                        "de:08118:5707",
	POST_HOHENSTEIN:                                        "de:08118:5725",
	POST_BOHMENKIRCH:                                       "de:08117:202",
	POSTAMT_NEBRINGEN:                                      "de:08115:4745",
	POSTAMT_WIESENSTEIG:                                    "de:08117:61",
	POSTDORFLE_STUTTGART:                                   "de:08111:2247",
	POSTDORFLE_FELLBACH:                                    "de:08119:2647",
	POSTPLATZ_BOBLINGEN:                                    "de:08115:3104",
	POSTSTRASSE_BEUTELSBACH:                                "de:08119:3765",
	POSTSTRASSLE_BIETIGHEIM:                                "de:08118:3401",
	POSTWEG_PLUDERHAUSEN:                                   "de:08119:5252",
	POTSDAMER_RING_BACKNANG:                                "de:08119:5333",
	PRAGFRIEDHOF_STUTTGART:                                 "de:08111:115",
	PRAGSATTEL_STUTTGART:                                   "de:08111:6113",
	PREVORST_FEUERWEHR_GRONAU:                              "de:08118:5596",
	PREVORST_LOWEN_GRONAU:                                  "de:08118:6798",
	PRINZEUGENPLATZ_GROSSHEPPACH:                           "de:08119:5169",
	PULVERDINGEN_B10_ENZWEIHINGEN:                          "de:08118:5743",
	PULVERDINGEN_ORT_ENZWEIHINGEN:                          "de:08118:5628",
	PULVERDINGER_STRASSE_HOCHDORF_EBERDINGEN:               "de:08118:6951",
	PUMPSTATION_HIRSCHLANDEN:                               "de:08118:3002",
	PUSTEBLUMENKREISEL_EBERSBACH_F:                         "de:08117:8008",
	QUADRIUM_WERNAU_N:                                      "de:08116:6710",
	QUELLENSTRASSE_BEINSTEIN:                               "de:08119:5953",
	QUELLENWEG_URBACH:                                      "de:08119:6626",
	QUERSPANGE_WAIBLINGEN:                                  "de:08119:5411",
	QUERSTRASSE_SALACH:                                     "de:08117:7605",
	RABENWIESEN_SUSSEN:                                     "de:08117:7609",
	RAICHBERG_SCHULZENTRUM_EBERSBACH_F:                     "de:08117:9881",
	RAIFFEISENBANK_LIPPOLDSWEILER:                          "de:08119:5385",
	RAIFFEISENSTRASSE_GROTZINGEN:                           "de:08116:190",
	RAIFFEISENSTRASSE_BONLANDEN:                            "de:08116:1997",
	RAIFFEISENSTRASSE_WINZERHAUSEN:                         "de:08118:5607",
	RAIFFEISENSTRASSE_BARTENBACH_GOPPINGEN:                 "de:08117:4011",
	RAINACKERSTRASSE_BONLANDEN:                             "de:08116:2052",
	RAINWIESEN_NECKARGRONINGEN:                             "de:08118:6826",
	RAITE_RENNINGEN:                                        "de:08115:3189",
	RAITELSBERG_STUTTGART:                                  "de:08111:222",
	RAMTEL_GERLINGEN:                                       "de:08118:2308",
	RANDENSTRASSE_HASLACH:                                  "de:08115:4744",
	RAPPACHSCHULE_STUTTGART:                                "de:08111:6792",
	RAPPACHSTRASSE_STUTTGART:                               "de:08111:2653",
	RAPPENBERG_KIRCHBERG_M:                                 "de:08119:7097",
	RAPPENHOF_LEONBERG:                                     "de:08115:2334",
	RASTATTER_STRASSE_STUTTGART:                            "de:08111:148",
	RATHAUS_PLATTENHARDT:                                   "de:08116:2028",
	RATHAUS_SIELMINGEN:                                     "de:08116:2036",
	RATHAUS_BERNHAUSEN:                                     "de:08116:2044",
	RATHAUS_ALTENRIET:                                      "de:08116:2066",
	RATHAUS_OEFFINGEN:                                      "de:08119:2372",
	RATHAUS_SCHMIDEN:                                       "de:08119:2377",
	RATHAUS_KAPPISHAUSERN:                                  "de:08116:2820",
	RATHAUS_SCHARNHAUSEN:                                   "de:08116:2917",
	RATHAUS_MAICHINGEN:                                     "de:08115:3209",
	RATHAUS_LUDWIGSBURG:                                    "de:08118:3419",
	RATHAUS_ERDMANNHAUSEN:                                  "de:08118:3515",
	RATHAUS_KIRCHBERG_M:                                    "de:08119:3527",
	RATHAUS_GUNDELBACH:                                     "de:08118:3590",
	RATHAUS_WALDREMS:                                       "de:08119:3615",
	RATHAUS_MAUBACH:                                        "de:08119:3669",
	RATHAUS_GERADSTETTEN:                                   "de:08119:3702",
	RATHAUS_ROMMELSHAUSEN:                                  "de:08119:3744",
	RATHAUS_ENDERSBACH:                                     "de:08119:3768",
	RATHAUS_ALTBACH:                                        "de:08116:3803",
	RATHAUS_BALTMANNSWEILER:                                "de:08116:3808",
	RATHAUS_DEIZISAU:                                       "de:08116:4057",
	RATHAUS_KONGEN:                                         "de:08116:4088",
	RATHAUS_AICHELBERG_AICHWALD:                            "de:08116:4166",
	RATHAUS_HOLZGERLINGEN:                                  "de:08115:4210",
	RATHAUS_NOTZINGEN:                                      "de:08116:4218",
	RATHAUS_HOCHDORF_ES:                                    "de:08116:4222",
	RATHAUS_UNTERLENNINGEN:                                 "de:08116:4284",
	RATHAUS_BEMPFLINGEN:                                    "de:08116:4316",
	RATHAUS_JESINGEN:                                       "de:08116:4323",
	RATHAUS_OHMDEN:                                         "de:08116:4391",
	RATHAUS_HEPSISAU:                                       "de:08116:4405",
	RATHAUS_BISSINGEN_AN_DER_TECK_T:                        "de:08116:4415",
	RATHAUS_OCHSENWANG:                                     "de:08116:4417",
	RATHAUS_SCHOPFLOCH_LENNINGEN:                           "de:08116:4424",
	RATHAUS_KLEINBETTLINGEN:                                "de:08116:4465",
	RATHAUS_TISCHARDT:                                      "de:08116:4466",
	RATHAUS_NECKARHAUSEN:                                   "de:08116:4480",
	RATHAUS_NECKARTAILFINGEN:                               "de:08116:4483",
	RATHAUS_NECKARTENZLINGEN:                               "de:08116:4487",
	RATHAUS_MOTZINGEN:                                      "de:08115:4509",
	RATHAUS_HEMMINGEN:                                      "de:08118:4526",
	RATHAUS_GEBERSHEIM:                                     "de:08115:4550",
	RATHAUS_DATZINGEN:                                      "de:08115:4585",
	RATHAUS_MUNKLINGEN:                                     "de:08115:4592",
	RATHAUS_GARTRINGEN:                                     "de:08115:4727",
	RATHAUS_ROHRAU:                                         "de:08115:4729",
	RATHAUS_OSCHELBRONN_GAUFELDEN:                          "de:08115:4747",
	RATHAUS_DOFFINGEN:                                      "de:08115:4770",
	RATHAUS_AIDLINGEN:                                      "de:08115:4777",
	RATHAUS_DEUFRINGEN:                                     "de:08115:4780",
	RATHAUS_BREITENSTEIN:                                   "de:08115:4788",
	RATHAUS_NEUWEILER:                                      "de:08115:4820",
	RATHAUS_KAYH:                                           "de:08115:4831",
	RATHAUS_DECKENPFRONN:                                   "de:08115:4838",
	RATHAUS_AFFSTATT:                                       "de:08115:4847",
	RATHAUS_TAILFINGEN:                                     "de:08115:4862",
	RATHAUS_HEGNACH:                                        "de:08119:4879",
	RATHAUS_GROSSERLACH:                                    "de:08119:4965",
	RATHAUS_UNTERBRUDEN:                                    "de:08119:4976",
	RATHAUS_ALTHUTTE:                                       "de:08119:4995",
	RATHAUS_DENKENDORF:                                     "de:08116:5008",
	RATHAUS_HEININGEN_BACKNANG:                             "de:08119:5031",
	RATHAUS_HEUTENSBACH:                                    "de:08119:5038",
	RATHAUS_HAUBERSBRONN:                                   "de:08119:5107",
	RATHAUS_NEUSTADT:                                       "de:08119:5148",
	RATHAUS_BEINSTEIN:                                      "de:08119:5164",
	RATHAUS_STRUMPFELBACH_WEINSTADT:                        "de:08119:5178",
	RATHAUS_BEUTELSBACH:                                    "de:08119:5179",
	RATHAUS_WEILER_SCHORNDORF:                              "de:08119:5193",
	RATHAUS_LEUTENBACH:                                     "de:08119:5313",
	RATHAUS_BACKNANG:                                       "de:08119:5324",
	RATHAUS_UNTERWEISSACH:                                  "de:08119:5375",
	RATHAUS_WAIBLINGEN:                                     "de:08119:5410",
	RATHAUS_SCHWAIKHEIM:                                    "de:08119:5419",
	RATHAUS_KORNWESTHEIM:                                   "de:08118:5462",
	RATHAUS_OTTMARSHEIM:                                    "de:08118:5579",
	RATHAUS_OBERSTENFELD:                                   "de:08118:5612",
	RATHAUS_HOCHDORF_EBERDINGEN:                            "de:08118:5621",
	RATHAUS_BISSINGEN_LB:                                   "de:08118:5644",
	RATHAUS_RUTESHEIM:                                      "de:08115:5836",
	RATHAUS_MALMSHEIM:                                      "de:08115:5848",
	RATHAUS_KOHLBERG:                                       "de:08116:5903",
	RATHAUS_HOFINGEN:                                       "de:08115:6045",
	RATHAUS_STUTTGART:                                      "de:08111:6074",
	RATHAUS_OBERLENNINGEN:                                  "de:08116:6715",
	RATHAUS_WEILER_ZUM_STEIN:                               "de:08119:6767",
	RATHAUS_BRETZENACKER:                                   "de:08119:6893",
	RATHAUS_OBERRIEXINGEN:                                  "de:08118:6934",
	RATHAUS_GERLINGEN:                                      "de:08118:7141",
	RATHAUS_GROSSBETTLINGEN:                                "de:08116:7413",
	RATHAUS_GEISLINGEN_STEIGE:                              "de:08117:2",
	RATHAUS_GINGEN_F:                                       "de:08117:20",
	RATHAUS_KUCHEN_F:                                       "de:08117:26",
	RATHAUS_REICHENBACH_IM_TALE_DEGGINGEN:                  "de:08117:51",
	RATHAUS_MUHLHAUSEN_IM_TALE:                             "de:08117:58",
	RATHAUS_WIESENSTEIG:                                    "de:08117:64",
	RATHAUS_BOHMENKIRCH:                                    "de:08117:203",
	RATHAUS_NENNINGEN:                                      "de:08117:329",
	RATHAUS_BIRENBACH:                                      "de:08117:1102",
	RATHAUS_HATTENHOFEN:                                    "de:08117:2052",
	RATHAUS_UHINGEN:                                        "de:08117:2201",
	RATHAUS_GRUIBINGEN:                                     "de:08117:3232",
	RATHAUS_OBERDRACKENSTEIN:                               "de:08117:3502",
	RATHAUS_EBERSBACH_F:                                    "de:08117:7412",
	RATHAUS_SUSSEN:                                         "de:08117:7616",
	RATHAUS_REICHENBACH_UNTER_RECHBERG_DONZDORF:            "de:08117:7631",
	RATHAUS_WINZINGEN:                                      "de:08117:7641",
	RATHAUS_STOTTEN:                                        "de:08117:7686",
	RATHAUS_GOPPINGEN:                                      "de:08117:9303",
	RATHAUS_HASLACH:                                        "de:08415:24319",
	RATHAUS_DONNSTETTEN:                                    "de:08415:28624",
	RATHAUS_GECHINGEN:                                      "de:08235:1510",
	RATHAUS_EBHAUSEN:                                       "de:08235:10207",
	RATHAUS_ALTENSTEIG:                                     "de:08235:7881",
	RATHAUS_LEHNINGEN:                                      "de:08236:3609",
	RATHAUS_NELLINGEN_UL:                                   "de:08425:2613",
	RATHAUS_NOTARIAT_WALDDORF:                              "de:08415:24320",
	RATHAUSDORFPLATZ_HOHENSTAUFEN:                          "de:08117:1304",
	RATHAUSSTADTHALLE_NECKARREMS:                           "de:08118:4887",
	RATHAUSPLATZ_HIRSCHLANDEN:                              "de:08118:3003",
	RATHAUSSTRASSE_EBERDINGEN:                              "de:08118:6469",
	RATTENHARZ_KAISERSTRKIRCHE_LORCH:                       "de:08136:7025",
	RATTENHARZ_PULZHOFWEG_LORCH:                            "de:08136:7026",
	RAUHER_KAPF_BOBLINGEN:                                  "de:08115:4674",
	RAUHWIESENSTRASSE_WINZINGEN:                            "de:08117:7640",
	REALSCHULE_TAMM:                                        "de:08118:3449",
	REALSCHULE_ZOLLBERG:                                    "de:08116:4179",
	REALSCHULE_WERNAU_N:                                    "de:08116:4251",
	REALSCHULE_NURTINGEN:                                   "de:08116:4302",
	REALSCHULE_KORNWESTHEIM:                                "de:08118:5441",
	REALSCHULE_GROSSSACHSENHEIM:                            "de:08118:5742",
	REALSCHULE_GERADSTETTEN:                                "de:08119:5802",
	REALSCHULE_IM_AURAIN_BIETIGHEIM:                        "de:08118:7585",
	REALSCHULE_REMSECK_PATTONVILLE:                         "de:08118:3576",
	REBENWEG_OWEN:                                          "de:08116:4281",
	REBHUHNWEG_LORCH:                                       "de:08136:7019",
	REBSTEIGE_MARKGRONINGEN:                                "de:08118:3325",
	RECHBERGSTRASSE_KORNWESTHEIM:                           "de:08118:5463",
	RECHBERGSTRASSE_WEITMARS:                               "de:08136:7038",
	REESENBANKLE_PLUDERHAUSEN:                              "de:08119:5248",
	REICHENBACH_REICHENBACH_BERGLEN:                        "de:08119:5292",
	REICHENBACH_F_REICHENBACH_F:                            "de:08116:4224",
	REICHENBACHER_STRASSE_BALTMANNSWEILER:                  "de:08116:3806",
	REICHENBACHSTRASSE_REICHENBACH_F:                       "de:08116:4507",
	REICHENBERG_OPPENWEILER:                                "de:08119:4902",
	REICHENHARDTSTRASSE_RECHBERGHAUSEN:                     "de:08117:4039",
	REICHERTSHALDE_LUDWIGSBURG:                             "de:08118:5524",
	REICHSDORFSTRASSE_HOHENSTAUFEN:                         "de:08117:1306",
	REINHARDTSTRASSE_WOLFSCHLUGEN:                          "de:08116:2920",
	REINHOLDMAIERPLATZ_SCHORNDORF:                          "de:08119:5978",
	REINHOLDMAIERSTRASSE_GRUNBACH:                          "de:08119:3704",
	REINHOLDSCHICKPLATZ_HERRENBERG:                         "de:08115:7013",
	REITERHOF_SINDELFINGEN:                                 "de:08115:4664",
	REITERHOF_BOBLINGEN:                                    "de:08115:4691",
	REITHALLE_SCHORNDORF:                                   "de:08119:6680",
	REITPRECHTS_AM_URSPRING_SCHWABISCH_GMUND:               "de:08136:3505",
	REITPRECHTS_NEUBRUNNENGASSE_SCHWABISCH_GMUND:           "de:08136:3518",
	REITSTEIGE_AURICH:                                      "de:08118:6965",
	REIZENWIESEN_WELZHEIM:                                  "de:08119:6997",
	REMBRANDTSTRASSE_STUTTGART:                             "de:08111:374",
	REMBRANDTWEG_KIRCHHEIM_T:                               "de:08116:4358",
	REMSMURRCENTER_FELLBACH:                                "de:08119:2497",
	REMSMURRKLINIK_SCHORNDORF:                              "de:08119:4497",
	REMSMURRKLINIKUM_WINNENDEN:                             "de:08119:6983",
	REMSBRUCKE_WINTERBACH:                                  "de:08119:3714",
	REMSECK_NECKARGRONINGEN:                                "de:08118:3573",
	REMSPARK_WAIBLINGEN:                                    "de:08119:3626",
	REMSSTRASSE_LORCH:                                      "de:08136:7047",
	REMSTALHALLE_WALDHAUSEN_B_SCHORND_LORCH:                "de:08136:7034",
	REMSTALSTRASSE_NECKARREMS:                              "de:08118:5453",
	RENNINGEN_RENNINGEN:                                    "de:08115:7302",
	RENNINGER_STRASSE_MAGSTADT:                             "de:08115:3304",
	RESIDENZSCHLOSS_LUDWIGSBURG:                            "de:08118:5454",
	RESSESTRASSE_STUTTGART:                                 "de:08111:2092",
	RETTERSBURG_RETTERSBURG:                                "de:08119:5269",
	REUSCHKIRCHE_GOPPINGEN:                                 "de:08117:4905",
	REUSSENSTEINSTRASSE_SCHOPFLOCH_LENNINGEN:               "de:08116:4425",
	REUSSENSTEINSTRASSE_BOBLINGEN:                          "de:08115:4685",
	REUSSENSTEINWEG_BONLANDEN:                              "de:08116:2020",
	REUSSENSTEINWEG_HOCHDORF_ES:                            "de:08116:3908",
	REUSTADT_HATTENHOFEN:                                   "de:08117:2156",
	REUTLINGER_STRASSE_SIELMINGEN:                          "de:08116:2144",
	REUTLINGER_STRASSE_ESSLINGEN_N:                         "de:08116:4031",
	REUTLINGER_STRASSE_KEMNAT:                              "de:08116:5019",
	REUTLINGER_STRASSE_STUTTGART:                           "de:08111:6567",
	REUTLINGER_STRASSE_GOPPINGEN:                           "de:08117:1007",
	REUTTE_FRICKENHAUSEN:                                   "de:08116:4467",
	RICHARDWAGNERSTRASSE_BACKNANG:                          "de:08119:6936",
	RICHTHOFENSTRASSE_GERLINGEN:                            "de:08118:6904",
	RIEDBRUNNEN_GARTRINGEN:                                 "de:08115:7151",
	RIEDENBERG_STUTTGART:                                   "de:08111:2548",
	RIEDSEE_STUTTGART:                                      "de:08111:6168",
	RIEDWEG_LEINFELDEN:                                     "de:08116:2099",
	RIEDWIESEN_AICH:                                        "de:08116:5910",
	RIEGELSTRASSE_NELLINGEN:                                "de:08116:2912",
	RIELINGSHAUSER_STRASSE_MARBACH_N:                       "de:08118:5363",
	RIEMENMUHLE_MERKLINGEN:                                 "de:08115:5838",
	RIENHARZ_PFAHLBRONN:                                    "de:08119:6647",
	RIENHARZER_STRSCHULE_PFAHLBRONN:                        "de:08119:6646",
	RIENZHOFER_MUHLE_BITTENFELD:                            "de:08119:5154",
	RIESENGEBIRGSTRASSE_FAURNDAU:                           "de:08117:133",
	RIETENAU_RIETENAU:                                      "de:08119:5359",
	RIETENAUER_WEG_BACKNANG:                                "de:08119:3659",
	RIETER_STRASSE_HOCHDORF_EBERDINGEN:                     "de:08118:6831",
	RIETER_TAL_RIET:                                        "de:08118:5747",
	RIETHMULLERHAUS_STUTTGART:                              "de:08111:2493",
	RIGIPARK_HOLZHEIM_GOPPINGEN:                            "de:08117:3010",
	RILKEREALSCHULE_STUTTGART:                              "de:08111:2474",
	RINGSTRASSE_MAGSTADT:                                   "de:08115:4596",
	RINGSTRASSE_WINNENDEN:                                  "de:08119:5210",
	RINGSTRASSE_GERLINGEN:                                  "de:08118:5916",
	RINGSTRASSE_SALACH:                                     "de:08117:6021",
	RISSHALDE_REICHENBACH_F:                                "de:08116:4240",
	RISSLERINSTRASSE_SCHORNDORF:                            "de:08119:5190",
	RITTERSTRASSE_SCHOCKINGEN:                              "de:08118:3008",
	ROBERTBOSCHCAMPUS_MALMSHEIM:                            "de:08115:5851",
	ROBERTBOSCHKRANKENHAUS_STUTTGART:                       "de:08111:2475",
	ROBERTBOSCHPLATZ_GERLINGEN:                             "de:08118:6906",
	ROBERTBOSCHSTRASSE_SCHWIEBERDINGEN:                     "de:08118:3323",
	ROBERTBOSCHSTRASSE_TAMM:                                "de:08118:3399",
	ROBERTBOSCHSTRASSE_NECKARTENZLINGEN:                    "de:08116:3819",
	ROBERTBOSCHSTRASSE_DETTINGEN_T:                         "de:08116:4272",
	ROBERTBOSCHSTRASSE_DARMSHEIM:                           "de:08115:4608",
	ROBERTBOSCHSTRASSE_PFLUGFELDEN:                         "de:08118:5492",
	ROBERTBOSCHSTRASSE_PLEIDELSHEIM:                        "de:08118:5560",
	ROBERTBOSCHSTRASSE_HOLZGERLINGEN:                       "de:08115:6534",
	ROBERTBOSCHSTRASSE_SCHORNDORF:                          "de:08119:7631",
	ROBERTFRANCKALLEE_LUDWIGSBURG:                          "de:08118:5521",
	ROBERTFRANCKSTRASSE_MURRHARDT:                          "de:08119:5136",
	ROBERTKOCHSTRASSE_PARKSIEDLUNG:                         "de:08116:2936",
	ROBERTKOCHSTRASSE_LUDWIGSBURG:                          "de:08118:7642",
	ROGGENMUHLE_EYBACH:                                     "de:08117:207",
	ROHR_SCHORNBACH:                                        "de:08119:5261",
	ROHR_STUTTGART:                                         "de:08111:6001",
	ROHR_MITTE_STUTTGART:                                   "de:08111:6014",
	ROHRACH_ALLMERSBACH_AM_WEINBERG_ASPACH:                 "de:08119:5361",
	ROHRACHWEG_SCHORNDORF:                                  "de:08119:5971",
	ROHRACKER_STUTTGART:                                    "de:08111:2524",
	ROHRACKER_KELTER_STUTTGART:                             "de:08111:2522",
	ROHRBRONN_ROHRBRONN:                                    "de:08119:5198",
	ROHRER_HOHE_STUTTGART:                                  "de:08111:2616",
	ROHRER_STRASSE_STEINENBRONN:                            "de:08115:7116",
	ROHRER_WEG_STUTTGART:                                   "de:08111:183",
	ROMERSTRASSE_LEONBERG:                                  "de:08115:2333",
	ROMERSTRASSE_DETTINGEN_T:                               "de:08116:4264",
	ROMERWEG_KIRCHBERG_M:                                   "de:08119:6856",
	ROMMELSHAUSEN_ROMMELSHAUSEN:                            "de:08119:7701",
	ROMMELSHAUSER_STRASSE_FELLBACH:                         "de:08119:5433",
	ROMMENTALER_STRASSE_SCHLAT:                             "de:08117:3022",
	RONTGENSTRASSE_OBERESSLINGEN:                           "de:08116:4045",
	RONTGENWEG_GOPPINGEN:                                   "de:08117:9202",
	ROSE_SCHOCKINGEN:                                       "de:08118:3005",
	ROSEGGERWEG_GOPPINGEN:                                  "de:08117:9020",
	ROSENACKERWEG_EGLOSHEIM:                                "de:08118:5507",
	ROSENBERGJOHANNESSTRASSE_STUTTGART:                     "de:08111:2204",
	ROSENBERGSEIDENSTRASSE_STUTTGART:                       "de:08111:6072",
	ROSENBERGPLATZ_STUTTGART:                               "de:08111:2205",
	ROSENPLATZ_GROSSBOTTWAR:                                "de:08118:5608",
	ROSENSTEINBRUCKE_STUTTGART:                             "de:08111:6253",
	ROSENSTEINPARK_STUTTGART:                               "de:08111:262",
	ROSENSTEINSTRASSE_STUTTGART:                            "de:08111:6296",
	ROSENSTRASSE_BERNHAUSEN:                                "de:08116:1999",
	ROSENSTRASSE_KEMNAT:                                    "de:08116:5018",
	ROSENSTRASSE_OBERBRUDEN:                                "de:08119:5380",
	ROSENSTRASSE_GERLINGEN:                                 "de:08118:6903",
	ROSSACKER_BEINSTEIN:                                    "de:08119:6637",
	ROSSBACHSTRASSE_STUTTGART:                              "de:08111:6147",
	ROSSBACHSTRASSE_GOPPINGEN:                              "de:08117:1033",
	ROSSBAUM_GERLINGEN:                                     "de:08118:142",
	ROSSBERGSTAFFEL_BACKNANG:                               "de:08119:5355",
	ROSSBERGSTRASSE_HARTHAUSEN:                             "de:08116:2080",
	ROSSBERGWEG_HOLZGERLINGEN:                              "de:08115:6528",
	ROSSDORF_NURTINGEN:                                     "de:08116:5858",
	ROSSDORF_BERLINER_STRASSE_NURTINGEN:                    "de:08116:4452",
	ROSSDORF_DURERPLATZ_NURTINGEN:                          "de:08116:4453",
	ROSSDORF_HOLBEINSTRASSE_NURTINGEN:                      "de:08116:2940",
	ROSSDORF_KLEEWEG_NURTINGEN:                             "de:08116:2811",
	ROSSDORF_LIEBERMANNSTRASSE_NURTINGEN:                   "de:08116:2965",
	ROSSDORF_SCHULE_NURTINGEN:                              "de:08116:2810",
	ROSSLE_SCHANBACH:                                       "de:08116:4159",
	ROSSMARKT_KIRCHHEIM_T:                                  "de:08116:4267",
	ROSSWAG_SPORTHALLE_ROSSWAG:                             "de:08118:5811",
	ROSSWALDER_STRASSE_HOCHDORF_ES:                         "de:08116:4221",
	ROTBAUMLESFELD_LUDWIGSBURG:                             "de:08118:3466",
	ROTBUHL_SINDELFINGEN:                                   "de:08115:2218",
	ROTE_WEIL_IM_SCHONBUCH:                                 "de:08115:4781",
	ROTE_DOFFINGEN:                                         "de:08115:6756",
	ROTEBUHLREINSBURGSTRASSE_STUTTGART:                     "de:08111:2262",
	ROTENBACH_NECKARTENZLINGEN:                             "de:08116:4488",
	ROTENBERG_STUTTGART:                                    "de:08111:2510",
	ROTENBERGPLATZ_NURTINGEN:                               "de:08116:4447",
	ROTER_BAUM_OBERRIEXINGEN:                               "de:08118:6517",
	ROTER_BERG_SCHONAICH:                                   "de:08115:4812",
	ROTES_KREUZ_LIEBERSBRONN:                               "de:08116:4071",
	ROTES_REUSCH_BARTENBACH_GOPPINGEN:                      "de:08117:4008",
	ROTESIEDLUNG_BURGSTALL_BURGSTETTEN:                     "de:08119:3507",
	ROTHWEG_DENKENDORF:                                     "de:08116:5013",
	ROTWEG_HOCHBERG:                                        "de:08118:5624",
	ROTWIESENSTRASSE_STUTTGART:                             "de:08111:2561",
	ROTZEIL_BONLANDEN:                                      "de:08116:2022",
	RUBEZAHLWEG_STUTTGART:                                  "de:08111:6174",
	RUDERSBERG_RUDERSBERG:                                  "de:08119:5036",
	RUDERSBERGER_STRASSE_MIEDELSBACH:                       "de:08119:7662",
	RUDOLFDIESELSTR_LAICHINGEN:                             "de:08425:2591",
	RUDOLFDIESELSTRASSE_ROMMELSHAUSEN:                      "de:08119:3740",
	RUDOLFSOPHIENSTIFT_STUTTGART:                           "de:08111:2209",
	RUDOLFSHOHE_AICH:                                       "de:08116:2000",
	RUHBANK_FERNSEHTURM_STUTTGART:                          "de:08111:6128",
	RUHRBERG_MUNCHINGEN:                                    "de:08118:7607",
	RUHRSTRASSE_LUDWIGSBURG:                                "de:08118:5472",
	RUIT_RUIT:                                              "de:08116:2968",
	RUMOLDREALSCHULE_ROMMELSHAUSEN:                         "de:08119:4936",
	RUNDSPORTHALLE_LUDWIGSBURG:                             "de:08118:3415",
	RUNDSPORTHALLE_WAIBLINGEN:                              "de:08119:5161",
	RUSSISCHE_KIRCHE_STUTTGART:                             "de:08111:6071",
	RUTESHEIM_LEONBERG:                                     "de:08115:1302",
	SAARSTRASSE_ASPERG:                                     "de:08118:3404",
	SAARSTRASSE_KIRCHHEIM_T:                                "de:08116:3948",
	SACHSENHEIM_GROSSSACHSENHEIM:                           "de:08118:5639",
	SACHSENHEIMER_STRASSE_OBERRIEXINGEN:                    "de:08118:5640",
	SACHSENHEIMER_WEG_BESIGHEIM:                            "de:08118:3506",
	SAERSTRASSE_NURTINGEN:                                  "de:08116:4444",
	SAGMUHLE_HEMMINGEN:                                     "de:08118:4524",
	SALACH_SALACH:                                          "de:08117:156",
	SALAMANDERSTRASSE_FAURNDAU:                             "de:08117:1220",
	SALAMANDERWEG_STUTTGART:                                "de:08111:145",
	SALIERSCHULE_WAIBLINGEN:                                "de:08119:6677",
	SALIERSTRASSE_WAIBLINGEN:                               "de:08119:5405",
	SALZACKER_STUTTGART:                                    "de:08111:360",
	SALZACKER_VAIHINGEN_E:                                  "de:08118:5757",
	SALZBURGER_STRASSE_MAUBACH:                             "de:08119:3664",
	SALZWIESENSTRASSE_STUTTGART:                            "de:08111:6105",
	SAMARITERSTIFT_LEONBERG:                                "de:08115:2319",
	SANGERSTRASSE_EISLINGEN_F:                              "de:08117:1024",
	SANKT_MARTIN_HERRENBERG:                                "de:08115:7076",
	SANKT_PAULS_KIRCHE_GOPPINGEN:                           "de:08117:3006",
	SATTLEREI_EISELE_HESSIGHEIM:                            "de:08118:5578",
	SAUERBRUNNEN_FAURNDAU:                                  "de:08117:7007",
	SAUERHOFLE_MURRHARDT:                                   "de:08119:5297",
	SAUSERHOF_GROSSBOTTWAR:                                 "de:08118:5610",
	SAUSTEIGE_GOPPINGEN:                                    "de:08117:9014",
	SCSTADION_GEISLINGEN_STEIGE:                            "de:08117:210",
	SCHADBERG_KAISERSBACH:                                  "de:08119:5283",
	SCHAFERSFELD_SCHULEN_LORCH:                             "de:08136:7027",
	SCHAFGARTENSTRASSE_MUSBERG:                             "de:08116:2158",
	SCHAFGASSE_BOBLINGEN:                                   "de:08115:4684",
	SCHAFHAUS_KLEINASPACH:                                  "de:08119:5362",
	SCHAFHAUSER_STRASSE_AIDLINGEN:                          "de:08115:4776",
	SCHAFHOF_KIRCHHEIM_T:                                   "de:08116:4354",
	SCHAFHOF_WELZHEIM:                                      "de:08119:5064",
	SCHAFSTRASSE_ROMMELSHAUSEN:                             "de:08119:4938",
	SCHAFWASCHE_HEIMSHEIM:                                  "de:08236:3509",
	SCHAICHHOFSIEDLUNG_WEIL_IM_SCHONBUCH:                   "de:08115:4185",
	SCHAICHHOFSTRASSE_WEIL_IM_SCHONBUCH:                    "de:08115:4145",
	SCHANZLE_WAIBLINGEN:                                    "de:08119:5957",
	SCHARNHAUSER_BRUCKE_STUTTGART:                          "de:08111:2040",
	SCHARNHAUSER_PARK_SCHARNHAUSER_PARK:                    "de:08116:2971",
	SCHARNHAUSER_STRASSE_RUIT:                              "de:08116:2914",
	SCHATTENGRUND_STUTTGART:                                "de:08111:2118",
	SCHAUBER_BESIGHEIM:                                     "de:08118:3505",
	SCHAUCHERT_HEMMINGEN:                                   "de:08118:4528",
	SCHEERSTRDRENGEL_REALSCHULE_EISLINGEN_F:                "de:08117:9812",
	SCHELLINGSTRASSE_NURTINGEN:                             "de:08116:2998",
	SCHELLINGSTRASSE_WAIBLINGEN:                            "de:08119:7092",
	SCHELMENACKER_HASLACH:                                  "de:08115:7071",
	SCHELMENSTRASSEFRIEDHOF_BARTENBACH_GOPPINGEN:           "de:08117:4012",
	SCHELMENWASEN_STUTTGART:                                "de:08111:363",
	SCHELMENWASEN_NURTINGEN:                                "de:08116:4450",
	SCHELMENWIESEN_GARTRINGEN:                              "de:08115:7152",
	SCHELZTOR_ESSLINGEN_N:                                  "de:08116:4007",
	SCHEMPPSTRASSE_STUTTGART:                               "de:08111:6130",
	SCHICKARDSTRASSE_BOBLINGEN:                             "de:08115:4718",
	SCHICKHARDTSCHULE_STUTTGART:                            "de:08111:2208",
	SCHIESSTAL_NECKARGRONINGEN:                             "de:08118:5549",
	SCHIESSTALE_HERRENBERG:                                 "de:08115:7021",
	SCHILDFARNWEG_STUTTGART:                                "de:08111:2526",
	SCHILDWACHTWEG_GEISLINGEN_STEIGE:                       "de:08117:100",
	SCHILLERREALSCHULE_GOPPINGEN:                           "de:08117:3386",
	SCHILLERHOHE_MARBACH_N:                                 "de:08118:3550",
	SCHILLERHOHE_BOSCH_GERLINGEN:                           "de:08118:2313",
	SCHILLERPLATZ_BACKNANG:                                 "de:08119:3608",
	SCHILLERPLATZ_ESSLINGEN_N:                              "de:08116:4040",
	SCHILLERPLATZ_NURTINGEN:                                "de:08116:4304",
	SCHILLERPLATZ_SCHORNDORF:                               "de:08119:5216",
	SCHILLERPLATZ_KIRCHHEIM_N:                              "de:08118:5724",
	SCHILLERPLATZ_BIETIGHEIM:                               "de:08118:7405",
	SCHILLERPLATZ_GOPPINGEN:                                "de:08117:1001",
	SCHILLERSCHULE_BISSINGEN_LB:                            "de:08118:7431",
	SCHILLERSCHULE_BITTENFELD:                              "de:08119:7549",
	SCHILLERSTRASSE_PLATTENHARDT:                           "de:08116:2051",
	SCHILLERSTRASSE_NEUHAUSEN_F:                            "de:08116:2907",
	SCHILLERSTRASSE_SCHONAICH:                              "de:08115:3217",
	SCHILLERSTRASSE_REICHENBACH_F:                          "de:08116:3905",
	SCHILLERSTRASSE_NECKARTAILFINGEN:                       "de:08116:3937",
	SCHILLERSTRASSE_KLEINGLATTBACH:                         "de:08118:5759",
	SCHILLERSTRASSE_RUTESHEIM:                              "de:08115:5826",
	SCHILLERSTRASSE_KOHLBERG:                               "de:08116:6707",
	SCHILLERSTRASSE_KUCHEN_F:                               "de:08117:14",
	SCHILLERSTRASSE_WASCHENBEUREN:                          "de:08117:1108",
	SCHILLERSTRASSE_ZELL_A:                                 "de:08117:2113",
	SCHILLERSTRASSE_GOPPINGEN:                              "de:08117:4037",
	SCHILLERSTRASSE_SCHWABISCH_GMUND:                       "de:08136:3029",
	SCHIPPERTSACKER_WAIBLINGEN:                             "de:08119:3690",
	SCHLACHTHAUSBRUCKE_ESSLINGEN_N:                         "de:08116:4008",
	SCHLACHTHOF_STUTTGART:                                  "de:08111:6223",
	SCHLATER_STRASSE_SUSSEN:                                "de:08117:7615",
	SCHLATTERHOHE_B465_SCHOPFLOCH_LENNINGEN:                "de:08116:4422",
	SCHLATTERHOHE_L_1212_SCHOPFLOCH_LENNINGEN:              "de:08116:3932",
	SCHLECHTBACH_SCHLECHTBACH:                              "de:08119:5792",
	SCHLEIERMACHERSTRASSE_LEONBERG:                         "de:08115:2309",
	SCHLICHTENER_STRASSE_SCHORNDORF:                        "de:08119:3759",
	SCHLIEFFENSTRASSE_LUDWIGSBURG:                          "de:08118:7643",
	SCHLIENZTOURS_ROMMELSHAUSEN:                            "de:08119:3739",
	SCHLIERBACHER_DREIECK_KIRCHHEIM_T:                      "de:08116:2826",
	SCHLIPFWEG_SCHORNDORF:                                  "de:08119:5259",
	SCHLOSS_SCHOCKINGEN:                                    "de:08118:3006",
	SCHLOSS_WALDENBUCH:                                     "de:08115:3119",
	SCHLOSS_URBACH:                                         "de:08119:6692",
	SCHLOSS_FAVORITE_LUDWIGSBURG:                           "de:08118:5506",
	SCHLOSS_FILSECK_FAURNDAU:                               "de:08117:5021",
	SCHLOSS_LAUTERECK_SULZBACH_M:                           "de:08119:4904",
	SCHLOSSJOHANNESSTRASSE_STUTTGART:                       "de:08111:6219",
	SCHLOSSBERGALLEE_BONNIGHEIM:                            "de:08118:6855",
	SCHLOSSBERGHALLE_DETTINGEN_T:                           "de:08116:7409",
	SCHLOSSERSTRASSE_NEUHAUSEN_F:                           "de:08116:2922",
	SCHLOSSFELD_BONNIGHEIM:                                 "de:08118:6840",
	SCHLOSSGARTEN_NEIDLINGEN:                               "de:08116:4409",
	SCHLOSSGYMNASIUM_KIRCHHEIM_T:                           "de:08116:4350",
	SCHLOSSHALDE_GEISLINGEN_STEIGE:                         "de:08117:106",
	SCHLOSSHALDENSTRASSE_HEMMINGEN:                         "de:08118:6932",
	SCHLOSSHOF_ALDINGEN:                                    "de:08118:4888",
	SCHLOSSHOF_FORNSBACH:                                   "de:08119:5304",
	SCHLOSSHOFSTRASSE_RECHBERGHAUSEN:                       "de:08117:4026",
	SCHLOSSLESFELD_LUDWIGSBURG:                             "de:08118:5502",
	SCHLOSSMARKT_RECHBERGHAUSEN:                            "de:08117:4027",
	SCHLOSSMUHLE_KIRCHENKIRNBERG:                           "de:08119:5986",
	SCHLOSSPLATZ_NEUHAUSEN_F:                               "de:08116:2906",
	SCHLOSSPLATZ_STUTTGART:                                 "de:08111:6022",
	SCHLOSSPLATZBAHNHOF_EISLINGEN_F:                        "de:08117:1049",
	SCHLOSSSTRASSE_SCHORNDORF:                              "de:08119:5969",
	SCHLOSSSTRASSE_KLEININGERSHEIM:                         "de:08118:6718",
	SCHLOTTERBECKSTRASSE_STUTTGART:                         "de:08111:268",
	SCHLOTWIESE_STUTTGART:                                  "de:08111:2465",
	SCHLUCHTWEG_WEILHEIM_T:                                 "de:08116:4325",
	SCHMELLBACHTAL_LEINFELDEN:                              "de:08116:2125",
	SCHMELLENHOF_WUSTENROT:                                 "de:08125:982",
	SCHMIDENER_STRASSE_WAIBLINGEN:                          "de:08119:3647",
	SCHMIDHAUSEN_ABZWEIG_BEILSTEIN:                         "de:08125:4677",
	SCHMIEDE_HERTMANNSWEILER:                               "de:08119:3622",
	SCHNALLENACKER_MALMSHEIM:                               "de:08115:3187",
	SCHNARRENBERG_STUTTGART:                                "de:08111:2471",
	SCHNECKEN_PLATTENHARDT:                                 "de:08116:2027",
	SCHNEIDER_NEUSTADT:                                     "de:08119:6676",
	SCHNEIDERACKERSTRASSE_STUTTGART:                        "de:08111:6276",
	SCHOCKINGER_STRASSE_HEMMINGEN:                          "de:08118:4535",
	SCHONAICHER_FIRST_BOBLINGEN:                            "de:08115:4672",
	SCHONAICHER_STRASSE_BOBLINGEN:                          "de:08115:3127",
	SCHONBERG_STUTTGART:                                    "de:08111:2575",
	SCHONBERG_HOLZGERLINGEN:                                "de:08115:4763",
	SCHONBERGSTRASSE_KEMNAT:                                "de:08116:3050",
	SCHONBLICK_REICHENBACH_F:                               "de:08116:4236",
	SCHONBLICK_GROTZINGEN:                                  "de:08116:4473",
	SCHONBLICKSTRASSE_WUSTENROT:                            "de:08125:44219",
	SCHONBRONN_GRAB:                                        "de:08119:7506",
	SCHONBUCHSTRASSE_PLATTENHARDT:                          "de:08116:2055",
	SCHONBUCHSTRASSE_BOBLINGEN:                             "de:08115:4694",
	SCHONBUCHSTRASSE_SPORTH_BOBLINGEN:                      "de:08115:4693",
	SCHONTALWEG_WIESENSTEIG:                                "de:08117:60",
	SCHOPFLENBERG_BEZGENRIET:                               "de:08117:2104",
	SCHOPFLOCH_ELTINGEN:                                    "de:08115:7030",
	SCHORNBACHER_WEG_SCHORNDORF:                            "de:08119:5255",
	SCHORNDORF_SCHORNDORF:                                  "de:08119:7703",
	SCHORNDORFER_STRASSE_GRUNBACH:                          "de:08119:3752",
	SCHORNDORFER_STRASSE_OBERESSLINGEN:                     "de:08116:4201",
	SCHORNDORFER_STRASSE_STEINENBERG:                       "de:08119:5111",
	SCHORNDORFER_STRASSE_REICHENBACH_F:                     "de:08116:6499",
	SCHORNDORFER_STRASSE_URBACH:                            "de:08119:6694",
	SCHORNDORFER_TOR_LUDWIGSBURG:                           "de:08118:4894",
	SCHOSSHOFE_SINDELFINGEN:                                "de:08115:4641",
	SCHOTTSTRASSE_STUTTGART:                                "de:08111:2192",
	SCHOZACHER_STRASSE_STUTTGART:                           "de:08111:6293",
	SCHOZACHSTRASSE_WALDREMS:                               "de:08119:3616",
	SCHRANNE_STUTTGART:                                     "de:08111:2603",
	SCHUBARTSTRASSE_BISSINGEN_LB:                           "de:08118:5687",
	SCHUBARTSTRASSE_FAURNDAU:                               "de:08117:7015",
	SCHUBERTSTRASSE_WOLFSCHLUGEN:                           "de:08116:2808",
	SCHULSPORTZENTRUM_SIELMINGEN:                           "de:08116:2042",
	SCHULBAD_GOPPINGEN:                                     "de:08117:9008",
	SCHULE_ALTENRIET:                                       "de:08116:2065",
	SCHULE_OBERENSINGEN:                                    "de:08116:2948",
	SCHULE_OBERSTENFELD:                                    "de:08118:3338",
	SCHULE_KIRCHBERG_M:                                     "de:08119:3530",
	SCHULE_HESSIGHEIM:                                      "de:08118:3564",
	SCHULE_ALTHUTTE:                                        "de:08119:3719",
	SCHULE_DEIZISAU:                                        "de:08116:4058",
	SCHULE_SCHANBACH:                                       "de:08116:4157",
	SCHULE_DATZINGEN:                                       "de:08115:4194",
	SCHULE_SULZGRIES:                                       "de:08116:4202",
	SCHULE_FRICKENHAUSEN:                                   "de:08116:4429",
	SCHULE_NECKARTAILFINGEN:                                "de:08116:4484",
	SCHULE_BEMPFLINGEN:                                     "de:08116:4490",
	SCHULE_SECHSELBERG:                                     "de:08119:4986",
	SCHULE_ALLMERSBACH_IM_TAL:                              "de:08119:5037",
	SCHULE_SPIEGELBERG:                                     "de:08119:5044",
	SCHULE_LEUTENBACH:                                      "de:08119:5319",
	SCHULE_GROSSINGERSHEIM:                                 "de:08118:5544",
	SCHULE_GRONAU:                                          "de:08118:5615",
	SCHULE_HOHENHASLACH:                                    "de:08118:5630",
	SCHULE_ALFDORF:                                         "de:08119:6640",
	SCHULE_NEUFFEN:                                         "de:08116:6775",
	SCHULE_SACHSENWEILER:                                   "de:08119:6935",
	SCHULE_MONCHBERG:                                       "de:08115:7025",
	SCHULE_GROSSERLACH:                                     "de:08119:7504",
	SCHULE_AUFHAUSEN_GEISLINGEN:                            "de:08117:10498",
	SCHULE_TREFFELHAUSEN:                                   "de:08117:17683",
	SCHULE_DETTENHAUSEN:                                    "de:08416:6354",
	SCHULE_WUSTENROT:                                       "de:08125:984",
	SCHULE_SCHALKSTETTEN:                                   "de:08425:2577",
	SCHULEHALLENBAD_WEIL_IM_SCHONBUCH:                      "de:08115:5948",
	SCHULERHOF_BURG:                                        "de:08119:6659",
	SCHULMEISTERBUCHE_HERRENBERG:                           "de:08115:4758",
	SCHULSTRASSE_TAILFINGEN:                                "de:08115:4834",
	SCHULSTRASSE_PLUDERHAUSEN:                              "de:08119:5763",
	SCHULSTRASSE_KIRCHBERG_M:                               "de:08119:6873",
	SCHULSTRASSE_UHINGEN:                                   "de:08117:2039",
	SCHULSTRASSE_WANGEN_GP:                                 "de:08117:5010",
	SCHULTHEISSSCHNEIDERSTRASSE_GEISLINGEN_STEIGE:          "de:08117:102",
	SCHULZEDELITZSCHSTRASSE_STUTTGART:                      "de:08111:2600",
	SCHULZENTRUM_SCHWIEBERDINGEN:                           "de:08118:3339",
	SCHULZENTRUM_GROSSBOTTWAR:                              "de:08118:3553",
	SCHULZENTRUM_REICHENBACH_F:                             "de:08116:4235",
	SCHULZENTRUM_WENDLINGEN_N:                              "de:08116:4296",
	SCHULZENTRUM_NELLINGEN:                                 "de:08116:5021",
	SCHULZENTRUM_RUDERSBERG:                                "de:08119:5121",
	SCHULZENTRUM_WELZHEIM:                                  "de:08119:5131",
	SCHULZENTRUM_BONNIGHEIM:                                "de:08118:5727",
	SCHULZENTRUM_RUTESHEIM:                                 "de:08115:5829",
	SCHULZENTRUM_STEINHEIM_M:                               "de:08118:6988",
	SCHULZENTRUM_DONZDORF:                                  "de:08117:7623",
	SCHULZENTRUM_BEILSTEIN:                                 "de:08125:32310",
	SCHULZENTRUM_LAICHINGEN:                                "de:08425:2653",
	SCHULZENTRUM_AM_BERG_WENDLINGEN_N:                      "de:08116:6779",
	SCHULZENTRUM_MURKENBACH_BOBLINGEN:                      "de:08115:4670",
	SCHULZENTRUM_NORD_SERACH:                               "de:08116:4135",
	SCHUMISBERG_LEONBERG:                                   "de:08115:2335",
	SCHURWALDHALLE_OBERBERKEN:                              "de:08119:5990",
	SCHURWALDHOHE_BALTMANNSWEILER:                          "de:08116:4076",
	SCHURWALDSTRASSE_BEUTELSBACH:                           "de:08119:5399",
	SCHURWALDWEG_BIRENBACH:                                 "de:08117:1100",
	SCHUTZENHOF_WASCHENBEUREN:                              "de:08117:1105",
	SCHUTZENSTRASSE_WENDLINGEN_N:                           "de:08116:4297",
	SCHUTZENSTRASSE_GOPPINGEN:                              "de:08117:1057",
	SCHUTZENWEG_DAGERSHEIM:                                 "de:08115:4633",
	SCHWABBEBELSTRASSE_STUTTGART:                           "de:08111:6206",
	SCHWABREINSBURGSTRASSE_STUTTGART:                       "de:08111:2207",
	SCHWABENGALERIE_STUTTGART:                              "de:08111:6729",
	SCHWABENLANDHALLE_FELLBACH:                             "de:08119:39",
	SCHWABENSTRASSE_SCHONAICH:                              "de:08115:4829",
	SCHWABISCH_GMUND_SCHWABISCH_GMUND:                      "de:08136:3077",
	SCHWABSTRASSE_BOBLINGEN:                                "de:08115:4667",
	SCHWABSTRASSE_STUTTGART:                                "de:08111:6052",
	SCHWABSTRASSE_WAIBLINGEN:                               "de:08119:7095",
	SCHWAIKHEIM_SCHWAIKHEIM:                                "de:08119:1603",
	SCHWALBENFLUG_GRAB:                                     "de:08119:7518",
	SCHWALBENHALDE_BESIGHEIM:                               "de:08118:7086",
	SCHWALBENWEG_SIRNAU:                                    "de:08116:4051",
	SCHWALBENWEG_PLUDERHAUSEN:                              "de:08119:5237",
	SCHWAMMHOF_MURRHARDT:                                   "de:08119:6689",
	SCHWANEN_WAIBLINGEN:                                    "de:08119:3646",
	SCHWANFELD_PLUDERHAUSEN:                                "de:08119:5247",
	SCHWARENBERGSTRASSE_STUTTGART:                          "de:08111:2059",
	SCHWARMERWEG_STUTTGART:                                 "de:08111:202",
	SCHWARZWALDSTRASSE_BIETIGHEIM:                          "de:08118:5799",
	SCHWARZWALDSTRASSE_GARTRINGEN:                          "de:08115:7053",
	SCHWEISSBRUCKE_ERDMANNHAUSEN:                           "de:08118:5370",
	SCHWERTMUHLE_OBERESSLINGEN:                             "de:08116:4042",
	SCHWERTSTRASSE_SINDELFINGEN:                            "de:08115:5818",
	SCHWIEBERDINGEN_SCHWIEBERDINGEN:                        "de:08118:5795",
	SCHWIEBERDINGER_STRASSE_MOGLINGEN:                      "de:08118:3443",
	SCHWIEBERDINGER_STRASSE_HEMMINGEN:                      "de:08118:3579",
	SCHWILKENHOFSTRASSE_STUTTGART:                          "de:08111:205",
	SEDANSTRASSE_SERSHEIM:                                  "de:08118:7421",
	SEE_BERKHEIM:                                           "de:08116:4144",
	SEE_BISSINGEN_AN_DER_TECK_T:                            "de:08116:4416",
	SEE_WEIL_IM_SCHONBUCH:                                  "de:08115:4790",
	SEE_HEIMSHEIM:                                          "de:08236:3510",
	SEE_SEESTR_WEIL_IM_SCHONBUCH:                           "de:08115:3144",
	SEEBRUCKENMUHLE_MUSBERG:                                "de:08116:7104",
	SEEBRUCKENMUHLE_KLARWERK_MUSBERG:                       "de:08116:7112",
	SEEGARTEN_LEONBERG:                                     "de:08115:2349",
	SEEGER_DOFFINGEN:                                       "de:08115:4771",
	SEEHALDE_URBACH:                                        "de:08119:6697",
	SEEHALDENWEG_HOFEN:                                     "de:08119:5205",
	SEEHOF_BACKNANG:                                        "de:08119:3680",
	SEEHOFWEG_BACKNANG:                                     "de:08119:5331",
	SEELACH_RUDERSBERG:                                     "de:08119:5034",
	SEEMUHLE_STRUMPFELBACH_WEINSTADT:                       "de:08119:5175",
	SEEMUHLE_ROSSWAG:                                       "de:08118:6849",
	SEEPLATZ_KORB:                                          "de:08119:5415",
	SEESTRASSE_HASLACH:                                     "de:08415:29040",
	SEESTRASSE_WINNENDEN:                                   "de:08119:3636",
	SEESTRASSE_ROMMELSHAUSEN:                               "de:08119:3743",
	SEESTRASSE_SCHONAICH:                                   "de:08115:4856",
	SEESTRASSE_HOLZMADEN:                                   "de:08116:6773",
	SEESTRASSE_KLEINBOTTWAR:                                "de:08118:6986",
	SEESTRASSE_UHINGEN:                                     "de:08117:2035",
	SEGELFALTERSTRASSE_STUTTGART:                           "de:08111:203",
	SEHNINGEN_BAD_BOLL:                                     "de:08117:3335",
	SEIBOLDSWEILER_WELZHEIM:                                "de:08119:5073",
	SEMINAR_BACKNANG:                                       "de:08119:5350",
	SENEFELDERSTRASSE_STUTTGART:                            "de:08111:2053",
	SENIORENZENTRUM_WAIBLINGEN:                             "de:08119:3699",
	SENIORENZENTRUM_HERRENBERG:                             "de:08115:4801",
	SERACHER_STRASSE_SERACH:                                "de:08116:4134",
	SERSHEIM_SERSHEIM:                                      "de:08118:5782",
	SEYFFERSTRASSE_STUTTGART:                               "de:08111:2263",
	SIEBERSBACH_SULZBACH_M:                                 "de:08119:4955",
	SIEDLUNG_GERLINGEN:                                     "de:08118:141",
	SIEDLUNG_AICH:                                          "de:08116:2003",
	SIEDLUNG_WOLFSCHLUGEN:                                  "de:08116:2921",
	SIEDLUNG_SCHANBACH:                                     "de:08116:4162",
	SIEDLUNG_OHMDEN:                                        "de:08116:4392",
	SIEDLUNG_HAUSEN_BAD_UBERKINGEN:                         "de:08117:91",
	SIEDLUNG_SCHNITTLINGEN:                                 "de:08117:318",
	SIEDLUNG_OTTENBACH:                                     "de:08117:6014",
	SIEDLUNG_BUNZWANGEN:                                    "de:08117:8005",
	SIEDLUNG_GUSSENSTADT:                                   "de:08135:1156",
	SIEGELSBERG_MURRHARDT:                                  "de:08119:4919",
	SIEGENBERGPLATZ_REICHENBACH_F:                          "de:08116:4388",
	SIEGLESTRASSE_STUTTGART:                                "de:08111:112",
	SIELMINGER_STRASSE_STETTEN_LEINFELDENECHT:              "de:08116:2139",
	SIEMENSSTRASSE_SCHWIEBERDINGEN:                         "de:08118:3324",
	SIEMENSSTRASSE_HOLZGERLINGEN:                           "de:08115:6535",
	SIEMENSSTRASSE_DECKENPFRONN:                            "de:08115:7048",
	SIEMENSSTRASSE_WANGEN_GP:                               "de:08117:1136",
	SIEMENSSTRASSE_SCHLIERBACH:                             "de:08117:7799",
	SIGMARINGER_STRASSE_STUTTGART:                          "de:08111:354",
	SILBERBURGREINSBURGSTRASSE_STUTTGART:                   "de:08111:2264",
	SILBERSTOLLEN_WUSTENROT:                                "de:08125:32451",
	SILBERWALD_STUTTGART:                                   "de:08111:129",
	SILBERWEG_BOBLINGEN:                                    "de:08115:4683",
	SILCHERSCHULE_EISLINGEN_F:                              "de:08117:1012",
	SILCHERSTRASSE_SCHOCKINGEN:                             "de:08118:3007",
	SILCHERSTRASSE_HERRENBERG:                              "de:08115:4850",
	SILCHERSTRASSE_SCHNAIT:                                 "de:08119:5402",
	SILCHERSTRASSE_BUNZWANGEN:                              "de:08117:8004",
	SILCHERWEG_FELLBACH:                                    "de:08119:2517",
	SILLENBUCH_STUTTGART:                                   "de:08111:6134",
	SILVANERSTRASSE_ROSSWAG:                                "de:08118:7842",
	SINDELFINGEN_SINDELFINGEN:                              "de:08115:3201",
	SINDELFINGER_STRASSE_BOBLINGEN:                         "de:08115:4205",
	SINDELFINGER_STRASSE_MAICHINGEN:                        "de:08115:4642",
	SINDLINGEN_AUSSIEDLERHOFE_UNTERJETTINGEN:               "de:08115:6931",
	SINDLINGEN_ORTSMITTE_UNTERJETTINGEN:                    "de:08115:4858",
	SINDLINGER_STRASSE_HASLACH:                             "de:08115:4742",
	SINDLINGER_STRASSE_OBERJETTINGEN:                       "de:08115:4861",
	SINDLINGER_STRASSE_NEBRINGEN:                           "de:08115:6745",
	SIRNAUER_BRUCKE_OBERESSLINGEN:                          "de:08116:4084",
	SIRNAUER_HOF_SIRNAU:                                    "de:08116:4050",
	SKILIFT_TREFFELHAUSEN:                                  "de:08117:7688",
	SKILIFT_LORCH:                                          "de:08136:7041",
	SOLITUDE_STUTTGART:                                     "de:08111:2316",
	SOLITUDEALLEE_LUDWIGSBURG:                              "de:08118:3420",
	SOLITUDESTRASSE_EISLINGEN_F:                            "de:08117:1043",
	SOLO_MAICHINGEN:                                        "de:08115:3210",
	SOMMERHALDE_WAIBLINGEN:                                 "de:08119:4940",
	SOMMERHOFENSTRASSE_SINDELFINGEN:                        "de:08115:4625",
	SOMMERRAIN_STUTTGART:                                   "de:08111:1300",
	SOMMERRAIN_BACKNANG:                                    "de:08119:5338",
	SONATENWEG_STUTTGART:                                   "de:08111:114",
	SONNE_HOCHDORF_REMSECK:                                 "de:08118:5592",
	SONNE_TREFFELHAUSEN:                                    "de:08117:317",
	SONNE_LORCH:                                            "de:08136:7014",
	SONNE_CENTER_GEISLINGEN_STEIGE:                         "de:08117:36",
	SONNENBERG_BONNIGHEIM:                                  "de:08118:5728",
	SONNENBERG_STUTTGART:                                   "de:08111:6167",
	SONNENBERGSCHULE_AIDLINGEN:                             "de:08115:7028",
	SONNENBERGSTRASSE_SINDELFINGEN:                         "de:08115:4661",
	SONNENBRUCKE_RECHBERGHAUSEN:                            "de:08117:4042",
	SONNENBRUNNEN_MOGLINGEN:                                "de:08118:3444",
	SONNENBUHL_STUTTGART:                                   "de:08111:2528",
	SONNENHALDE_MUSBERG:                                    "de:08116:2138",
	SONNENHALDE_SCHARNHAUSEN:                               "de:08116:5011",
	SONNENHALDE_HOFEN_BONNIGHEIM:                           "de:08118:6700",
	SONNENHOF_KLEINASPACH:                                  "de:08119:6843",
	SONNENWEG_NOTZINGEN:                                    "de:08116:3913",
	SONNENWEG_LORCH:                                        "de:08136:7020",
	SOS_KINDERDORF_OBERBERKEN:                              "de:08119:5989",
	SPAICHINGER_STRASSE_STUTTGART:                          "de:08111:2131",
	SPECHTSHOF_REICHENBACH_BERGLEN:                         "de:08119:5291",
	SPIELKARTENMUSEUM_LEINFELDEN:                           "de:08116:7103",
	SPIELPLATZ_KIRCHHEIM_T:                                 "de:08116:4334",
	SPITTLER_STIFT_SCHORNDORF:                              "de:08119:5214",
	SPITZACKER_NECKARTENZLINGEN:                            "de:08116:4485",
	SPITZACKER_PLATTENHARDT:                                "de:08116:6736",
	SPITZHOLZSTRASSE_SINDELFINGEN:                          "de:08115:2220",
	SPORTGELANDE_AISCHBACH_PEROUSE:                         "de:08115:7831",
	SPORTGELANDE_TENNWENGERT_OEFFINGEN:                     "de:08119:2383",
	SPORTHALLE_WENDLINGEN_N:                                "de:08116:7578",
	SPORTHALLE_BUHL_RUTESHEIM:                              "de:08115:5936",
	SPORTPARK_FEUERBACH_STUTTGART:                          "de:08111:153",
	SPORTPARK_WEIL_WEIL:                                    "de:08116:4064",
	SPORTPLATZ_FLACHT:                                      "de:08115:4565",
	SPORTPLATZ_RIELINGSHAUSEN:                              "de:08118:5367",
	SPORTPLATZ_PFLUGFELDEN:                                 "de:08118:5494",
	SPORTPLATZ_ALTENRIET:                                   "de:08116:6708",
	SPORTPLATZ_GEMMRIGHEIM:                                 "de:08118:6890",
	SPORTPLATZ_LOCHGAU:                                     "de:08118:7598",
	SPORTPLATZ_BARTENBACH_GOPPINGEN:                        "de:08117:4914",
	SPORTPLATZ_WINZINGEN:                                   "de:08117:7642",
	SPORTPLATZE_GOPPINGEN:                                  "de:08117:9000",
	SPORTZENTRUM_MUNCHINGEN:                                "de:08118:4520",
	SPORTZENTRUM_LEINFELDEN:                                "de:08116:7177",
	SSBZENTRUM_STUTTGART:                                   "de:08111:350",
	ST_GOTTH_WEILERBACHWEG_GOPPINGEN:                       "de:08117:3011",
	STAATSGALERIE_STUTTGART:                                "de:08111:6024",
	STADION_SINDELFINGEN:                                   "de:08115:2221",
	STADION_WALDENBUCH:                                     "de:08115:3130",
	STADION_KIRCHHEIM_T:                                    "de:08116:4320",
	STADION_SCHONAICH:                                      "de:08115:4824",
	STADION_DENKENDORF:                                     "de:08116:5006",
	STADION_KORNWESTHEIM:                                   "de:08118:5469",
	STADIONHALLE_MARBACH_N:                                 "de:08118:3549",
	STADIONSTRASSE_ECHTERDINGEN:                            "de:08116:7180",
	STADTBADALTENWOHNANLAGE_DITZINGEN:                      "de:08118:3010",
	STADTBAHNHOF_VAIHINGEN_E:                               "de:08118:5932",
	STADTBIBLIOTHEK_STUTTGART:                              "de:08111:6116",
	STADTBIBLIOTHEK_EBERSBACH_F:                            "de:08117:7637",
	STADTBUCHEREI_LEINFELDEN:                               "de:08116:2147",
	STADTFRIEDHOF_HERRENBERG:                               "de:08115:7075",
	STADTGARTEN_WEIL_DER_STADT:                             "de:08115:4574",
	STADTGRENZE_STUTTGART:                                  "de:08111:4013",
	STADTHALLE_SINDELFINGEN:                                "de:08115:2227",
	STADTHALLE_KORNTAL:                                     "de:08118:2627",
	STADTHALLE_KIRCHHEIM_T:                                 "de:08116:4266",
	STADTHALLE_MURRHARDT:                                   "de:08119:5294",
	STADTHALLE_VAIHINGEN_E:                                 "de:08118:5733",
	STADTHALLE_DONZDORF:                                    "de:08117:7622",
	STADTHALLE_GOPPINGEN:                                   "de:08117:9007",
	STADTKIRCHE_GEISLINGEN_STEIGE:                          "de:08117:1",
	STADTMITTE_WENDLINGEN_N:                                "de:08116:4258",
	STADTMITTE_WAIBLINGEN:                                  "de:08119:5391",
	STADTMITTE_STUTTGART:                                   "de:08111:6056",
	STADTMITTE_ASPERG:                                      "de:08118:6766",
	STADTMITTE_EISLINGEN_F:                                 "de:08117:1029",
	STADTMITTE_EBERSBACH_F:                                 "de:08117:8014",
	STADTMUHLE_BACKNANG:                                    "de:08119:3638",
	STADTPLATZ_WERNAU_N:                                    "de:08116:4228",
	STADTWERKE_WAIBLINGEN:                                  "de:08119:3629",
	STADTWERKE_LUDWIGSBURG:                                 "de:08118:5466",
	STADTWERKE_HERRENBERG:                                  "de:08115:7009",
	STADTWERKE_GOPPINGEN:                                   "de:08117:1034",
	STADTZENTRUM_FREIBERG_N:                                "de:08118:3521",
	STAFFEL_HOHENHASLACH:                                   "de:08118:5709",
	STAFFLENBERGSTRASSE_STUTTGART:                          "de:08111:123",
	STAHLACKERWEG_NECKARHALDE:                              "de:08116:4120",
	STAIGACKER_BACKNANG:                                    "de:08119:6687",
	STAMMHEIM_STUTTGART:                                    "de:08111:100",
	STAMMHEIM_FRIEDHOF_STUTTGART:                           "de:08111:206",
	STAMMHEIM_SPORTHALLE_STUTTGART:                         "de:08111:196",
	STANGEN_ECHTERDINGEN:                                   "de:08116:7186",
	STANGENBACH_WUSTENROT:                                  "de:08125:981",
	STANGENBACH_HAUS_WALDESRUH_WUSTENROT:                   "de:08125:32453",
	STARENWEG_FREIBERG_N:                                   "de:08118:3559",
	STATTMANNSTRASSE_OBERBOIHINGEN:                         "de:08116:4300",
	STAUCHLE_HOLZGERLINGEN:                                 "de:08115:4881",
	STAUFENECKSCHULE_SALACH:                                "de:08117:9862",
	STAUFENSTRASSE_ERKENBRECHTSWEILER:                      "de:08116:6945",
	STAUFENSTRASSE_ALBERSHAUSEN:                            "de:08117:2015",
	STAUFERPARK_GOPPINGEN:                                  "de:08117:9016",
	STAUFERSCHULE_WAIBLINGEN:                               "de:08119:4874",
	STAUFERSCHULE_WASCHENBEUREN:                            "de:08117:9911",
	STAUFERSCHULE_LORCH:                                    "de:08136:7049",
	STAUFERSTRASSE_ALTBACH:                                 "de:08116:3816",
	STAUFERSTRASSE_WAIBLINGEN:                              "de:08119:5158",
	STAUFERSTRASSE_SCHMIDEN:                                "de:08119:5425",
	STAUFERWEG_MAITIS:                                      "de:08117:1226",
	STAUFFENBERGSTRASSE_OBERESSLINGEN:                      "de:08116:4023",
	STAUFFENBERGSTRASSE_KORNWESTHEIM:                       "de:08118:5442",
	STECKFELD_STUTTGART:                                    "de:08111:2041",
	STEIGE_ERDMANNHAUSEN:                                   "de:08118:3598",
	STEIGE_HOHENHASLACH:                                    "de:08118:5710",
	STEIGE_NECKARREMS:                                      "de:08118:6853",
	STEIGHOF_ABZWEIG_TREFFELHAUSEN_LAUTERSTEIN:             "de:08117:315",
	STEIGSTRASSE_FELLBACH:                                  "de:08119:2505",
	STEINACHBRUCKE_NURTINGEN:                               "de:08116:4451",
	STEINACKER_STUTTGART:                                   "de:08111:2547",
	STEINACKERSTRASSE_REICHENBACH_F:                        "de:08116:3904",
	STEINBACHBRUCKE_LOCHGAU:                                "de:08118:7564",
	STEINBASS_SCHONAICH:                                    "de:08115:4811",
	STEINBEISE_WELZHEIM:                                    "de:08119:5130",
	STEINBEISSTRASSE_FELLBACH:                              "de:08119:2641",
	STEINBEISSTRASSE_ZELL_ESSLINGEN:                        "de:08116:3850",
	STEINBEISSTRASSE_BOBLINGEN:                             "de:08115:4702",
	STEINBEISSTRASSE_LUDWIGSBURG:                           "de:08118:5491",
	STEINBEISSTRASSE_VAIHINGEN_E:                           "de:08118:5875",
	STEINBEISSTRASSE_MARKGRONINGEN:                         "de:08118:6978",
	STEINBEISSTRASSE_SCHORNDORF:                            "de:08119:6993",
	STEINBEISSTRASSE_EISLINGEN_F:                           "de:08117:1009",
	STEINBEISWEG_NURTINGEN:                                 "de:08116:2987",
	STEINBERG_MURRHARDT:                                    "de:08119:5062",
	STEINBERG_FRITZHOF_MURRHARDT:                           "de:08119:5079",
	STEINBERG_HALDE_MURRHARDT:                              "de:08119:6966",
	STEINBRUCH_SCHOCKINGEN:                                 "de:08118:3031",
	STEINBRUCK_WELZHEIM:                                    "de:08119:5053",
	STEINBUTTSTRASSE_STUTTGART:                             "de:08111:2482",
	STEINEN_NELLINGEN:                                      "de:08116:2937",
	STEINENKIRCH_STEINENKIRCH:                              "de:08117:206",
	STEINGARTENSCHULE_DONZDORF:                             "de:08117:310",
	STEINGAUSTRASSE_KIRCHHEIM_T:                            "de:08116:4384",
	STEINGRUBE_WAIBLINGEN:                                  "de:08119:3722",
	STEINHALDE_OBERESSLINGEN:                               "de:08116:4002",
	STEINHALDENFELD_STUTTGART:                              "de:08111:2490",
	STEINHEIMER_STRASSE_POPPENWEILER:                       "de:08118:5557",
	STEINHEIMER_STRASSE_BIETIGHEIM:                         "de:08118:7407",
	STEINIGE_HALDE_SPARWIESEN:                              "de:08117:2204",
	STEINREINACH_KORB:                                      "de:08119:5416",
	STEINSTRASSE_HOLZHEIM_GOPPINGEN:                        "de:08117:3025",
	STEINWERK_VAIHINGEN_E:                                  "de:08118:6981",
	STELLE_STUTTGART:                                       "de:08111:6127",
	STERN_AICH:                                             "de:08116:2007",
	STERN_UNTERENSINGEN:                                    "de:08116:4290",
	STERNBERGSTRASSE_FRICKENHAUSEN:                         "de:08116:5878",
	STERNHAULE_STUTTGART:                                   "de:08111:6351",
	STERNKREUZUNG_GOPPINGEN:                                "de:08117:2001",
	STERNPLATZ_ZOH_GEISLINGEN_STEIGE:                       "de:08117:8",
	STETTENBEINSTEIN_ENDERSBACH:                            "de:08119:1701",
	STETTENER_STRASSE_ENDERSBACH:                           "de:08119:5166",
	STETTINER_RING_BACKNANG:                                "de:08119:5334",
	STETTINER_STRASSE_NURTINGEN:                            "de:08116:2983",
	STETTINER_STRASSE_BOBLINGEN:                            "de:08115:6744",
	STETTINER_STRASSE_SCHWIEBERDINGEN:                      "de:08118:6958",
	STIEGELPLATZ_MUNCHINGEN:                                "de:08118:4523",
	STIEGELWIESE_SUSSEN:                                    "de:08117:7612",
	STIFTSGYMNASIUM_SINDELFINGEN:                           "de:08115:5857",
	STOCKACH_STUTTGART:                                     "de:08111:26",
	STOCKACHSCHULE_WINNENDEN:                               "de:08119:5286",
	STOCKENHOF_OSCHELBRONN_BERGLEN:                         "de:08119:5271",
	STORCHEN_GOPPINGEN:                                     "de:08117:1059",
	STRASSENACKER_EGLOSHEIM:                                "de:08118:5545",
	STRASSENMEISTEREI_KIRCHHEIM_T:                          "de:08116:4270",
	STRASSENMEISTEREI_HERRENBERG:                           "de:08115:4855",
	STRASSLE_MARKGRONINGEN:                                 "de:08118:3315",
	STRAUSSSTAFFEL_STUTTGART:                               "de:08111:2311",
	STREICH_VORDERWEISSBUCH:                                "de:08119:5263",
	STRESEMANNSTRASSE_BACKNANG:                             "de:08119:3658",
	STROHGAUSTRASSE_LEONBERG:                               "de:08115:6054",
	STROHHOF_ALTHUTTE:                                      "de:08119:4997",
	STROMBERGGYMNASIUM_VAIHINGEN_E:                         "de:08118:6836",
	STRUBELMUHLE_ALFDORF:                                   "de:08119:5098",
	STRUDELBACHHALLE_WEISSACH_BB:                           "de:08115:5918",
	STRUMPFELBACH_STRUMPFELBACH_BACKNANG:                   "de:08119:5983",
	STRUTSTRASSE_SCHLICHTEN:                                "de:08119:6502",
	STUIFENSTRASSE_PLOCHINGEN:                              "de:08116:3940",
	STUIFENSTRASSE_MAITIS:                                  "de:08117:1225",
	STUMPENHOF_PLOCHINGEN:                                  "de:08116:4233",
	STUTTGARTER_PLATZ_FELLBACH:                             "de:08119:6042",
	STUTTGARTER_STRASSE_BOBLINGEN:                          "de:08115:3110",
	STUTTGARTER_STRASSE_FREIBERG_N:                         "de:08118:3524",
	STUTTGARTER_STRASSE_PLIENSAUVORSTADT:                   "de:08116:4035",
	STUTTGARTER_STRASSE_ALTDORF_ES:                         "de:08116:4489",
	STUTTGARTER_STRASSE_OSCHELBRONN_GAUFELDEN:              "de:08115:4746",
	STUTTGARTER_STRASSE_WEILER_ZUM_STEIN:                   "de:08119:5315",
	STUTTGARTER_STRASSE_VAIHINGEN_E:                        "de:08118:5865",
	STUTTGARTER_STRASSE_SCHORNDORF:                         "de:08119:5974",
	STUTTGARTER_STRASSE_REICHENBACH_F:                      "de:08116:6495",
	STUTTGARTER_STRASSE_ASPERG:                             "de:08118:6948",
	STUTTGARTER_STRASSEEICHHOLZ_MAICHINGEN:                 "de:08115:4186",
	STUTTGARTER_WEG_MARKGRONINGEN:                          "de:08118:3316",
	STUTTGARTERMARIENSTR_LORCH:                             "de:08136:7021",
	STUTTGARTERWESTSTRASSE_GRUNBACH:                        "de:08119:3715",
	SUD_HASLACH:                                            "de:08415:29041",
	SUD_RENNINGEN:                                          "de:08115:3195",
	SUD_KIRCHHEIM_T:                                        "de:08116:5777",
	SUDBAHNHOF_BOBLINGEN:                                   "de:08115:6742",
	SUDETENSTRASSE_JEBENHAUSEN:                             "de:08117:2151",
	SUDETENSTRASSE_DONZDORF:                                "de:08117:7614",
	SUDHEIMER_PLATZ_STUTTGART:                              "de:08111:13",
	SUDSTRASSE_BISSINGEN_LB:                                "de:08118:7424",
	SULZBACH_M_SULZBACH_M:                                  "de:08119:4952",
	SULZBACHER_STRASSE_SPIEGELBERG:                         "de:08119:4958",
	SULZBACHTAL_SCHONAICH:                                  "de:08115:4813",
	SULZGRIESER_STEIGE_ESSLINGEN_N:                         "de:08116:4109",
	SUSSEN_SUSSEN:                                          "de:08117:157",
	SUTTNERSTRASSE_STUTTGART:                               "de:08111:288",
	SUWAG_PLEIDELSHEIM:                                     "de:08118:5572",
	SYNERGIEPARK_STUTTGART:                                 "de:08111:2598",
	TACHENBERGSTRASSE_KORNTAL:                              "de:08118:2628",
	TACHENHAUSERSTRASSE_OBERBOIHINGEN:                      "de:08116:3825",
	TAFERHALDE_UNTERWEISSACH:                               "de:08119:6509",
	TAGFALTERWEG_STUTTGART:                                 "de:08111:204",
	TAILFINGER_STRASSE_OSCHELBRONN_GAUFELDEN:               "de:08115:6892",
	TALLANDHAUSSTRASSE_STUTTGART:                           "de:08111:6078",
	TALOSTENDSTRASSE_STUTTGART:                             "de:08111:306",
	TALACKERSTRASSE_WARMBRONN:                              "de:08115:4549",
	TALACKERSTRASSE_WENDLINGEN_N:                           "de:08116:3833",
	TALAUE_BIRKMANNSWEILER:                                 "de:08119:5203",
	TALBACHBRUCKE_HOCHDORF_ES:                              "de:08116:4220",
	TALDORFER_STRASSE_STUTTGART:                            "de:08111:2574",
	TALE_KAISERSBACH:                                       "de:08119:6916",
	TALHOFE_ENZWEIHINGEN:                                   "de:08118:5756",
	TALSTRASSE_HOLZGERLINGEN:                               "de:08115:3220",
	TALSTRASSE_EHNINGEN:                                    "de:08115:3250",
	TALSTRASSE_HOF_UND_LEMBACH:                             "de:08118:3335",
	TALSTRASSE_HOPFIGHEIM:                                  "de:08118:3591",
	TALSTRASSE_WAIBLINGEN:                                  "de:08119:3651",
	TALSTRASSE_ST_BERNHARDT:                                "de:08116:4137",
	TALSTRASSE_RAIDWANGEN:                                  "de:08116:4462",
	TALSTRASSE_BOBLINGEN:                                   "de:08115:4619",
	TALSTRASSE_LUDWIGSBURG:                                 "de:08118:5504",
	TALSTRASSE_SCHORNDORF:                                  "de:08119:5962",
	TALSTRASSE_KORNWESTHEIM:                                "de:08118:6834",
	TALSTRASSE_EISLINGEN_F:                                 "de:08117:1020",
	TALSTRASSE_LORCH:                                       "de:08136:7039",
	TAMM_TAMM:                                              "de:08118:7404",
	TAMMER_STRASSE_BISSINGEN_LB:                            "de:08118:3386",
	TAMMER_STRASSE_EGLOSHEIM:                               "de:08118:5512",
	TANNBACHBRUCKE_MIEDELSBACH:                             "de:08119:6512",
	TANNENBERG_BOBLINGEN:                                   "de:08115:4689",
	TANNENBERGSTRASSE_KIRCHHEIM_T:                          "de:08116:4331",
	TANNENHOFE_SCHLIERBACH:                                 "de:08117:7800",
	TANNHOF_PFAHLBRONN:                                     "de:08119:6979",
	TAPACHSTRASSE_STUTTGART:                                "de:08111:6291",
	TAUBENHOF_WELZHEIM:                                     "de:08119:7524",
	TAUNUSSTRASSE_BOBLINGEN:                                "de:08115:4673",
	TAUSGYMNASIUM_BACKNANG:                                 "de:08119:5339",
	TAUSSCHULE_BACKNANG:                                    "de:08119:3731",
	TECHNISCHE_AKADEMIE_NELLINGEN:                          "de:08116:2973",
	TECHNOLOGIEPARK_BOBLINGEN:                              "de:08115:3140",
	TECHNOLOGIEPARK_MARBACH_N:                              "de:08118:3531",
	TECKHALLE_OWEN:                                         "de:08116:4279",
	TECKSTRASSE_KIRCHHEIM_T:                                "de:08116:4319",
	TEINACHER_STRASSE_EGLOSHEIM:                            "de:08118:3421",
	TENNISPLATZ_STUTTGART:                                  "de:08111:2344",
	TENNISPLATZ_LUDWIGSBURG:                                "de:08118:5520",
	TENNISPLATZE_MOGLINGEN:                                 "de:08118:3434",
	TERLANERSTRASSE_STUTTGART:                              "de:08111:2531",
	TESSINER_STRASSE_DARMSHEIM:                             "de:08115:4640",
	TEUCHELWEG_MAICHINGEN:                                  "de:08115:3185",
	TEXTILZENTRUM_SINDELFINGEN:                             "de:08115:4654",
	THHEUSSSTRASSE_HOCHDORF_EBERDINGEN:                     "de:08118:6950",
	THSTORMSTRSIEDLUNG_STRASSDORF:                          "de:08136:3513",
	THALESPLATZ_DITZINGEN:                                  "de:08118:3035",
	THEATER_IM_BAHNHOF_RECHBERGHAUSEN:                      "de:08117:4015",
	THEOLORCHWERKSTATTEN_BIETIGHEIM:                        "de:08118:3395",
	THEOLORCHWERKSTATTEN_LUDWIGSBURG:                       "de:08118:7438",
	THEODORHEUSSGYMNASIUM_OBERESSLINGEN:                    "de:08116:4197",
	THEODORHEUSSPLATZ_SCHELMENHOLZ:                         "de:08119:6662",
	THEODORHEUSSSTRASSE_HILDRIZHAUSEN:                      "de:08115:4796",
	THEODORHEUSSSTRASSE_OBERRIEXINGEN:                      "de:08118:6513",
	THERMALBAD_HEXENBUCKEL_BOBLINGEN:                       "de:08115:3135",
	THERMALBAD_STUTTGARTER_STR_BOBLINGEN:                   "de:08115:3102",
	THINGSTRASSE_STUTTGART:                                 "de:08111:2611",
	THOMASHARDTER_STRASSE_HEGENLOHE:                        "de:08116:3812",
	TIEFENBACHSCHULE_STUTTGART:                             "de:08111:2523",
	TIEFENMAD_KIRCHENKIRNBERG:                              "de:08119:5083",
	TIEFGARAGE_HOHENGEHREN:                                 "de:08116:4079",
	TIGERSTRASSE_STUTTGART:                                 "de:08111:5852",
	TIROLER_STRASSE_ELTINGEN:                               "de:08115:2353",
	TISCHARDTER_STRASSE_NURTINGEN:                          "de:08116:2990",
	TISCHARDTER_STRASSE_KOHLBERG:                           "de:08116:3820",
	TOBELKREUZUNG_SUSSEN:                                   "de:08117:7611",
	TOBELWASEN_WEILHEIM_T:                                  "de:08116:4306",
	TORFGRUBE_SCHOPFLOCH_LENNINGEN:                         "de:08116:3824",
	TORGAUER_WEG_HERRENBERG:                                "de:08115:7069",
	TORLENSWEG_LEONBERG:                                    "de:08115:2355",
	TORSTRASSE_NEUSTADT:                                    "de:08119:7597",
	TORSTRASSE_BAHNHOF_DETTENHAUSEN:                        "de:08416:26357",
	TRANKE_STUTTGART:                                       "de:08111:3309",
	TRAUBE_NEUHAUSEN_F:                                     "de:08116:2908",
	TRAUBE_STRUMPFELBACH_WEINSTADT:                         "de:08119:5176",
	TRAUBE_LIPPOLDSWEILER:                                  "de:08119:5382",
	TRAUBENSTRASSE_NECKARREMS:                              "de:08118:6854",
	TRAUZENBACH_GRAB:                                       "de:08119:5142",
	TREIBERSTRASSE_STUTTGART:                               "de:08111:2438",
	TRIPSDRILL_ERLEBNISPARK_CLEEBRONN:                      "de:08125:5444",
	TROPPEL_WEIL_IM_SCHONBUCH:                              "de:08115:4192",
	TUBINGER_ALLEE_SINDELFINGEN:                            "de:08115:3245",
	TUBINGER_STRASSE_ECHTERDINGEN:                          "de:08116:3113",
	TUBINGER_STRASSE_BOBLINGEN:                             "de:08115:3128",
	TUBINGER_STRASSE_DECKENPFRONN:                          "de:08115:4839",
	TUBINGER_STRASSE_NECKARTAILFINGEN:                      "de:08116:5820",
	TUBINGER_STRASSE_DETTENHAUSEN:                          "de:08416:6352",
	TULPENSTRASSE_SINDELFINGEN:                             "de:08115:4647",
	TUNNEL_OSTPORTAL_STUTTGART:                             "de:08111:2406",
	TUNNELSTRASSE_STUTTGART:                                "de:08111:371",
	TURKHEIM_TURKHEIM:                                      "de:08117:258",
	TURMUHRENMUSEUM_MAINHARDT:                              "de:08127:28266",
	TURNACKERSTRASSE_BERNHAUSEN:                            "de:08116:2089",
	TURNERSCHAFT_STADION_GOPPINGEN:                         "de:08117:4906",
	TURNHALLE_HIRSCHLANDEN:                                 "de:08118:3004",
	TURNHALLE_ERDMANNHAUSEN:                                "de:08118:3516",
	TURNHALLE_STEINENBERG:                                  "de:08119:5122",
	TURNHALLE_REICHENBACH_IM_TALE_DEGGINGEN:                "de:08117:50",
	TURNHALLE_UHINGEN:                                      "de:08117:2029",
	TUV_BERNHAUSEN:                                         "de:08116:2072",
	TUXER_WEG_MAUBACH:                                      "de:08119:6845",
	UDITORIUM_UHINGEN:                                      "de:08117:2011",
	UFERSTRASSE_HOHENECK:                                   "de:08118:5527",
	UFFKIRCHHOF_STUTTGART:                                  "de:08111:32",
	UHINGEN_UHINGEN:                                        "de:08117:152",
	UHLANDGYMNASIUM_KIRCHHEIM_T:                            "de:08116:4356",
	UHLANDSTRASSE_LUDWIGSBURG:                              "de:08118:3416",
	UHLANDSTRASSE_HEININGEN_GP:                             "de:08117:3309",
	UHLANDSTRASSE_FAURNDAU:                                 "de:08117:5002",
	UHLANDSTRASSE_OBERWALDEN:                               "de:08117:5013",
	UHLBACH_STUTTGART:                                      "de:08111:2530",
	UHLBERG_PLATTENHARDT:                                   "de:08116:6728",
	ULMENSTRASSE_STUTTGART:                                 "de:08111:376",
	ULMENSTRASSE_MAICHINGEN:                                "de:08115:7444",
	ULMER_STRASSE_RAMTEL:                                   "de:08115:2368",
	ULMER_STRASSE_PLOCHINGEN:                               "de:08116:4225",
	ULMER_STRASSE_REICHENBACH_F:                            "de:08116:6498",
	ULMERHEININGER_STR_GOPPINGEN:                           "de:08117:1005",
	ULRICHSTRASSE_BESIGHEIM:                                "de:08118:3510",
	ULRICHSTRASSE_BEUTELSBACH:                              "de:08119:5398",
	ULRICHWEG_DOFFINGEN:                                    "de:08115:6754",
	UMGELTERWEG_STUTTGART:                                  "de:08111:2417",
	UMSPANNWERK_SUSSEN:                                     "de:08117:7618",
	UNGEHEUERHOF_BACKNANG:                                  "de:08119:4975",
	UNGEHEUERHOF_ABZWEIG_BACKNANG:                          "de:08119:4973",
	UNIVERSITAT_STUTTGART:                                  "de:08111:6008",
	UNIVERSITAT_SCHLEIFE_STUTTGART:                         "de:08111:6021",
	UNIVERSITAT_HOHENHEIM_STUTTGART:                        "de:08111:2555",
	UNTER_DEN_ARKADEN_ROMMELSHAUSEN:                        "de:08119:4939",
	UNTERAICHEN_LEINFELDEN:                                 "de:08116:174",
	UNTERBERKEN_HERRENBACHSTRASSE_OBERBERKEN:               "de:08119:4948",
	UNTERBOHRINGEN_UNTERBOHRINGEN:                          "de:08117:92",
	UNTERDORF_WEILER_OB_HELFENSTN:                          "de:08117:108",
	UNTERDORFSTRASSE_KONGEN:                                "de:08116:4086",
	UNTERDRACKENSTEIN_KIRCHE_DRACKENSTEIN:                  "de:08117:3503",
	UNTERE_BRUCK_BORTLINGEN:                                "de:08117:4031",
	UNTERE_BUHLSTRASSE_ALFDORF:                             "de:08119:5100",
	UNTERE_HALDE_STETTEN_LEINFELDENECHT:                    "de:08116:2121",
	UNTERE_HALDE_WEIL_IM_SCHONBUCH:                         "de:08115:7037",
	UNTERE_MARBACHER_STRASSE_LUDWIGSBURG:                   "de:08118:5484",
	UNTERE_MAYENNER_STRASSE_WAIBLINGEN:                     "de:08119:3691",
	UNTERE_SCHLOSSSTRASSE_ALFDORF:                          "de:08119:5103",
	UNTERE_STRASSE_ERKENBRECHTSWEILER:                      "de:08116:4428",
	UNTERENSINGER_STRASSE_ZIZISHAUSEN:                      "de:08116:4456",
	UNTERLENNINGEN_UNTERLENNINGEN:                          "de:08116:5781",
	UNTERM_GREUT_SCHOPFLOCH_LENNINGEN:                      "de:08116:3823",
	UNTERNEUSTETTEN_KIRCHENKIRNBERG:                        "de:08119:5085",
	UNTERRIEDEN_GYMNASIUM_MAICHINGEN:                       "de:08115:5856",
	UNTERSCHONTAL_BACKNANG:                                 "de:08119:5353",
	UNTERTURKHEIM_STUTTGART:                                "de:08111:6085",
	UNTERTURKHEIM_ALTER_FRIEDHOF_STUTTGART:                 "de:08111:2512",
	UNTERTURKHEIM_FRIEDHOF_STUTTGART:                       "de:08111:2513",
	UNTERTURKHEIM_KELTERPLATZ_STUTTGART:                    "de:08111:2509",
	UNTERTURKHEIMER_STRASSE_FELLBACH:                       "de:08119:2503",
	URACHER_WEG_NEUFFEN:                                    "de:08116:7082",
	URACHSTRASSE_STUTTGART:                                 "de:08111:2058",
	URBACH_URBACH:                                          "de:08119:5789",
	URBAN_HARBOR_LUDWIGSBURG:                               "de:08118:7645",
	URBANSTRASSE_BEUTELSBACH:                               "de:08119:3762",
	URSENWANG_ESCHENSTR_GOPPINGEN:                          "de:08117:3015",
	URSENWANG_GRUNDSCHULE_GOPPINGEN:                        "de:08117:9833",
	URSENWANG_KIEFERNSTEIGE_GOPPINGEN:                      "de:08117:3014",
	URSENWANG_LINDENPLATZ_GOPPINGEN:                        "de:08117:3013",
	VAIHINGEN_STUTTGART:                                    "de:08111:6002",
	VAIHINGEN_E_VAIHINGEN_E:                                "de:08118:5879",
	VAIHINGEN_ALTER_FRIEDHOF_STUTTGART:                     "de:08111:2602",
	VAIHINGEN_BF_OST_STUTTGART:                             "de:08111:6006",
	VAIHINGEN_RATHAUS_STUTTGART:                            "de:08111:2595",
	VAIHINGEN_SCHILLERPLATZ_STUTTGART:                      "de:08111:6004",
	VAIHINGEN_VIADUKT_STUTTGART:                            "de:08111:6",
	VAIHINGER_ECK_VAIHINGEN_E:                              "de:08118:5745",
	VAIHINGER_STRASSE_STUTTGART:                            "de:08111:170",
	VAIHINGER_STRASSE_SERSHEIM:                             "de:08118:5651",
	VAIHINGER_STRASSE_NUSSDORF:                             "de:08118:6962",
	VARNBULERSTRASSE_HOFINGEN:                              "de:08115:6815",
	VEILCHENSTRASSE_HOLZHAUSEN_UHINGEN:                     "de:08117:2032",
	VEILCHENWEG_LEINFELDEN:                                 "de:08116:2127",
	VENDELAUSTRASSE_NURTINGEN:                              "de:08116:2985",
	VERDISTRASSE_NECKARHAUSEN:                              "de:08116:4479",
	VESOULER_STRASSE_GERLINGEN:                             "de:08118:6905",
	VIADUKT_ENDERSBACH:                                     "de:08119:5168",
	VIEHTRIEB_MALMSHEIM:                                    "de:08115:7098",
	VIEHWEIDE_SINDELFINGEN:                                 "de:08115:4665",
	VIER_PAPPELN_SCHORNDORF:                                "de:08119:4498",
	VIERGIEBELWEG_STUTTGART:                                "de:08111:2242",
	VIKTORKOCHLWEG_STUTTGART:                               "de:08111:2191",
	VILLASTRASSE_STUTTGART:                                 "de:08111:372",
	VILLENEUVESTRASSE_KORNWESTHEIM:                         "de:08118:5438",
	VINZENZ_THERME_BAD_DITZENBACH:                          "de:08117:55",
	VOGELSANG_STUTTGART:                                    "de:08111:217",
	VOGELSANG_NAGOLD:                                       "de:08235:10070",
	VOGELSANGSTRASSE_DENKENDORF:                            "de:08116:5010",
	VOGGENHOF_ABZWEIG_ALTHUTTE:                             "de:08119:5027",
	VOGTHESSSCHULE_HERRENBERG:                              "de:08115:4851",
	VOLKERBOHRINGERWEG_WIFLINGSHAUSEN:                      "de:08116:4139",
	VOLKSBANK_DETTINGEN_T:                                  "de:08116:4274",
	VOLKSBANK_ASPERG:                                       "de:08118:5479",
	VOLKSBANK_AUENDORF:                                     "de:08117:3226",
	VOLKSBANK_ADELBERG:                                     "de:08117:7656",
	VOLKSHOCHSCHULE_ESSLINGEN_N:                            "de:08116:4009",
	VOLLMWEGBIRKENWEG_NAGOLD:                               "de:08235:10343",
	VOLLZUGSANSTALT_HEIMSHEIM:                              "de:08236:3508",
	VORDERBUCHELBERG_SPIEGELBERG:                           "de:08119:3730",
	VORDERBUCHELBERG_ABZWEIG_WUSTENROT:                     "de:08125:2024",
	VORDERE_KARLSTRASSE_GOPPINGEN:                          "de:08117:1060",
	VORDERE_SCHMIEDGASSE_SCHWABISCH_GMUND:                  "de:08136:3065",
	VORDERER_BERG_JEBENHAUSEN:                              "de:08117:2152",
	VORDERHUNDSBERG_WELZHEIM:                               "de:08119:7528",
	VORDERSTEINENBERG_VORDERSTEINENBERG:                    "de:08119:5095",
	VORDERWEISSBUCH_VORDERWEISSBUCH:                        "de:08119:5264",
	VORDERWESTERMURR_MURRHARDT:                             "de:08119:5299",
	VORSTADT_NURTINGEN:                                     "de:08116:5863",
	VOSSSTRASSE_GEISLINGEN_STEIGE:                          "de:08117:211",
	W__W_KORNWESTHEIM:                                      "de:08118:3361",
	WAAGE_ERDMANNHAUSEN:                                    "de:08118:3517",
	WAAGE_AFFALTERBACH:                                     "de:08118:3602",
	WAGENBURGSTRASSE_STUTTGART:                             "de:08111:2202",
	WAGNERSTRASSE_ROMMELSHAUSEN:                            "de:08119:3736",
	WAGON_NAGOLD:                                           "de:08235:9332",
	WAGRAINACKER_STUTTGART:                                 "de:08111:274",
	WAHLENHEIM_VORDERSTEINENBERG:                           "de:08119:6797",
	WAIBLINGEN_WAIBLINGEN:                                  "de:08119:7604",
	WAIBLINGER_STRASSE_BEINSTEIN:                           "de:08119:5162",
	WAIBLINGERSCHORNDORFER_STR_FELLBACH:                    "de:08119:5434",
	WALA_BAD_BOLL:                                          "de:08117:2119",
	WALD_KIRCHHEIM_T:                                       "de:08116:5768",
	WALDACKERWEG_WIFLINGSHAUSEN:                            "de:08116:4141",
	WALDALLEE_HOCHBERG:                                     "de:08118:6824",
	WALDAU_STUTTGART:                                       "de:08111:130",
	WALDBURGSTRASSE_STUTTGART:                              "de:08111:2",
	WALDBURGSTRASSE_BOBLINGEN:                              "de:08115:4675",
	WALDCAFE_WANNENHOF_GOPPINGEN:                           "de:08117:1302",
	WALDECK_STUTTGART:                                      "de:08111:6010",
	WALDECKSCHULE_JEBENHAUSEN:                              "de:08117:2155",
	WALDECKSTRASSE_ALBERSHAUSEN:                            "de:08117:2016",
	WALDECKSTRASSE_GOPPINGEN:                               "de:08117:9203",
	WALDENBRONN_KREUZUNG_HOHENKREUZ:                        "de:08116:4125",
	WALDENBUCHER_PLATZ_STUTTGART:                           "de:08111:6568",
	WALDENSERSTRASSE_PEROUSE:                               "de:08115:5915",
	WALDENWEILER_SECHSELBERG:                               "de:08119:4987",
	WALDFRIEDHOF_LEONBERG:                                  "de:08115:2336",
	WALDFRIEDHOF_MAICHINGEN:                                "de:08115:3211",
	WALDFRIEDHOF_KIRCHHEIM_T:                               "de:08116:4355",
	WALDFRIEDHOF_NURTINGEN:                                 "de:08116:4449",
	WALDFRIEDHOF_BOBLINGEN:                                 "de:08115:4715",
	WALDFRIEDHOF_HERRENBERG:                                "de:08115:4800",
	WALDFRIEDHOF_BACKNANG:                                  "de:08119:5322",
	WALDFRIEDHOF_STUTTGART:                                 "de:08111:6102",
	WALDFRIEDHOF_GERLINGEN:                                 "de:08118:6907",
	WALDFRIEDHOF_KOHLBERG:                                  "de:08116:6923",
	WALDHAUSEN_WALDHAUSEN_B_SCHORND_LORCH:                  "de:08136:7031",
	WALDHEIM_SINDELFINGEN:                                  "de:08115:2223",
	WALDHOF_BIETIGHEIM:                                     "de:08118:5692",
	WALDHORN_HOHENGEHREN:                                   "de:08116:4077",
	WALDHORN_EISLINGEN_F:                                   "de:08117:1031",
	WALDHORNSTRASSE_ECHTERDINGEN:                           "de:08116:7179",
	WALDLUST_NAGOLD:                                        "de:08235:10069",
	WALDORFSCHULE_BONLANDEN:                                "de:08116:2046",
	WALDORFSCHULE_BOBLINGEN:                                "de:08115:3109",
	WALDSCHULE_BISSINGEN_LB:                                "de:08118:5685",
	WALDSEE_FORNSBACH:                                      "de:08119:6921",
	WALDSIEDLUNG_GERLINGEN:                                 "de:08118:2318",
	WALDSTRASSE_PLATTENHARDT:                               "de:08116:2029",
	WALDSTRASSE_GROSSBOTTWAR:                               "de:08118:5605",
	WALHEIM_WALHEIM:                                        "de:08118:5785",
	WALKERSBACH_PLUDERHAUSEN:                               "de:08119:5765",
	WALKERSBACHER_TAL_LORCH:                                "de:08136:7035",
	WALLGRABEN_STUTTGART:                                   "de:08111:355",
	WALLSTRASSE_WINNENDEN:                                  "de:08119:5209",
	WALTERFLEXSTRASSE_OSSWEIL:                              "de:08118:6882",
	WALTERHELMESWEG_LEONBERG:                               "de:08115:2364",
	WALTERSBERG_MURRHARDT:                                  "de:08119:6690",
	WANDERPARKPLATZ_ROTER_STICH_HOSSLINSWART:               "de:08119:7459",
	WANDERWEG_MUSBERG:                                      "de:08116:7105",
	WANGEN_KELTER_STUTTGART:                                "de:08111:377",
	WANGEN_MARKTPLATZ_STUTTGART:                            "de:08111:225",
	WANGENER_STRASSE_RECHBERGHAUSEN:                        "de:08117:4038",
	WANGENER_STRASSE_GOPPINGEN:                             "de:08117:4908",
	WANGENERLANDHAUSSTRASSE_STUTTGART:                      "de:08111:80",
	WANKELSTRASSE_MALMSHEIM:                                "de:08115:5847",
	WANNENHOFE_AUFHAUSEN_GEISLINGEN:                        "de:08117:497",
	WANNENRAIN_WEIL:                                        "de:08116:4065",
	WARNENBERG_OBERBOIHINGEN:                               "de:08116:7957",
	WARTBERGSTRASSE_MOGLINGEN:                              "de:08118:3431",
	WARTEHALLE_JESINGEN:                                    "de:08116:4322",
	WARTEHALLE_GERSTETTEN:                                  "de:08135:1138",
	WARTH_OTLINGEN:                                         "de:08116:4344",
	WASCHERHOFSTRASSE_WASCHENBEUREN:                        "de:08117:1205",
	WASENFELD_ALLMERSBACH_IM_TAL:                           "de:08119:6874",
	WASENHOF_EISLINGEN_F:                                   "de:08117:1015",
	WASENSTRASSE_STUTTGART:                                 "de:08111:84",
	WASENSTRASSE_NECKARGRONINGEN:                           "de:08118:4889",
	WASENSTRASSE_URBACH:                                    "de:08119:5226",
	WASENSTRASSE_JEBENHAUSEN:                               "de:08117:2154",
	WASSERBERG_UHINGEN:                                     "de:08117:2202",
	WASSERBERGSTRASSE_GOPPINGEN:                            "de:08117:3005",
	WASSERBERGWEG_PLUDERHAUSEN:                             "de:08119:5239",
	WASSERSTUBENWEG_WAIBLINGEN:                             "de:08119:7090",
	WASSERTURM_BESIGHEIM:                                   "de:08118:5732",
	WASSERWERK_NECKARTAILFINGEN:                            "de:08116:6709",
	WEBEREI_REICHENBACH_F:                                  "de:08116:4223",
	WEBERSTRASSE_NURTINGEN:                                 "de:08116:5898",
	WEIDACH_STETTEN_LEINFELDENECHT:                         "de:08116:2123",
	WEIDACHTAL_STUTTGART:                                   "de:08111:6572",
	WEIDENWEG_BACKNANG:                                     "de:08119:5340",
	WEIHERSTRASSE_OBERESSLINGEN:                            "de:08116:4026",
	WEIL_WEIL:                                              "de:08116:4017",
	WEIL_DER_STADT_WEIL_DER_STADT:                          "de:08115:1303",
	WEILDORF_BONDORF:                                       "de:08115:4749",
	WEILEMER_PFAD_STUTTGART:                                "de:08111:2637",
	WEILENBERGER_HOF_UHINGEN:                               "de:08117:2203",
	WEILER_BB_SCHAFHAUSEN:                                  "de:08115:4583",
	WEILER_R_WEILER_SCHORNDORF:                             "de:08119:1704",
	WEILER_HELFENSTEIN_WEILER_OB_HELFENSTN:                 "de:08117:107",
	WEILER_LINDENHOF_GEISLINGEN_STEIGE:                     "de:08117:111",
	WEILERHAU_PLATTENHARDT:                                 "de:08116:2047",
	WEILHEIMER_STRASSE_ZELL_A:                              "de:08117:2116",
	WEILIMDORF_BF_STUTTGART:                                "de:08111:2270",
	WEILIMDORF_BF_UNTERFUHRUNG_STUTTGART:                   "de:08111:2624",
	WEILIMDORF_LOWENMARKT_STUTTGART:                        "de:08111:6149",
	WEILSTRASSE_PLIENSAUVORSTADT:                           "de:08116:2801",
	WEIMARSTRASSE_MARBACH_N:                                "de:08118:3595",
	WEIMARSTRASSE_LUDWIGSBURG:                              "de:08118:5542",
	WEINBERGTECKSTRASSE_REICHENBACH_F:                      "de:08116:3901",
	WEINBERGSTRASSE_HOCHDORF_ES:                            "de:08116:5871",
	WEINBERGSTRASSE_REICHENBACH_F:                          "de:08116:6713",
	WEINGARTENTEGELBERG_GEISLINGEN_STEIGE:                  "de:08117:125",
	WEINGARTENSTRASSE_MONCHBERG:                            "de:08115:4867",
	WEINGARTENSTRASSE_GAMMELSHAUSEN:                        "de:08117:3324",
	WEINGARTSTRASSE_DENKENDORF:                             "de:08116:5889",
	WEINSTEIGE_STUTTGART:                                   "de:08111:163",
	WEINSTEIGE_BEUTELSBACH:                                 "de:08119:3770",
	WEINSTEIGE_WEILHEIM_T:                                  "de:08116:3895",
	WEINSTRASSE_ASPERG:                                     "de:08118:3405",
	WEINSTRASSE_METTINGEN:                                  "de:08116:4012",
	WEISSACHER_STRASSE_HEIMERDINGEN:                        "de:08118:4530",
	WEISSACHSTRASSE_OBERWEISSACH:                           "de:08119:4978",
	WEISSENHOF_OST_LOCHGAU:                                 "de:08118:5693",
	WEISSER_STEIN_PLOCHINGEN:                               "de:08116:4074",
	WEIZENSTRASSE_KLEINGLATTBACH:                           "de:08118:3352",
	WELLARIUM_STEINHEIM_M:                                  "de:08118:3504",
	WELLINGEN_KIRCHLE_NOTZINGEN:                            "de:08116:4365",
	WELLINGEN_SCHLIERBACHER_STR_NOTZINGEN:                  "de:08116:4366",
	WELLINGER_STRASSE_NOTZINGEN:                            "de:08116:4217",
	WELTESTRASSE_EISLINGEN_F:                               "de:08117:1011",
	WELZHEIMER_STRASSE_SCHORNDORF:                          "de:08119:3708",
	WELZHEIMER_STRASSE_PFAHLBRONN:                          "de:08119:5091",
	WELZHEIMER_STRASSE_UNTERWEISSACH:                       "de:08119:5374",
	WELZHEIMER_STRASSE_GSCHWEND:                            "de:08136:6514",
	WENDELKONIG_WAIBLINGEN:                                 "de:08119:3644",
	WENDEPLATTE_NECKARWEIHINGEN:                            "de:08118:3428",
	WENDEPLATTE_HOCHDORF_REMSECK:                           "de:08118:5548",
	WENDEPLATTE_MAINHARDT:                                  "de:08127:28270",
	WENDLINGEN_N_WENDLINGEN_N:                              "de:08116:4257",
	WENGERTSTRASSE_SINDELFINGEN:                            "de:08115:3239",
	WERASTRASSE_GEISLINGEN_STEIGE:                          "de:08117:41",
	WERNAU_N_WERNAU_N:                                      "de:08116:4241",
	WERNER_UND_PFLEIDERER_STUTTGART:                        "de:08111:4402",
	WERNERSTRASSE_FELLBACH:                                 "de:08119:2640",
	WERNERSTRASSE_LUDWIGSBURG:                              "de:08118:7641",
	WERT_DEIZISAU:                                          "de:08116:4055",
	WERTSTOFFHOF_GOPPINGEN:                                 "de:08117:1061",
	WESERSTRASSE_STUTTGART:                                 "de:08111:279",
	WESPENWEG_STUTTGART:                                    "de:08111:197",
	WEST_REUDERN:                                           "de:08116:4379",
	WESTBAHNHOF_STUTTGART:                                  "de:08111:6249",
	WESTBAHNHOF_SCHLEIFE_STUTTGART:                         "de:08111:6250",
	WESTENDSTRASSE_BISSINGEN_LB:                            "de:08118:7428",
	WESTFALENSTRASSE_OSSWEIL:                               "de:08118:4892",
	WESTRANDALLEE_SCHONAICH:                                "de:08115:3254",
	WESTSTADTKIRCHE_LUDWIGSBURG:                            "de:08118:5541",
	WETTBACHPLATZ_SINDELFINGEN:                             "de:08115:3230",
	WETTE_LOCHGAU:                                          "de:08118:5704",
	WETTEMARKT_OSSWEIL:                                     "de:08118:5516",
	WETTERWARTE_SCHNITTLINGEN:                              "de:08117:319",
	WETZSTEINSTOLLEN_JUX:                                   "de:08119:7581",
	WHG_GOPPINGEN:                                          "de:08117:9895",
	WIDDERSTALL_MERKLINGEN:                                 "de:08425:2625",
	WIDDUMHOF_RUTESHEIM:                                    "de:08115:5937",
	WIEGEHALLE_PLEIDELSHEIM:                                "de:08118:5571",
	WIEGEHALLERIEDBACHSTRASSE_PLEIDELSHEIM:                 "de:08118:3470",
	WIELANDSHOHE_STUTTGART:                                 "de:08111:342",
	WIELANDSTRASSE_STUTTGART:                               "de:08111:215",
	WIELANDSTRASSEBF_HOLZHEIM_GOPPINGEN:                    "de:08117:3008",
	WIESENACKERSTRASSE_HEIMERDINGEN:                        "de:08118:4531",
	WIESENSTEIGER_STRASSE_GOSBACH:                          "de:08117:71",
	WIESENSTR_SCHORNDORF:                                   "de:08119:7576",
	WIESENSTRASSE_ROMMELSHAUSEN:                            "de:08119:3738",
	WIESENTAL_WASCHENBEUREN:                                "de:08117:7017",
	WIESENTALSTRASSE_SCHNAIT:                               "de:08119:6665",
	WIESENTALSTRASSE_HERTMANNSWEILER:                       "de:08119:7455",
	WIESLAUFWEG_SCHORNDORF:                                 "de:08119:5215",
	WILDENSTEINSTRASSE_STUTTGART:                           "de:08111:2424",
	WILDGANSWEG_STUTTGART:                                  "de:08111:6277",
	WILDPARADIES_TRIPSDRILL_CLEEBRONN:                      "de:08125:4714",
	WILHELMOLGASTRASSE_STUTTGART:                           "de:08111:2198",
	WILHELMFEINSTRASSE_LUDWIGSBURG:                         "de:08118:7644",
	WILHELMFOLLSTRASSE_STRUMPFELBACH_BACKNANG:              "de:08119:7452",
	WILHELMGEIGERPLATZ_STUTTGART:                           "de:08111:6180",
	WILHELMHASPELSTRASSE_SINDELFINGEN:                      "de:08115:4662",
	WILHELMKOPPSTRASSE_PEROUSE:                             "de:08115:7832",
	WILHELMMAIERSTRASSE_KONGEN:                             "de:08116:4288",
	WILHELMPFITZERSTRASSE_FELLBACH:                         "de:08119:2642",
	WILHELMRAABESTRASSE_WENDLINGEN_N:                       "de:08116:3829",
	WILHELMA_STUTTGART:                                     "de:08111:254",
	WILHELMSBAU_STUTTGART:                                  "de:08111:2061",
	WILHELMSHILFE_GOPPINGEN:                                "de:08117:1056",
	WILHELMSHILFE_BARTENBACH_GOPPINGEN:                     "de:08117:4014",
	WILHELMSHOHE_GEISLINGEN_STEIGE:                         "de:08117:101",
	WILHELMSPLATZ_ASPERG:                                   "de:08118:3406",
	WILHELMSPLATZ_HOCHDORF_REMSECK:                         "de:08118:6821",
	WILHELMSTRASSE_NEUHAUSEN_F:                             "de:08116:2909",
	WILHELMSTRASSE_WEIL_IM_SCHONBUCH:                       "de:08115:3143",
	WILHELMSTRASSE_KLEINSACHSENHEIM:                        "de:08118:3562",
	WILHELMSTRASSE_BISSINGEN_AN_DER_TECK_T:                 "de:08116:4413",
	WILHELMSTRASSE_GEISLINGEN_STEIGE:                       "de:08117:7",
	WILHELMSTRASSE_LORCH:                                   "de:08136:7042",
	WILLIBLEICHERSTRASSE_KIRCHHEIM_T:                       "de:08116:3954",
	WIMPFENER_STRASSE_STUTTGART:                            "de:08111:104",
	WINDHALMWEG_STUTTGART:                                  "de:08111:2554",
	WINKELS_GROSSSACHSENHEIM:                               "de:08118:7083",
	WINNENDEN_WINNENDEN:                                    "de:08119:7605",
	WINNENDER_STRASSE_WAIBLINGEN:                           "de:08119:5409",
	WINNENDER_STRASSE_SCHWAIKHEIM:                          "de:08119:5428",
	WINNENDER_STRASSE_LEUTENBACH:                           "de:08119:6911",
	WINTERBACH_WINTERBACH:                                  "de:08119:1705",
	WIRTSCHAFTSSCHULE_LAICHINGEN:                           "de:08425:2655",
	WITHAU_STUTTGART:                                       "de:08111:194",
	WITHAUWEG_STUTTGART:                                    "de:08111:6624",
	WITTENBERGER_STRASSE_HERRENBERG:                        "de:08115:7068",
	WITTUMHALLE_URBACH:                                     "de:08119:5227",
	WOHRTOURSDEPOT_WEISSACH_BB:                             "de:08115:5911",
	WOLFHIRTHSTRASSE_SIRNAU:                                "de:08116:6758",
	WOLFBUSCH_STUTTGART:                                    "de:08111:147",
	WOLFENMUHLE_SCHONAICH:                                  "de:08115:6716",
	WOLFSKLINGENWEG_WINNENDEN:                              "de:08119:5991",
	WOLFSOLDEN_AFFALTERBACH:                                "de:08118:3603",
	WOLFSTRASSE_ALTENHEIM_SINDELFINGEN:                     "de:08115:4649",
	WOLKBAD_GEISLINGEN_STEIGE:                              "de:08117:117",
	WOLLINSTRASSE_STUTTGART:                                "de:08111:94",
	WUNNEBAD_WINNENDEN:                                     "de:08119:3642",
	WUNNENSTEINSTRASSE_STUTTGART:                           "de:08111:2405",
	WURMBRUCKE_SCHAFHAUSEN:                                 "de:08115:4582",
	WURMBRUCKE_EHNINGEN:                                    "de:08115:4756",
	WURMLINGER_STRASSE_STUTTGART:                           "de:08111:2567",
	WURMSTRASSE_ALTDORF_BB:                                 "de:08115:4772",
	WURMTALSTRASSE_HAUSEN_WEIL_DER_STADT:                   "de:08115:4593",
	WURTTEMBERGSTRASSE_DITZINGEN:                           "de:08118:3030",
	WURTTEMBERGSTRASSE_WEIL:                                "de:08116:6841",
	WUSTENROT_LUDWIGSBURG:                                  "de:08118:5518",
	YHAUSER_GEISLINGEN_STEIGE:                              "de:08117:103",
	YITZHAKRABINSTRASSE_STUTTGART:                          "de:08111:2469",
	ZABERGAUSTRASSE_NEUHAUSEN_F:                            "de:08116:2939",
	ZABERGAUSTRASSE_STUTTGART:                              "de:08111:6266",
	ZABERSTRASSE_OBERJESINGEN:                              "de:08115:4842",
	ZACHERSMUHLE_ADELBERG:                                  "de:08117:7653",
	ZAHNNOPPERSTRASSE_STUTTGART:                            "de:08111:106",
	ZAHNKLINIK_BISSINGEN_LB:                                "de:08118:5689",
	ZAHNRADBAHNHOF_STUTTGART:                               "de:08111:6340",
	ZAHRINGER_STRASSE_LINDORF:                              "de:08116:3943",
	ZAUNACKER_ECHTERDINGEN:                                 "de:08116:2130",
	ZAZENHAUSEN_STUTTGART:                                  "de:08111:1200",
	ZAZENHAUSEN_KRONE_STUTTGART:                            "de:08111:4400",
	ZAZENHAUSEN_STEIGLE_STUTTGART:                          "de:08111:6269",
	ZAZENHAUSEN_VIADUKT_STUTTGART:                          "de:08111:6989",
	ZEHNTSCHEUER_SCHONAICH:                                 "de:08115:4816",
	ZEHNTSCHEUER_POPPENWEILER:                              "de:08118:5559",
	ZEIL_KLEINSACHSENHEIM:                                  "de:08118:5760",
	ZEISIGWEG_HERRENBERG:                                   "de:08115:7059",
	ZEISIGWEG_EISLINGEN_F:                                  "de:08117:1019",
	ZEISSSTRASSE_DITZINGEN:                                 "de:08118:3032",
	ZELL_ZELL_ESSLINGEN:                                    "de:08116:1803",
	ZELL_OPPENWEILER:                                       "de:08119:7486",
	ZELL_RECHBERGHAUSER_STR_BORTLINGEN:                     "de:08117:4030",
	ZELLER_STRASSE_OHMDEN:                                  "de:08116:4303",
	ZELLER_STRASSE_HATTENHOFEN:                             "de:08117:2050",
	ZELLERSTRASSE_STUTTGART:                                "de:08111:2413",
	ZENTRUM_SULZGRIES:                                      "de:08116:4113",
	ZENTRUM_ZIEGELFELD_HERRENBERG:                          "de:08115:3252",
	ZEPPELINKIRCHSTRASSE_DEIZISAU:                          "de:08116:6876",
	ZEPPELINSCHULE_FELLBACH:                                "de:08119:2498",
	ZEPPELINSTRASSE_OBERESSLINGEN:                          "de:08116:4044",
	ZEPPELINSTRASSE_KEMNAT:                                 "de:08116:5020",
	ZEPPELINSTRASSE_KORNWESTHEIM:                           "de:08118:5439",
	ZEPPELINSTRASSE_DEIZISAU:                               "de:08116:6878",
	ZIEGELBACHKLINGE_EISLINGEN_F:                           "de:08117:1018",
	ZIEGELBACHSTRASSE_EISLINGEN_F:                          "de:08117:1016",
	ZIEGELBERGWEG_UHINGEN:                                  "de:08117:2012",
	ZIEGELEISTRASSE_FRICKENHAUSEN:                          "de:08116:4403",
	ZIEGELEISTRASSE_MUNCHINGEN:                             "de:08118:4522",
	ZIEGELHOF_HOCHDORF_ES:                                  "de:08116:3906",
	ZIEGELHUTTE_KAISERSBACH:                                "de:08119:6943",
	ZIEGELSTRASSE_RECHBERGHAUSEN:                           "de:08117:4021",
	ZIEGELSTRASSE_SALACH:                                   "de:08117:6018",
	ZIEGELWASEN_KIRCHHEIM_T:                                "de:08116:4363",
	ZIEGELWASENSTRASSE_GRAFENBERG:                          "de:08415:28280",
	ZIEGERHOF_LENGLINGEN:                                   "de:08117:1207",
	ZIMMERBACHSTRASSE_OBERESSLINGEN:                        "de:08116:4027",
	ZIMMERER_PFAD_GROSSSACHSENHEIM:                         "de:08118:5890",
	ZIMMERSCHLAG_BOBLINGEN:                                 "de:08115:4713",
	ZIMMERSTRASSE_SINDELFINGEN:                             "de:08115:2224",
	ZINSHOLZ_RUIT:                                          "de:08116:2969",
	ZOB_SINDELFINGEN:                                       "de:08115:3200",
	ZOB_ESSLINGEN_N:                                        "de:08116:7801",
	ZOB_NAGOLD:                                             "de:08235:9340",
	ZOB_LAICHINGEN:                                         "de:08425:2638",
	ZOLLBERG_ZOLLBERG:                                      "de:08116:2803",
	ZOLLERNLICHTENSTEINSTRASSE_DITZINGEN:                   "de:08118:3017",
	ZOLLERNPLATZ_ZOLLBERG:                                  "de:08116:4177",
	ZOLLHAUSWEG_ZOLLBERG:                                   "de:08116:4178",
	ZUCKERBERG_STUTTGART:                                   "de:08111:2491",
	ZUFFENHAUSEN_STUTTGART:                                 "de:08111:6465",
	ZUFFENHAUSEN_FRIEDHOF_STUTTGART:                        "de:08111:109",
	ZUFFENHAUSEN_KELTERPLATZ_STUTTGART:                     "de:08111:6294",
	ZUFFENHAUSEN_PORSCHE_STUTTGART:                         "de:08111:6462",
	ZUFFENHAUSEN_RATHAUS_STUTTGART:                         "de:08111:108",
	ZUFFENHAUSEN_REIBEDANZ_STUTTGART:                       "de:08111:4401",
	ZUFFENHAUSER_HEIDE_STUTTGART:                           "de:08111:2478",
	ZUM_KLEINEN_FELDLE_STETTEN_I_R:                         "de:08119:4932",
	ZUR_ANHOHE_STUTTGART:                                   "de:08111:2568",
	ZURICHER_STRASSE_STUTTGART:                             "de:08111:2494",
	ZURICHER_STRASSE_BACKNANG:                              "de:08119:3663",
	ZWERCHWEG_HERRENBERG:                                   "de:08115:7202",
	ZWERENBERG_SULZBACH_M:                                  "de:08119:7511",
	ZWINGELH_GROSSASPACHER_STRASSE_KIRCHBERG_M:             "de:08119:3547",
	ZWINGELHAUSEN_ABZWEIG_GROSSASPACH:                      "de:08119:5365",
	ZWINGELHAUSEN_KIRCHBERGER_STR_KIRCHBERG_M:              "de:08119:3548",
	ZWINGELHAUSER_STRASSE_KIRCHBERG_M:                      "de:08119:6861",
}

func StationNameToGlobalId(name string) (string, error) {
	id, ok := stationNameIdMap[name]
	if !ok {
		return "", fmt.Errorf("station does not exist: %s", name)
	}

	return id, nil
}
