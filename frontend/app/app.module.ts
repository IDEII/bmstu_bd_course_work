import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

// import { RegisterComponent } from './register/register.component';
// import { RegistrationService } from './registration.service';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { provideAnimationsAsync } from '@angular/platform-browser/animations/async';
import { TournamentComponent } from './tournament/tournament.component';
import { RegistrationComponent } from './registration/registration.component';
import { AuthComponent } from './auth/auth.component';
import { AddClubComponent } from './add-club/add-club.component';
import { CreateTournamentComponent } from './create-tournament/create-tournament.component';
import { RequestOnTournamentComponent } from './request-on-tournament/request-on-tournament.component';
import { DrawComponent } from './draw/draw.component';
import { MatchComponent } from './match/match.component';
import { LoginComponent } from './login/login.component';
import { ReactiveFormsModule } from '@angular/forms';
import { AuthService } from './auth.service';
import { LogoutComponent } from './logout/logout.component';
import { HomeComponent } from './home/home.component';
import { FighterComponent } from './fighter/fighter.component';
import { ClubsComponent } from './clubs/clubs.component';
import { FightersComponent } from './fighters/fighters.component';
import { ClubProfileComponent } from './club-profile/club-profile.component';
import { TournamentListComponent } from './tournament-list/tournament-list.component';
import { OrganazerListComponent } from './organazer-list/organazer-list.component';
import { OrganazerProfileComponent } from './organazer-profile/organazer-profile.component';
import { AddOrganazerComponent } from './add-organazer/add-organazer.component';
import { AddFighterComponent } from './add-fighter/add-fighter.component';
import { MembersComponent } from './members/members.component';
import { InvateOnTournamentComponent } from './invate-on-tournament/invate-on-tournament.component';
import { RequestManagerComponent } from './request-manager/request-manager.component';
import { SelectedMembersComponent } from './selected-members/selected-members.component';
import { ResultsTableComponent } from './results-table/results-table.component';
import { DatePipe } from '@angular/common';
import { UserProfileComponent } from './user-profile/user-profile.component';
import { UserService } from './user.service'

@NgModule({
  declarations: [
    AppComponent,
    TournamentComponent,
    RegistrationComponent,
    AuthComponent,
    AddClubComponent,
    CreateTournamentComponent,
    RequestOnTournamentComponent,
    DrawComponent,
    MatchComponent,
    LoginComponent,
    LogoutComponent,
    HomeComponent,
    FighterComponent,
    ClubsComponent,
    FightersComponent,
    ClubProfileComponent,
    TournamentListComponent,
    OrganazerListComponent,
    OrganazerProfileComponent,
    AddOrganazerComponent,
    AddFighterComponent,
    MembersComponent,
    InvateOnTournamentComponent,
    RequestManagerComponent,
    SelectedMembersComponent,
    ResultsTableComponent,
    UserProfileComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule,
  ],
  providers: [
    AuthService,
    DatePipe
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }


