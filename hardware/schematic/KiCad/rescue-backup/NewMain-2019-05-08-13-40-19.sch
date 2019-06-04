EESchema Schematic File Version 4
EELAYER 29 0
EELAYER END
$Descr A4 11693 8268
encoding utf-8
Sheet 1 1
Title ""
Date ""
Rev ""
Comp ""
Comment1 ""
Comment2 ""
Comment3 ""
Comment4 ""
$EndDescr
$Comp
L transistors:BC807 Q1
U 1 1 5B08F43C
P 7350 1900
F 0 "Q1" H 7550 1975 50  0000 L CNN
F 1 "BC807" H 7550 1900 50  0000 L CNN
F 2 "TO_SOT_Packages_SMD:SOT-23" H 7550 1825 50  0001 L CIN
F 3 "" H 7350 1900 50  0001 L CNN
	1    7350 1900
	-1   0    0    1   
$EndComp
$Comp
L transistors:BC807 Q2
U 1 1 5B08F4F9
P 7500 2150
F 0 "Q2" H 7700 2225 50  0000 L CNN
F 1 "BC807" H 7700 2150 50  0000 L CNN
F 2 "TO_SOT_Packages_SMD:SOT-23" H 7700 2075 50  0001 L CIN
F 3 "" H 7500 2150 50  0001 L CNN
	1    7500 2150
	1    0    0    1   
$EndComp
$Comp
L Device:R R3
U 1 1 5B08F67D
P 7600 1700
F 0 "R3" V 7680 1700 50  0000 C CNN
F 1 "5R6" V 7600 1700 50  0000 C CNN
F 2 "Resistors_SMD:R_0603_HandSoldering" V 7530 1700 50  0001 C CNN
F 3 "" H 7600 1700 50  0001 C CNN
	1    7600 1700
	1    0    0    -1  
$EndComp
$Comp
L Device:R R1
U 1 1 5B08F6E8
P 7050 2150
F 0 "R1" V 7130 2150 50  0000 C CNN
F 1 "2K2" V 7050 2150 50  0000 C CNN
F 2 "Resistors_SMD:R_0603_HandSoldering" V 6980 2150 50  0001 C CNN
F 3 "" H 7050 2150 50  0001 C CNN
	1    7050 2150
	0    -1   -1   0   
$EndComp
$Comp
L Device:R R4
U 1 1 5B08F7FA
P 7600 2600
F 0 "R4" V 7680 2600 50  0000 C CNN
F 1 "680R" V 7600 2600 50  0000 C CNN
F 2 "Resistors_SMD:R_0603_HandSoldering" V 7530 2600 50  0001 C CNN
F 3 "" H 7600 2600 50  0001 C CNN
	1    7600 2600
	1    0    0    -1  
$EndComp
$Comp
L conn:Conn_02x03_Counter_Clockwise J1
U 1 1 5B08F9D8
P 2450 2050
F 0 "J1" H 2500 2250 50  0000 C CNN
F 1 "MDB" H 2500 1850 50  0000 C CNN
F 2 "Connectors_Molex:Molex_MegaFit_2x03x5.70mm_Angled" H 2450 2050 50  0001 C CNN
F 3 "" H 2450 2050 50  0001 C CNN
	1    2450 2050
	1    0    0    -1  
$EndComp
$Comp
L power:+24V #PWR01
U 1 1 5B08FA4F
P 2200 1900
F 0 "#PWR01" H 2200 1750 50  0001 C CNN
F 1 "+24V" H 2200 2040 50  0000 C CNN
F 2 "" H 2200 1900 50  0001 C CNN
F 3 "" H 2200 1900 50  0001 C CNN
	1    2200 1900
	1    0    0    -1  
$EndComp
Text GLabel 6000 2150 0    60   Input ~ 0
TTL-TX
Text GLabel 7750 2400 2    60   Input ~ 0
MDB-MT
Text GLabel 2850 2050 2    60   Input ~ 0
MDB-MT
Text GLabel 2850 2150 2    60   Input ~ 0
MDB-MR
Text GLabel 7100 3900 0    60   Input ~ 0
TTL-RX
$Comp
L Device:R R2
U 1 1 5B090100
P 7350 3900
F 0 "R2" V 7430 3900 50  0000 C CNN
F 1 "10K" V 7350 3900 50  0000 C CNN
F 2 "Resistors_SMD:R_0603_HandSoldering" V 7280 3900 50  0001 C CNN
F 3 "" H 7350 3900 50  0001 C CNN
	1    7350 3900
	0    1    1    0   
$EndComp
$Comp
L Device:R R5
U 1 1 5B090179
P 7550 3700
F 0 "R5" V 7630 3700 50  0000 C CNN
F 1 "680R" V 7550 3700 50  0000 C CNN
F 2 "Resistors_SMD:R_0603_HandSoldering" V 7480 3700 50  0001 C CNN
F 3 "" H 7550 3700 50  0001 C CNN
	1    7550 3700
	1    0    0    -1  
$EndComp
Text GLabel 7600 3900 2    60   Input ~ 0
MDB-MR
Wire Wire Line
	7600 1850 7600 1900
Wire Wire Line
	7550 1900 7600 1900
Connection ~ 7600 1900
Wire Wire Line
	7300 2150 7250 2150
Wire Wire Line
	7250 2100 7250 2150
Connection ~ 7250 2150
Wire Wire Line
	7600 2350 7600 2400
Wire Wire Line
	7600 2400 7750 2400
Connection ~ 7600 2400
Wire Wire Line
	2850 2150 2750 2150
Wire Wire Line
	2850 2050 2750 2050
Wire Wire Line
	7550 3850 7550 3900
Wire Wire Line
	7500 3900 7550 3900
Connection ~ 7550 3900
Wire Wire Line
	7200 3900 7150 3900
NoConn ~ 2250 2150
Text GLabel 4950 3600 0    60   Input ~ 0
TTL-TX
Text GLabel 4950 3850 0    60   Input ~ 0
TTL-RX
Wire Wire Line
	2750 1950 2800 1950
Wire Wire Line
	2800 1950 2800 2200
$Comp
L power:+5V #PWR02
U 1 1 5B093514
P 7550 3550
F 0 "#PWR02" H 7550 3400 50  0001 C CNN
F 1 "+5V" H 7550 3690 50  0000 C CNN
F 2 "" H 7550 3550 50  0001 C CNN
F 3 "" H 7550 3550 50  0001 C CNN
	1    7550 3550
	1    0    0    -1  
$EndComp
Wire Wire Line
	2250 2050 2200 2050
Wire Wire Line
	2200 2050 2200 2250
Wire Wire Line
	7600 1550 7250 1550
Wire Wire Line
	7250 1500 7250 1550
$Comp
L power:GND #PWR03
U 1 1 5B0943E7
P 7600 2850
F 0 "#PWR03" H 7600 2600 50  0001 C CNN
F 1 "GND" H 7600 2700 50  0000 C CNN
F 2 "" H 7600 2850 50  0001 C CNN
F 3 "" H 7600 2850 50  0001 C CNN
	1    7600 2850
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR04
U 1 1 5B09447A
P 2200 2250
F 0 "#PWR04" H 2200 2000 50  0001 C CNN
F 1 "GND" H 2200 2100 50  0000 C CNN
F 2 "" H 2200 2250 50  0001 C CNN
F 3 "" H 2200 2250 50  0001 C CNN
	1    2200 2250
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR05
U 1 1 5B0944A7
P 2800 2200
F 0 "#PWR05" H 2800 1950 50  0001 C CNN
F 1 "GND" H 2800 2050 50  0000 C CNN
F 2 "" H 2800 2200 50  0001 C CNN
F 3 "" H 2800 2200 50  0001 C CNN
	1    2800 2200
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR06
U 1 1 5B094546
P 4950 4100
F 0 "#PWR06" H 4950 3850 50  0001 C CNN
F 1 "GND" H 4950 3950 50  0000 C CNN
F 2 "" H 4950 4100 50  0001 C CNN
F 3 "" H 4950 4100 50  0001 C CNN
	1    4950 4100
	1    0    0    -1  
$EndComp
Wire Wire Line
	7600 2750 7600 2850
$Comp
L power:PWR_FLAG #FLG07
U 1 1 5B094A97
P 9250 1900
F 0 "#FLG07" H 9250 1975 50  0001 C CNN
F 1 "PWR_FLAG" H 9250 2050 50  0000 C CNN
F 2 "" H 9250 1900 50  0001 C CNN
F 3 "" H 9250 1900 50  0001 C CNN
	1    9250 1900
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR08
U 1 1 5B094AFA
P 9250 1900
F 0 "#PWR08" H 9250 1650 50  0001 C CNN
F 1 "GND" H 9250 1750 50  0000 C CNN
F 2 "" H 9250 1900 50  0001 C CNN
F 3 "" H 9250 1900 50  0001 C CNN
	1    9250 1900
	1    0    0    -1  
$EndComp
$Comp
L power:+5V #PWR09
U 1 1 5B094D2B
P 7250 1500
F 0 "#PWR09" H 7250 1350 50  0001 C CNN
F 1 "+5V" H 7250 1640 50  0000 C CNN
F 2 "" H 7250 1500 50  0001 C CNN
F 3 "" H 7250 1500 50  0001 C CNN
	1    7250 1500
	1    0    0    -1  
$EndComp
Connection ~ 7250 1550
Wire Wire Line
	2200 1900 2200 1950
Wire Wire Line
	2200 1950 2250 1950
$Comp
L power:+24V #PWR010
U 1 1 5B094E40
P 9250 1450
F 0 "#PWR010" H 9250 1300 50  0001 C CNN
F 1 "+24V" H 9250 1590 50  0000 C CNN
F 2 "" H 9250 1450 50  0001 C CNN
F 3 "" H 9250 1450 50  0001 C CNN
	1    9250 1450
	1    0    0    -1  
$EndComp
$Comp
L power:PWR_FLAG #FLG011
U 1 1 5B094EE4
P 9250 1450
F 0 "#FLG011" H 9250 1525 50  0001 C CNN
F 1 "PWR_FLAG" H 9250 1600 50  0000 C CNN
F 2 "" H 9250 1450 50  0001 C CNN
F 3 "" H 9250 1450 50  0001 C CNN
	1    9250 1450
	-1   0    0    1   
$EndComp
$Comp
L power:+5V #PWR012
U 1 1 5B095280
P 4950 3400
F 0 "#PWR012" H 4950 3250 50  0001 C CNN
F 1 "+5V" H 4950 3540 50  0000 C CNN
F 2 "" H 4950 3400 50  0001 C CNN
F 3 "" H 4950 3400 50  0001 C CNN
	1    4950 3400
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR013
U 1 1 5B09563E
P 2350 4950
F 0 "#PWR013" H 2350 4700 50  0001 C CNN
F 1 "GND" H 2350 4800 50  0000 C CNN
F 2 "" H 2350 4950 50  0001 C CNN
F 3 "" H 2350 4950 50  0001 C CNN
	1    2350 4950
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR014
U 1 1 5B095685
P 3350 4950
F 0 "#PWR014" H 3350 4700 50  0001 C CNN
F 1 "GND" H 3350 4800 50  0000 C CNN
F 2 "" H 3350 4950 50  0001 C CNN
F 3 "" H 3350 4950 50  0001 C CNN
	1    3350 4950
	1    0    0    -1  
$EndComp
$Comp
L power:+5V #PWR015
U 1 1 5B095758
P 3350 4550
F 0 "#PWR015" H 3350 4400 50  0001 C CNN
F 1 "+5V" H 3350 4690 50  0000 C CNN
F 2 "" H 3350 4550 50  0001 C CNN
F 3 "" H 3350 4550 50  0001 C CNN
	1    3350 4550
	1    0    0    -1  
$EndComp
$Comp
L power:+24V #PWR016
U 1 1 5B0957B3
P 2350 4550
F 0 "#PWR016" H 2350 4400 50  0001 C CNN
F 1 "+24V" H 2350 4690 50  0000 C CNN
F 2 "" H 2350 4550 50  0001 C CNN
F 3 "" H 2350 4550 50  0001 C CNN
	1    2350 4550
	1    0    0    -1  
$EndComp
$Comp
L conn:Conn_01x01 J2-1
U 1 1 5B0BD3DF
P 5150 3400
F 0 "J2-1" H 5150 3500 50  0000 C CNN
F 1 "Conn_01x01" H 5150 3300 50  0000 C CNN
F 2 "Connectors:Pin_d0.7mm_L6.5mm_W1.8mm_FlatFork" H 5150 3400 50  0001 C CNN
F 3 "" H 5150 3400 50  0001 C CNN
	1    5150 3400
	1    0    0    -1  
$EndComp
$Comp
L conn:Conn_01x01 J2-2
U 1 1 5B0BD4D2
P 5150 3600
F 0 "J2-2" H 5150 3700 50  0000 C CNN
F 1 "Conn_01x01" H 5150 3500 50  0000 C CNN
F 2 "Connectors:Pin_d0.7mm_L6.5mm_W1.8mm_FlatFork" H 5150 3600 50  0001 C CNN
F 3 "" H 5150 3600 50  0001 C CNN
	1    5150 3600
	1    0    0    -1  
$EndComp
$Comp
L conn:Conn_01x01 J2-3
U 1 1 5B0BD530
P 5150 3850
F 0 "J2-3" H 5150 3950 50  0000 C CNN
F 1 "Conn_01x01" H 5150 3750 50  0000 C CNN
F 2 "Connectors:Pin_d0.7mm_L6.5mm_W1.8mm_FlatFork" H 5150 3850 50  0001 C CNN
F 3 "" H 5150 3850 50  0001 C CNN
	1    5150 3850
	1    0    0    -1  
$EndComp
$Comp
L conn:Conn_01x01 J2-4
U 1 1 5B0BD597
P 5150 4100
F 0 "J2-4" H 5150 4200 50  0000 C CNN
F 1 "Conn_01x01" H 5150 4000 50  0000 C CNN
F 2 "Connectors:Pin_d0.7mm_L6.5mm_W1.8mm_FlatFork" H 5150 4100 50  0001 C CNN
F 3 "" H 5150 4100 50  0001 C CNN
	1    5150 4100
	1    0    0    -1  
$EndComp
$Comp
L transistors:2N7002 Q3
U 1 1 5B2C9902
P 6550 2050
F 0 "Q3" H 6750 2125 50  0000 L CNN
F 1 "2N7002" H 6750 2050 50  0000 L CNN
F 2 "TO_SOT_Packages_SMD:SOT-23" H 6750 1975 50  0001 L CIN
F 3 "" H 6550 2050 50  0001 L CNN
	1    6550 2050
	0    1    1    0   
$EndComp
$Comp
L Device:R R6
U 1 1 5B2C9A5A
P 6350 1800
F 0 "R6" V 6430 1800 50  0000 C CNN
F 1 "10k" V 6350 1800 50  0000 C CNN
F 2 "" V 6280 1800 50  0001 C CNN
F 3 "" H 6350 1800 50  0001 C CNN
	1    6350 1800
	0    1    1    0   
$EndComp
Wire Wire Line
	6000 2150 6150 2150
Wire Wire Line
	6150 2150 6150 1800
Wire Wire Line
	6150 1800 6200 1800
Wire Wire Line
	6550 1850 6550 1800
Wire Wire Line
	6600 1800 6550 1800
$Comp
L power:+3.3V #PWR?
U 1 1 5B2CA167
P 6600 1800
F 0 "#PWR?" H 6600 1650 50  0001 C CNN
F 1 "+3.3V" H 6600 1940 50  0000 C CNN
F 2 "" H 6600 1800 50  0001 C CNN
F 3 "" H 6600 1800 50  0001 C CNN
	1    6600 1800
	0    1    1    0   
$EndComp
Connection ~ 6550 1800
Wire Wire Line
	6900 2150 6750 2150
Connection ~ 6150 2150
$Comp
L Device:R R7
U 1 1 5B2CAF3E
P 7150 4150
F 0 "R7" V 7230 4150 50  0000 C CNN
F 1 "20K" V 7150 4150 50  0000 C CNN
F 2 "" V 7080 4150 50  0001 C CNN
F 3 "" H 7150 4150 50  0001 C CNN
	1    7150 4150
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR?
U 1 1 5B2CAFCD
P 7150 4400
F 0 "#PWR?" H 7150 4150 50  0001 C CNN
F 1 "GND" H 7150 4250 50  0000 C CNN
F 2 "" H 7150 4400 50  0001 C CNN
F 3 "" H 7150 4400 50  0001 C CNN
	1    7150 4400
	1    0    0    -1  
$EndComp
Wire Wire Line
	7150 4000 7150 3900
Connection ~ 7150 3900
Wire Wire Line
	7150 4300 7150 4400
Wire Wire Line
	7600 1900 7600 1950
Wire Wire Line
	7250 2150 7200 2150
Wire Wire Line
	7600 2400 7600 2450
Wire Wire Line
	7550 3900 7600 3900
Wire Wire Line
	7250 1550 7250 1700
Wire Wire Line
	6550 1800 6500 1800
Wire Wire Line
	6150 2150 6350 2150
Wire Wire Line
	7150 3900 7100 3900
$Comp
L Converter_DCDC:dc24-dc5-module psu?
U 1 1 5CD320C8
P 2850 4750
F 0 "psu?" H 2850 5217 50  0000 C CNN
F 1 "dc24-dc5-module" H 2850 5126 50  0000 C CNN
F 2 "" H 2850 5150 50  0001 C CNN
F 3 "" H 2850 5150 50  0001 C CNN
	1    2850 4750
	1    0    0    -1  
$EndComp
$EndSCHEMATC
